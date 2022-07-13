package main

import "fmt"

func main() {
	var total int
	fmt.Scan(&total)

	numberList := make(map[int]int, total)

	for i := 0; i < total; i++ {

		var tmp int
		fmt.Scan(&tmp)

		if numberList[tmp] == 0 {
			numberList[tmp]++
			continue

		} else if numberList[tmp] == 1 {
			delete(numberList, tmp)
			continue
		}

		numberList[tmp]--
	}

	fmt.Println(len(numberList))
}
