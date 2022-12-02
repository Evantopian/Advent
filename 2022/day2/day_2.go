package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// every possible unique combination
	mapPt1 := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	mapPt2 := map[string]int{
		"A X": 3,
		"B X": 1,
		"C X": 2,
		"A Y": 4,
		"B Y": 5,
		"C Y": 6,
		"A Z": 8,
		"B Z": 9,
		"C Z": 7,
	}

	scanner := bufio.NewScanner(file)
	pt1, pt2 := 0, 0

	for scanner.Scan() {
		match := string(scanner.Text())
		pt1 += mapPt1[match]
		pt2 += mapPt2[match]

	}
	fmt.Println(pt1, pt2)
}
