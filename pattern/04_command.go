package pattern

/*
Паттерн «Команда» инкапсулирует запрос как объект, позволяя задавать параметры для обработки,
ставить запросы в очередь, а также отменять или повторять операции.
Он отделяет отправителя запроса от его получателя.

Когда применять:
	1. Отделение отправителя от получателя. Полезно, если вам нужно отделить объект, инициирующий запрос, от объекта, который его обрабатывает.
	2. История операций. Если требуется сохранять команды для повторного выполнения или отмены.
	3. Динамическая конфигурация. Когда команды нужно настраивать в зависимости от ситуации.

+:
	1. Отделяет отправителя от получателя, снижая связанность.
	2. Поддерживает отмену и повтор операций.
	3. Удобен для работы с очередями команд.

-:
	1. Усложняет код из-за создания множества дополнительных классов или структур для команд.
	2. Может привести к увеличению количества объектов.

Реальные примеры использования:
	1. Управление транзакциями в базах данных: Операции commit/rollback.
	2. Системы управления устройствами: Например, умный дом, пульты управления телевизорами.
	3. GUI-приложения: Связывание действий кнопок (например, "Undo", "Redo") с логикой приложения.
*/

import "fmt"

// Command Interface.
type Command interface {
	Execute()
	Undo()
}

// Light - recipient of the command.
type Light struct{}

func (l *Light) TurnOn() {
	fmt.Println("Light is turned ON")
}

func (l *Light) TurnOff() {
	fmt.Println("Light is turned OFF")
}

// TurnOnLightCommand - specific command to turn on the light.
type TurnOnLightCommand struct {
	light *Light
}

func NewTurnOnLightCommand(light *Light) *TurnOnLightCommand {
	return &TurnOnLightCommand{light: light}
}

func (c *TurnOnLightCommand) Execute() {
	c.light.TurnOn()
}

func (c *TurnOnLightCommand) Undo() {
	c.light.TurnOff()
}

// TurnOffLightCommand - specific command to turn off the light.
type TurnOffLightCommand struct {
	light *Light
}

func NewTurnOffLightCommand(light *Light) *TurnOffLightCommand {
	return &TurnOffLightCommand{light: light}
}

func (c *TurnOffLightCommand) Execute() {
	c.light.TurnOff()
}

func (c *TurnOffLightCommand) Undo() {
	c.light.TurnOn()
}

// RemoteControl : Sender - control panel.
type RemoteControl struct {
	history []Command
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{history: []Command{}}
}

func (r *RemoteControl) PressButton(command Command) {
	command.Execute()
	r.history = append(r.history, command)
}

func (r *RemoteControl) PressUndo() {
	if len(r.history) == 0 {
		fmt.Println("Nothing to undo")
		return
	}
	lastCommand := r.history[len(r.history)-1]
	r.history = r.history[:len(r.history)-1]
	lastCommand.Undo()
}

// main - example.
func main() {
	light := &Light{}
	turnOnCommand := NewTurnOnLightCommand(light)
	turnOffCommand := NewTurnOffLightCommand(light)

	remote := NewRemoteControl()

	// Executing commands.
	remote.PressButton(turnOnCommand)
	remote.PressButton(turnOffCommand)

	// Cancel last command.
	remote.PressUndo()
	remote.PressUndo()
}

/*
 - Output: -
Light is turned ON
Light is turned OFF
Light is turned ON
Light is turned OFF
*/
