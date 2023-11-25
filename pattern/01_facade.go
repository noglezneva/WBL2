package main

import "fmt"

/*
Паттерн "Фасад" является структурным паттерном проектирования, который позволяет предоставить
упрощенный интерфейс к сложной системе объектов.
Он выступает в роли упаковки для других классов, скрывая их сложность
и предоставляя простой интерфейс для взаимодействия с системой.

Применимость:

* Когда нужно предоставить простой интерфейс для сложной подсистемы.
* Когда необходимо уменьшить зависимости между клиентским кодом и подсистемой.
* Когда нужно разложить сложную подсистему на отдельные уровни абстракции.


Основной плюс - изолирует клиентов от поведения сложной системы

Основной минус - интерфейс фасада может стать супер-объектом (то есть
мы будем привязаны к этой системе, все последующие функции будут проходить через
этот супер-объект, который хранит в себе «слишком много» или делает «слишком много» ).
Если подсистема функционально разнообразна, фасад может стать слишком объемным и сложным для поддержки.

Примеры использования паттерна "Фасад" на практике:

* Библиотеки API: Фасадный класс может быть создан для облегчения использования сложного программного интерфейса (API) библиотеки. Фасад предоставляет более простую и понятную абстракцию API, скрывая сложность внутренних процессов библиотеки.
* Фреймворки: Многие фреймворки используют фасады для облегчения взаимодействия с различными частями фреймворка. Например, фасады могут быть использованы для упрощения доступа к базе данных, управления сессиями пользователя или взаимодействия с внешними сервисами.
С использованием паттерна "Фасад" разработчики могут сократить сложность кода, упростить интерфейсы и улучшить подход к обслуживанию программного обеспечения.

*/

/*
Реализация: пусть у нас есть магазин, оторый предлагает
широкий ассортимент товаров, начиная от электроники до
продуктов питания. Когда клиент делает заказ, система доставки
должна заботиться о различных аспектах доставки, таких как логистика,
упаковка и отправка.
есть три компонента: Logistics (логистика), Packaging (упаковка) и
DeliveryServiceFacade (фасад доставки).
Фасад DeliveryServiceFacade предоставляет упрощенный интерфейс
DeliverProduct для взаимодействия с логистикой и упаковкой.
*/

// Логистика
type Logistics struct{}

func (l *Logistics) CreateShippingLabel(location string) {
	fmt.Printf("Создание метки для места доставки: %s\n", location)
}

func (l *Logistics) SchedulePickup(location string) {
	fmt.Printf("Запланирована сборка товара из места доставки: %s\n", location)
}

// Упаковка
type Packaging struct{}

func (p *Packaging) PackItems(items []string) {
	fmt.Printf("Упаковка товаров: %v\n", items)
}

// Фасад
type DeliveryServiceFacade struct {
	logistics *Logistics
	packaging *Packaging
}

func NewDeliveryServiceFacade() *DeliveryServiceFacade {
	return &DeliveryServiceFacade{
		logistics: &Logistics{},
		packaging: &Packaging{},
	}
}

// Упрощенный интерфейс для доставки товара
func (d *DeliveryServiceFacade) DeliverProduct(items []string, location string) {
	d.packaging.PackItems(items)
	d.logistics.CreateShippingLabel(location)
	d.logistics.SchedulePickup(location)
}

func main() {
	facade := NewDeliveryServiceFacade()

	items := []string{"Телевизор", "Микроволновая печь", "Шампунь"}
	location := "Дом"

	facade.DeliverProduct(items, location)
}

/*
Этот пример демонстрирует, как паттерн "Фасад" может быть применен для упрощения
процесса доставки товаров, скрывая сложность взаимодействия с логистикой и упаковкой за простым интерфейсом.
*/