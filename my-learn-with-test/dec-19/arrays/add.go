package arrays

func sum(arr []int) int {
	s := 0
	for _, el := range arr {
		s += el
	}
	return s
}

func sumAll(arrs [][]int) []int {
	var results []int
	for _, arr := range arrs {
		results = append(results, sum(arr))
	}
	return results
}

func sumAllTails(arrs ...[]int) []int {
	var results []int
	for _, arr := range arrs {
		if len(arr) == 0 {
			results = append(results, 0)
		} else {
			results = append(results, sum(arr[1:]))
		}
	}
	return results
}
