package main

import "fmt"

func main() {
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println("___________")
	for j := 4; j <= 10; j++ {
		fmt.Println(j)
	}
	fmt.Println("___________")
	const n = 10
	for l := range n {
		fmt.Println(l)
	}
	fmt.Println("___________")
	for {
		fmt.Println("break")
		break
	}
	fmt.Println("___________")
	for k := range 10 {
		if k%2 == 0 {
			continue
		} else {
			fmt.Println(k)
		}
	}

}
