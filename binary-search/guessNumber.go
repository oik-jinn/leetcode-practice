package main

import (
	"fmt"
	"os"
)

func main() {
	A := 10

	fmt.Println(guessNumber(A))
	os.Exit(1)
}

func guessNumber(x int) int {
	st := 1
	end := x

	for st < end {
		mid := st + (end - st) / 2 
		if guess(mid) > 0 {
			st = mid + 1
		} else if guess(mid) < 0 {
			end = mid - 1
		} else {
			 return mid
		}
	}

	return st
}

func guess(x int) int {
	if x > 6 {
		return -1
	} else if x < 6 {
		return 1
	} else {
		return 0
	}
}