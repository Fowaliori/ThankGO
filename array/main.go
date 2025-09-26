package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)
	arr[2] = 25
	fmt.Println(arr, len(arr))
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	var arr23 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			arr23[i][j] = i + j
		}
	}
	fmt.Println(arr23)
}
