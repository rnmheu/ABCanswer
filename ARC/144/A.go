package main

import "fmt"

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)

	start := string(S[0])
	end := string(S[N-1])

	if (start == "A" && end == "B") || N == 2 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
