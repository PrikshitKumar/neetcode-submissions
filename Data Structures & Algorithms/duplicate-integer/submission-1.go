func hasDuplicate(nums []int) bool {
    hashMap := make(map[int]int)
    for _, val := range nums {
        hashMap[val]++;
        if hashMap[val] != 1 {
            return true
        }
    }
    return false
}
