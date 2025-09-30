package main

import (
	"fmt"
	"math"
	"unicode"
)

type person struct {
	age  int
	name string
}

type book struct {
	title  string
	author *person
}

type rect struct {
	a, b int
}

type circle struct {
	radius int
}

type counter struct {
	value uint
}

func (c *counter) increment() {
	c.value++
}

type method struct {
	source string
	counter
}

type site struct {
	website string
	counter
}

type inn string

type ID int

type ID2 = int

func (id inn) isValid() bool {
	if len(id) != 12 {
		return false
	}
	for _, char := range id {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func (c *circle) area() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2)
}

func (r rect) area() int {
	return r.a * r.b
}

func (r *rect) scale() {
	r.a *= 2
	r.b *= 3
}

func newPerson(name string, age int) *person {
	p := person{age, name}
	return &p
}

func main() {
	bair := person{age: 25, name: "Bair"}
	fmt.Println(bair)
	vlad := newPerson("Vlad", 23)
	fmt.Println(vlad.age, vlad.name)
	kirill := &person{24, "Kirill"}
	fmt.Println(kirill, kirill.age, (*kirill).age)
	b := book{"title book", &bair}
	fmt.Println(b)
	r := rect{4, 6}
	fmt.Println(r.area())
	c := circle{5}
	fmt.Println(c.area())
	r.scale()
	fmt.Println(r)
	id := inn("123456789101")
	fmt.Println(id.isValid())
	m := method{source: "Slice"}
	m.counter.increment()
	m.counter.increment()
	m.increment()
	fmt.Println(m.counter.value)
	s := site{"google", counter{125}}
	s.increment()
	fmt.Println(s.counter.value)
	sd := new(counter)
	inc := sd.increment
	inc()
	inc()
	fmt.Println(sd.value)
	var n int = 10
	var idi ID = 10
	//id = n
	fmt.Println(n, idi)
	var id2 ID2 = 10
	id2 = n
	fmt.Println(id2)
}
