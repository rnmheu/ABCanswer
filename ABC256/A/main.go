package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var N float64
	fmt.Scan(&N)

	ans := strconv.FormatFloat(math.Exp2(N), 'f', 0, 64)

	fmt.Println(ans)

}
