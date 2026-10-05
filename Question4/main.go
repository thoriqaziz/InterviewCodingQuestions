// ABC repeatedly in sequence using three goroutines
// cycle 6 -> ABCABCABCABCABCABC

package main

import (
	"fmt"
	"sync"
)

func main() {
	const cycle = 6
	chA := make(chan bool)
	chB := make(chan bool)
	chC := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()

		for i := 0; i < cycle; i++ {
			<-chA
			fmt.Print("A")
			chB <- true
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < cycle; i++ {
			<-chB
			fmt.Print("B")
			chC <- true
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < cycle; i++ {
			<-chC
			fmt.Print("C")
			if i < cycle-1 {
				chA <- true
			}
		}
	}()

	chA <- true
	wg.Wait()
}
