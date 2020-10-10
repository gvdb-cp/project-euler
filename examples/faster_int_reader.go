package examples

import (
	"bufio"
	"os"
)

func fasterIntReader() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n := toInt(s.Bytes())
	println(n)
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}
