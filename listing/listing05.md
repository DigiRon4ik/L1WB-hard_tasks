Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error  

customError реализует интерфейс error, поэтому присваивание err = test() прошло успешно.  
При проверке err != nil условие выполняется, так как интерфейс err имеет значение nil, но тип не является nil.  
```
