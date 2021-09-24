package main

import (
	"fmt"
)

func main() {
	if test := 3; test >4 {
		fmt.Println("ok")
	} else if test=6;test > 5 {
		fmt.Println("xx")
	} else {
		fmt.Println("aa")
	}
}
