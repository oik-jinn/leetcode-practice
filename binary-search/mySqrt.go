package main

import (
	"fmt"
	"os"
)

func main() {
	A := 1

	fmt.Println(mySqrt(A))
	os.Exit(1)
}

func mySqrt(x int) int {
	if x < 2 {
		 return x
	}
	st := 1
	end := x/2

	for st < end {
		mid := st + (end - st) / 2 +1
		temp := mid * mid
		if temp > x {
			end = mid -1
		} else {
			st = mid
		}
	}

	return end
}