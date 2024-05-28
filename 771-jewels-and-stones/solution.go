package _71_jewels_and_stones

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))
}

func numJewelsInStones(jewels string, stones string) int {
	total := 0
	for _, jewel := range jewels {
		total += strings.Count(stones, string(jewel))
	}
	return total
}
