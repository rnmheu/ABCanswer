package main

import "fmt"

func main() {
	var totalPaper, totalYen int
	fmt.Scan(&totalPaper, &totalYen)

	TenM := -1
	FiveM := -1
	M := -1

	for i := 0; i <= totalPaper; i++ {
		for j := 0; j+i <= totalPaper; j++ {

			k := totalPaper - (i + j)

			total := 10000*i + 5000*j + 1000*k

			if total == totalYen {

				TenM = i
				FiveM = j
				M = k

				break
			}
		}
	}
	fmt.Println("\n", TenM, FiveM, M)
}
