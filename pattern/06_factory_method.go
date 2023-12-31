package main

import "fmt"

/*
	Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс
	для создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.

	Паттерн "Фабричный метод" используется, когда у нас есть базовый интерфейс или абстрактный класс,
	и мы хотим, чтобы каждая конкретная реализация создавалась через свой собственный фабричный метод.
	Это позволяет делегировать создание конкретных объектов наследникам базового класса или интерфейса.

	Плюсы:
	* Устраняет жесткую привязку между клиентским кодом и конкретными классами, так как клиент работает через абстрактный интерфейс или класс базового типа.
	* Обеспечивает расширяемость кода путем добавления новых классов, реализующих интерфейс или наследующих базовый класс.
	* Позволяет гибко настраивать создание объектов в зависимости от определенных условий или параметров.
	Минусы:
	* Может привести к увеличению количества классов из-за введения отдельных фабричных методов для каждой конкретной реализации.
	* Клиентский код должен знать о конкретных классах фабрик, что может нарушить принцип инверсии зависимостей.

	Один из практических примеров использования паттерна "Фабричный метод" - это создание различных типов логгеров в логгирующей библиотеке.
	Предположим, у нас есть абстрактный класс Logger и две конкретные реализации: FileLogger и ConsoleLogger. Мы можем создать фабричные методы
	для каждой реализации, которые будут создавать соответствующие объекты логгеров в зависимости от некоторых параметров, например, типа журнала.

*/
// Интерфейс фабрики создания логгеров
type LoggerFactory interface {
	CreateLogger() Logger
}

// Абстрактный класс логгера
type Logger interface {
	Log(message string)
}

// Фабрика для создания файловых логгеров
type FileLoggerFactory struct{}

func (factory FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

// Фабрика для создания консольных логгеров
type ConsoleLoggerFactory struct{}

func (factory ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}

// Конкретная реализация файлового логгера
type FileLogger struct{}

func (logger *FileLogger) Log(message string) {
	fmt.Println("Запись в файл:", message)
}

// Конкретная реализация консольного логгера
type ConsoleLogger struct{}

func (logger *ConsoleLogger) Log(message string) {
	fmt.Println("Вывод в консоль:", message)
}

func main() {
	// Создание фабрик
	fileLoggerFactory := FileLoggerFactory{}
	consoleLoggerFactory := ConsoleLoggerFactory{}

	// Создание логгеров с помощью фабричных методов
	fileLogger := fileLoggerFactory.CreateLogger()
	consoleLogger := consoleLoggerFactory.CreateLogger()

	// Использование логгеров
	fileLogger.Log("Запись в файл")
	consoleLogger.Log("Вывод в консоль")
}

/*
у нас есть интерфейс Logger и две реализации: FileLogger и ConsoleLogger.
Для каждой реализации мы создаем соответствующую фабрику (FileLoggerFactory и ConsoleLoggerFactory),
которая реализует интерфейс LoggerFactory и имеет метод CreateLogger(), возвращающий соответствующий объект логгера.

Клиентский код использует фабричные методы для создания объектов логгеров, при этом он не зависит от конкретных классов логгеров.

Таким образом, паттерн "Фабричный метод" позволяет гибко создавать объекты, основываясь на различных фабричных методах,
и разделять процесс создания от использования объектов.
*/
