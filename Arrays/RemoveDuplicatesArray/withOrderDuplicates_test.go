package main

import (
	"testing"
	"reflect"
)

func TestRemoveDuplicates (t *testing.T) {
		input := []int{1,1,2,2,3,3,4,4,5,5,6,6,7,7,8,8,9,9}
		expectedResult := []int{1,2,3,4,5,6,7,8,9}
		result := removeDuplicates(input)
		if !reflect.DeepEqual(expectedResult,result){
			t.Errorf("Result not as expected")
		}
}

