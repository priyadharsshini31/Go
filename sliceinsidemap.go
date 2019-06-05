package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"priya_dharsshini": []string{"playing", "chess"},
		"Divya_darsini": []string{"cheating", "idiot"},
	}
	for _, v := range m {
		for i, v := range v {
			fmt.Println(i, v)
		}
	}

}
