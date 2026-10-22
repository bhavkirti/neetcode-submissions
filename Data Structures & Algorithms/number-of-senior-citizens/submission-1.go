func countSeniors(details []string) int {
    res := 0
	for _, d := range details {
		ten := int(d[11]-'0')
		one := int(d[12]-'0')
		age := one + 10*ten
		if age > 60 {
			res++
		}
	}
	return res
}