package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	numbers := float64(0)
	for _, input := range arguments {
		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers += number
	}
	fmt.Printf("Average: %0.2f", numbers/float64(len(arguments)))
}
