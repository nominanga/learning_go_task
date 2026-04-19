// Package learning_go_task provides basic math utilities.
package learning_go_task

import "golang.org/x/exp/constraints"

// Number is an interface that combines Integer and Float types.
type Number interface {
    constraints.Integer | constraints.Float
}

// Add takes two Number values and returns their sum.
// For more information on addition, see:
// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a T, b T) T {
    return a + b
}
