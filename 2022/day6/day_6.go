package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	pt1, pt2 := 0, 0

	pt1 = solve(input, 4)
	pt2 = solve(input, 14)
	fmt.Printf("pt1: %d, pt2: %d", pt1, pt2)
	fmt.Println()
}

// instead of scanner, a byte is always of uint8,
// so we can use thast for ascii conversion.
func solve(arr []byte, size int) int {
	ch := make([]int, 26)
	marker, i := 0, 0
	for ; i-marker != size && marker < len(arr)-size-1; i++ {
		pos := ch[arr[i]-97]
		if pos != 0 && pos >= marker {
			marker = pos + 1
		}
		ch[arr[i]-97] = i
	}
	return i
}
