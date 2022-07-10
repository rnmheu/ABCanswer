// 提出日時 	2022-07-11 00:46:36

// 問題 	B - OddString
// ユーザ 	heutronica
// 言語 	Go (1.14.1)

// 得点 	200
// コード長 	193 Byte

// 結果
// 実行時間 	212 ms
// メモリ 	7920 KB

package main

import "fmt"

func main() {
	var S string
	fmt.Scanf("%s", &S)

	var word string
	for i := 0; i < len(S); i += 2 {
		word += S[i : i+1]
	}
	fmt.Printf("%s\n", word)
}
