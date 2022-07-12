package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)

	wordList := [4]string{"dream", "dreamer", "erase", "eraser"}

	for j := 0; j < 4; j++ {

		if strings.HasSuffix(S, wordList[j]) {
			S = strings.TrimSuffix(S, wordList[j])
			j = -1
		}

		if len(S) == 0 {
			fmt.Println("YES")
			break
		}
	}

	if len(S) > 0 {
		fmt.Println("NO")
	}
}
