func isAnagram(s string, t string) bool {
    if (len(s) != len(t)) {
        return false
    }
    hashMap := make(map[byte]int)
    for i:=0; i<len(s); i++ {
        hashMap[s[i]]++
    }
    for v := range t {
        if(hashMap[t[v]] == 0) {
            return false
        }
        hashMap[t[v]]--
    }

    return true
}