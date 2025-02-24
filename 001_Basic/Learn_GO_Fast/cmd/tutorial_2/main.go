package main

import (
	"fmt"
	"strings"
)

func main() {
	var strSlice = []string{"a", "b", "c", "d", "e"}
	var strBulder strings.Builder

	for i := range strSlice {
		strBulder.WriteString(strSlice[i])
	}
	var catStr = strBulder.String()
	fmt.Printf("\n%v", catStr)

}
