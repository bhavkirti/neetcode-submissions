func topKFrequent(nums []int, k int) []int {
    hashMap := make(map[int]int)
    freq := make([][]int, len(nums)+1)

    for _, num := range nums {
        hashMap[num]++
    }

    for num, cnt := range hashMap {
        freq[cnt] = append(freq[cnt], num)
    }

    result := []int{}

    for i := len(freq) - 1; i > 0; i-- {
        for _, num := range freq[i] {
            result = append(result, num)
            if len(result) == k {
                return result
            }
        }
    }
    return result
}
