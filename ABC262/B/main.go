package main

import "fmt"

func main() {
	var POINT, EDGE int
	fmt.Scan(&POINT, &EDGE)

	EDGE_LIST := make([][2]int, EDGE)

	for i := range EDGE_LIST {
		var tmp [2]int
		fmt.Scan(&tmp[0], &tmp[1])
		EDGE_LIST[i] = tmp
	}

	var count int
	for i := range EDGE_LIST {
		for j := range EDGE_LIST {
			if i == j {
				continue
			}
			countNum := make(map[int]int)
			countNum[EDGE_LIST[i][0]]++
			countNum[EDGE_LIST[i][1]]++
			countNum[EDGE_LIST[j][0]]++
			countNum[EDGE_LIST[j][1]]++

			if len(countNum) != 3 {
				continue
			}

			for k := range EDGE_LIST {
				if j == k || k == i {
					continue
				}

				_, ok1 := countNum[EDGE_LIST[k][0]]
				_, ok2 := countNum[EDGE_LIST[k][1]]

				if ok1 == true && ok2 == true {
					count++
				}
			}
		}
	}
	fmt.Println(count / 6)
}
