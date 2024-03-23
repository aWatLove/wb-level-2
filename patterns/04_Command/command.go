package main

import "fmt"

/*
	Реализовать паттерн "Command"
https://ru.wikipedia.org/wiki/Команда_(шаблон_проектирования)

Паттерн Command - поведенченский шаблон, представляющий действие.
Объект комманды заключает в себе само действие и его параметры.
Это включает множество плюсов, таких как:
+ Добавление посредничества между отправителем действия и получаетелем, что позволяет динамически менять реализацию команд
+ Возможность реализации отмены
+ Команда - отдельный объект, в который мы добавляем нужную логику, не меняя оригинальной структуры, мы следуем принципу открытости - закрытости
+ Возможность агрегации команд из более простых

Минусы:
- Усложнение кода из-за ввода доп. структур

*/

// Интерфейс команды
type Command interface {
	Execute()
}

// структура лампочки
type Light struct {
}

// методы включения и выключения лампочки
func (l Light) turnOn() {
	fmt.Println("lights turn on!")
}
func (l Light) turnOff() {
	fmt.Println("lights turn off!")
}

// конкретная реализация команды, полем которой является Light
type OnCommand struct {
	light Light
}

func (o OnCommand) Execute() {
	o.light.turnOn()
}

// конкретная реализация команды, полем которой является Light
type OffCommand struct {
	light Light
}

func (o OffCommand) Execute() {
	o.light.turnOff()
}

// Пример использования
func main() {
	light := Light{}

	onCommand := OnCommand{light}
	offCommand := OffCommand{light}

	onCommand.Execute()
	offCommand.Execute()
}
