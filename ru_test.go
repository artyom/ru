package ru

import "fmt"

func ExamplePluralize() {
	for _, i := range []int{0, 1, 2, 3, 4, 9, 11, 12, 21, 25, -1, -3, -11, -21, 132, 3795} {
		fmt.Printf("%d %s\n", i, Pluralize(i, "штук", "штука", "штуки"))
	}
	// Output:
	// 0 штук
	// 1 штука
	// 2 штуки
	// 3 штуки
	// 4 штуки
	// 9 штук
	// 11 штук
	// 12 штук
	// 21 штука
	// 25 штук
	// -1 штука
	// -3 штуки
	// -11 штук
	// -21 штука
	// 132 штуки
	// 3795 штук
}