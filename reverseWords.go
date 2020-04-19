package main

import "fmt"
import "os"

func main() {
	s := "the sky is blue"
	string := "  hello world!  " 

	fmt.Println(reverseWords(s))
	fmt.Println(reverseWords(string))
	os.Exit(1)
}

func reverseWords(s string) string {
	stringMap := make(map[int]string)
	i := 0
	flag := false
	for _, r := range s {
		c := string(r)
		if c != " " {
			stringMap[i] += c
			flag = true
		} else {
			if flag {
				i++
			}
			flag = false
		}
	}

	result := ""
	for j := i; j >=0; j-- {
		if stringMap[j] != "" {
			result += stringMap[j]
			if j!=0 {
				result += " "
			}
		}
	}

	return result
}