package main

import "fmt"

func main() {

	var tefuda [5]int
	count := map[int]int{}
	for i := range tefuda {
		fmt.Scan(&tefuda[i])
	}

	var flag1, flag2 bool
	for i := range tefuda {
		count[tefuda[i]]++
	}
	for _, v := range count {
		if v == 2 {
			flag1 = true
		}
		if v == 3 {
			flag2 = true
		}
	}
	if flag2 == true && flag1 == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
