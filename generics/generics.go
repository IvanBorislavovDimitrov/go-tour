package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type StrangeMap[K comparable, V constraints.Ordered] map[K]V
type AnotherMap[K constraints.Ordered, V constraints.Complex] map[K]V

func AddWithConstraints[T constraints.Ordered](a, b T) T {
	return a + b
}

func Add[T int | float64](a, b T) T {
	return a + b
}

func main() {
	result := Add(1.2, 2.2)
	fmt.Println(result)
}
