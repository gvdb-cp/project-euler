package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		// sum divisible by 3
		arith3 := finiteArithmethicSum(n-1, 3)
		// sum divisible by 5
		arith5 := finiteArithmethicSum(n-1, 5)
		// sum of duplicates
		arith15 := finiteArithmethicSum(n-1, 15)
		// sum = 3 plus 5 minus duplicates
		sum := arith3 + arith5 - arith15
		fmt.Printf("%d\n", sum)
	}
}

// Sum of all numbers in range 1..limit divisable by step
func finiteArithmethicSum(limit int, step int) int {
	n := limit / step
	return step * ((n * (n + 1)) / 2)
}
