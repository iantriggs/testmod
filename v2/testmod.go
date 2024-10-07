// This is a test package for Chapter 10 of Learning Go.
// Not much more to say, really.
package testmod 

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// This function will add two numbers together and return the result.
// It follows the rules specified by https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
