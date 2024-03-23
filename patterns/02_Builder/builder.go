package main

import "fmt"

/*
	Реализовать паттерн "Builder"
https://ru.wikipedia.org/wiki/Строитель_(шаблон_проектирования)

Паттерн Builder - порождающий шаблон, который позволяет поэтапно создавать сложные объекты с помощью
четко определенной последовательности действий
*/

type Human struct {
	FirstName string
	LastName  string
	Age       int
	Job       string
}

type HumanBuilder struct {
	firstName string
	lastName  string
	age       int
	job       string
}

func (b *HumanBuilder) FirstName(val string) *HumanBuilder {
	b.firstName = val
	return b
}

func (b *HumanBuilder) LastName(val string) *HumanBuilder {
	b.lastName = val
	return b
}

func (b *HumanBuilder) Age(val int) *HumanBuilder {
	b.age = val
	return b
}

func (b *HumanBuilder) Job(val string) *HumanBuilder {
	b.job = val
	return b
}

func (b HumanBuilder) Build() Human {
	return Human{
		FirstName: b.firstName,
		LastName:  b.lastName,
		Age:       b.age,
		Job:       b.job,
	}
}

func main() {
	hb := HumanBuilder{}
	hb.FirstName("Vladislav").
		LastName("Suvorov").
		Age(21).
		Job("go developer")
	hum := hb.Build()
	fmt.Println(hum)
}
