package main

import (
	"fmt"
	"sort"
)

func sorting(arr1 []int, arr2 []int){
	var finalArray []int
	var tempArray []int
	m := make(map[int]int)
	for _, val := range arr1 {
		if _, ok := m[val]; ok {
			m[val]++
		} else {
			m[val] =1 
		}
	}
	
	for _, val := range arr2 {
		if value, ok := m[val]; ok {
			for i := 0 ; i < value; i++{
				finalArray = append(finalArray, val)
				delete(m,val)
			}
		}
	}

	for key , value:= range m{
		if value == 1 {
			tempArray = append(tempArray,key)
		}
	}
	sort.Ints(tempArray)
	
	finalArray = append(finalArray, tempArray...)
	fmt.Println(finalArray)
	
}
		
func main() {
	arr1 := []int{2, 1, 2, 5, 7, 1, 9, 3, 6, 8, 8}
	arr2 := []int{2, 1, 8, 3}
	sorting(arr1, arr2)
}
