package palindrome_test

import (
	"testing"

	"github.com/ebiiim/palindrome"
)

func TestCheck(t *testing.T) {
	cases := []struct {
		name  string
		input interface{}
		want  bool
	}{
		{"str_yes", "wow", true},
		{"str_no", "abc", false},
		{"int_yes", 1234543210123454321, true},
		{"int_no", 1234555550123454321, false},
		{"i64_yes", int64(1234543210123454321), true},
		{"i64_no", int64(1234555550123454321), false},
		{"u64_yes", uint64(12345432100123454321), true},
		{"u64_no", uint64(12345555500123454321), false},
		{"invalid", struct{}{}, false},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if got := palindrome.Check(c.input); got != c.want {
				t.Error("failed")
			}
		})
	}
}

func TestCheckString(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"yes", "wow", true},
		{"no", "abc", false},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if got := palindrome.CheckStr(c.input); got != c.want {
				t.Error("failed")
			}
		})
	}
}

func TestCheckInt(t *testing.T) {
	cases := []struct {
		name  string
		input int
		want  bool
	}{
		{"yes_1", 1122332211, true},
		{"yes_2", 555000555, true},
		{"yes_3", 1, true},
		{"no", 123456789, false},
		{"long", 1234543210123454321, true},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if got := palindrome.CheckInt1(c.input); got != c.want {
				t.Error("CheckInt1")
			}
			if got := palindrome.CheckInt2(c.input); got != c.want {
				t.Error("CheckInt2")
			}
		})
	}
}

func BenchmarkCheckInt1_long(b *testing.B) {
	x := 1234543210123454321
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		palindrome.CheckInt1(x)
	}
}

func BenchmarkCheckInt2_long(b *testing.B) {
	x := 1234543210123454321
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		palindrome.CheckInt2(x)
	}
}

func BenchmarkCheckInt1_short(b *testing.B) {
	x := 505
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		palindrome.CheckInt1(x)
	}
}

func BenchmarkCheckInt2_short(b *testing.B) {
	x := 505
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		palindrome.CheckInt2(x)
	}
}
