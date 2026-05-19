func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, v := range nums {
		compliment := target - v
		if idx, found := hashMap[compliment]; found {
			return []int{idx, i}
		}
		hashMap[v] = i
	}

	return nil    
}
