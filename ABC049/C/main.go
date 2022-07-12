package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)

	Slength := len(S)
	flag := "NO"

	wordList := [4]string{"dream", "dreamer", "erase", "eraser"}
	wordListCount := 0

	for Slength > 0 {

		searchRange := Slength - len(wordList[wordListCount])

		if searchRange < 0 {

			wordListCount++

			if wordListCount > 3 {
				break
			}

			continue
		}

		tmp := S[searchRange:]

		if tmp != wordList[wordListCount] {

			wordListCount++
			if wordListCount > 3 {
				break
			}

			continue
		}

		wordListCount = 0

		S = S[:searchRange]
		Slength = len(S)

	}

	if Slength == 0 {
		flag = "YES"
	}
	fmt.Println(flag)

}
