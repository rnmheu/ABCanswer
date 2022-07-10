// 2022/07/11 MON 00:25
package main

import (
	"fmt"
	"sort"
)

func main() {
	var length int
	fmt.Scanf("%d", &length)

	numList := make([]int, length)

	for i := range numList {
		fmt.Scanf("%d", &numList[i])
	}
	sort.Slice(numList, func(i, j int) bool {
		return (numList[i] < numList[j])
	})
	fmt.Printf("%d\n", numList[length-1]-numList[0])
}
