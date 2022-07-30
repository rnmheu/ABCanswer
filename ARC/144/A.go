package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	const (
		initialBufSize = 10000
		maxBufSize     = 1000000
	)

	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, initialBufSize)
	sc.Buffer(buf, maxBufSize)

	var N int
	var S string

	if sc.Scan() {
		N, _ = strconv.Atoi(sc.Text())
	}

	if sc.Scan() {
		S = sc.Text()
	}

	start := string(S[0])
	end := string(S[N-1])

	if (start == "A" && end == "B") || S == "BA" || S == "AB" {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
