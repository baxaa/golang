package main

import (
	"fmt"
	"math"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	var first string = strs[0]
	last := strs[len(strs)-1]
	var ans string = ""
	var min1 float64 = float64(len(first))
	var min2 float64 = float64(len(last))
	var size int = int(math.Min(min1, min2))
	for i := 0; i < size; i++ {
		if first[i] != last[i] {
			return ans
		}
		ans = ans + string(first[i])
	}
	return ans
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
