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

	fmt.Println(lines[0])
	fmt.Println(lines[1])
	fmt.Println(lines[2])	
	
	j := 0
	rlines := lines
	
	for i := (len(lines) - 1); i > 0 ; i-- {
	   j = j + 1	
	   rlines[j] = stringutil.Reverse(lines[i])	
	}


	fmt.Println(rlines[0])
	fmt.Println(rlines[1])	
	fmt.Println(rlines[2])	
	fmt.Println(rlines[3])	
	
}
