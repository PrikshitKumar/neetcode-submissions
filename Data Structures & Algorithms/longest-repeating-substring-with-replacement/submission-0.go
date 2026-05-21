func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	left, maxFreq, maxLen := 0, 0, 0

	for right := 0; right < len(s); right++ {
		count[s[right]]++

		if count[s[right]] > maxFreq {
			maxFreq = count[s[right]]
		}

		for (right - left + 1) - maxFreq > k {
			count[s[left]]--
			left++
		} 

		if right - left + 1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen

}
