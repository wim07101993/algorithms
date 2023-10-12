package selection_sort

func Sort(list []int) {
	for i := range list {
		jmin := i
		for j := i; j < len(list); j++ {
			if list[j] < list[jmin] {
				jmin = j
			}
		}
		list[i], list[jmin] = list[jmin], list[i]
	}
}
