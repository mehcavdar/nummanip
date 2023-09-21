package calc

import "errors"

func Add(numbers ...int) (error, int) {
	sum := 0

	if len(numbers) < 2 {
		return errors.New("more tahn 2 parameter"), sum
	} else {
		for _, num := range numbers {

			sum += num
		}

	}
	return nil, sum
}
