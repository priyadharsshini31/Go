package main

import ( 
	"testing"
	"reflect"
)

func TestFindLongestValidParanthesis(t *testing.T){
	input := "(()))"
	Result := findLongestValidParanthesis(input)
	expectedResult := 4
	if reflect.DeepEqual(Result, expectedResult) == false {
		t.Errorf("Not working")
	}
}