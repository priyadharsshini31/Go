package main

import "testing"

func TestTrappingRainWater (t *testing.T){
		input := []int{2,0,2}
		expectedResult := 2
		result := trappingRainWater(input)
		if result != expectedResult{
			t.Errorf("Not as expected")
		}
}