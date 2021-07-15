package main

import "testing"
import "reflect"
func TestRotate(t *testing.T) {
	input := []int{1,2,3,4,5,6,7,8}	
	Result := rotate(input, 2)
	expectedResult := []int{3,4,5,6,7,8,1,2}
	if reflect.DeepEqual(Result, expectedResult) == false {
		t.Errorf("Rotation is not working properly")
	}
}
