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
	
	slice := fizzBuzzSlice(fizzRange)
	fmt.Println(slice)
}

func fizzBuzzSlice(fizzRange int) []string {

	var slice = make([]string,fizzRange)

	for  i:= 0 ; i < fizzRange ; i++ {

		if i % 3 == 0 && i % 5 == 0 {
			slice[i] = "fizzbuzz"
		} else if i % 3 == 0 {
			slice[i] = "fizz"
		} else if i % 5 == 0 {
			slice[i] = "buzz"
		} else {
			slice[i] = strconv.Itoa(i)
		}
	}
	
	return slice
}