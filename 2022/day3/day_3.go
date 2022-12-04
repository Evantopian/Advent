package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var alpha = make(map[string]int)

func sim(arr []string, size int) int {
	sim_map := make(map[rune]int)
	keys := ""

	for i, item := range arr {
		for _, ch := range item {
			if sim_map[ch] == 0 {
				keys += string(ch)
			}
			if sim_map[ch] == i {
				sim_map[ch]++
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
			//fmt.Println(alpha[string(key)])
			return alpha[string(key)]
		}
	}

	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < 26; i++ {
		alpha[string(rune(i+97))] = 1 + i
		alpha[string(rune(i+65))] = 27 + i
	}

	scanner := bufio.NewScanner(file)
	three_sack := make([]string, 3)
	var p1, p2, rep = 0, 0, 0

	for scanner.Scan() {
		sack := make([]string, 2)

		text := scanner.Text()
		sack[0] = text[0 : len(text)/2]
		sack[1] = text[len(text)/2:]

		// 97 (lower)
		three_sack[rep] = text
		rep++
		p1 += sim(sack, 2)
		if (rep)%3 == 0 {
			p2 += sim(three_sack, 3)
			three_sack = three_sack[:3]
			rep = 0

		}

	}
	fmt.Printf("pt1: %d\npt2: %d", p1, p2)

}
