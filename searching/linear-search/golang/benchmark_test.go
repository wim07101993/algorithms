package linear_search

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

var numberCounts = []int{
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
}

func BenchmarkLinearSearch(b *testing.B) {
	for _, c := range numberCounts {
		list := generateNumbers(c)
		runBenchmark(list, 0, b)
		runBenchmark(list, len(list)/2, b)
		runBenchmark(list, len(list)/3*2, b)
	}
}

func runBenchmark(list []int, position int, b *testing.B) {
	b.Run(fmt.Sprintf("number_count_%d_solution_position_%d", len(list), position), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LinearSearch(list, list[position])
		}
	})
}

func generateNumbers(n int) []int {
	is := make([]int, n)
	for i := 0; i < len(is); i++ {
		is[i] = rand.Int()
	}
	sort.Ints(is)
	return is
}
