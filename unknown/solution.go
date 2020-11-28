package unknown

import (
	"math"
	"math/rand"
)

func PrimePalindrome(N int) int {
	var oneDigitPrimes = []int{2, 3, 5, 7, 11}

	for _, x := range oneDigitPrimes {
		if N <= x {
			return x
		}
	}

	for x := 1; ; x++ {
		pal := genPalindrome(x)
		if pal >= N && checkPrime(pal) {
			return pal
		}
	}

	return 0
}

func myPow(a, b, m int) int {
	result := 1
	n := a
	y := b

	for y > 0 {
		if y%2 == 1 {
			result = ((result % m) * (n % m)) % m
		}

		y = y >> 1
		n = ((n % m) * (n % m)) % m
	}

	return result % m
}

func checkFermat(n int) bool {
	a := 2 + rand.Intn(n-3)

	if myPow(a, n-1, n) != 1 {
		return false
	}

	return true
}

func checkPrime(a int) bool {
	n := a
	if n%2 == 0 {
		return false
	}

	if n == 3 || n == 5 || n == 7 || n == 11 {
		return true
	}

	for i := 0; i < 5; i++ {
		if !checkFermat(n) {
			return false
		}
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func genPalindrome(a int) int {
	result := a
	n := a / 10
	for n > 0 {
		rem := n % 10
		result = result*10 + rem
		n /= 10
	}

	return result
}
