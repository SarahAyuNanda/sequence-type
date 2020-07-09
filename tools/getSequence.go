package tools

import (
	"fmt"
	"strings"

	"github.com/golang/sarahayunanda/goday4/sequence_type/model"
)

// GetSequence function
func GetSequence(typeSequence string, firstNumb, lastNumb int) {
	if firstNumb <= 0 {
		fmt.Println("The lower bound must be positive number!")
	} else {
		switch typeSequence {
		case strings.ToLower("odd"):
			{
				odd := model.OddSequence(firstNumb, lastNumb)
				fmt.Println("Odd sequence is", odd)
			}
		case strings.ToLower("even"):
			{
				even := model.EvenSequence(firstNumb, lastNumb)
				fmt.Println("Even sequence is", even)
			}
		case strings.ToLower("prime"):
			{
				prime := model.PrimeSequence(firstNumb, lastNumb)
				fmt.Println("Prime sequence is", prime)
			}
		default:
			fmt.Println("No other sequences!")
		}
	}

}
