func isValid(s string) bool {
    stack := []rune{}
	closeToOpenMap := map[rune]rune{')':'(', ']':'[', '}':'{'}

	for _, c := range s {
			if open, exists := closeToOpenMap[c];exists {
				if len(stack) > 0 {
					top := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					if open != top {
						return false
					}
				} else{
					return false
				}
			} else {
				stack = append(stack, c)
			}
	}
	
	return len(stack) == 0
}