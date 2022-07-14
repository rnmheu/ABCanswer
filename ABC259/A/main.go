package main

import "fmt"

func main() {
	var N, M, X, T, D int
	fmt.Scan(&N, &M, &X, &T, &D)

	base := T - D*X

	var ans int
	if M <= X {
		ans = base + M*D
	} else if X < M {
		ans = base + D*X
	}

	fmt.Println(ans)

}
