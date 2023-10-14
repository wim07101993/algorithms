package quick_sort

func Sort(list []int) {
	SortPartition(list, 0, len(list)-1)
}

func SortPartition(list []int, start, end int) {
	if end-start < 2 {
		return
	}
	middle := partition(list, start, end)
	SortPartition(list, start, middle)
	SortPartition(list, middle+1, end)
}

func partition(list []int, start, end int) (middle int) {
	pivot := list[(end-start)/2+start]
	print(pivot)
	i := start
	j := end
	for {
		for list[i] < pivot {
			i++
		}
		for list[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}

		list[i], list[j] = list[j], list[i]
	}
}
