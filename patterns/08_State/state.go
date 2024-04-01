package main

import (
	"errors"
	"fmt"
	"log"
)

/*
	Реализовать паттерн "State"
https://ru.wikipedia.org/wiki/Состояние_(шаблон_проектирования)

Паттерн State - поведенческий шаблон, используется, когда во время выполнения программы,
объект должен менять свое поведение в зависимости от своего состояния.

Плюсы:
+ Избавление от больщих условных конструкций (или свитчей)
+ Упрощение читабельности. Весь код относящийся к одному состоянию объекта находится в одном месте

Минусы:
- Оверкилл для маленького количества состояний
*/

// Интерфейс состояния
type State interface {
	requestItem() error
	addItem(int) error
	insertMoney(money int) error
	dispenseItem() error
}

// Структура сервиса со всеми возможными состояниями
type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	currentState  State
	itemCount     int
	itemPrice     int
}

// Конструктор сервиса
func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{v: v}
	itemRequestedState := &ItemRequestedState{v: v}
	hasMoneyState := &HasMoneyState{v: v}
	noItemState := &NoItemState{v: v}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

// методы сервиса
func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendingMachine) incrementItemCount(count int) {
	fmt.Println("Adding items:", count)
	v.itemCount += count
}

// Ниже реализации конкретных состояний с реализацией методов интерфейса State
type NoItemState struct {
	v *VendingMachine
}

func (i *NoItemState) requestItem() error {
	return errors.New("item out of stock")
}

func (i *NoItemState) addItem(count int) error {
	i.v.incrementItemCount(count)
	i.v.setState(i.v.hasItem)
	return nil
}

func (i *NoItemState) insertMoney(money int) error {
	return errors.New("item out of stock")
}

func (i *NoItemState) dispenseItem() error {
	return errors.New("item out of stock")
}

// -------------------
type HasItemState struct {
	v *VendingMachine
}

func (i *HasItemState) requestItem() error {
	if i.v.itemCount == 0 {
		i.v.setState(i.v.noItem)
		return errors.New("no item")
	}
	fmt.Println("Item requested")
	i.v.setState(i.v.itemRequested)
	return nil
}

func (i *HasItemState) addItem(count int) error {
	fmt.Println(count, "items added")
	i.v.incrementItemCount(count)
	return nil
}

func (i *HasItemState) insertMoney(money int) error {
	return errors.New("no item has been selected")
}

func (i *HasItemState) dispenseItem() error {
	return errors.New("no item has been selected")
}

// -------------------
type ItemRequestedState struct {
	v *VendingMachine
}

func (i *ItemRequestedState) requestItem() error {
	return errors.New("item already requested")
}

func (i *ItemRequestedState) addItem(count int) error {
	return errors.New("item dispense in progress")
}

func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.v.itemPrice {
		return errors.New(fmt.Sprintf("inserted money is less. Please insert %d", i.v.itemPrice))
	}
	fmt.Println("Money entered is ok")
	i.v.setState(i.v.hasMoney)
	return nil
}
func (i *ItemRequestedState) dispenseItem() error {
	return errors.New("please insert money first")
}

// -------------------
type HasMoneyState struct {
	v *VendingMachine
}

func (i *HasMoneyState) requestItem() error {
	return errors.New("item dispense in progress")
}

func (i *HasMoneyState) addItem(count int) error {
	return errors.New("item dispense in progress")
}

func (i *HasMoneyState) insertMoney(money int) error {
	return errors.New("item out of stock")
}
func (i *HasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing item")
	i.v.itemCount = i.v.itemCount - 1
	if i.v.itemCount == 0 {
		i.v.setState(i.v.noItem)
	} else {
		i.v.setState(i.v.hasItem)
	}
	return nil
}

// -------------------
// Пример использования
func main() {
	vm := NewVendingMachine(1, 50)

	err := vm.requestItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vm.insertMoney(50)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vm.dispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	// добавим предметов
	err = vm.addItem(5)
	if err != nil {
		log.Fatal(err.Error())
	}

	// повторим операции
	err = vm.requestItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vm.insertMoney(50)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vm.dispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}

}
