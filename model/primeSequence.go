package model

// PrimeSequence function
func PrimeSequence(first, last int) (store []int) {
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