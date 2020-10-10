package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	var inputs []int64
	for i := 0; i < t; i++ {
		var n int64
		fmt.Scan(&n)
		inputs = append(inputs, n)
	}
	max := max(inputs)
	fib := evenFib(max)
	for _, input := range inputs {
		sum := sumLimited(fib, input)
		fmt.Printf("%d\n", sum)
	}
}

func evenFib(limit int64) (even []int64) {
	var t1 int64 = 1
	var t2 int64 = 2
	even = append(even, 2)
	i := 1
	for {
		n := t1 + t2
		if n > limit {
			return
		}
		t1 = t2
		t2 = n
		if n%2 == 0 {
			even = append(even, n)
			i++
		}
	}
}

func max(arr []int64) (m int64) {
	for _, n := range arr {
		if n > m {
			m = n
		}
	}
	return
}

func sumLimited(arr []int64, limit int64) (s int64) {
	for _, n := range arr {
		if n > limit {
			return
		}
		s += n
	}
	return
}
