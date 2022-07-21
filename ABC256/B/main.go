package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	array := make([]int, N)

	for i := range array {
		fmt.Scan(&array[i])
	}

	if N > 3 {
		array = array[N-4:]
	}

	mass := make([]int, len(array))
	var ans int

	for i := range mass {
		for j := range array {
			mass[i] += array[j]
		}

		array = array[1:]

		if mass[i] > 3 {
			ans++
		}
	}

	fmt.Println(N - len(mass) + ans)
}
