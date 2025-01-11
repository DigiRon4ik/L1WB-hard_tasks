package pattern

/*
Паттерн «Состояние» позволяет объекту изменять своё поведение в зависимости от внутреннего состояния.
При этом создаётся впечатление, что объект меняет свой класс во время выполнения.

Когда применять:
	1. Объект имеет несколько состояний. Разное поведение объекта в зависимости от текущего состояния.
	2. Логика переключения сложная. Если логика переходов между состояниями может расти, и её нужно изолировать.
	3. Избегание большого количества условий. Вместо громоздких if/else или switch, состояния разносятся по классам.

+:
	1. Упрощает код за счёт разделения состояний на отдельные классы.
	2. Изолирует логику, связанную с конкретным состоянием.
	3. Упрощает добавление новых состояний.

-:
	1. Увеличивает количество классов.
	2. Может усложнить понимание кода из-за большого количества переключений.

Реальные примеры использования:
	1. Банкоматы: Разные состояния (ожидание карты, ввод PIN, выдача денег).
	2. Игры: Поведение персонажа в зависимости от состояния (например, бег, прыжок, отдых).
	3. Торговые автоматы: Логика работы с монетами, выбором товаров и выдачей сдачи.
*/

import "fmt"

// State - Status Interface.
type State interface {
	InsertCoin()
	SelectDrink()
	DispenseDrink()
}

// VendingMachine : Context - a vending machine with drinks.
type VendingMachine struct {
	currentState State
}

func (vm *VendingMachine) SetState(state State) {
	vm.currentState = state
}

func (vm *VendingMachine) InsertCoin() {
	vm.currentState.InsertCoin()
}

func (vm *VendingMachine) SelectDrink() {
	vm.currentState.SelectDrink()
}

func (vm *VendingMachine) DispenseDrink() {
	vm.currentState.DispenseDrink()
}

// WaitingForCoinState - Condition: Waiting for a coin.
type WaitingForCoinState struct {
	vendingMachine *VendingMachine
}

func (s *WaitingForCoinState) InsertCoin() {
	fmt.Println("Coin inserted. Please select your drink.")
	s.vendingMachine.SetState(&WaitingForSelectionState{vendingMachine: s.vendingMachine})
}

func (s *WaitingForCoinState) SelectDrink() {
	fmt.Println("Please insert a coin first.")
}

func (s *WaitingForCoinState) DispenseDrink() {
	fmt.Println("Please insert a coin first.")
}

// WaitingForSelectionState - Condition: Drink selection.
type WaitingForSelectionState struct {
	vendingMachine *VendingMachine
}

func (s *WaitingForSelectionState) InsertCoin() {
	fmt.Println("Coin already inserted. Please select your drink.")
}

func (s *WaitingForSelectionState) SelectDrink() {
	fmt.Println("Drink selected. Dispensing drink.")
	s.vendingMachine.SetState(&DispensingState{vendingMachine: s.vendingMachine})
}

func (s *WaitingForSelectionState) DispenseDrink() {
	fmt.Println("Please select a drink first.")
}

// DispensingState - Condition: Beverage dispensing.
type DispensingState struct {
	vendingMachine *VendingMachine
}

func (s *DispensingState) InsertCoin() {
	fmt.Println("Please wait, dispensing drink.")
}

func (s *DispensingState) SelectDrink() {
	fmt.Println("Please wait, dispensing drink.")
}

func (s *DispensingState) DispenseDrink() {
	fmt.Println("Here is your drink. Thank you!")
	s.vendingMachine.SetState(&WaitingForCoinState{vendingMachine: s.vendingMachine})
}

// main - example.
func main() {
	vendingMachine := &VendingMachine{}
	initialState := &WaitingForCoinState{vendingMachine: vendingMachine}
	vendingMachine.SetState(initialState)

	// Example of use.
	vendingMachine.SelectDrink()   // Waiting for a coin
	vendingMachine.InsertCoin()    // Coin insertion
	vendingMachine.SelectDrink()   // Drink selection
	vendingMachine.DispenseDrink() // Drink dispensing
	vendingMachine.InsertCoin()    // Ready for the next user
	vendingMachine.DispenseDrink() // Waiting for a coin
}

/*
 - Output: -
Please insert a coin first.
Coin inserted. Please select your drink.
Drink selected. Dispensing drink.
Here is your drink. Thank you!
Coin inserted. Please select your drink.
Please select a drink first.
*/
