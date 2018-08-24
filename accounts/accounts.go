package accounts

import (
	"errors"
	"strings"
	"sync"
	"time"

	pb "github.com/maurodelazeri/elliptor/api"
	"github.com/sirupsen/logrus"

	"github.com/decillion/go-stm"
)

// UserAccounts holds all active accounts in memory
var UserAccounts *Account

// Account ...
type Account struct {
	mutex         sync.RWMutex
	account       map[pb.Venue]map[string]*stm.TVar
	TotalAccounts int
}

//https://play.golang.org/p/ib-dfXjPDy
// sync.Map might be faster, need to test it later
// Initialize ...
func init() {
	UserAccounts = new(Account)
	UserAccounts.mutex.RLock()
	defer UserAccounts.mutex.RUnlock()
	UserAccounts.account = make(map[pb.Venue]map[string]*stm.TVar)
	UserAccounts.LoadDataFromDB()
}

// LoadDataFromDB get the accounts from the database and put it in memory
func (m *Account) LoadDataFromDB() {

	logrus.Info("Loading accounts balances to memory")

	for i := 0; i < 1000; i++ {
		data := &pb.Account{
			Id:       time.Now().String(),
			User:     time.Now().String(),
			Venue:    pb.Venue_BINANCE,
			Active:   true,
			Balances: make(map[string]*pb.Balance),
		}

		data.Balances[pb.Symbol_BTC.String()] = &pb.Balance{
			Available: 500000,
			Hold:      300,
		}
		data.Balances[pb.Symbol_USD.String()] = &pb.Balance{
			Available: 500000,
			Hold:      300,
		}

		m.LoadDataToMemory(data)
	}

	data := &pb.Account{
		Id:       "XXXXXX",
		User:     "XXXXXX",
		Venue:    pb.Venue_BINANCE,
		Active:   true,
		Balances: make(map[string]*pb.Balance),
	}

	data.Balances[pb.Symbol_BTC.String()] = &pb.Balance{
		Available: 500000,
		Hold:      300,
	}
	data.Balances[pb.Symbol_USD.String()] = &pb.Balance{
		Available: 500000,
		Hold:      300,
	}

	m.LoadDataToMemory(data)
}

// LoadDataToMemory data onto memory
func (m *Account) LoadDataToMemory(data *pb.Account) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	m.account[data.GetVenue()] = make(map[string]*stm.TVar)
	newAccount := stm.New(data)
	m.account[data.GetVenue()][string(data.GetId())] = newAccount
	m.TotalAccounts++
}

// ValidateAndUpdateBalances data onto memory
func (m *Account) ValidateAndUpdateBalances(venue pb.Venue, product pb.Product, account string, amount float64) (*pb.Account, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	var err error
	// Here is before trading, if things go well, this needs to update again the hold and available to the correct amount based on the execution
	if venue, ok := m.account[venue]; ok {
		if accountNumber, ok := venue[account]; ok {
			transfer := func(rec *stm.TRec) interface{} {
				account := rec.Load(accountNumber).(*pb.Account)
				symbols := strings.Split(product.String(), "_")
				if balance, ok := account.Balances[symbols[1]]; ok {
					if balance.Available >= amount && amount > 0 { // we dont want users sending negative amounts
						balance.Available = balance.Available - amount
						balance.Hold = balance.Hold + amount
						rec.Store(accountNumber, account)
					} else {
						err = errors.New("Balance is not enough to execute this operation")
						return new(pb.Account)
					}
					return account
				} else {
					err = errors.New("Venue does not support this product")
					return new(pb.Account)
				}
			}
			return stm.Atomically(transfer).(*pb.Account), err
		} else {
			err = errors.New("Account was not found")
			return new(pb.Account), err
		}
	} else {
		err = errors.New("Balances not found for this venue")
		return new(pb.Account), err
	}
}
