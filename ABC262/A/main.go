package main

import "fmt"

func main() {
	var YEAR int
	fmt.Scan(&YEAR)
	for {
		if YEAR%4 == 2 {
			fmt.Println(YEAR)
			break
		}
		YEAR++
	}
}
