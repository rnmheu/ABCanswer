package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, deg float64
	fmt.Scan(&x, &y, &deg)

	var ansX, ansY float64

	deg = deg * math.Pi / 180

	ansX = (x * math.Cos(deg)) - (y * math.Sin(deg))
	ansY = (y * math.Cos(deg)) + (x * math.Sin(deg))

	fmt.Println(ansX, ansY)
}
