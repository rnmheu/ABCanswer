package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	max := math.Sqrt(float64(N))

	a := N
	b := 1

	for i := 1; i < int(max); i++ {
		if max-math.Floor(max) == 0 {
			a = int(max)
			b = int(max)
			break
		}

		tmp := N / i

		if N == tmp*i {
			a = i
			b = tmp
		}
	}

	lengthA := len(strconv.Itoa(a))
	lengthB := len(strconv.Itoa(b))

	length := math.Max(float64(lengthA), float64(lengthB))

	fmt.Println(length)

}
