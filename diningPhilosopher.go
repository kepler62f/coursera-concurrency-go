package main

import (
	"fmt"
	"sync"
)

// https://hackernoon.com/dancing-with-go-s-mutexes-92407ae927bf
// https://stackoverflow.com/questions/37135193/how-to-set-default-values-in-go-structs

type Host struct {
	sync.Mutex
	guestCount int
}

func (h *Host) addGuest() {
	h.Lock()
	h.guestCount++
	h.Unlock()
}

func (h *Host) removeGuest() {
	h.Lock()
	if h.guestCount > 0 {
		h.guestCount--
	}
	h.Unlock()

}

func (h *Host) permitNewGuest() bool {
	if h.guestCount < 2 {
		return true
	}
	return false
}

func (h *Host) getGuestCount() int {
	return h.guestCount
}

func NewHost() *Host {
	// Mutex does not need init -- https://stackoverflow.com/questions/45744165/do-mutexes-need-initialization-in-go
	return &Host{guestCount: 0}
}

type ChopS struct{ sync.Mutex }  // Declaring types: https://www.golang-book.com/books/intro/9

type Philo struct {
	leftCS, rightCS *ChopS
	// leftCS and rightCs is declared as pointers to type ChopS
	// Pointer to type: https://medium.com/learning-the-go-programming-language/pointing-to-go-the-go-pointer-type-a3c3f587592f
	eatCount int
	i int // index of philosopher
}

func (p Philo) eat(h *Host) {
	// p is the receiver of function eat();
	// https://stackoverflow.com/questions/34031801/function-declaration-syntax-things-in-parenthesis-before-function-name
	// http://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/

	for p.eatCount < 3 {

		if h.permitNewGuest() {

			h.addGuest()

			p.leftCS.Lock()
			p.rightCS.Lock()

			fmt.Println("starting to eat", p.i+1)  // Philosopher numbered from 1 to 5 but indexed 0 to 4
			p.eatCount++

			fmt.Println("finishing eating", p.i+1)

			p.rightCS.Unlock()
			p.leftCS.Unlock()

			h.removeGuest()
		}
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {

	host := NewHost()

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i],CSticks[(i+1)%5], 0, i}
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(host)
	}
	wg.Wait()
}
