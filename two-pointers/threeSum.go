package main

import (
	"os"
	"fmt"
	"sort"
)

func main() {

	nums := [] int {-1, 0, 1, 2, -1, -4}
	nums = [] int {0,0,0,0}
	// nums = [] int {1,-1,-1,0}
	nums = [] int {-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}

	result := threeSum(nums)
	fmt.Println(result)
	os.Exit(1)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	for left := 0; left < len(nums) && nums[left] <= 0; left++ {
		middle := left + 1
		right := len(nums)-1
		tagert := 0-nums[left]

		if left > 0 && nums[left] == nums[left-1] {
			continue
		}
		for middle < right {
			if nums[middle]+nums[right] < tagert {
				middle ++
			} else if nums[middle]+nums[right] > tagert {
				right--
			} else {
				temp := []int{nums[left], nums[middle], nums[right]}
				result = append(result, temp)
				middle++
				right--
				for middle < len(nums) && nums[middle] == nums[middle-1] {
					middle++
					continue
				}

				for right >0 && nums[right] == nums[right+1] {
					right--
					continue
				}
			}
		}
	}

	return result	
}