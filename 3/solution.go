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
	for _, input := range inputs {
		out := highestPrimeFactor(input)
		fmt.Printf("%d\n", out)
	}
}

func highestPrimeFactor(n int64) int64 {
	var pf []int64
	var d int64 = 2
	for n > 1 {
		for n%d == 0 {
			pf = append(pf, d)
			n /= d
		}
		d++
		if d*d > n {
			if n > 1 {
				pf = append(pf, n)
			}
			break
		}
	}
	return pf[len(pf)-1]
}
