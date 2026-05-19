func productExceptSelf(nums []int) []int {
    zeroCount := 0
    total := 1

    for _, val := range nums {
        if val == 0 {
            zeroCount++
        } else {
            total *= val
        }
    }

    result := make([]int, len(nums)) 

    for i, val := range nums {
        if zeroCount > 1 {
            result[i] = 0
        } else if zeroCount == 1 {
            if val == 0 {
                result[i] = total
            } else {
                result[i] = 0
            }
        } else {
            result[i] = total / val
        }
    }

    return result
}
