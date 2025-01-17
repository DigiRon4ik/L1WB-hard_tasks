Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Создаются два канала с помощью функции asChan, которая добавляет значения в канал с случайной задержкой.
Эти два канала объединяются в канал c через функцию merge, после чего производится чтение из канала c.
Однако, как только один из каналов полностью опустеет, начнётся спам нулевых значений.

Чтобы этого избежать, можно изменить select в функции merge следующим образом:
select {
	case v, ok := <-a:
		if !ok {
			return
		}
		c <- v
	case v, ok := <-b:
		if !ok {
			return
		}
		c <- v
}
```
