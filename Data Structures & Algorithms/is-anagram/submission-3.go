func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    hashMap := make(map[rune]int)
    for _, v := range s {
        hashMap[v]++
    }    
    for _, v := range t {
        if hashMap[v] == 0 {
            return false
        }
        hashMap[v]--
    }
    return true
}
