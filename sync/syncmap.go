package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// fmt.Println("done")
	// unsafeWrite()
	safeWrite()
	time.Sleep(time.Second)
}

type SafeMap struct {
	safeMap map[int]int
	sync.Mutex
}

func safeWrite() {
	fmt.Println("start safe write...")

	s := SafeMap{
		safeMap: map[int]int{},
		Mutex:   sync.Mutex{},
	}

	for i := 0; i < 100; i++ {
		go func() {
			s.Write(1, 1)
		}()
	}

}

func (s *SafeMap) Write(k, v int) (int, bool) {
	// fmt.Println(k, v)
	// s.Lock()
	// defer s.Unlock()
	result, ok := s.safeMap[k]
	return result, ok
}

func unsafeWrite() {
	fmt.Println("start unsafe write...")

	myMap := map[int]int{}
	for i := 0; i < 100; i++ {
		go func() {
			myMap[1] = i
			// fmt.Println(i)
		}()
		// fmt.Println(i)
	}
}
