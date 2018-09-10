package accounts

import (
	"errors"
	"log"
	"strings"
	"sync"
	"time"

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
	account       map[pbAPI.Venue]map[string]*stm.TVar
	TotalAccounts int
}

// Balances ...
type AccountBalances struct {
	Symbol    string  `db:"symbol"`
	VenueID   int     `db:"venue_id"`
	AccountID int     `db:"account_id"`
	UsersID   int     `db:"users_id"`
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
	UserAccounts.account = make(map[pbAPI.Venue]map[string]*stm.TVar)
	UserAccounts.LoadDataFromDB()
}

// LoadDataFromDB get the accounts from the database and put it in memory
func (m *Account) LoadDataFromDB() {

	logrus.Info("Loading accounts balances to memory")

	balances := []AccountBalances{}
	if err := postgres.PostgresDB.Select(&balances, "SELECT s.name as symbol,b.venue_id,b.account_id,a.users_id,b.hold,b.available FROM balance b, symbol s, account a WHERE b.symbol_id=s.symbol_id AND a.account_id=b.account_id"); err != nil {
		log.Fatal(err)
	}
	logrus.Info("HERE ", balances)

	for i := 0; i < 1000; i++ {
		data := &pbAPI.Account{
			Id:       time.Now().String(),
			User:     time.Now().String(),
			Venue:    pbAPI.Venue_BINANCE,
			Active:   true,
			Balances: make(map[string]*pbAPI.Balance),
		}

		data.Balances[pbAPI.Symbol_BTC.String()] = &pbAPI.Balance{
			Available: 500000,
			Hold:      300,
		}
		data.Balances[pbAPI.Symbol_USD.String()] = &pbAPI.Balance{
			Available: 500000,
			Hold:      300,
		}

		m.LoadDataToMemory(data)
	}

	data := &pbAPI.Account{
		Id:       "XXXXXX",
		User:     "XXXXXX",
		Venue:    pbAPI.Venue_COINBASEPRO,
		Active:   true,
		Balances: make(map[string]*pbAPI.Balance),
	}

	data.Balances[pbAPI.Symbol_BTC.String()] = &pbAPI.Balance{
		Available: 500000,
		Hold:      300,
	}
	data.Balances[pbAPI.Symbol_USD.String()] = &pbAPI.Balance{
		Available: 500000,
		Hold:      300,
	}

	m.LoadDataToMemory(data)
}

// LoadDataToMemory data onto memory
func (m *Account) LoadDataToMemory(data *pbAPI.Account) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	m.account[data.GetVenue()] = make(map[string]*stm.TVar)
	newAccount := stm.New(data)
	m.account[data.GetVenue()][string(data.GetId())] = newAccount
	m.TotalAccounts++
}

// ValidateAndUpdateBalances data onto memory
func (m *Account) ValidateAndUpdateBalances(venue pbAPI.Venue, product pbAPI.Product, account string, amount float64) (*pbAPI.Account, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	var err error
	// Here is before trading, if things go well, this needs to update again the hold and available to the correct amount based on the execution
	if venue, ok := m.account[venue]; ok {
		if accountNumber, ok := venue[account]; ok {
			transfer := func(rec *stm.TRec) interface{} {
				account := rec.Load(accountNumber).(*pbAPI.Account)
				symbols := strings.Split(product.String(), "_")
				if balance, ok := account.Balances[symbols[1]]; ok {

					if balance.Available >= amount && amount > 0 { // we dont want users sending negative amounts
						balance.Available = balance.Available - amount
						balance.Hold = balance.Hold + amount
						rec.Store(accountNumber, account)
					} else {
						err = errors.New("Balance is not enough to execute this operation")
						return new(pbAPI.Account)
					}
					return account

				} else {
					err = errors.New("Venue does not support this product")
					return new(pbAPI.Account)
				}
			}
			return stm.Atomically(transfer).(*pbAPI.Account), err
		} else {
			err = errors.New("Account was not found")
			return new(pbAPI.Account), err
		}
	} else {
		err = errors.New("Balances not found for this venue")
		return new(pbAPI.Account), err
	}
}

// RefundAccountValues balances to the user account
func (m *Account) RefundAccountValues(venue pbAPI.Venue, product pbAPI.Product, account string, amount float64) (*pbAPI.Account, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	var err error
	// Here is before trading, if things go well, this needs to update again the hold and available to the correct amount based on the execution
	if venue, ok := m.account[venue]; ok {
		if accountNumber, ok := venue[account]; ok {
			transfer := func(rec *stm.TRec) interface{} {
				account := rec.Load(accountNumber).(*pbAPI.Account)
				symbols := strings.Split(product.String(), "_")
				if balance, ok := account.Balances[symbols[1]]; ok {

					// Refund values
					balance.Available = balance.Available + amount
					balance.Hold = balance.Hold - amount
					rec.Store(accountNumber, account)

					return account

				} else {
					err = errors.New("Venue does not support this product")
					return new(pbAPI.Account)
				}
			}
			return stm.Atomically(transfer).(*pbAPI.Account), err
		} else {
			err = errors.New("Account was not found")
			return new(pbAPI.Account), err
		}
	} else {
		err = errors.New("Balances not found for this venue")
		return new(pbAPI.Account), err
	}
}
