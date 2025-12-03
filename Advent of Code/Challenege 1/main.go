package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Fields(string(file))

	currentPosition := 50
	zeroCount := 0

	// Loop over input
	for _, value := range lines {

		direction := string(value[0])
		rotations, err := strconv.Atoi(strings.TrimSpace(value[1:]))

		if err != nil {
			fmt.Println("Error: ", err)
		}

		if direction == "L" {
			currentPosition = ((currentPosition - rotations)%100 + 100) % 100
		}

		if direction == "R" {
			currentPosition = (currentPosition + rotations) % 100
		}

		if currentPosition == 0 {
			zeroCount++
		}
	} 
	fmt.Println(zeroCount)
}