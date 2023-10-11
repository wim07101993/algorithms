package selection_sort

func Sort[T interface{}](list []T, comparer func(a, b T) int) []T {
	ret := make([]T, len(list))
	low := list[0]
	for i := 0; i < len(list); i++ {
		for _, t := range list {
			if i > 0 && comparer(ret[i-1], t) > 0 {
				continue
			}
			if comparer(low, t) > 0 {
				low = t
			}
		}
		ret[i] = low
	}
	return ret
}
