func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    temp := []byte(t)
    for i := 0; i < len(s); i++ {
        charPresent := false
        for j := 0; j < len(temp) && !charPresent; j++ {
            if s[i] == temp[j] {
                temp[j] = 0
                charPresent = true
            }
        }
        if !charPresent {
            return false
        }
    }
    return true
}
