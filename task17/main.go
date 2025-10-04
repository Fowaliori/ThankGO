package main

import "fmt"

// начало решения

// ZipMap возвращает карту, где ключи - элементы из keys, а значения - из vals.
func ZipMap[K comparable, V any](keys []K, vals []V) map[K]V {
	size := 0
	if len(keys) > len(vals) {
		size = len(vals)
	} else {
		size = len(keys)
	}
	m := make(map[K]V, size)
	for i := 0; i < size; i++ {
		m[keys[i]] = vals[i]
	}
	return m
}

// конец решения

func main() {
	keys := []string{"one", "two", "thr"}
	vals := []int{11, 22, 33}

	m := ZipMap(keys, vals)
	fmt.Println(m)
}
