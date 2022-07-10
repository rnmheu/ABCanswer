// 2022/07/10 SUN 23:43 - 24:00

package main

import "fmt"

func main() {
	var totalNum int
	fmt.Scan(&totalNum)

	intList := make([]int, totalNum)
	for i := range intList {
		fmt.Scan(&intList[i])
	}

	var count int
	flag := false

	for flag != true {
		for i := range intList {
			if intList[i]%2 != 0 {
				flag = true
				count--
				break
			}
			intList[i] /= 2
		}
		count++
	}
	fmt.Println(count)
}
