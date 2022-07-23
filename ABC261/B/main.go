package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var result = map[string]int{"-": 0, "W": 0, "L": 0, "D": 0}

	fmt.Scan(&N)

	for i := 0; i < N; i++ {

		var tmp string
		fmt.Scan(&tmp)
		splitTmp := strings.Split(tmp, "")

		for j := range splitTmp {
			for k := range result {
				if splitTmp[j] == k {
					result[k]++
					break
				}
			}
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
