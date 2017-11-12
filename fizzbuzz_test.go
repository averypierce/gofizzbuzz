package main

import "testing"

func Test(t *testing.T) {
	t.Log("Testing fizzbuzz with range 6")

	slice := fizzBuzzSlice(6)      

	correct := [6]string{"fizzbuzz", "1", "2", "fizz", "4", "buzz"}

	for i:= 0; i < len(correct) ; i++ {
		if slice[i] != correct[i]{
			t.Error("derp darp")
		}
	}
}