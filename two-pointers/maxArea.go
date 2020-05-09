package main

import (
	"fmt"
	"os"
)

func main() {
	A := []int{1,8,6,2,5,4,8,3,7}

	fmt.Println(maxArea(A))
	os.Exit(1)
}

func maxArea(height []int) int {
	area := 0
	begin :=0
	end := len(height)-1

	for begin < end {
		temp := minNum(height[begin], height[end]) * (end-begin)
		if temp > area {
			area = temp
		}

		if height[begin] < height[end] {
			begin++
		} else {
			end --
		}
	}

	return area
}

func minNum(a int, b int) int {
	if (a < b) {
		return a
	} else {
		return b
	}
}