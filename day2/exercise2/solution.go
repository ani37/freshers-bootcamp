package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func allocate(noOfStudents int, ratings chan int) {
	for i := 0; i < noOfStudents; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		rating := rand.Intn(10) + 1
		ratings <- rating
	}
	close(ratings)
}

func calculate(wg *sync.WaitGroup,ratings chan int, average chan int){

	output :=0

	for rating := range ratings {
		output += rating
	}

	output /= cap(ratings)
	average <- output
	wg.Done()
}

func main(){

	var ratings = make(chan int, 200)
	var average = make(chan int, 1)

	go allocate(200, ratings)

	var wg sync.WaitGroup
	wg.Add(1)

	go calculate(&wg, ratings, average)

	wg.Wait()

	result := <- average
	fmt.Println("the average rating given from the students is", result)
}