package pattern

/*
Паттерн «Стратегия» определяет семейство алгоритмов, инкапсулирует их в отдельные классы и
позволяет подменять их во время выполнения.
Это помогает разделить поведение объекта и его реализацию, делая систему более гибкой.

Когда применять:
	1. Разные варианты поведения. Когда объект должен менять своё поведение в зависимости от условий.
	2. Избегание большого количества условий (if/else или switch). Когда есть много альтернативных алгоритмов, и их лучше структурировать.
	3. Гибкость и расширяемость. Когда нужно легко добавлять новые алгоритмы без изменения существующего кода.

+:
	1. Позволяет подменять алгоритмы во время выполнения.
	2. Устраняет необходимость большого количества условных конструкций.
	3. Соблюдает принцип открытости/закрытости.

-:
	1. Увеличивает количество классов, так как каждая стратегия — это отдельный класс.
	2. Клиенту нужно знать о различиях между стратегиями.

Реальные примеры использования:
	1. Сортировка данных: Разные алгоритмы сортировки (быстрая, пузырьковая, слиянием), выбираемые в зависимости от размера данных.
	2. Оплата: Выбор способа оплаты в интернет-магазине (карта, PayPal, криптовалюта).
	3. Игры: Реализация поведения персонажей (агрессивное, оборонительное, нейтральное).
*/

import "fmt"

// PaymentStrategy - Strategy Interface.
type PaymentStrategy interface {
	Pay(amount float64)
}

// CreditCardPayment - Specific Strategy.
type CreditCardPayment struct {
	CardNumber string
}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card: %s\n", amount, c.CardNumber)
}

// PayPalPayment - Specific Strategy.
type PayPalPayment struct {
	Email string
}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using PayPal: %s\n", amount, p.Email)
}

// PaymentProcessor : Context - The class that uses the strategy.
type PaymentProcessor struct {
	strategy PaymentStrategy
}

func (pp *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
	pp.strategy = strategy
}

func (pp *PaymentProcessor) ProcessPayment(amount float64) {
	if pp.strategy == nil {
		fmt.Println("No payment strategy set!")
		return
	}
	pp.strategy.Pay(amount)
}

// main - example.
func main() {
	// Create a payment processor.
	processor := &PaymentProcessor{}

	// Choosing a credit card payment strategy.
	creditCard := &CreditCardPayment{CardNumber: "1234-5678-9012-3456"}
	processor.SetStrategy(creditCard)
	processor.ProcessPayment(100.50)

	// Changing the strategy for PayPal.
	payPal := &PayPalPayment{Email: "user@example.com"}
	processor.SetStrategy(payPal)
	processor.ProcessPayment(200.75)
}

/*
 - Output: -
Paid 100.50 using Credit Card: 1234-5678-9012-3456
Paid 200.75 using PayPal: user@example.com
*/
