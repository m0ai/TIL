package main

type (
	// Calculator is the interface we want to expose
	Calculator interface {
		Calculate(...int) int
	}

	// calculatorImpl contains the implementation we want to use
	calculatorImpl struct {
		options CalculatorOptions
	}

	// CalculatorOptions is the dataset we can use for the
	// implementation
	CalculatorOptions struct {
		BaseNumber int
	}
)

// NewCalculator is the Factory Method that returns our implementation
func NewCalculator(options CalculatorOptions) Calculator {
	return &calculatorImpl{
		options: options,
	}
}

// Calculate is the implementation of the interface method
func (c *calculatorImpl) Calculate(numbers ...int) (output int) {
	output = c.options.BaseNumber
	for i, number := range numbers {
		output += i * number
	}
	return
}

func main() {

}
