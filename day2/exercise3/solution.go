package main
import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

var balance = 500

func deposit(amount int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	balance += amount
	fmt.Println("Successfully deposit performed, current account balance", balance)
	m.Unlock()
	wg.Done()
}

func withdrawal(amount int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	if balance >= amount{
		balance -= amount
		fmt.Println("Successfully withdrawal performed, current account balance", balance)
	} else{
		log.Fatal("lack of balance in account...")
	}

	m.Unlock()
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex

	w.Add(2)
	go deposit( rand.Intn(1000), &w, &m)
	go withdrawal( rand.Intn(1000), &w, &m)

	w.Wait()
	fmt.Println("final account balance:", balance)
}