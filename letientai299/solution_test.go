package letientai299

import (
	"fmt"
	"testing"
)

func Test_nextPotentialPalindromeNum(t *testing.T) {
	ds := digits(12822)
	for i := 0; i < 100; i++ {
		ds = nextPotentialPalindromeNum(ds)
		fmt.Println(ds)
	}
}

func Benchmark_isPrime(b *testing.B) {
	nums := []int{
		999853,
		999863,
		999883,
		999907,
		999917,
		999931,
		999953,
		999959,
		999961,
		999979,
		999983,
		12822,
		22822,
		100000,
		1000000,
		10000000,
	}
	tests := []struct {
		name string
		fn   func(n int) bool
	}{
		{name: "isPrime", fn: isPrime},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				for _, n := range nums {
					tt.fn(n)
				}
			}
		})
	}
}
