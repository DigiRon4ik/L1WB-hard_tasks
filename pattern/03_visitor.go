package pattern

/*
Паттерн «Посетитель» позволяет добавить новую операцию к объектам без изменения их классов.
Он разделяет алгоритмы и структуры данных, позволяя внедрять новое поведение через отдельный объект.

Когда применять:
	1. Добавление новых операций без изменения классов. Это удобно, когда структуры стабильны, но алгоритмы меняются.
	2. Разделение ответственности. Паттерн позволяет отделить логику обработки объектов от их структуры.
	3. Объектная структура имеет иерархию. Когда нужно работать с объектами разного типа в одной иерархии.

+:
	1. Позволяет добавлять новые операции без изменения классов.
	2. Упрощает добавление сложной логики, не засоряя классы данных.
	3. Удобен при работе со сложными иерархиями объектов.

-:
	1. Нарушает принцип инкапсуляции, так как посетитель часто должен знать детали реализации объекта.
	2. Добавление новых типов объектов в структуру требует изменения посетителей.

Реальные примеры использования:
	1. Компиляторы: Выполнение различных операций (например, синтаксический анализ, оптимизация кода) над абстрактным синтаксическим деревом.
	2. Системы расчётов: Добавление операций над структурами данных (расчёт площади, стоимости, веса и т.д.).
	3. UI и рендеринг: Применение действий к элементам интерфейса, например, добавление обработки событий или анимации.
*/

import "fmt"

// Shape - The interface of the element that receives the visitor.
type Shape interface {
	Accept(visitor ShapeVisitor)
}

// Circle - specific element.
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor ShapeVisitor) {
	visitor.VisitCircle(c)
}

// Square - specific element.
type Square struct {
	Side float64
}

func (s *Square) Accept(visitor ShapeVisitor) {
	visitor.VisitSquare(s)
}

// ShapeVisitor - Visitor Interface.
type ShapeVisitor interface {
	VisitCircle(circle *Circle)
	VisitSquare(square *Square)
}

// AreaCalculator - specific visitor.
type AreaCalculator struct{}

func (a *AreaCalculator) VisitCircle(circle *Circle) {
	area := 3.14 * circle.Radius * circle.Radius
	fmt.Printf("Area of Circle: %.2f\n", area)
}

func (a *AreaCalculator) VisitSquare(square *Square) {
	area := square.Side * square.Side
	fmt.Printf("Area of Square: %.2f\n", area)
}

// PerimeterCalculator - specific visitor.
type PerimeterCalculator struct{}

func (p *PerimeterCalculator) VisitCircle(circle *Circle) {
	perimeter := 2 * 3.14 * circle.Radius
	fmt.Printf("Perimeter of Circle: %.2f\n", perimeter)
}

func (p *PerimeterCalculator) VisitSquare(square *Square) {
	perimeter := 4 * square.Side
	fmt.Printf("Perimeter of Square: %.2f\n", perimeter)
}

// main - example.
func main() {
	shapes := []Shape{
		&Circle{Radius: 5},
		&Square{Side: 4},
	}

	areaCalculator := &AreaCalculator{}
	perimeterCalculator := &PerimeterCalculator{}

	fmt.Println("Calculating areas:")
	for _, shape := range shapes {
		shape.Accept(areaCalculator)
	}

	fmt.Println("\nCalculating perimeters:")
	for _, shape := range shapes {
		shape.Accept(perimeterCalculator)
	}
}

/*
 - Output: -
Calculating areas:
Area of Circle: 78.50
Area of Square: 16.00

Calculating perimeters:
Perimeter of Circle: 31.40
Perimeter of Square: 16.00
*/
