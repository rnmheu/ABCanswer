package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var N int
	var result = map[string]int{"-": 0, "W": 0, "L": 0, "D": 0}

	if sc.Scan() {
		N, _ = strconv.Atoi(sc.Text())
	}

	for i := 0; i < N; i++ {

		var tmp string
		if sc.Scan() {
			tmp = sc.Text()
		}

		for j := range result {
			count := strings.Count(tmp, j)
			result[j] += count
		}
	}

	ans := "correct"

	if result["-"] != N {
		ans = "incorrect"
	}

	if result["W"] != result["L"] {
		ans = "incorrect"
	}

	if result["D"]%2 != 0 {
		ans = "incorrect"
	}

	fmt.Println(ans)
}
