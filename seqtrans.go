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
	
	seq := ""
	
	for i := 1; i < (len(lines)-1); i++ {
	   seq = seq + lines[i]
    }
	
	fmt.Println(lines[0] + "\n" + seq)

    sr := stringutil.Reverse(seq)
	
	fmt.Println(lines[0] + " (reversed)\n" + sr)	
	
}

