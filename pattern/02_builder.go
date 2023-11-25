package main

import "fmt"

/*
	Паттерн "Строитель" (Builder) представляет собой порождающий паттерн проектирования,
	который позволяет создавать сложные объекты пошагово.
	Он используется, чтобы абстрагировать процесс создания сложного объекта от его представления,
	таким образом, что один и тот же процесс строительства может создавать разные представления объекта.

	Применимость паттерна "Строитель":

	* Когда требуется создание объекта, состоящего из более чем одного шага или этапа, и требуется независимость конкретных шагов строительства.
	* Когда нужно создавать различные представления сложного объекта, используя один и тот же процесс строительства.
	* Когда требуется разделение процесса создания объекта от его представления.

	Преимущества:
	* Позволяет создавать продукты пошагово.
	* Позволяет использовать один и тот же код для создания различных продуктов.
	* Изолирует сложный код сборки продукта от его основной бизнес-логики.
Недостатки:
	* Усложняет код программы из-за введения дополнительных классов.
	* Клиент будет привязан к конкретным классам строителей,
	так как в интерфейсе директора может не быть метода получения результата.

	Паттерн "Строитель" помогает управлять сложными процессами создания объектов,
	облегчает поддержку различных представлений и повышает расширяемость. Он находит
	применение во многих областях программирования, где необходимо пошаговое создание
	сложных объектов.

*/
// Тип продукта, который мы собираем
type House struct {
	Walls   string
	Doors   string
	Windows string
}

// Интерфейс строителя для создания различных компонентов дома
type HouseBuilder interface {
	BuildWalls()
	BuildDoors()
	BuildWindows()
	GetHouse() House
}

// Конкретный строитель, строит дом с красными стенами и деревянными дверями
type RedWoodenHouseBuilder struct {
	house House
}

func (b *RedWoodenHouseBuilder) BuildWalls() {
	b.house.Walls = "Red"
}

func (b *RedWoodenHouseBuilder) BuildDoors() {
	b.house.Doors = "Wooden"
}

func (b *RedWoodenHouseBuilder) BuildWindows() {
	b.house.Windows = "Glass"
}

func (b *RedWoodenHouseBuilder) GetHouse() House {
	return b.house
}

// Конкретный строитель, строит дом с белыми стенами и металлическими дверями
type WhiteMetalHouseBuilder struct {
	house House
}

func (b *WhiteMetalHouseBuilder) BuildWalls() {
	b.house.Walls = "White"
}

func (b *WhiteMetalHouseBuilder) BuildDoors() {
	b.house.Doors = "Metal"
}

func (b *WhiteMetalHouseBuilder) BuildWindows() {
	b.house.Windows = "Metal"
}

func (b *WhiteMetalHouseBuilder) GetHouse() House {
	return b.house
}

// Руководитель, управляет процессом строительства
type Director struct {
	builder HouseBuilder
}

func (d *Director) SetBuilder(builder HouseBuilder) {
	d.builder = builder
}

func (d *Director) ConstructHouse() House {
	d.builder.BuildWalls()
	d.builder.BuildDoors()
	d.builder.BuildWindows()
	return d.builder.GetHouse()
}

// Пример использования
func main() {
	director := Director{}

	builder1 := &RedWoodenHouseBuilder{}
	director.SetBuilder(builder1)
	house1 := director.ConstructHouse()
	fmt.Println("House with red walls and wooden doors:", house1)

	builder2 := &WhiteMetalHouseBuilder{}
	director.SetBuilder(builder2)
	house2 := director.ConstructHouse()
	fmt.Println("House with white walls and metal doors:", house2)
}
