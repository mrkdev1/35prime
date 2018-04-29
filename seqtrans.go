package main

import (
	"os"
	"fmt"
	"strings"
	"log"
    "io/ioutil"
    "github.com/mrkdev1/stringutil"

)

func main() {
    file := string(os.Args[1])
	
    content, err := ioutil.ReadFile(file)
    if err != nil {
        log.Fatal(err)
    }
    lines := strings.Split(string(content), "\n")
	
	s := ""
	
	for i := 1; i < (len(lines)-1); i++ {
	   s = s + lines[i]
    }
	
	fmt.Println(lines[0] + "\n" + s)

    sr := stringutil.Reverse(s)
	
	fmt.Println(lines[0] + " (reversed)\n" + sr)		
	fmt.Println(lines[0] + " (complement)\n" + comp(s))		
	fmt.Println(lines[0] + " (complement reversed)\n" + comp(sr))		
}

func comp(s string) string {
	r := []rune(s)
	
    for i := range r {
	    t := r[i]
		switch {
		case t==65:
		    r[i] = 84
		case t==84:
            r[i] = 65
        case t==71:
            r[i] = 67
        case t==67:
            r[i] = 71
        default:
            r[i] = 78		
	    }		
    }
	return string(r)
}