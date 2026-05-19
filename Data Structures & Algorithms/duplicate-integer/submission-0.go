func hasDuplicate(nums []int) bool {
    hashMap := make(map[int]int);
    for i:=0; i<len(nums); i++ {
        hashMap[nums[i]]++; 
    }
    for _, v := range hashMap {
        if v > 1 {
            return true;
        }
    }

    return false
}
