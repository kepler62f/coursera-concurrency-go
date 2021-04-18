package main

/*
Race condition is when multiple program codes run concurrently and
are executed in an interleaving fashion. In such case, the outcome of
the program execution is non-deterministic, meaning that the ouput can
be different depending on the interleaving execution of codes.
 */

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var x int = 0

func main() {

	/*
	Runs two cycles; during each cycle the counter increments from 0 to 15 at one second interval
	and concurrently a checker checks the counter value at one-second interval. The output of
	the two cycles show different sequence of values, showing that the checker in each cycle does not get the same
	counter value each time.
	 */
	for i := 0; i < 2; i++ {
		fmt.Println("Increment Cycle #", i+1)
		wg.Add(1)
		go incrementCounter()
		wg.Add(1)
		go checkCounter()
		wg.Wait()
		x = 0
	}

}

func incrementCounter() {
	for i := 0; i < 15; i++ {
		x++
		time.Sleep(time.Second)
	}
	wg.Done()
}

func checkCounter() {
	for i := 0; i < 15; i++ {
		fmt.Print(x, " ")
		time.Sleep(time.Second)
	}
	fmt.Println()
	wg.Done()
}