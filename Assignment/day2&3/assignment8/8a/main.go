package main

import (
	"fmt"
	"os"
)

func removeFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := removeFile("a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = removeFile("b.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Both files removed")
}
