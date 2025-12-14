package main

func isSubsequence(s string, t string) bool {
	characterPointer := 0

	for _, char := range []byte(t) {
		if characterPointer == len(s) {
			return true
		}

		if char == s[characterPointer] {
			characterPointer++
		}
	}

	return characterPointer == len(s)
}
