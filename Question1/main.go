// Write a program that prints nmbers from 1 to 10 using two goroutines in sequence
// odd-even-problem

package main

import "fmt"

func main() {
	oddCh := make(chan bool)
	evenCh := make(chan bool)
	doneCh := make(chan bool)

	go func() {
		for i := 1; i <= 9; i += 2 {
			<-oddCh
			fmt.Println(i)
			evenCh <- true
		}
	}()

	go func() {
		for i := 2; i <= 10; i += 2 {
			<-evenCh
			fmt.Println(i)
			if i == 10 {
				doneCh <- true
			}
			oddCh <- true
		}
	}()

	oddCh <- true

	<-doneCh
}
