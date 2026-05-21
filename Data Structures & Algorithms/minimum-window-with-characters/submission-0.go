func minWindow(s string, t string) string {
    if len(s) < len(t) {
		return ""
	}

	countT := make(map[byte]int) 
	window := make(map[byte]int)

	for _, val := range t {
		countT[byte(val)]++
	}

	have, need := 0, len(countT)
	left := 0
	minLen := len(s) + 1
	start := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++

		if countT[c] > 0 && window[c] == countT[c] {
			have++
		}

		for have == need {
			if right - left + 1 < minLen {
				minLen = right - left + 1
				start = left
			}

			leftChar := s[left]
			window[leftChar]--

			if countT[leftChar] > 0 && window[leftChar] < countT[leftChar] {
				have--
			}

			left++
		}
	}

	if minLen == len(s) + 1 {
		return ""
	}

	return s[start : start + minLen]

}
