package main

import (
	"fmt"
	"os"
)

func main() {
	input := [] string {"flower","flow","flight"}

	fmt.Println(longestCommonPrefix(input))
	fmt.Println(longestCommonPrefix1(input))
	os.Exit(1)
}

func longestCommonPrefix(strs []string) string {
	var minLength int
	var minStr string
	for i, str := range strs {
		if i == 0 {
			minLength = len(str)
			minStr = str
		} else {
			currentLength := len(str)
			if currentLength < minLength {
				minStr = str
				minLength = currentLength
			}
		}
	}

	var prefix string
	for i := minLength; i > 0; i-- {
		prefix = minStr[0:i]
		prefixLength := len(prefix)
		flag := true
		for _, str := range strs {
			if prefix != str[0:prefixLength] {
				flag = false
				break
			}
		}

		if flag {
			return prefix
		}
	}

	return ""
}

func longestCommonPrefix1(strs []string) string {
	if len(strs) ==0 {
		return ""
	}
	var prefix string
	prefix = strs[0]
	for i := 1; i < len(strs); i++ {
		curentStr := strs[i]
		var minStr, compareStr string
		if len(prefix) > len(curentStr) {
			minStr = curentStr
			compareStr = prefix
		} else {
			minStr = prefix
			compareStr = curentStr
		}

		flag := false
		for i := len(minStr); i > 0; i-- {
			if minStr[0:i] == compareStr[0:i] {
				prefix = minStr[0:i]
				flag = true
				break
			}
		}

		if !flag {
			prefix = ""
			break
		}
	}

	return prefix
}