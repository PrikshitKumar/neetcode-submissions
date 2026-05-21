func lengthOfLongestSubstring(s string) int {
	left, maxLen := 0, 0
	hashMap := make(map[byte]bool, 0)

	for right := 0; right < len(s); right++ {
		for hashMap[s[right]] {
			delete(hashMap, s[left])
			left++
		} 

		hashMap[s[right]] = true

		if right - left + 1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen

}
