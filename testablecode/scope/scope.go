package scope

var multiplier = 5

// Multiply uses external state (the multiplier variable above)
// to calculate its response.  If any other function changes the multiplier
// value, this method becomes fragile and hard to test.
// Fix with a constant or by adding the multiplier to the func params.
func Multiply(operand int) int {
	return operand * multiplier
}

// IntCalculator is a simple calculator for integers
type IntCalculator struct {
	Result int
}

// Multiply stores the result of a * b into IntCalculator.Result
func (i *IntCalculator) Multiply(a, b int) {
	i.Result = a * b
}

// Clear resets the Result of the calculator
func (i *IntCalculator) Clear() {
	i.Result = 0
}
