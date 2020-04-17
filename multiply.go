package main

import "fmt"
import "os"
import "strconv"

func main() {
	num1 := "123" 
	num2 := "456"

	fmt.Println(multiply(num1, num2))
	os.Exit(1)
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	len1 := len(num1)
	len2 := len(num2)

	var result = make([]int, len1+len2)
	for i := len1-1; i >=0; i-- {
		for j:= len2-1; j>=0; j-- {
			temp := int(num1[i]-48) * int(num2[j]-48) + result[i+j+1]
			if temp >=10 {
				result[i+j+1] = temp % 10
				result[i+j] += temp / 10
			} else {
				result[i+j+1] = temp
			}
		}
	}

	var str string
	for i, v := range result {
		if !(i==0 && v==0) {
			str += strconv.Itoa(v)
		}
	}

	return str
}