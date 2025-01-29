package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	// two pointers technique
	even, odd := 0, 1

	for even < len(nums) && odd < len(nums) {
		// calc odd num
		if nums[even]%2 == 1 {
			nums[even], nums[odd] = nums[odd], nums[even]
			odd += 2
		} else {
			even += 2
		}
	}

	return nums
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
	nums2 := []int{1, 4, 5, 2, 3, 6, 7, 8}

	fmt.Println(sortArrayByParityII(nums))  // [0 1 2 3 4 5 6 7]
	fmt.Println(sortArrayByParityII(nums2)) // [4 1 2 5 6 3 8 7]
}
