package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	a, b float64
}

type circle struct {
	radius float64
}

type counter struct {
	val uint
}

type usage struct {
	service string
	Counter
}

type Counter interface {
	increment()
	value() uint
}

func (c *counter) increment() {
	c.val++
}

func (c *counter) value() uint {
	return c.val
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (r rect) area() float64 {
	return r.a * r.b
}

func (r rect) perim() float64 {
	return 2*r.a + 2*r.b
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{4, 3}
	fmt.Println(r.area())
	fmt.Println(r.perim())
	c := circle{3.5}
	measure(c)
	u := usage{"dsada", &counter{}}
	u.increment()
	u.increment()
	fmt.Println(u.value())
	var a any = "rsada"
	switch v := a.(type) {
	case string:
		fmt.Println("string", v)
	case int:
		fmt.Println("int", v)
	case float64:
		fmt.Println("float64", v)
	default:
		fmt.Println("unkown", v)
	}
}
