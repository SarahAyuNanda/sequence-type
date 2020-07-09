package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstNumb, lastNumb int
	fmt.Print("First number\t: ")
	fmt.Scan(&firstNumb)
	fmt.Print("Last number\t: ")
	fmt.Scan(&lastNumb)

	var typeSequence string
	fmt.Println("Type of sequence")
	fmt.Println("1. Odd")
	fmt.Println("2. Even")
	fmt.Println("3. Prime")
	fmt.Print("Choose type of sequence\t: ")
	fmt.Scan(&typeSequence)

	if firstNumb <= 0 {
		fmt.Println("The lower bound must be positive number!")
	} else {
		switch typeSequence {
		case strings.ToLower("odd"):
			{
				odd := oddSequence(firstNumb, lastNumb)
				fmt.Println("Odd sequence is", odd)
			}
			// fallthrough
		case strings.ToLower("even"):
			{
				even := evenSequence(firstNumb, lastNumb)
				fmt.Println("Even sequence is", even)
			}
			// fallthrough
		case strings.ToLower("prime"):
			{
				prime := primeSequence(firstNumb, lastNumb)
				fmt.Println("Prime sequence is", prime)
			}
		default:
			fmt.Println("No other sequences!")
		}
	}
}

func oddSequence(first, last int) (store []int) {
	for i := first; i <= last; i++ {
		if i%2 != 0 {
			store = append(store, i)
		}
	}
	return store
}

func evenSequence(first, last int) (store []int) {
	for i := first; i <= last; i++ {
		if i%2 == 0 {
			store = append(store, i)
		}
	}
	return store
}

func primeSequence(first, last int) (store []int) {
	for numb := first; numb <= last; numb++ {
		factor := 0
		for check := 1; check <= numb; check++ {
			if numb%check == 0 {
				factor++
			}
		}
		if factor == 2 {
			store = append(store, numb)
		}
	}
	return store
}
