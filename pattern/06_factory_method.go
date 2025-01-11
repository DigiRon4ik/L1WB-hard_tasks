package pattern

/*
Паттерн «Фабричный метод» предоставляет интерфейс для создания объектов,
позволяя подклассам решать, какой класс создавать.
Это позволяет делегировать создание объектов подклассам, не нарушая принципы открытости/закрытости.

Когда применять:
	1. Необходима гибкость создания объектов. Когда нужно создавать объекты разных типов, не изменяя код клиента.
	2. Классы заранее неизвестны. Если тип создаваемого объекта определяется во время выполнения.
	3. Логика создания сложная. Если создание объектов включает дополнительную логику.

+:
	1. Делает код гибким и расширяемым.
	2. Убирает привязку к конкретным классам в клиентском коде.
	3. Поддерживает принцип открытости/закрытости.

-:
	1. Увеличивает количество классов.
	2. Усложняет код из-за добавления новых фабрик и иерархий.

Реальные примеры использования:
	1. Подключение к базе данных: Фабричный метод выбирает соответствующий драйвер (PostgreSQL, MySQL, SQLite).
	2. Веб-фреймворки: Создание HTTP-запросов и ответов через фабрику, в зависимости от типа данных или формата.
	3. Логирование: Создание логгеров для разных целей (консоль, файл, база данных).
*/

import "fmt"

// Transport - the product is an interface for vehicles.
type Transport interface {
	Deliver()
}

// Truck - Specific Product.
type Truck struct{}

func (t *Truck) Deliver() {
	fmt.Println("Delivering cargo by truck")
}

// Ship - Specific Product.
type Ship struct{}

func (s *Ship) Deliver() {
	fmt.Println("Delivering cargo by ship")
}

// TransportFactory : Creator - Factory Interface.
type TransportFactory interface {
	CreateTransport() Transport
}

// TruckFactory - Specific Factory.
type TruckFactory struct{}

func (tf *TruckFactory) CreateTransport() Transport {
	return &Truck{}
}

// ShipFactory - Specific Factory.
type ShipFactory struct{}

func (sf *ShipFactory) CreateTransport() Transport {
	return &Ship{}
}

// main - example.
func main() {
	// Choosing a factory.
	var factory TransportFactory

	// Use the truck factory.
	factory = &TruckFactory{}
	transport := factory.CreateTransport()
	transport.Deliver()

	// Use the ship factory.
	factory = &ShipFactory{}
	transport = factory.CreateTransport()
	transport.Deliver()
}

/*
 - Output: -
Delivering cargo by truck
Delivering cargo by ship
*/
