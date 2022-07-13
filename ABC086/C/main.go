package main

import (
	"fmt"
	"math"
)

func main() {
	var totalTravel int
	fmt.Scan(&totalTravel)

	travelPointList := make([][3]int, totalTravel+1)
	travelPointList[0] = [3]int{0, 0, 0}

	for i := 1; i < totalTravel+1; i++ {
		var t, x, y int
		fmt.Scan(&t, &x, &y)
		travelPointList[i] = [3]int{t, x, y}
	}

	flag := "Yes"

	for i := 0; i < totalTravel; i++ {
		time := travelPointList[i+1][0] - travelPointList[i][0]
		dist := dist(travelPointList[i][1], travelPointList[i][2], travelPointList[i+1][1], travelPointList[i+1][2])

		if dist == 0 || time < dist {
			flag = "No"
			break
		}
		if time%2 != dist%2 {
			flag = "No"
			break
		}
	}
	fmt.Println(flag)
}

func dist(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x2)-float64(x1)) + math.Abs(float64(y2)-float64(y1)))
}
