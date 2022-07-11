package main

import (
	"fmt"
)

// a == rune(97)

func main() {
	var word string
	fmt.Scan(&word)

	runeList := make([]int, 26)

	for i := 0; i < len(word); i++ {
		tmp := rune(word[i]) - 97
		runeList[tmp] += 1
	}

	for i := range runeList {

		if i >= len(runeList)-1 && runeList[i] != 0 {
			fmt.Println("None")
			break
		}

		if runeList[i] == 0 {
			tmp := i + 97
			fmt.Println(string(rune(tmp)))
			break
		}
	}
}
