package main

import (
	"github.com/golang/sarahayunanda/goday4/sequence_type/input"
	"github.com/golang/sarahayunanda/goday4/sequence_type/sequence"
	"github.com/golang/sarahayunanda/goday4/sequence_type/tools"
)

func main() {
	number1, number2 := input.Input()
	choose := sequence.ChooseSequence()
	tools.GetSequence(choose, number1, number2)
}
