package main

import (
	"fmt"
)

func main() {
	var source string
	var times int
	// гарантируется, что значения корректные
	fmt.Scan(&source, &times)

	// возьмите строку `source` и повторите ее `times` раз
	// запишите результат в `result`
	// ...
	i := 0
	result := ""
	for i = range times {
		result += source
		i++
	}
	fmt.Println(result)
}
