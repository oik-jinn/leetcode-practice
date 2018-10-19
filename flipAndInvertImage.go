package main

import "fmt"

func main() {
	A := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}

	fmt.Println(flipAndInvertImage(A))
}

func flipAndInvertImage(A [][]int) [][]int {
	length := len(A)
	B := make([][]int, length)
	for i, val := range A {
		T := make([]int, length)
		for j, _ := range val {
			temp := A[i][length-j-1]
			if temp == 0 {
				T[j] = 1
			} else {
				T[j] = 0
			}
		}
		B[i] = T
	}

	return B
}
