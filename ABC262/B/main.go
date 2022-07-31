package main

import "fmt"

func main() {
	var POINT, EDGE int
	fmt.Scan(&POINT, &EDGE)

	EDGE_LIST := make([][]bool, POINT)

	for i := range EDGE_LIST {
		uv := make([]bool, POINT)
		EDGE_LIST[i] = uv
	}

	for i := 0; i < EDGE; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		u--
		v--
		EDGE_LIST[u][v] = true
		EDGE_LIST[v][u] = true
	}

	var count int
	for i := 0; i < POINT; i++ {
		for j := i; j < POINT; j++ {
			for k := j; k < POINT; k++ {
				if EDGE_LIST[i][j] == EDGE_LIST[j][k] &&
					EDGE_LIST[j][k] == EDGE_LIST[k][i] &&
					EDGE_LIST[k][i] == true {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
