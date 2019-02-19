package main

import (
	"fmt"
)
func main() {
    	var m = make(map[string]int)
	    m["priya"] = 24
	    m["Divya"] = 15
	    m["Alice"] = 12
	    var keys []string
	    for k := range m{
	      keys = append(keys,k)
	    }
	    for _,key:= range keys {
	        fmt.Println(key, m[key])
	    }
	}

