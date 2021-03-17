package main

import (
	"fmt"
	"log"
	"datafile"
)

func main()  {
	numbers, err := floats.GetFloats("./goText.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}