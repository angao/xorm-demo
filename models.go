package main

import (
	"log"
    _ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"errors"
)

type Account struct {
	Id int64
	Name string `xorm:"unique"`
	Balance float64
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
	Version int64 `xorm:"version"`
}

var x *xorm.Engine

func init() {
    var err error
	x, err = xorm.NewEngine("mysql", "root:123@tcp(192.168.31.138:3306)/test?charset=utf8")
	if err != nil {
		log.Fatalf("db can't connect: %v\n", err)
	}
	err = x.Sync(new(Account))
	if err != nil {
		log.Fatal(err)
	}
	// x.ShowSQL(true)
}

func NewAccount(name string, balance float64) error {
	var a = &Account{
		Name: name,
		Balance: balance,
	}
	_, err := x.Insert(a)
	return err
}

func GetAccount(id int64) (*Account, error) {
	a := &Account{}
	has, err := x.Id(id).Get(a)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("account not found")
	}
	return a, err
}

// MakeDeposit 存款
func MakeDeposit(id int64, deposit float64) (*Account, error) {
	account, err := GetAccount(id)
	if err != nil {
		return nil, err
	}
	account.Balance += deposit

	_, err = x.Update(account)
	return account, err
}

// MakeWithdraw 取款
func MakeWithdraw(id int64, withdraw float64) (*Account, error) {
	account, err := GetAccount(id)
	if err != nil {
		return nil, err
	}
	if account.Balance < withdraw {
		return nil, errors.New("balance not enough")
	}
	account.Balance -= withdraw

	_, err = x.Update(account)
	return account, err
}

// MakeTransfer 转账
func MakeTransfer(out int64, money float64, in int64) error {
	outAcc, err := GetAccount(out)
	if err != nil {
		return err
	}

	inAcc, err := GetAccount(in)
	if err != nil {
		return err
	}

	if outAcc.Balance < money {
		return errors.New("out balance not enough")
	}

	outAcc.Balance -= money
	inAcc.Balance += money

	session := x.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}

	if _, err := session.Update(outAcc); err != nil {
		session.Rollback()
		return err
	}
	if _, err := session.Update(inAcc); err != nil {
		session.Rollback()
		return err
	}

	session.Commit()
	return nil
}

// DeleteAccount 删除
func DeleteAccount(id int64) error {
	 _,err := x.Delete(&Account{Id: id})
	 return err
}

// ListAccounts 
func ListAccounts()(as []*Account, err error)  {
	err = x.Find(&as)
	return as, err
}