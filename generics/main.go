package main

import (
	"cmp"
	"fmt"
)

type Named interface {
	Name() string
}

type Person struct {
	name    string
	country string
}

func (p Person) Name() string {
	return p.name
}

type Pet struct {
	name string
}

func (p Pet) Name() string {
	return p.name
}

func makeMap[K comparable, V any](size int) map[K]V {
	m := make(map[K]V, size)
	return m
}

func Reverse[T any](s []T) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func Find[T comparable](slice []T, elem T) int {
	for i, v := range slice {
		if v == elem {
			return i
		}
	}
	return -1
}

func FindInterface[T Named](slice []T, name string) int {
	for i, v := range slice {
		if v.Name() == name {
			return i
		}
	}
	return -1
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	str := []string{"a", "b", "c", "d", "e"}
	fmt.Println(nums)
	fmt.Printf("idx 6 = %v\n", Find(nums, 6))
	Reverse(nums)
	fmt.Println(nums)
	fmt.Println(str)
	Reverse(str)
	fmt.Println(str)
	fmt.Printf("idx 6 = %v\n", Find(nums, 6))
	a := 6
	b := 4
	fmt.Printf("max(%v, %v) = %v\n", a, b, Max(a, b))
	people := []Person{
		{name: "Dmitriy", country: "Russia"},
		{name: "Alice", country: "USA"},
	}
	fmt.Printf("idx Alice = %v\n", FindInterface(people, "Alice"))
	pets := []Pet{
		{name: "Rex"},
		{name: "Felya"},
	}
	fmt.Printf("idx Rex = %v\n", FindInterface(pets, "Rex"))
	m := makeMap[string, int](5)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Println(m)
}
