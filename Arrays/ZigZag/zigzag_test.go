package main
import (
	"testing"
	"reflect"
)

func TestZigZag(t *testing.T) {
	input := []int{4, 3, 7, 8, 6, 2, 1} 
	expectedResult := []int{3, 7, 4, 8, 2, 6, 1}
	Result := zigzag(input)
	if !reflect.DeepEqual(Result,expectedResult) {
		t.Errorf("They are not equal")
	}
}
