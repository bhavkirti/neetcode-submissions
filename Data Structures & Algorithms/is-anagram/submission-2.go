func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    charMap := [26]int{}

    for i:=0; i < len(s); i++ {
        charMap[s[i]-'a']++
        charMap[t[i]-'a']--
    }

    for _, val := range charMap {
        if val != 0 {
            return false
        }
    }
    return true
}

