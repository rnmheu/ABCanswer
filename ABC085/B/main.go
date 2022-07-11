package main

import (
	"fmt"
)

func main() {
	var totalMochi int
	fmt.Scan(&totalMochi)

	mochiList := make(map[int]int, totalMochi)

	var mochiSize int

	for i := 0; i < totalMochi; i++ {
		fmt.Scan(&mochiSize)
		mochiList[mochiSize] += 1
	}

	delete(mochiList, 0)

	fmt.Println(len(mochiList))
}
