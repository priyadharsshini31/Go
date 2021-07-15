package main

import ("testing")

func TestFindFirstRepeating(t *testing.T) {
		input := []int{1,2,3,4,5,6,7,3,4,5,6,9,0}
		expectedResult := 3
		result := findFirstRepeating(input)
		if expectedResult != result {
				t.Errorf("Not working as expected")
		}
}