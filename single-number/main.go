package main

func singleNumber(numbers []int) int {
	number := 0 

	for _, n := range numbers {
		number ^= n
	}

	return number
}
