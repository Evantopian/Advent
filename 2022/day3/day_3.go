package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func sim(arr []string, size int) int {
	keys := ""
	for _, item := range arr {
		for _, ch := range item {
			if !strings.ContainsAny(keys, string(ch)) {
				keys += string(ch)
			}
		}
	}
	for _, key := range keys {
		found := true
		for i := 0; i < size; i++ {
			if !strings.ContainsRune(arr[i], key) {
				found = false
			}
		}
		if found == true {
			return alpha(key)
		}
	}
	return 0
}

func alpha(ch rune) int {
	if unicode.IsUpper(ch) {
		return (int(ch) - 65) + 27
	} else {
		return (int(ch) - 97) + 1
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	three_sack := make([]string, 3)
	var p1, p2, rep = 0, 0, 0

	for scanner.Scan() {
		sack := make([]string, 2)
		text := scanner.Text()

		sack[0] = text[0 : len(text)/2]
		sack[1] = text[len(text)/2:]
		p1 += sim(sack, 2)

		three_sack[rep] = text
		rep++

		if (rep)%3 == 0 {
			p2 += sim(three_sack, 3)
			three_sack = three_sack[:3]
			rep = 0
		}
	}
	fmt.Printf("pt1: %d\npt2: %d", p1, p2)
}
