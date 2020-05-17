package main

import (
	"fmt"
	"os"
)

func main() {
	nums := []int{4,5,6,7,0,1,2}
	// nums = []int{1}
	target := 0
	fmt.Println(search(nums, target))
	os.Exit(1)
}

func search(nums []int, target int) int {
	st := 0
	end := len(nums) -1

	for st <= end {
		mid := st + (end -st) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[0] {
			if target >= nums[0] && target <= nums[mid] {
				temp := append(nums[:0:0], nums[0:mid]...)
				return binarySearch(temp, target)
			} else {
				temp := append(nums[:0:0], nums[mid+1:]...)
				result := search(temp, target)
				if result == -1 {
					return -1
				}
				return result + mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[end] {
				temp := append(nums[:0:0], nums[mid+1:]...)
				result := binarySearch(temp, target)
				if result == -1 {
					return -1
				}
				return result + mid + 1
			} else {
				temp := append(nums[:0:0], nums[0:mid]...)
				return search(temp, target)
			}
		}
	}

	return -1
}

func binarySearch(a []int , target int) int {
	st := 0 
	ed := len(a) -1
	if target < a[st] || target > a[ed] {
		return -1
	}

	for st < ed {
		mid := st + (ed-st) / 2
		if a[mid] < target {
			st = mid + 1
		} else {
			ed = mid
		}
	}

	if a[st] != target {
		return -1
	}
	return st
}