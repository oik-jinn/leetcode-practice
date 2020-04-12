package main

import "fmt"
import "os"

func main() {
	input := [] string {"flower","flow","flight"}

	fmt.Println(longerstCommonPrefix(input))
}

func longerstCommonPrefix(strs []string) string {
	length := len(strs)
	if (length == 0) {
		return ""
	}

	return commonPrefix(strs, 0, length-1)
}

func commonPrefix(strs []string, left int, right int) string {
	if left == right {
		return strs[left]
	}

	var mid int
	mid := (left + right) / 2
	leftStr := commonPrefix(strs, left, mid)
	rightStr := commonPrefix(strs, mid+1, right)

	return compareString(leftStr, rightStr)
}

func compareString(left string, right string) string {
	
}