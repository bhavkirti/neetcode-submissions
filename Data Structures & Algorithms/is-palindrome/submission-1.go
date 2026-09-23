
func isPalindrome(s string) bool {
   
    var clean []rune
    for _, char := range s {
        if unicode.IsLetter(char) || unicode.IsDigit(char) {
            clean = append(clean, unicode.ToLower(char))
        }
    }
    n := len(clean)

    for i := 0; i < n/2; i++ {
        if clean[i] != clean[n-1-i]{
            return false
        }
    }
    return true
}
