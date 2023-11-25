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

Возвращаемая ошибка не nil, потому что в переменной err хранится информация о типе. Для функций, которые возвращают ошибки, рекомендуется всегда использовать тип error в своей сигнатуре, а не конкретный тип, такой как *customError, чтобы гарантировать правильное создание ошибки.
```