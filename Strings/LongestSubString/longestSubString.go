package main

import "fmt"

func main() {
	input := "abbdefga"
	result := longestSubString(input)
	fmt.Println(result)
}

func longestSubString(input string) string {
	temp := ""
	result := ""
	tempMap := make(map[string]bool)
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			if _, ok := tempMap[string(input[j])]; !ok {
				tempMap[string(input[j])] = true
				temp = temp + string(input[j])

			} else {
				break
			}
		}

		if len(result) < len(temp) {
			result = temp
		}
		temp = ""
		for k := range tempMap {
			delete(tempMap, k)
		}

	}
	return result
}
