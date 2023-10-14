package main

import "fmt"

func main() {
	var studentGrades = make(map[string]float32)
	studentGrades["Sandeep"] = 9.2
	studentGrades["Satyam"] = 9.9
	studentGrades["Tarun"] = 8.6
	fmt.Println("The grade of each student")
	for _, v := range studentGrades {
		fmt.Println(v)
	}
	delete(studentGrades, "Satyam")
	fmt.Println("After removing", studentGrades)
}
