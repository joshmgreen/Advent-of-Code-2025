package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file into a slice of strings
	var lines []string

	// Create scanner for file
	scanner := bufio.NewScanner(file)

	// Loop over input file and store in slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Close the file
	file.Close()

	// zeroCount := 0
	direction := ""
	rotations := 0
	startingPosition := 50
	zeroCount := 0

	// Loop over input
	for _, value := range lines {
		direction = value[:1]
		rotations, err = strconv.Atoi(value[1:])
		if direction == "L" {
			startingPosition = startingPosition - rotations
		} else {
			startingPosition = startingPosition + rotations
		}
		if startingPosition == 0 {
			zeroCount++
		}
		if startingPosition > 99 {
			difference := startingPosition - 100
			startingPosition = startingPosition - difference
		}
		if startingPosition < 0 {
			difference := startingPosition + 100
			startingPosition = startingPosition + difference
		}
	} 
	fmt.Println(zeroCount)
}