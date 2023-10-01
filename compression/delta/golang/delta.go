package delta

func Encode(list []int) []int {
	encoded := make([]int, len(list))
	encoded[0] = list[0]
	for i := 1; i < len(list); i++ {
		encoded[i] = list[i] - list[i-1]
	}
	return encoded
}

func Decode(list []int) []int {
	decoded := make([]int, len(list))
	decoded[0] = list[0]
	for i := 1; i < len(list); i++ {
		decoded[i] = list[i] + decoded[i-1]
	}
	return decoded
}
