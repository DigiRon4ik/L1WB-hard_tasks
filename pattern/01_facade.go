package pattern

/*
Паттерн «Фасад» предоставляет упрощённый интерфейс к сложной системе, скрывая её внутреннюю структуру
и детали реализации. Он создаёт "оболочку", которая объединяет несколько сложных подсистем
в простой и понятный интерфейс.

Когда применять:
	1. Сложные системы. Позволяет скрыть многочисленные компоненты системы за простым интерфейсом.
	2. Упрощение использования. Фасад помогает снизить сложность кода клиента.
	3. Поддержка и изменения. Упрощает замену или обновление подсистем, не влияя на клиентский код.

+:
	1. Упрощает взаимодействие с системой.
	2. Скрывает детали реализации, защищая клиента от изменений в подсистемах.
	3. Делает код клиента чище и понятнее.

-:
	1. Может скрывать важную функциональность.
	2. Увеличивает вероятность дублирования кода, если подсистемы используются напрямую без фасада.

Реальные примеры использования:
	1. REST API. Создание единого интерфейса для работы с множеством микросервисов.
	2. ORM (например, GORM). Упрощённый доступ к базе данных через методы библиотеки.
	3. Работа с файловой системой. Упрощение операций чтения, записи и обработки файлов.
*/

import "fmt"

// SubsystemA - Complex subsystems.
type SubsystemA struct{}

func (s *SubsystemA) OperationA() string {
	return "SubsystemA: OperationA executed"
}

// SubsystemB - Complex subsystems.
type SubsystemB struct{}

func (s *SubsystemB) OperationB() string {
	return "SubsystemB: OperationB executed"
}

// SubsystemC - Complex subsystems.
type SubsystemC struct{}

func (s *SubsystemC) OperationC() string {
	return "SubsystemC: OperationC executed"
}

// Facade - example of a facade.
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
	subsystemC *SubsystemC
}

// NewFacade - constructor.
func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
		subsystemC: &SubsystemC{},
	}
}

// SimplifiedOperation - main func for Facade.
func (f *Facade) SimplifiedOperation() {
	fmt.Println("Starting SimplifiedOperation via Facade:")
	fmt.Println(f.subsystemA.OperationA())
	fmt.Println(f.subsystemB.OperationB())
	fmt.Println(f.subsystemC.OperationC())
}

// main - example.
func main() {
	facade := NewFacade()
	facade.SimplifiedOperation()
}

/*
 - Output: -
Starting SimplifiedOperation via Facade:
SubsystemA: OperationA executed
SubsystemB: OperationA executed
SubsystemC: OperationC executed
*/
