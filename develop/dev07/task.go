package main

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизвестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or-каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fine after %v”, time.Since(start))
*/

import (
	"fmt"
	"time"
)

// or combines several done channels into one.
// If at least one of the channels is closed, then the resulting channel is also closed.
func or(channels ...<-chan interface{}) <-chan interface{} {
	// Step 1: Process the base cases.
	switch len(channels) {
	case 0:
		return nil // If there are no channels, return nil.
	case 1:
		return channels[0] // If there is only one channel, we simply return it.
	}

	// Step 2: Create the resulting channel.
	orDone := make(chan interface{})

	go func() { // Launch the goroutine.
		defer close(orDone) // The goroutine will close the orDone channel upon completion.

		// Step 3: Divide the channels into groups and listen to them.
		switch len(channels) {
		case 2: // If there are 2 channels, listen to them using select.
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: // If there are more channels, we divide them in half.
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...): // Call it recursively for the first half.
			case <-or(channels[m:]...): // Call it recursively for the second half.
			}
		}
	}()

	return orDone // Return the resulting channel.
}

func main() {
	// An example of the sig helper function.
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{}) // Create a channel.
		go func() {
			defer close(c)    // Close the channel when we finish work.
			time.Sleep(after) // “Fall asleep” for a specified time.
		}()
		return c
	}

	// Testing the or function.
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v\n", time.Since(start))
}

/*
 - Output: -
Done after 1.0004678s
*/
