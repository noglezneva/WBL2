package main

import "fmt"

/*
	Паттерн "Состояние" — это поведенческий паттерн проектирования,
	который позволяет объектам менять поведение в зависимости от своего состояния.
	Извне создаётся впечатление, что изменился класс объекта.
	Паттерн "Состояние" применим, когда в объекте есть различные
	варианты поведения в зависимости от его текущего состояния.
	Он позволяет объекту изменять свое поведение, делая его более
	гибким и легким для поддержки.

	Плюсы:
	* Позволяет объекту изменять свое поведение в зависимости от состояния без использования множественных условных операторов.
	* Обеспечивает четкую структуру состояний и переходов между ними, что делает код более понятным и легким для поддержки.
	* Упрощает добавление новых состояний и изменение существующих без изменения клиентского кода.

	Минусы:
	* Может привести к увеличению количества классов из-за необходимости создания отдельных классов для каждого состояния.
	* Усложняет код, если количество состояний и переходов между ними очень большое.

	Один из практических примеров использования паттерна "Состояние" -
	это реализация автомата состояний. Предположим, у нас есть объект
	Printer, который может находиться в различных состояниях, например,
	"Включен", "Выключен" и "Режим печати". Каждое из состояний определяет
	свое поведение метода печати.
*/

// Интерфейс состояния
type State interface {
	Print()
}

// Конкретное состояние "Включен"
type OnState struct{}

func (state OnState) Print() {
	fmt.Println("Печать включена.")
	// Логика печати в режиме "Включен"
}

// Конкретное состояние "Выключен"
type OffState struct{}

func (state OffState) Print() {
	fmt.Println("Печать выключена.")
	// Логика при выключенном состоянии
}

// Конкретное состояние "Режим печати"
type PrintState struct{}

func (state PrintState) Print() {
	fmt.Println("Печать в режиме печати.")
	// Логика печати в режиме "Режим печати"
}

// Контекст, использующий состояние
type Printer struct {
	state State
}

func (printer Printer) SetState(state State) {
	printer.state = state
}

func (printer Printer) Print() {
	printer.state.Print()
}

func main() {
	printer := Printer{}
	printer.SetState(OnState{})
	printer.Print()

	printer.SetState(PrintState{})
	printer.Print()

	printer.SetState(OffState{})
	printer.Print()
}

/*
у нас есть интерфейс State и три конкретных реализации:
OnState, OffState и PrintState, которые представляют состояния принтера.
Контекст Printer содержит поле для хранения текущего состояния и методы
SetState() и Print(), которые делегируют вызов метода Print() текущего состояния.

Клиентский код создает экземпляр принтера Printer и изменяет его состояние
с помощью метода SetState(). Затем клиентский код вызывает метод Print(),
который выполняет печать с учетом текущего состояния.

Таким образом, паттерн "Состояние" позволяет объекту изменять
свое поведение в зависимости от его состояния, делая его более гибким
и легким для поддержки.
*/
