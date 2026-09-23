func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
   sRune, tRune := []rune(s), []rune(t)

   sort.Slice(sRune, func(i, j int) bool {
    return sRune[i] < sRune[j]
   })
   sort.Slice(tRune, func(i, j int) bool {
    return tRune[i] < tRune[j]
   })

   for i:=0 ; i < len(sRune); i++ {
    if sRune[i] != tRune[i]{
        return false
    }
   }
   return true
}
