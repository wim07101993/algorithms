package binary_search

func BinarySearch(list []int, toSearch int) (index int) {
	min := 0
	max := len(list)

	for min <= max {
		// maths optimisation:
		//
		// maxi - mini          maxi - mini   2 * mini   maxi + mini
		// ----------- + mini = ----------- + -------- = -----------
		//      2                    2           2            2
		//
		// warning this can cause integer overflow
		index = (max + min) / 2

		val := list[index]

		if val == toSearch {
			return index
		}

		if val > toSearch {
			max = index - 1
			continue
		}

		if val < toSearch {
			min = index + 1
			continue
		}
	}

	return -1
}
