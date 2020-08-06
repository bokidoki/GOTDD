package arraytest

// Sum ...
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SizeOf...
func SizeOf(slice []int) int {
	return cap(slice)
}
