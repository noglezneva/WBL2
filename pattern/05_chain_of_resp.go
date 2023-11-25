package main

import (
	"fmt"
	"strings"
)

/*
	Паттерн "Цепочка обязанностей" — это поведенческий паттерн проектирования,
	который позволяет передавать запросы последовательно
	по цепочке обработчиков. Каждый последующий обработчик решает,
	может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

	Паттерн "Цепочка вызовов" полезен в ситуациях, когда есть несколько
	объектов, способных обработать определенный запрос, и порядок их
	обработки неизвестен заранее. При использовании этого паттерна
	каждый объект принимает решение о том, может ли он обработать запрос,
	и передает его следующему объекту в цепочке, если не может.
	Таким образом, паттерн создает динамическую цепочку объектов для
	обработки запросов.

	Плюсы:
	* Уменьшение связанности между отправителем запроса и получателем, так как отправитель не знает точно, какой объект обработает запрос.
	* Позволяет гибко настраивать обработку запросов и менять порядок или состав объектов в цепочке без изменения клиентского кода.
	* Упрощает добавление новых обработчиков в систему.

	Минусы:
	* Запрос может остаться необработанным, если ни один из объектов в цепочке не может его обработать.
	* Повышение сложности отладки из-за динамической природы цепочки вызовов.
	* Некорректная настройка цепочки может привести к бесконечным петлям запросов или зацикливанию.

	Один из практических примеров использования паттерна "Цепочка вызовов"
	- это обработка запросов в веб-фреймворках. В таких фреймворках
	обычно есть несколько слоев или модулей, отвечающих за различные
	аспекты запроса, такие как аутентификация, авторизация, валидация
	данных и обработка бизнес-логики. Каждый слой может быть представлен
	отдельным объектом в цепочке, способным обработать или передать
	дальше запрос. Это позволяет гибко настраивать обработку запроса и
	добавлять новые слои без изменения кода фреймворка.

*/
// Интерфейс обработчика
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request string)
}

// Базовая реализация обработчика
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

func (b *BaseHandler) HandleRequest(request string) {
	if b.next != nil {
		b.next.HandleRequest(request)
	}
}

// Конкретные обработчики
type AuthenticationHandler struct {
	BaseHandler
}

func (a *AuthenticationHandler) HandleRequest(request string) {
	if strings.HasPrefix(request, "auth ") {
		fmt.Println("AuthenticationHandler: обработка запроса", request)
	} else {
		a.BaseHandler.HandleRequest(request)
	}
}

type AuthorizationHandler struct {
	BaseHandler
}

func (a *AuthorizationHandler) HandleRequest(request string) {
	if strings.HasPrefix(request, "authz ") {
		fmt.Println("AuthorizationHandler: обработка запроса", request)
	} else {
		a.BaseHandler.HandleRequest(request)
	}
}

type ValidationHandler struct {
	BaseHandler
}

func (v *ValidationHandler) HandleRequest(request string) {
	if strings.HasPrefix(request, "validate ") {
		fmt.Println("ValidationHandler: обработка запроса", request)
	} else {
		v.BaseHandler.HandleRequest(request)
	}
}

func main() {
	// Создание обработчиков
	authenticationHandler := &AuthenticationHandler{}
	authorizationHandler := &AuthorizationHandler{}
	validationHandler := &ValidationHandler{}

	// Настройка цепочки вызовов
	authenticationHandler.SetNext(authorizationHandler)
	authorizationHandler.SetNext(validationHandler)

	// Обработка запросов
	authenticationHandler.HandleRequest("auth user:password")
	authenticationHandler.HandleRequest("authz user=admin")
	authenticationHandler.HandleRequest("validate data")
	authenticationHandler.HandleRequest("other request")
}

/*
В этом примере, каждый конкретный обработчик (AuthenticationHandler,
AuthorizationHandler, ValidationHandler) определяет свои условия для
обработки запроса. Если запрос соответствует условиям, обработчик
обрабатывает его. Если нет, запрос передается следующему обработчику
в цепочке с помощью вызова b.BaseHandler.HandleRequest(request).
Настройка цепочки вызовов определяет порядок обработки запросов.
*/
