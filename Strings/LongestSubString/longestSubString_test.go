package main

import "testing"

func TestLongestSubString(t *testing.T){
		input := "ababcdgfbsfg"
		expected := "abcdgf"
		result := longestSubString(input)
		if expected != result{
				t.Errorf("Not as expected")
		}
}