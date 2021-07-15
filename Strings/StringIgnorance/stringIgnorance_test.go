package main

import "testing"

func TestIgnoreEvenoccurance (t *testing.T){
		input := "Geeks for geeks"
		expected := "Geks fore"
		result := ignoreEvenoccurance(input)
		if expected != result{
			t.Errorf("Not working as expected")
		}
}