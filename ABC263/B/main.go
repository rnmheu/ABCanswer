package main

import "fmt"

//Piの数字がなにを示すのか理解出来ずにおわった

func main() {
	var TOTAL_PEOPLE int
	fmt.Scan(&TOTAL_PEOPLE)
	FAMILY_TREE_LIST := make([]int, TOTAL_PEOPLE-1)

	for i := range FAMILY_TREE_LIST {
		fmt.Scan(&FAMILY_TREE_LIST[i])
	}

	var count int
	current := TOTAL_PEOPLE

	for {
		if current == 1 {
			break
		}
		count++
		current = FAMILY_TREE_LIST[current-2]
	}

	fmt.Println(count)
}
