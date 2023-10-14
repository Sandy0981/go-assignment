package person

import (
	"8c/model"
	"fmt"
)

func PrintPerson(P model.Person) {
	fmt.Println(P.Name, "\n", P.Age)
}
