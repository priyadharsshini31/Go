package main

import "testing"

func TestFindLastIndex(t *testing.T){
		input := []int{1,2,3,4,5,6,7,8,9,0,1,2,1,1,1,3,4,5,1}
		expectedResult := 18
		result := findLastIndex(input, 1)
		if expectedResult != result {
			t.Errorf("Not as Expected")
		}
}