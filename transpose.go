package main

import "fmt"

func main() {
	A := [][]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(transpose(A))
}

func transpose(A [][]int) [][]int {
	var length int
	lengthA := len(A)
	lengthB := len(A[0])
	if lengthA > lengthB {
		length := lengthA
	} else {
		length := lengthB
	}
	R := make([][]int, length)

	for i := 0; i < lengthB; i++ {
		for j := 0; j < lengthA; j++ {
			T := make([]int, lengthA)
			T = A[j][i]
		}
		R[j] = T
	}
	// for i, val := range A {
	// 	T := make([]int, length)
	// 	for j, _ := range val {
	// 		temp := A[i][length-j-1]
	// 		if temp == 0 {
	// 			T[j] = 1
	// 		} else {
	// 			T[j] = 0
	// 		}
	// 	}
	// 	B[i] = T
	// }

	return T
}
