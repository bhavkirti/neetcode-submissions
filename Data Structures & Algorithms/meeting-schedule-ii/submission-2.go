/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	n := len(intervals)
	start, end := make([]int, n), make([]int, n)
	for i:= 0; i < n; i++{
		start[i] = intervals[i].start
		end[i] = intervals[i].end
	}

	sort.Ints(start)
	sort.Ints(end)
	res, count := 0, 0
	s, e := 0 ,0
	for s < n {
		if start[s] < end[e] {
			s++
			count++
		} else {
			e++
			count--
		}
		res = max(res, count)
	}
	return res
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
