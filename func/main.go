package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func divide(a, b int) (int, int) {
	q := a / b
	q2 := a % b
	return q, q2
}

func summa(nums ...int) int {
	s := 0
	for i := range nums {
		s += nums[i]
	}
	return s
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	a := 5
	b := 6
	fmt.Println(sum(a, b))
	fmt.Println(divide(b, a))
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(summa(nums...))
	r := intSeq()
	r2 := intSeq()
	fmt.Println("r", r())
	fmt.Println("r", r())
	fmt.Println("r2", r2())
}
