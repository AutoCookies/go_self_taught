// file: benchmark_large_test.go
package branchless

import (
	"math/rand"
	"testing"
)

const N = 10_000_000

var dataA = make([]int, N)
var dataB = make([]int, N)

func init() {
	for i := 0; i < N; i++ {
		dataA[i] = rand.Intn(1_000_000)
		dataB[i] = rand.Intn(1_000_000)
	}
}

func MaxIfSlice(a, b []int) int {
	max := 0
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			max += a[i]
		} else {
			max += b[i]
		}
	}
	return max
}

func MaxBranchlessSlice(a, b []int) int {
	max := 0
	for i := 0; i < len(a); i++ {
		max += a[i]*(boolToInt(a[i] >= b[i])) + b[i]*(boolToInt(b[i] > a[i]))
	}
	return max
}

func BenchmarkMaxIfSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MaxIfSlice(dataA, dataB)
	}
}

func BenchmarkMaxBranchlessSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MaxBranchlessSlice(dataA, dataB)
	}
}
