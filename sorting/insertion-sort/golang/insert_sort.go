package insert_sort

func Sort(list []int) {
	for next := 1; next < len(list); next++ {
		previous := next - 1
		current := next
		for previous >= 0 {
			if list[current] >= list[previous] {
				break
			}
			list[current], list[previous] = list[previous], list[current]
			previous--
			current--
		}
	}
}
