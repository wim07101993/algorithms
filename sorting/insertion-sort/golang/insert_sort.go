package insert_sort

func Sort[T interface{}](list []T, comparer func(a, b T) int) []T {
	var ret []T
	for i, a := range list {
		added := false
		for j, b := range ret {
			cmp := comparer(a, b)
			if cmp < 0 {
				ret = append(ret[:j+1], ret[j:]...)
				ret[j] = list[i]
				added = true
				break
			}
		}
		if !added {
			ret = append(ret, list[i])
		}
	}
	return ret
}
