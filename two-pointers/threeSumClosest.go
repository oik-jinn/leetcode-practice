package main

import (
	"os"
	"fmt"
	"sort"
)

func main () {
	nums := []int {-1,2,1,-4}
	// nums = []int {1,1,1,0}
	nums = []int {1,2,5,10,11}
	target := 12
	result := threeSumClosest(nums, target)
	fmt.Println(result)
	os.Exit(1)
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var distance int
	for left := 0; left < len(nums)-2; left++ {
		mid := left+1
		right := len(nums) -1

		search := target - nums[left]
		for mid < right {
			temp := nums[mid] + nums[right]
			distanceT := search - temp
			if (mid == 1 && right == len(nums)-1)  || abs(distanceT) < abs(distance) {
				distance = distanceT
			}

			// fmt.Println(nums[left], nums[mid], nums[right], search, distanceT, distance)
			if distanceT > 0 {
				mid++
			} else if distanceT < 0 {
				right--
			} else {
				return target
			}
		}
	}

	// fmt.Println(distance)
	return target - distance
}

func abs(num int) int {
	if num >= 0 {
		return num 
	} else {
		return 0-num
	}
}