package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var output = 0

func allocate(wg *sync.WaitGroup) {

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	rating := rand.Intn(10) + 1
	output += rating
	wg.Done()
}

func main(){


	var noOfStudents = 200
	var wg sync.WaitGroup

	for i := 0; i < noOfStudents; i++ {
		wg.Add(1)
		go allocate(&wg)
	}

	wg.Wait()
	result := output / noOfStudents

	fmt.Println("the average rating given from the students is", result)
}