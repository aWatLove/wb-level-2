package main

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн "Factory method"
https://ru.wikipedia.org/wiki/Фабричный_метод_(шаблон_проектирования)

Паттерн Factory method - порождающий шаблон, предоставляющий подклассам интерфейс для создания экземпляров некоторого класса.

Применимость:
- Необходимость разделения создания структур от их использования. Когда мы не знаем какую конкретно структуру будем использовать
- Возможность расширения, создавая новые типы, не ломая старой логики создания объектов
- ПОвторное использование уже существующих объектов


Так как в Go нет классического наследования, можно реализовать данный паттерн с использованием встраивания
*/

// интерфейс банковского аккаунта
type IBankAccount interface {
	setBalance(float64)
	getBalance() float64
	setMoneyLimit(float64)
	getMoneyLimit() float64
	setAnnualInterest(float64)
	getAnnualInterest() float64
}

// структура банковского аккаунта с нужными полями
type BankAccount struct {
	balance        float64
	moneyLimit     float64
	annualInterest float64
}

// реализация методов интерфейса
func (b *BankAccount) setBalance(f float64) {
	b.moneyLimit = f
}

func (b *BankAccount) getBalance() float64 {
	return b.balance
}

func (b *BankAccount) setMoneyLimit(f float64) {
	b.moneyLimit = f
}

func (b *BankAccount) getMoneyLimit() float64 {
	return b.moneyLimit
}

func (b *BankAccount) setAnnualInterest(f float64) {
	b.annualInterest = f
}

func (b *BankAccount) getAnnualInterest() float64 {
	return b.annualInterest
}

// структура дебетового аккаунта со встраиванием анонимного поля типа BankAccount
type DebitAccount struct {
	BankAccount
}

// Метод, возвращающий конкретный объект дебетового аккаунта
func newDebitAccount(balance, limit, percent float64) IBankAccount {
	return &DebitAccount{BankAccount{
		balance:        balance,
		moneyLimit:     limit,
		annualInterest: percent,
	}}
}

// структура кредитного аккаунта со встраиванием анонимного поля типа BankAccount
type CreditAccount struct {
	BankAccount
}

// Метод, возвращающий конкретный объект дебетового аккаунта
func newCreditAccount(balance, limit, percent float64) IBankAccount {
	return &CreditAccount{BankAccount{
		balance:        balance,
		moneyLimit:     limit,
		annualInterest: percent,
	}}
}

func getBankAccount(accountType string, balance, limit, percent float64) (IBankAccount, error) {
	if accountType == "debitAccount" {
		return newDebitAccount(balance, limit, percent), nil
	}
	if accountType == "creditAccount" {
		return newCreditAccount(balance, limit, percent), nil
	}
	return nil, errors.New("wrong account type")
}

// пример использования
func main() {
	creditAcc, _ := getBankAccount("creditAccount", 0, 15000, 15)
	fmt.Println(creditAcc)

}
