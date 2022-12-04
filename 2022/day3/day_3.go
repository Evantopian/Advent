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

	lower := make(map[string]int)

	for i := 0; i < 26; i++ {
		lower[string(rune(i+97))] = 1 + i
		lower[string(rune(i+65))] = 27 + i
		//fmt.Println(string(rune(i+97)), lower[string(rune(i+97))])
	}

	scanner := bufio.NewScanner(file)

	k := 0
	for scanner.Scan() {
		ruck_sack := make(map[rune]int)

		text := scanner.Text()
		first := text[0 : len(text)/2]
		second := text[len(text)/2:]

		// 97 (lower)
		for _, ch := range first {
			ruck_sack[ch] = lower[string(ch)]
		}

		for _, ch := range second {
			if ruck_sack[ch] != 0 {
				k += lower[string(ch)]
				fmt.Print(" - : ", lower[string(ch)])
				break
			}
		}
		fmt.Println()
	}
	fmt.Println("k", k)

}
