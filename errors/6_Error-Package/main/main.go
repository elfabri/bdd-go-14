/*
Textio's software architects may have overcomplicated the requi-
rements from the last coding assignment... oops. All we needed
was a new generic error message that returns the string no dividing
by 0 when a user attempts to get us to perform the taboo.

Complete the divide function. Use the errors.New() function to return
an error when y == 0 that reads "no dividing by 0".
*/

package main

import (
	"errors"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("no dividing by 0")
	}
	return x / y, nil
}
