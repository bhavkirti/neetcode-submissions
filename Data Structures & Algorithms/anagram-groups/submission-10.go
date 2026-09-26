import "slices"
func groupAnagrams(strs []string) [][]string {
    var result [][]string
    res := make(map[string][]string)

    for _, str := range strs {
        runes := []rune(str) 
        slices.Sort(runes)
        st := string(runes)
        res[st] = append(res[st], str)
    }

    for _, group := range res {
        result = append(result, group)
    }
    return result
}
