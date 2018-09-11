package accounts

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/decillion/go-stm"
	"github.com/maurodelazeri/lion/postgres"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/sirupsen/logrus"
)

// UserAccounts holds all active accounts in memory
var UserAccounts *Account

// Account ...
type Account struct {
	mutex         sync.RWMutex
	account       map[pbAPI.Venue]map[uint32]*stm.TVar
	TotalAccounts int
}

// UserAccount ...
type UserAccount struct {
	Active    bool   `db:"active"`
	AccountID uint32 `db:"account_id"`
	UsersID   uint32 `db:"users_id"`
}

// AccountBalances ...
type AccountBalances struct {
	Symbol    uint32  `db:"symbol_id"`
	VenueID   uint32  `db:"venue_id"`
	AccountID uint32  `db:"account_id"`
	Hold      float64 `db:"hold"`
	Available float64 `db:"available"`
}

//https://play.golang.org/p/ib-dfXjPDy
// sync.Map might be faster, need to test it later
// Initialize ...
func init() {
	UserAccounts = new(Account)
	UserAccounts.mutex.RLock()
	defer UserAccounts.mutex.RUnlock()
	UserAccounts.account = make(map[pbAPI.Venue]map[uint32]*stm.TVar)
	UserAccounts.LoadDataFromDB()
}

// LoadDataFromDB get the accounts from the database and put it in memory
func (m *Account) LoadDataFromDB() {
	logrus.Info("Loading accounts balances to memory")
	accounts := []UserAccount{}
	if err := postgres.PostgresDB.Select(&accounts, "SELECT account_id,users_id,active FROM account WHERE active=true AND account_mode_id="+strconv.FormatInt(int64(pbAPI.AccountMode_value[os.Getenv("MODE")]), 10)); err != nil {
		log.Fatal(err)
	}
	balances := []AccountBalances{}
	if err := postgres.PostgresDB.Select(&balances, "SELECT b.symbol_id,b.venue_id,b.account_id,b.hold,b.available FROM balance b,account a WHERE a.account_id=b.account_id AND a.active=true ORDER BY account_id ASC"); err != nil {
		log.Fatal(err)
	}
	for _, account := range accounts {
		data := &pbAPI.Account{
			AccountId: account.AccountID,
			UserId:    account.UsersID,
			Active:    account.Active,
			Balances:  make(map[string]*pbAPI.Balance),
		}
		for i := 0; i < len(balances); i++ {
			if balances[i].AccountID == account.AccountID {
				data.Venue = pbAPI.Venue(balances[i].VenueID)
				data.Balances[pbAPI.Symbol(balances[i].Symbol).String()] = &pbAPI.Balance{
					Available: balances[i].Available,
					Hold:      balances[i].Hold,
				}
			}
		}
		m.LoadDataToMemory(data)
	}
}

// LoadDataToMemory data onto memory
func (m *Account) LoadDataToMemory(data *pbAPI.Account) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if _, ok := m.account[data.GetVenue()]; !ok {
		m.account[data.GetVenue()] = make(map[uint32]*stm.TVar)
	}
	newAccount := stm.New(data)
	m.account[data.GetVenue()][data.GetAccountId()] = newAccount
	m.TotalAccounts++
}

// Order Type
// Buy 0 2 4 6
// Sell 1 3 5 7

// ValidateAndUpdateBalances data onto memory
func (m *Account) ValidateAndUpdateBalances(orderType pbAPI.OrderType, venue pbAPI.Venue, product pbAPI.Product, account uint32, amount float64, price float64, mode string) (*pbAPI.Account, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	var err error
	if venue, ok := m.account[venue]; ok {
		if accountNumber, ok := venue[account]; ok {
			transfer := func(rec *stm.TRec) interface{} {
				account := rec.Load(accountNumber).(*pbAPI.Account)
				symbols := strings.Split(product.String(), "_")
				if pbAPI.OrderType_value[orderType.String()] == 0 || pbAPI.OrderType_value[orderType.String()] == 2 || pbAPI.OrderType_value[orderType.String()] == 4 || pbAPI.OrderType_value[orderType.String()] == 6 {
					// Buying
					if balance, ok := account.Balances[symbols[1]]; ok {
						if balance.Available >= price && price > 0 { // we dont want users sending negative prices
							switch mode {
							case "transaction-hold":
								balance.Available = balance.Available - price
								balance.Hold = balance.Hold + price
								rec.Store(accountNumber, account)
							case "transaction-done":
								balance.Hold = balance.Hold - price
								if balanceQuote, ok := account.Balances[symbols[0]]; ok {
									balanceQuote.Available = balanceQuote.Available + amount
								}
								rec.Store(accountNumber, account)
							case "transaction-refund":
								balance.Available = balance.Available + price
								balance.Hold = balance.Hold - price
								rec.Store(accountNumber, account)
							default:
								logrus.Error("ValidateAndUpdateBalances - Mode is not valid ", mode)
							}
						} else {
							err = errors.New("Balance is not enough to execute this operation")
							return new(pbAPI.Account)
						}
						return account
					}
				} else if pbAPI.OrderType_value[orderType.String()] == 1 || pbAPI.OrderType_value[orderType.String()] == 3 || pbAPI.OrderType_value[orderType.String()] == 5 || pbAPI.OrderType_value[orderType.String()] == 7 {
					// Selling
					if balance, ok := account.Balances[symbols[0]]; ok {
						if balance.Available >= amount && amount > 0 { // we dont want users sending negative amounts
							switch mode {
							case "transaction-hold":
								balance.Available = balance.Available - amount
								balance.Hold = balance.Hold + amount
								rec.Store(accountNumber, account)
							case "transaction-done":
								balance.Hold = balance.Hold - amount
								if balanceQuote, ok := account.Balances[symbols[1]]; ok {
									balanceQuote.Available = balanceQuote.Available + price
								}
								rec.Store(accountNumber, account)
							case "transaction-refund":
								balance.Available = balance.Available + amount
								balance.Hold = balance.Hold - amount
								rec.Store(accountNumber, account)
							default:
								logrus.Error("ValidateAndUpdateBalances - Mode is not valid ", mode)
							}
						} else {
							err = errors.New("Balance is not enough to execute this operation")
							return new(pbAPI.Account)
						}
						return account
					}
				}
				err = errors.New("Venue does not support this product")
				return new(pbAPI.Account)
			}
			return stm.Atomically(transfer).(*pbAPI.Account), err
		}
		err = errors.New("Account was not found")
		return new(pbAPI.Account), err
	}
	err = errors.New("Balances not found for this venue")
	return new(pbAPI.Account), err
}
