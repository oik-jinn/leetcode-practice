package main

import "fmt"
import "os"
import "strings"

func main() {
	path := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(path))
	os.Exit(1)
}

func simplifyPath(path string) string {
	pathMap := make(map[int] string)
	temp := strings.Split(path, "/")
	i := 0
	for _, v := range temp {
		if v != "" && v != ".." && v != "." {
			pathMap[i] = v
			i++
		} else if v == ".." {
			i--
			delete(pathMap, i)
			if i < 0 {
				i = 0
			}
		}
	}

	result := "/"
	for j := 0; j < i; j++ {
		str := pathMap[j]
		if str != "" {
			result += str
		}
		val, ok := pathMap[j+1]
		if ok && val != "" {
			result += "/"
		}
	}

	return result    
}