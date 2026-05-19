func longestConsecutive(nums []int) int {
    hashmap := make(map[int]int)
    result := 0

    for _, val := range nums {
        hashmap[val]++
    }

    for i := 0; i < len(nums); i++ {
        count := 0
        curr := nums[i]
        for hashmap[curr] != 0 {
            count++
            curr++
            if result < count {
                result = count
            }
        }
    }

    return result
}
