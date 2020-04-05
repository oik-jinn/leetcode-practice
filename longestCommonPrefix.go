package main

import "fmt"

func main() {
	input := [] string {"flower","flow","flight"}

	fmt.Println(longestCommonPrefix(input))
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
