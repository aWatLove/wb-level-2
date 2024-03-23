package main

import "fmt"

/*
	Реализовать паттерн "Chain of responsibility"
https://ru.wikipedia.org/wiki/Цепочка_обязанностей

Паттерн Chain of responsibility - поведенченский шаблон, для огранизации в системе уровней ответственности.

*/

type WorkaroundSheet struct {
	text           string
	isAccounting   bool
	isHRDep        bool
	isSafetyDep    bool
	isHouseholdDep bool
	isBoss         bool
}

type Check interface {
	Execute(*WorkaroundSheet)
	SetNext(Check)
}

// Первый обработчик - бухгалтерия, поле типа Check, указывающего на следующего обработчика
type Accounting struct {
	next Check
}

func (a *Accounting) Execute(sheet *WorkaroundSheet) {
	if sheet.isAccounting { // если этап уже пройден -> переходим к следующему
		a.next.Execute(sheet)
		return
	}
	sheet.isAccounting = true // в противном случае этап успешно завершен
	a.next.Execute(sheet)     // также переходим к следующему
}

func (a *Accounting) SetNext(check Check) {
	a.next = check
}

// обработчик - HR department
type HRDep struct {
	next Check
}

func (a *HRDep) Execute(sheet *WorkaroundSheet) {
	if sheet.isHRDep { // если этап уже пройден -> переходим к следующему
		a.next.Execute(sheet)
		return
	}
	sheet.isHRDep = true  // в противном случае этап успешно завершен
	a.next.Execute(sheet) // также переходим к следующему
}

func (a *HRDep) SetNext(check Check) {
	a.next = check
}

// обработчик - Safety department
type SafetyDep struct {
	next Check
}

func (a *SafetyDep) Execute(sheet *WorkaroundSheet) {
	if sheet.isSafetyDep { // если этап уже пройден -> переходим к следующему
		a.next.Execute(sheet)
		return
	}
	sheet.isSafetyDep = true // в противном случае этап успешно завершен
	a.next.Execute(sheet)    // также переходим к следующему
}

func (a *SafetyDep) SetNext(check Check) {
	a.next = check
}

// обработчик - Household department
type HouseholdDep struct {
	next Check
}

func (a *HouseholdDep) Execute(sheet *WorkaroundSheet) {
	if sheet.isHouseholdDep { // если этап уже пройден -> переходим к следующему
		a.next.Execute(sheet)
		return
	}
	sheet.isHouseholdDep = true // в противном случае этап успешно завершен
	a.next.Execute(sheet)       // также переходим к следующему
}

func (a *HouseholdDep) SetNext(check Check) {
	a.next = check
}

// последний обработчик - Boss
type Boss struct {
	next Check
}

func (a *Boss) Execute(sheet *WorkaroundSheet) {
	if sheet.isBoss { // если этап уже пройден -> выходим из фукнции: следующего этапа нет
		return
	}
	sheet.isBoss = true
}

func (a *Boss) SetNext(check Check) {
	a.next = check
}

// Пример использования
func main() { // создаем сущности с конца, т.к. необходимо в поля обработчиков при их инициализации указывать сущность для следующей обработки
	boss := Boss{}

	householdDep := HouseholdDep{}
	householdDep.SetNext(&boss)

	safetyDep := SafetyDep{}
	safetyDep.SetNext(&householdDep)

	hrDep := HRDep{}
	hrDep.SetNext(&safetyDep)

	accounting := Accounting{}
	accounting.SetNext(&hrDep)

	ws := WorkaroundSheet{text: "Some text in the workaround sheet"}
	accounting.Execute(&ws)

	fmt.Println(ws)
}
