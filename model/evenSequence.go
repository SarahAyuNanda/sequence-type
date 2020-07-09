package model

// EvenSequence function
func EvenSequence(first, last int) (store []int) {
	for i := first; i <= last; i++ {
		if i%2 == 0 {
			store = append(store, i)
		}
	}
	return store
}