package sequence

import "fmt"

// ChooseSequence function
func ChooseSequence() (typeSequence string) {
	fmt.Println("Type of sequence")
	fmt.Println("1. Odd")
	fmt.Println("2. Even")
	fmt.Println("3. Prime")
	fmt.Print("Choose type of sequence\t: ")
	fmt.Scan(&typeSequence)
	return
}
