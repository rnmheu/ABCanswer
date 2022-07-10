package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scanf("%s", &S)

	word := make([]string, (len(S)+1)/2)
	for i := 0; i < len(S); i += 2 {
		word[i/2] = S[i : i+1]
	}
	answer := strings.Join(word, "")
	fmt.Printf("%s\n", answer)
}
