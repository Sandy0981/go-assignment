package main

import "fmt"

func main() {
	var studentData = make(map[string]map[string]interface{})
	studentData["Sandeep"] = map[string]interface{}{
		"Age":   29,
		"Grade": 9.2,
		"City":  "Jodhpur",
	}
	studentData["Satyam"] = map[string]interface{}{
		"Age":   30,
		"Grade": 8.2,
		"City":  "Varanasi",
	}
	studentData["Tarun"] = map[string]interface{}{
		"Age":   25,
		"Grade": 8.9,
		"City":  "Bangalore",
	}
	for i, v := range studentData {
		fmt.Println(i, v)
	}
	fmt.Println(studentData)
}
