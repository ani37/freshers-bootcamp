package main
import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

var balance = 500
var wg sync.WaitGroup
var m sync.Mutex

func deposit(amount int) {
	m.Lock()
	balance += amount
	fmt.Println("Successfully deposit performed, current account balance", balance)
	m.Unlock()
	wg.Done()
}

func withdrawal(amount int) {
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

	wg.Add(2)
	go deposit( rand.Intn(1000))
	go withdrawal( rand.Intn(1000))

	wg.Wait()
	fmt.Println("final account balance:", balance)
}