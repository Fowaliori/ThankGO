package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type error interface {
	Error() string
}

type lookupError struct {
	src    string
	substr string
}

func (e lookupError) Error() string {
	return fmt.Sprintf("'%s' not found in '%s'", e.substr, e.src)
}

func indexOF(src string, substr string) (int, error) {
	idx := strings.Index(src, substr)
	if idx == -1 {
		return -1, lookupError{src, substr}
	}
	return idx, nil
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("x должен быть > 0")
	}
	return math.Sqrt(x), nil
}

func main() {
	nums := []float64{1.4, 6.2, -4.2, 7.8, -1.2, -2.9}
	for _, v := range nums {
		if res, err := sqrt(v); err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("sqrt(%v) = %v\n", v, res)
		}
	}
	src := "Hello I am king"
	str := []string{"ell", "am", "kigg", "I am"}
	for _, v := range str {
		if res, err := indexOF(src, v); err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("indexOF(%v, %v) = %v\n", src, v, res)
		}
	}
}
