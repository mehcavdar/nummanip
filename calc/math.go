package calc

import (
	"errors"

	"github.com/fatih/color"
)

func Add(numbers ...int) (int, error) {

	sum := 0

	if len(numbers) < 2 {
		errorMessage := color.RedString("2 den fazla sayı sagla")
		return sum, errors.New(errorMessage)
	} else {

		for _, num := range numbers {

			sum = sum + num
		}
		return sum, nil
	}
}
