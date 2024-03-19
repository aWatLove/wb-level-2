package main

import "fmt"

/*
	Реализовать паттерн "Фасад"
https://en.wikipedia.org/wiki/Facade_pattern

Паттерн фасад - структурный шаблон, позволяющий скрыть сложность системы путем сведения
всех внешних вызовов к однму объекту.

*/

type DB struct {
}

func (d DB) connect() {
	fmt.Println("Start DB")
}

type Cache struct {
}

func (c Cache) startCache() {
	fmt.Println("Cache is started!")
}

func (c Cache) updateData(db DB) {
	fmt.Println("Cache is updated! from db:")
}

type Repo struct {
	db    DB
	cache Cache
}

func NewRepo() *Repo {
	return &Repo{db: DB{}, cache: Cache{}}
}

func (r Repo) start() {
	r.db.connect()
	r.cache.startCache()
	r.cache.updateData(r.db)
}

func main() {
	r := Repo{}
	r.start()
}
