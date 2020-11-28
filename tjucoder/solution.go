package tjucoder

import (
	"math"
	"strconv"
)

func PrimePalindrome(N int) int {
	if N <= 2 {
		return 2
	}
	number := N
	for {
		adjustNumber(&number)
		if isPalindrome(number) && isPrime(number) {
			return number
		}
		number += 2
	}
	return -1
}

func adjustNumber(number *int) {
	/*
	   Even digits palindrome property:  Always divisible by 11

	   1st digit(odd th) == last digit(even th)
	   2nd digit(even th) == last 2nd digit(odd th)

	   11 Divisibility Test : <https://www.math.hmc.edu/funfacts>.

	   Take the alternating sum of the digits in the number, read from left to right.
	   If that is divisible by  11, so is the original number

	   22 --> 2-2 = 0
	   1221 --> 1-2+2-1 = 0
	   134431 --> 1-3+4-4+3-1=0

	   Now since 11 is a prime number itself ; it will qualify but others will not
	*/

	// 11 divisibility test and adjustment
	inputNumber := *number
	digits := 0
	for inputNumber > 0 {
		digits++
		inputNumber /= 10
	}
	if digits%2 == 0 && *number > 11 {
		newNumber := 1
		for i := 0; i < digits; i++ {
			newNumber *= 10
		}
		*number = newNumber + 1
	}
	// 2 divisibility test and adjustment
	if *number%2 == 0 {
		*number++
	}

	//Note : Can probably add tests and adjustments for other prime numbers but need draw a line somewhere
}

func isPrime(number int) bool {
	root := int(math.Sqrt(float64(number)))
	for i := 2; i <= root; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func isPalindrome(number int) bool {
	s := strconv.Itoa(number)
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
