package main

import "testing"

func TestLargestAndSmallest (t *testing.T){
	input := []int{1,0,2,5,8,10,100}
	largest, smallest := largestAndSmallest(input)
	expectedLargest, expectedSmallest := 100,0
	if largest != expectedLargest || smallest != expectedSmallest {
		t.Errorf("Not the expected result")
	}
}