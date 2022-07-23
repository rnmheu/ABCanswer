package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var N int

	if sc.Scan() {
		N, _ = strconv.Atoi(sc.Text())
	}

	array := make(map[string]int, N)

	for i := 0; i < N; i++ {
		var tmp string
		if sc.Scan() {
			tmp = sc.Text()
		}

		if _, ok := array[tmp]; ok {
			array[tmp]++
			num := strconv.Itoa(array[tmp])
			fmt.Println(tmp + "(" + num + ")")
		} else {
			array[tmp] = 0
			fmt.Println(tmp)
		}

	}
}
