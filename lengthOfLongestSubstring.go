package main

import (
	"fmt"
	"os"
)

func main() {
	// https://leetcode-cn.com/explore/featured/card/bytedance/242/string/1012/
	var s string
	s = "abcabcbb"
	// s = "bbbb"
	// s = "abbac"
	// s = "pwwkew"
	// s = ""
	// s = "ab"

	count := lengthOfLongestSubstring(s)
	fmt.Println(count)
	os.Exit(1)
}

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	count := 0
	tempCount := 0
	label := 0
	for i := 0; i < length; i++ {
		current := int(s[i])
		if tempCount > count {
			count = tempCount
		}
		tempCount = 1
		for j := i - 1; j >= label; j-- {
			previous := int(s[j])
			if current == previous {
				label = j + 1
				break
			} else {
				tempCount++
			}
		}
	}

	if tempCount > count {
		count = tempCount
	}

	return count
}
