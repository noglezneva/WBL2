package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Паттерн "Посетитель" (Visitor) — это поведенческий паттерн проектирования,
	который позволяет добавить функционал к существующей структуре, не изменяя ее структуру.

	Применимость:
	* Когда необходимо выполнить операции над объектами разных классов с разными интерфейсами, и вы не хотите изменять эти классы, добавляя новые методы.
	* Когда необходимо выполнить набор связанных операций над группой объектов, независимо от их конкретных классов.
	* Когда классы, над которыми нужно выполнять операции, находятся в разных иерархиях наследования и вы хотите избежать использования множества проверок типов.

	Плюсы:
	* Упрощает добавление операций, работающих со сложными структурами объектов.
	* Легкость добавления новых операций при работе с существующими классами.
	* Посетитель может накапливать состояние при обходе структуры элементов.
	* Чистота и гибкость кода, поскольку каждая операция находится в отдельном методе.

	Минусы:
	* Добавление новых классов в систему будет требовать соответствующего изменения интерфейсов и соотношений существующих классов.
	* Паттерн не оправдан, если иерархия элементов часто меняется.
	* Может привести к нарушению инкапсуляции элементов.

	Паттерн "Посетитель" очень полезен, когда необходимо добавлять новые операции к существующей иерархии классов, не изменяя эти классы. Он позволяет легко добавлять новую функциональность и упрощает расширение системы в будущем.
*/

func main() {
	shapes := NewShapes()
	visitorShape := VisitorShapeToPrint{}
	shapes.Visit(visitorShape)
}

// IVisitorShape Интерфейс для будущих посетителей
type IVisitorShape interface {
	VisitRectangle(rectangle *Rectangle)
	VisitCircle(circle *Circle)
}

// VisitorShapeToPrint Структура реализующая интерфейс выше
type VisitorShapeToPrint struct{}

// VisitRectangle посещает объект типа Rectangle
func (VisitorShapeToPrint) VisitRectangle(rectangle *Rectangle) {
	fmt.Printf("Visitor visited Rectangle: %+v\n", *rectangle)
}

// VisitCircle посещает объект типа Circle
func (VisitorShapeToPrint) VisitCircle(circle *Circle) {
	fmt.Printf("Visitor visited Circle: %+v\n", *circle)
}

// Shapes наша структура с возможностью приема посетителей
type Shapes struct {
	rectangle Rectangle
	circle    Circle
}

// Rectangle объект прямоугольника
type Rectangle struct{ width, height float32 }

// Circle объект круга
type Circle struct{ radius float32 }

// Visit 	---> Принимает интерфейс посетителя <---
func (s *Shapes) Visit(v IVisitorShape) {
	v.VisitRectangle(&s.rectangle)
	v.VisitCircle(&s.circle)
}

// NewShapes простой генератор случайных фигур
func NewShapes() *Shapes {
	s := new(Shapes)
	rand.Seed(time.Now().UnixNano())
	s.rectangle.width = rand.Float32() * 10
	s.rectangle.height = rand.Float32() * 10
	s.circle.radius = rand.Float32() * 10
	return s
}
