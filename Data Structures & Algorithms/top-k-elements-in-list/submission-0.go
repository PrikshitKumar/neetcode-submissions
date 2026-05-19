func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	for _, val := range nums {
		hashMap[val]++
	}

	type kv struct {
		key int
		value int
	}

	var list []kv
	for k, v := range hashMap {
		list = append(list, kv{k, v})
	}

	sort.Slice(list, func(i, j int) bool {
        return list[i].value > list[j].value
    })

	var result []int
	for i := 0; i < k; i++ {
		result = append(result, list[i].key)
	}

	return result

}
