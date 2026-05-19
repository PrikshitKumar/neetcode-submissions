func hasDuplicate(nums []int) bool {
    seen := make(map[int]struct{})
    for _, val := range nums {
        if _, exists := seen[val]; exists {
            return true
        }
        seen[val] = struct{}{}
    }
    return false
}
