package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func print(freq map[string]int)  {
	// Convert structs to JSON.
	data, err := json.MarshalIndent(freq,""," ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

func main(){

	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	freq := make(map[string] int)

	for i := 0; i < len(words); i++ {
		done := make(chan bool)

		go func() {
			for _, ch := range words[i] {
				freq[string(ch)]++
			}
			done <- true
		}()

		<- done
	}

	print(freq)

}