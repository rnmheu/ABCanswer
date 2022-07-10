package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if N == i*j {
				fmt.Printf("%s\n", "Yes")
				return
			}
		}
	}
	fmt.Printf("%s\n", "No")
}
