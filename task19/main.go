package main

import "fmt"

// начало решения

// Map - карта "ключ-значение".
type Map[K comparable, V any] map[K]V

// Set устанавливает значение для ключа.
func (m Map[K, V]) Set(key K, val V) {
	m[key] = val
}

// Get возвращает значение по ключу.
func (m Map[K, V]) Get(key K) V {
	return m[key]
}

// Keys возвращает срез ключей карты.
// Порядок ключей неважен, и не обязан совпадать
// с порядком значений из метода Values.
func (m Map[K, V]) Keys() []K {
	s := []K{}
	for k := range m {
		s = append(s, k)
	}
	return s
}

// Values возвращает срез значений карты.
// Порядок значений неважен, и не обязан совпадать
// с порядком ключей из метода Keys.
func (m Map[K, V]) Values() []V {
	s := []V{}
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

// конец решения

func main() {
	m := Map[string, int]{}
	m.Set("one", 1)
	m.Set("two", 2)

	fmt.Println(m.Get("one")) // 1
	fmt.Println(m.Get("two")) // 2

	fmt.Println(m.Keys())   // [one two]
	fmt.Println(m.Values()) // [1 2]
}
