package main

import (
	"fmt"
)

func main() {

	waitByChanel()
	// fmt.Println("waitgroup...")
	// for i := 0; i < 100; i++ {
	// 	go fmt.Println(i)
	// }
	// time.Sleep(time.Second)
}

func waitByChanel() {
	c := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-c
	}

}
