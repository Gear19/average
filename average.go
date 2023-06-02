package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
func main() {
	var count []float64
	for _, number := range os.Args[1:] {
		num, err := strconv.ParseFloat(number, 64)
		if err != nil {
			log.Fatal(err)
		}
		count = append(count, num)
	}
	arguments := average(count...)

	fmt.Printf("Average: %0.2f", arguments)
}
