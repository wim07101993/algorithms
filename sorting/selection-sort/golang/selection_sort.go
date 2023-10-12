package selection_sort

func Sort[T interface{}](list []T, less func(a, b T) bool) {
	for i := 0; i < len(list); i++ {
		jmin := i
		for j := i; j < len(list); j++ {
			if less(list[j], list[jmin]) {
				jmin = j
			}
		}
		list[i], list[jmin] = list[jmin], list[i]
	}
}
