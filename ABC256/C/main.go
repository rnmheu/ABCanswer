package main

import "fmt"

func main() {
	var H [3]int
	var W [3]int

	for i := range H {
		fmt.Scan(&H[i])
	}
	for j := range W {
		fmt.Scan(&W[j])
	}

	ans := 0

	for a := 1; a <= 30; a++ {
		for b := 1; b <= 30; b++ {
			for d := 1; d <= 30; d++ {
				for e := 1; e <= 30; e++ {

					c := H[0] - a - b
					f := H[1] - e - d
					g := W[0] - a - d
					h := W[1] - b - e
					i := W[2] - c - f

					if (c > 0 && f > 0 && g > 0 && h > 0 && i > 0) && H[2] == g+h+i {
						ans++
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
