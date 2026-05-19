func longestConsecutive(nums []int) int {
    set := make(map[int]bool)

    for _, val := range nums {
        set[val] = true
    }

    result := 0

    for num := range set {
        if !set[num-1] {
            current := num
            count := 1

            for set[current+1] {
                current++
                count++
            }

            if count > result {
                result = count
            }
        }
    }

    return result
}
