package pattern

/*
Паттерн «Цепочка обязанностей» (Chain of Responsibility) позволяет передавать запрос последовательно
через цепочку обработчиков до тех пор, пока запрос не будет обработан.
Каждый обработчик принимает запрос, либо передаёт его следующему обработчику в цепочке.

Когда применять:
	1. Гибкая обработка запросов. Когда нужно передать запрос по цепочке обработчиков без жёсткой привязки к их порядку.
	2. Динамическая конфигурация. Если порядок или количество обработчиков могут меняться.
	3. Разделение обязанностей. Когда обработка запроса должна быть разделена между несколькими обработчиками.

+:
	1. Ослабляет связанность между отправителем и получателями.
	2. Упрощает добавление новых обработчиков.
	3. Позволяет динамически настраивать цепочку.

-:
	1. Может быть сложно отследить, где запрос обработан.
	2. Если цепочка слишком длинная, это может повлиять на производительность.

Реальные примеры использования:
	1. Логирование: Цепочка обработчиков, которая обрабатывает сообщения разного уровня (Debug, Info, Error).
	2. Обработка HTTP-запросов: Middleware в веб-фреймворках (например, gin или net/http в Go).
	3. Системы безопасности: Последовательная проверка авторизации, прав доступа и других ограничений.
*/

import "fmt"

// Handler Interface.
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string)
}

// BaseHandler - basic structure for chain implementation.
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) Handler {
	b.next = handler
	return handler
}

func (b *BaseHandler) Handle(request string) {
	if b.next != nil {
		b.next.Handle(request)
	}
}

// AuthHandler : Specific handler - Login check.
type AuthHandler struct {
	BaseHandler
}

func (a *AuthHandler) Handle(request string) {
	if request == "auth" {
		fmt.Println("AuthHandler: Authentication successful")
	} else {
		fmt.Println("AuthHandler: Passing to the next handler")
		a.BaseHandler.Handle(request)
	}
}

// PermissionHandler : Specific handler - Permission check.
type PermissionHandler struct {
	BaseHandler
}

func (p *PermissionHandler) Handle(request string) {
	if request == "permission" {
		fmt.Println("PermissionHandler: Access granted")
	} else {
		fmt.Println("PermissionHandler: Passing to the next handler")
		p.BaseHandler.Handle(request)
	}
}

// LoggingHandler : Specific handler - Logging.
type LoggingHandler struct {
	BaseHandler
}

func (l *LoggingHandler) Handle(request string) {
	fmt.Println("LoggingHandler: Logging the request")
	l.BaseHandler.Handle(request)
}

// main - example.
func main() {
	// Create handlers.
	authHandler := &AuthHandler{}
	permissionHandler := &PermissionHandler{}
	loggingHandler := &LoggingHandler{}

	// Building a chain.
	authHandler.SetNext(permissionHandler).SetNext(loggingHandler)

	// Forward requests.
	fmt.Println("Request: auth")
	authHandler.Handle("auth")

	fmt.Println("\nRequest: permission")
	authHandler.Handle("permission")

	fmt.Println("\nRequest: unknown")
	authHandler.Handle("unknown")
}

/*
 - Output: -
Request: auth
AuthHandler: Authentication successful

Request: permission
AuthHandler: Passing to the next handler
PermissionHandler: Access granted

Request: unknown
AuthHandler: Passing to the next handler
PermissionHandler: Passing to the next handler
LoggingHandler: Logging the request
*/
