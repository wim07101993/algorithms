package linear_search

func LinearSearch(list []int, toSearch int) (index int) {
	for i, v := range list {
		if v == toSearch {
			return i
		}
	}

	return -1
}
