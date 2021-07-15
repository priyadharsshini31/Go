package main

import "testing"

func TestRotation(t *testing.T){
		input := "GeeksforGeeks"
		input2 := "forGeeksGeeks"
		expected := true
		result := rotation(input, input2)
		if result != expected{
			t.Errorf("Not working as expected")
		}
}