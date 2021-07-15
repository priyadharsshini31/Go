package main

import "testing"

func TestSaveIronMan(t *testing.T){
	input := "Ab?/Ba "
	expected := true
	result := saveIronMan(input)
	if expected != result {
		t.Errorf("Not working as expected")
	}

	
}