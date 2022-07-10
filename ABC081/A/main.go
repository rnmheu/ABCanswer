// 2022/07/10 SUN 23:35 -23:39

package main

import (
	"fmt"
	"strings"
)

func main() {
	var num string
	fmt.Scan(&num)
	total := strings.Count(num, "1")
	fmt.Println(total)
}
