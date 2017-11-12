package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var fizzRange = 50

	if len(os.Args) > 1 {

		r,err := strconv.Atoi(os.Args[1])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fizzRange = r
	}
	
	for  i:= 0; i < fizzRange ; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}