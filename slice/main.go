package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Printf("%#v\n", s)
	fmt.Println(len(s))
	s2 := []int{1, 2, 3}
	fmt.Println(s2)
	s2 = append(s2, 4)
	fmt.Println(s2)
	carr := make([]int, len(s2))
	copy(carr, s2)
	fmt.Println(carr, s2)
	s3 := s2[1:]
	fmt.Println(s3)
	str := "ававав"
	bytes := []byte(str)
	fmt.Println(bytes)
	fmt.Println([]rune(str))
	fmt.Println(str[1])
}
