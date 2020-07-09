package palindrome

import (
	"strconv"
)

// Check checks if the given string or integer is a palindrome.
// The type of `v` must be [ string | int | int64 | uint64 ].
func Check(v interface{}) bool {
	switch v.(type) {
	case string:
		s := v.(string)
		return CheckStr(s)
	case int:
		s := strconv.Itoa(v.(int))
		return reverseStr(s) == s
	case int64:
		s := strconv.FormatInt(v.(int64), 10)
		return reverseStr(s) == s
	case uint64:
		s := strconv.FormatUint(v.(uint64), 10)
		return reverseStr(s) == s
	default:
		return false
	}
}

// CheckStr checks if the given string is a palindrome.
func CheckStr(s string) bool {
	return reverseStr(s) == s
}

// CheckInt1 checks if the given int is a palindrome by calculation.
func CheckInt1(n int) bool {
	return reverseInt(n) == n
}

// CheckInt2 checks if the given int is a palindrome by comparing it with the reversed string.
func CheckInt2(n int) bool {
	s := strconv.Itoa(n)
	return reverseStr(s) == s
}

func reverseStr(s string) string {
	r := []rune(s)
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}
	return string(r)
}

func pow(a, b int) int {
	c := 1
	for i := 0; i < b; i++ {
		c *= a
	}
	return c
}

func countDigit(n int) int {
	count := 1
	for {
		if n/pow(10, count) != 0 {
			count++
		} else {
			return count
		}
	}
}

func reverseInt(n int) int {
	digit := countDigit(n)
	x := 0
	for d := 0; d < digit; d++ {
		x += ((n / pow(10, d)) % 10) * (pow(10, digit-1-d))
	}
	return x
}
