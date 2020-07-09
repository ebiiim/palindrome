# Palindrome Checker

[![GoDoc](https://godoc.org/github.com/ebiiim/palindrome?status.svg)](https://godoc.org/github.com/ebiiim/palindrome)
[![Build Status](https://travis-ci.org/ebiiim/palindrome.svg?branch=master)](https://travis-ci.org/ebiiim/palindrome)
[![Go Report Card](https://goreportcard.com/badge/github.com/ebiiim/palindrome)](https://goreportcard.com/report/github.com/ebiiim/palindrome)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org)

```go
palindrome.Check("nishioishin") // true
palindrome.Check("Was it a cat I saw?") // false
palindrome.Check(500050005) // true
palindrome.Check(1) // true
palindrome.Check(123) // false
palindrome.Check(struct{}) // false
```

---

### Note: String comparison is faster way in Go!

There are two main ways to check whether an integer is a palindrome.

Method 1: Calculation

```go
reverseInt(n) == n
```

Method 2: String comparison

```go
s := strconv.Itoa(n)
reverseStr(s) == s
```

After trying both, we found that Method 2 was faster.

```text
BenchmarkCheckInt1_long-12       2025841               592 ns/op               0 B/op          0 allocs/op
BenchmarkCheckInt2_long-12       5093236               228 ns/op              64 B/op          2 allocs/op
BenchmarkCheckInt1_short-12     15779865                82.3 ns/op             0 B/op          0 allocs/op
BenchmarkCheckInt2_short-12     14655747                83.8 ns/op            10 B/op          2 allocs/op
```
