package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	t_plus_1 := 0
	t_plus_2 := 1
	return func() int {
		t := t_plus_1
		t_plus_1 = t_plus_2
		t_plus_2 = t + t_plus_1
		return t
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
