func isPalindrome(s string) bool {
	var chars []rune
	for _, char := range s {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			chars = append(chars, unicode.ToLower(char))
		}
	}

	for i := 0; i < len(chars)/2; i++ {
		if chars[i] != chars[len(chars)-1-i] {
			return false
		}
	}
	return true
}
