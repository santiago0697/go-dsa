package main

func moveZeroes(nums []int) []int{
	pointer := 0

	for k, n := range nums {
		if n != 0 {
			tmp := nums[pointer]
			nums[pointer] = n
			nums[k] = tmp
			pointer++
		}
	}

	return nums
}
