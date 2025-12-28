package dysons_array_solutions

func ReverseInPlace(list []int, start int, end int) []int {
	if len(list) == 0 {
		return list
	}
	for {
		if end < start {
			break
		}
		list[end], list[start] = list[start], list[end]
		end--
		start++
	}
	return list
}
