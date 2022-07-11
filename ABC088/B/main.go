package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	cardList := make([]int, N)

	for i := range cardList {
		fmt.Scanf("%d", &cardList[i])
	}
	sort.Slice(cardList, func(i, j int) bool {
		return cardList[j] < cardList[i]
	})

	var bob, alice int
	for i := range cardList {
		if i%2 != 0 {
			bob += cardList[i]
		} else {
			alice += cardList[i]
		}
	}
	fmt.Println(alice - bob)
}
