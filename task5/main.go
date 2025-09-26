package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...
	res := ""
	if width >= len(text) {
		fmt.Println(text)
	} else {
		res = text[:width]
		res = res + "..."
	}
	fmt.Println(res)
}
