package main

import "fmt"

func main() {
	var A, B, C, X int
	fmt.Scanf("%d", &A) // 500yen
	fmt.Scanf("%d", &B) // 100yen
	fmt.Scanf("%d", &C) // 50yen
	fmt.Scanf("%d", &X) // totalNum

	var count int
	for i := 0; i <= A; i++ {
		for j := 0; j <= B; j++ {
			for k := 0; k <= C; k++ {
				total := 500*i + 100*j + 50*k
				if total == X {
					count++
				}
			}
		}
	}
	fmt.Printf("%d", count)
}

// 早期Breakがよさそう
