package main

import (
	"fmt"
	"sort"
)

func main() {
	var totalMochi int
	fmt.Scan(&totalMochi)

	mochiList := make([]int, totalMochi)

	count := 1

	for i := range mochiList {
		fmt.Scan(&mochiList[i])
	}

	sort.Slice(mochiList, func(i, j int) bool {
		return mochiList[i] < mochiList[j]
	})

	for i := 0; i < len(mochiList)-1; i++ {
		if mochiList[i] != mochiList[i+1] {
			count++
		}
	}
	fmt.Println(count)
}
