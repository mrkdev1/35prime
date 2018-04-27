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

	src := []rune(sr)
	
    // Loop over range of indexes in the slice.
    for i := range src {
	    if sr[i]==65 {
            src[i] = 84
	    }
	    if sr[i]==84 {
            src[i] = 65
	    }
	    if sr[i]==71 {
            src[i] = 67
	    }
	    if sr[i]==67 {
            src[i] = 71
	    }		
    }
	
	sc := []rune(s)
	
    // Loop over range of indexes in the slice.
    for i := range src {
	    if s[i]==65 {
            sc[i] = 84
	    }
	    if s[i]==84 {
            sc[i] = 65
	    }
	    if s[i]==71 {
            sc[i] = 67
	    }
	    if s[i]==67 {
            sc[i] = 71
	    }		
    }
	fmt.Println(lines[0] + " (complement)\n" + string(sc))		
	fmt.Println(lines[0] + " (complement reversed)\n" + string(src))		
}

