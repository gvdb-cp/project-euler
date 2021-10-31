package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	var inputs []int
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		inputs = append(inputs, n)
	}
	palindromes := findPalindromeProducts()
	for _, input := range inputs {
		out := largestPalindromeProductLessThan(palindromes, input)
		fmt.Printf("%d\n", out)
	}
}

func findPalindromeProducts() (p []int) {
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			x := i * j
			if isPalindrome(x) {
				p = append(p, x)
			}
		}
	}
	return
}

func largestPalindromeProductLessThan(palindromes []int, n int) int {
	if isPalindrome(n) {
		n--
	}
	largest := 0
	for _, p := range palindromes {
		if p > n {
			continue
		}
		if p > largest {
			largest = p
		}
	}
	return largest
}

func isPalindrome(n int) bool {
	s := fmt.Sprint(n)
	if len(s) < 6 {
		return false
	}
	return s[:3] == reverse(s[3:])
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
