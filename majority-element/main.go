package main

func majorityElement(nums []int) int {
	numsSeen := make(map[int]int, len(nums))
	maxNumber := numsSeen[0]
	maxCounter := 0

	for _, n := range nums {
		numsSeen[n] += 1

		if numsSeen[n] > maxCounter {
			maxNumber = n
			maxCounter = numsSeen[n]
		}
	}

	return maxNumber

	// countMap := make(map[int]int)
	// n := len(nums)
	// for _, num := range nums {
	// 	countMap[num] += 1
	// 	// If an element's count exceeds n/2, return it
	// 	if countMap[num] > n/2 {
	// 		return num
	// 	}
	// }
	// return -1 // Shouldn't reach here if input is valid
}

