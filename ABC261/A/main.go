package main

import "fmt"

func main() {
	var array [4]int

	//しまった！！！！  0<= x <= 100の長さは101だ！！！
	var count [101]int
	for i := range array {
		fmt.Scan(&array[i])
	}

	ans := 0

	for i := range count {
		if i >= array[0] && i <= array[1] {
			count[i]++
		}
		if i >= array[2] && i <= array[3] {
			count[i]++
		}
		if count[i] == 2 {
			ans++
		}
	}

	if ans > 0 {
		ans--
	}

	fmt.Println(ans)
}
