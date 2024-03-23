package main

import "fmt"

/*
	Реализовать паттерн "Strategy"
https://ru.wikipedia.org/wiki/Стратегия_(шаблон_проектирования)

Паттерн Strategy - поведенческий шаблон, позволяет менять выбранную стратегию выполнения алгоритма, обеспечивая их взаимозаменяемость.

Плюсы:
+ Динамическая замена поведения
+ Изоляция логики от остальных объектов
+ Принцип открытости - закрытости

Минусы:
- Усложнение кода добавлением новых структур
- Клиент должен знать, какую именно стратегию применять
*/

// интерфейс стратегии
type Strategy interface {
	execute(a, b int) int
}

// структура Executor, выполняющая стратегию
type Executor struct {
	strategy Strategy
}

// метод, который устанавливает стратегию
func (e *Executor) setStrategy(strategy Strategy) {
	e.strategy = strategy
}

// метод выполняющий стратегию
func (e Executor) executeStrategy(a, b int) int {
	return e.strategy.execute(a, b)
}

// Конкретная реализация стратегии
type ConcreteStrategySum struct {
}

// метод, суммирующий две переменные типа int
func (s ConcreteStrategySum) execute(a, b int) int {
	return a + b
}

// Конкретная реализация стратегии
type ConcreteStrategyMul struct {
}

// метод, перемножающий две переменные типа int
func (s ConcreteStrategyMul) execute(a, b int) int {
	return a * b
}

// Пример использования
func main() {
	// объявляем стратегии
	sum := ConcreteStrategySum{}
	mul := ConcreteStrategyMul{}

	// объявляем Executor
	executor := Executor{}

	// присваиваем стратегию суммирования
	executor.setStrategy(sum)
	fmt.Println(executor.executeStrategy(2, 3))

	// присваиваем стратегию умножения
	executor.setStrategy(mul)
	fmt.Println(executor.executeStrategy(2, 3))
}
