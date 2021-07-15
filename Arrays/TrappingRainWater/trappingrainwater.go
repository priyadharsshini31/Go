package main

import "fmt"

func main() {
	input := []int{3, 0, 2, 0, 4}
	result := trappingRainWater(input)
	fmt.Println(result)
}

func trappingRainWater(input []int) int {

	result := 0
	for i := 1; i < len(input); i++ {
		left := 0
		right := 0

		for k := i; k >= 0; k-- {
			if input[k] > left {
				left = input[k]
			}
		}

		for j := i; j < len(input); j++ {
			if input[j] > right {
				right = input[j]
			}
		}
		if left < right {

			result = result + left - input[i]
		} else {
			result = result + right - input[i]
		}

	}
	return result

}
