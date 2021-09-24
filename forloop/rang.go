package main

import (
	"fmt"
)

func main() {

	var name string = "cadmus"

	fmt.Println(name)

	for index, char := range name {
		fmt.Println(index)
		fmt.Println(string(char))
	}

}
