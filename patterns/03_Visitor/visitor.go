package main

import "fmt"

/*
	Реализовать паттерн "Visitor"
https://ru.wikipedia.org/wiki/Посетитель_(шаблон_проектирования)

Паттерн Visitor - поведенченский шаблон, описывающий операцию, которая выполняется над объектами других структур.
К примеру, при добавлении нового функционала, нет необходимости изменять обслуживаемую структуру
*/

// Новостной сервис
type NewsService struct {
	url   string
	news  string
	count int
}

// Реализуемые методы интерфейса Service
func (n *NewsService) accept(v Visitor) {
	v.visitForNewsService(n)
}

// Рекламный сервис
type AdsService struct {
	url string
	ads int
}

// Реализуемые методы интерфейса Service
func (a *AdsService) accept(v Visitor) {
	v.visitForAdsService(a)
}

// Сервис журнала
type MagazineService struct {
	url     string
	article int
}

// Реализуемые методы интерфейса Service
func (m *MagazineService) accept(v Visitor) {
	v.visitForMagazineService(m)
}

// Интерфейс сервисов
type Service interface {
	// метод для реализации шаблона "Visitor"
	accept(v Visitor)
	// some else methods...
}

// Интерфейс Visitor
type Visitor interface {
	visitForNewsService(*NewsService)
	visitForAdsService(*AdsService)
	visitForMagazineService(*MagazineService)
}

// Структура с новым функционалом для сервисов
// В нашем примере, просто выводит int переменную из сервиса
type ServicesCalculator struct {
}

// Реализация методов интерфейса Visitor
func (s ServicesCalculator) visitForNewsService(service *NewsService) {
	fmt.Println("NewsService:", service.count)
}

func (s ServicesCalculator) visitForAdsService(service *AdsService) {
	fmt.Println("AdsService:", service.ads)
}

func (s ServicesCalculator) visitForMagazineService(service *MagazineService) {
	fmt.Println("MagazinesService:", service.article)
}

// Пример использования
func main() {
	// Объявляем сервисв
	// NewsService
	ns := NewsService{url: "https://some.news", news: "good news!!!", count: 4}
	// AdsService
	as := AdsService{url: "http://some.ads.ru", ads: 3}
	// MagazineService
	ms := MagazineService{url: "https://some.magazine.ru", article: 12}

	// объявляем ServiceCalculator
	var sc ServicesCalculator
	// вызываем методы accept
	ns.accept(sc)
	as.accept(sc)
	ms.accept(sc)
}
