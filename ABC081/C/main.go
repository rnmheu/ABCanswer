package main

import (
	"fmt"
	"sort"
)

func main() {

	// init

	var totalBall, totalBallNum int
	fmt.Scan(&totalBall, &totalBallNum)

	ballCountBucket := make(map[int]int, 200000)

	for i := 0; i < totalBall; i++ {
		var tmpNum int
		fmt.Scanf("%d", &tmpNum)

		ballCountBucket[tmpNum]++
	}

	// sort

	totalBallByNum := len(ballCountBucket)

	ballList := make([]int, 0, totalBallByNum)

	for _, value := range ballCountBucket {
		ballList = append(ballList, value)
	}

	sort.Slice(ballList, func(i, j int) bool {
		return ballList[i] < ballList[j]
	})

	// count rewrite

	count := 0

	for i := 0; i < totalBallByNum; i++ {

		if len(ballList)-i <= totalBallNum {
			fmt.Println(count)
			break
		}
		count += ballList[i]
	}
}
