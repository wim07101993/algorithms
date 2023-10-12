package merge_sort

func Sort[T interface{}](list []T, compare func(a, b T) int) []T {
	a, b := split(list)
	return sortSplit(a, b, compare)
}

func sortSplit[T interface{}](a, b []T, compare func(a, b T) int) []T {
	ret := make([]T, len(a)+len(b))
	if len(a) > 1 || len(b) > 1 {
		a1, a2 := split(a)
		a = sortSplit(a1, a2, compare)
		b1, b2 := split(b)
		b = sortSplit(b1, b2, compare)
	}
	var ia, ib int
	for i := 0; i < len(ret); i++ {
		if ia > len(a)-1 {
			ret[i] = b[ib]
			ib++
			continue
		}
		if ib > len(b)-1 {
			ret[i] = a[ia]
			ia++
			continue
		}

		if compare(a[ia], b[ib]) < 0 {
			ret[i] = a[ia]
			ia++
		} else {
			ret[i] = b[ib]
			ib++
		}
	}
	return ret
}

func split[T interface{}](list []T) ([]T, []T) {
	mid := len(list) / 2
	return list[:mid], list[mid:]
}
