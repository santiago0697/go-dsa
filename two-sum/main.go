package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for i, n := range nums {
		expected := target - n
		if pos, ok := seen[expected]; ok {
			return []int{i, pos}
		}
		seen[n] = i
	}
	return []int{}
}

func main() {
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 9))
}
