func groupAnagrams(strs []string) [][]string {
	hashmap := make(map[string][]string)
	for _, val := range strs {
		chars := []rune(val)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		key := string(chars)
		hashmap[key] = append(hashmap[key], val)
	}

	var result [][]string
	for _, group := range hashmap {
		result = append(result, group)
	}

	return result
}
