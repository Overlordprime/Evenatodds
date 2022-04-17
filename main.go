package main

import "fmt"

type numb []int

func main() {
	ints := createInts(10)
	ints.evenAtOdd()
}

func createInts(max int) numb {
	ints := numb{}
	seq := 0

	for max >= seq {
		ints = append(ints, seq)
		seq += 1
	}
	return ints
}

func (n numb) evenAtOdd() {
	for val := range n {
		if val%2 == 0 {
			fmt.Printf("%v is even\n", val)
		} else {
			fmt.Printf("%v is odd\n", val)
		}
	}
}

//go mod
