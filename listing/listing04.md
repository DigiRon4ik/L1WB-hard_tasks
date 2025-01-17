Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
0  
1  
2  
3  
...  
8  
9  
deadlock  

Цикл for n := range ch считывает данные из канала до его закрытия.  
Когда горутина завершает запись данных, но канал остаётся открытым, возникает deadlock.
```
