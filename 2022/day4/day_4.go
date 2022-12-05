package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	count := 0
	touch := 0
	for scanner.Scan() {
		line := strings.Split(strings.Replace(scanner.Text(), ",", "-", 4), "-")
		x := make([]int, len(line))

		for i, num := range line {
			x[i], err = strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
		}

		// pt 1: if within interval, increment count
		// pt 2: if a touches b, increment touch
		if x[0] >= x[2] && x[1] <= x[3] {
			count++
		} else if x[0] <= x[2] && x[1] >= x[3] {
			count++
		} else if x[1] >= x[2] && x[1] <= x[3] {
			touch++
		} else if x[0] >= x[2] && x[0] <= x[3] {
			touch++
		}

	}
	fmt.Printf("Pt1: %d\nPt2: %d", count, count+touch)
	fmt.Println()
}
