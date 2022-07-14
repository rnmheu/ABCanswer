package main

import (
	"fmt"
	"strconv"
)

func main() {
	var K int
	fmt.Scan(&K)

	var h, m, time string
	nowTime := 21*60 + K

	if nowTime < 24*60 {
		h = strconv.Itoa(nowTime / 60)
	} else {
		h = strconv.Itoa(24 - (nowTime / 60))
	}
	m = strconv.Itoa(nowTime % 60)

	if len(m) == 1 {
		m = "0" + m
	}
	if len(h) == 1 {
		h = "0" + h
	}

	time = h + ":" + m

	fmt.Println(time)
}
