package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	result, step := 0, 0

	return func() int {
		tmp := result
		result = step
		step = result + tmp

		if result == 0 {
			step = 1
			return 0
		} else if result == 1 {
			return 1
		}

		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
