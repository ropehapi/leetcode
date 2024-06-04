package two_sum

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))

	fmt.Println(twoSumOpt([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumOpt([]int{3, 2, 4}, 6))
	fmt.Println(twoSumOpt([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	var values []int
	for i, number := range nums {
		for i2, number2 := range nums {
			if i != i2 && number+number2 == target {
				return append(values, i, i2)
			}
		}
	}
	return values
}

func twoSumOpt(nums []int, target int) []int {
	table := make(map[int]int)
	answer := make([]int, 2)

	for key, value := range nums {
		diff := target - value
		if indice, ok := table[diff]; ok {
			answer[0] = indice
			answer[1] = key
			break
		}

		table[value] = key
	}

	return answer
}
