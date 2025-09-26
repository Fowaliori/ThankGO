package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["key"] = 6
	m["qww"] = 7
	fmt.Println(m)
	fmt.Println(m["key"])
	delete(m, "qww")
	fmt.Println(m)
	_, ok := m["key"]
	fmt.Println(ok)
}
