package main

import "testing"

func TestFindFirstNonRepeating(t *testing.T){
	input := []int{1,2,3,4,5,1,2,3,4,5,6,7,8}
	expectedResult := 6
	result := findFirstNonRepeating(input)
	if expectedResult != result{
		t.Errorf("Not working as expected")
	}
}