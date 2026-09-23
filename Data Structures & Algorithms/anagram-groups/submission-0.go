import "slices"
func groupAnagrams(strs []string) [][]string {
    var result [][]string
    res := make(map[string][]string)

    for i := 0; i < len(strs); i++ {
        runes := []rune(strs[i])
        slices.Sort(runes)
        str := string(runes)
        res[str] = append(res[str], strs[i])
    }

    for _, group := range res{
        result = append(result, group)
    }

    return result

}
