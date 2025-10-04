package main

import "fmt"

// начало решения

// Avg - накопительное среднее значение.
type Avg[T int | float64] struct {
	sum   T
	count int
	val   T
}

// Add пересчитывает среднее значение с учетом val.
func (a *Avg[T]) Add(val T) *Avg[T] {
	a.count++
	a.sum += val
	a.val = a.sum / T(a.count)
	return a
}

// Val возвращает текущее среднее значение.
func (a *Avg[T]) Val() T {
	return a.val
}

// конец решения

func main() {
	intAvg := Avg[int]{}
	intAvg.Add(6).Add(2).Add(2)
	fmt.Println(intAvg.Val())
}
