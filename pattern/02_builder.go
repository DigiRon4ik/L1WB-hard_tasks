package pattern

/*
Паттерн «Строитель» используется для пошагового создания сложных объектов.
Он отделяет процесс построения объекта от его представления, позволяя использовать один и тот же процесс
для создания различных видов объектов.

Когда применять:
	1. Сложные объекты с множеством параметров. Паттерн позволяет задавать параметры пошагово, избегая громоздких конструкторов.
	2. Гибкость в создании. Вы можете легко менять процесс создания, добавлять новые параметры или типы объектов.
	3. Разные представления одного объекта. Например, создание спортивной машины, грузовика или седана через одну и ту же логику.

+:
	1. Позволяет избежать громоздких конструкторов с множеством аргументов.
	2. Обеспечивает гибкость и контроль над процессом построения.
	3. Упрощает создание объектов с разными конфигурациями.

-:
	1. Может усложнить код при создании простых объектов.
	2. Требует дополнительных классов и структур (строитель, директор).

Реальные примеры использования:
	1. Создание объектов UI: Построение сложных интерфейсов с множеством элементов, например, формы или панели управления.
	2. Создание SQL-запросов: Построение запросов с параметрами, как это делают библиотеки вроде sqlx или gorm.
	3. Конфигураторы машин или оборудования: Упрощение создания конфигураций через цепочку вызовов.
*/

import "fmt"

// Car - The product is a complex object.
type Car struct {
	Brand        string
	Engine       string
	Color        string
	Transmission string
}

// CarBuilder - Builder interface.
type CarBuilder interface {
	SetBrand(brand string) CarBuilder
	SetEngine(engine string) CarBuilder
	SetColor(color string) CarBuilder
	SetTransmission(transmission string) CarBuilder
	Build() Car
}

// SportsCarBuilder - Concrete builder.
type SportsCarBuilder struct {
	car Car
}

// NewSportsCarBuilder - construction set for sports car builder.
func NewSportsCarBuilder() *SportsCarBuilder {
	return &SportsCarBuilder{}
}

func (b *SportsCarBuilder) SetBrand(brand string) CarBuilder {
	b.car.Brand = brand
	return b
}

func (b *SportsCarBuilder) SetEngine(engine string) CarBuilder {
	b.car.Engine = engine
	return b
}

func (b *SportsCarBuilder) SetColor(color string) CarBuilder {
	b.car.Color = color
	return b
}

func (b *SportsCarBuilder) SetTransmission(transmission string) CarBuilder {
	b.car.Transmission = transmission
	return b
}

func (b *SportsCarBuilder) Build() Car {
	return b.car
}

// Director - manages the construction process.
type Director struct {
	builder CarBuilder
}

// NewDirector - director constructor.
func NewDirector(builder CarBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) BuildSportsCar() Car {
	return d.builder.
		SetBrand("Porsche").
		SetEngine("V8").
		SetColor("Red").
		SetTransmission("Automatic").
		Build()
}

// main - example.
func main() {
	// Using a builder through a director.
	builder := NewSportsCarBuilder()
	director := NewDirector(builder)

	sportsCar := director.BuildSportsCar()
	fmt.Printf("Built car: %+v\n", sportsCar)
}

/*
 - Output: -
Built car: {Brand:Porsche Engine:V8 Color:Red Transmission:Automatic}
*/
