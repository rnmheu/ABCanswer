package main

import (
	"fmt"
)

func main() {
	var totalMochi int
	fmt.Scan(&totalMochi)

	mochiList := make(map[int]int, totalMochi)

	for i := 0; i < totalMochi; i++ {
		var mochiSize int
		fmt.Scan(&mochiSize)
		mochiList[mochiSize] += 1
	}

	fmt.Println(len(mochiList))
}
