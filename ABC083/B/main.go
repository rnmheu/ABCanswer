package main

import "fmt"

func main() {
	var N, min, max int
	fmt.Scan(&N, &min, &max)

	var ans, tmpTotal int

	for i := 1; i <= N; i++ {
		tmpTotal = 0

		for tmp := i; tmp > 0; tmp /= 10 {
			tmpTotal += tmp % 10
		}
		if tmpTotal >= min && tmpTotal <= max {
			ans += i
		}
	}
	fmt.Printf("%d\n", ans)
}
