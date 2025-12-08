package main

func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		c := s[right]

		if idx, ok := lastSeen[c]; ok && idx >= left {
			left = idx + 1
		}

		lastSeen[c] = right

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}
