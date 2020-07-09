package input

import "fmt"

// Input function
func Input() (firstNumb, lastNumb int){
	fmt.Print("First number\t: ")
	fmt.Scan(&firstNumb)
	fmt.Print("Last number\t: ")
	fmt.Scan(&lastNumb)
	return
}
