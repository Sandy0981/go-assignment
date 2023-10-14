package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	s := Student{
		Name: "Sandeep",
		Age:  23,
	}
	update(&s)
	fmt.Println(s)
}
func update(s *Student) {
	s.Age += 1
}
