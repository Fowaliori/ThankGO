package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("not found")

var lang languages = languages{
	"go":         "golang",
	"python":     "it python",
	"javascript": "js is cool",
}

type languages map[string]string

func (l languages) describe(lang string) (string, error) {
	descr, err := getValue(l, lang)
	if err != nil {
		return "", err
	}
	return descr, nil
}

func getValue(m map[string]string, key string) (string, error) {
	val, ok := m[key]
	if !ok {
		return "", errNotFound
	}
	return val, nil
}

func getchar(str string, idx int) byte {
	return str[idx]
}

func protect(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("Good")
		}
	}()
	fn()
}

func main() {
	protect(func() {
		fmt.Println(getchar("home", 10))
	})
	protect(func() {
		fmt.Println(getchar("home", 2))
	})
	descr, err := lang.describe("java")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(descr)
}
