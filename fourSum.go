package main

import (
	"os"
	"fmt"
	"sort"
)

func main() {

	nums := [] int {1,0,-1,0,-2,2}
	nums = [] int {0,0,0,0,0}
	target := 0

	// nums = [] int {-1,0,1,2,-1,-4}
	// target = -1

	result := fourSum(nums, target)
	fmt.Println(result)
	os.Exit(1)
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for start:=0; start<len(nums)-3; start++ {
		search := target - nums[start]

		if start > 0 && nums[start] == nums[start-1] {
			continue
		}
		for left := start+1; left < len(nums)-2; left++ {
			secondSearch := search-nums[left]

			mid := left+1
			right := len(nums)-1
			if left >start+1 && nums[left] == nums[left-1]  {
				continue
			}

			for mid < right {
				if nums[mid] + nums[right] < secondSearch {
					mid++
				} else if  nums[mid] + nums[right] > secondSearch {
					right--
				} else {
					temp := []int {nums[start], nums[left], nums[mid], nums[right]}
					result = append(result, temp)

					for mid < len(nums)-2 && nums[mid] == nums[mid+1] {
						mid++
					}
					for right > 0 && nums[right] == nums[right-1] {
						right--
					}
					mid++
					right--
				}
			}
		}
	}
	
	return result
}