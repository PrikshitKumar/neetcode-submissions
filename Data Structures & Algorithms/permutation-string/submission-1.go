func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var s1Count, window [26]int

	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
	}

	left := 0

	for right := 0; right < len(s2); right++ {
		window[s2[right]-'a']++

		if right - left + 1 > len(s1) {
			window[s2[left]-'a']--
			left++
		}

		if window == s1Count {
			return true
		}

	}

	return false
}
