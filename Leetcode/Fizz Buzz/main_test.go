package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	a := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	b := fizzBuzz(15)
	for i := range b {
		if a[i] != b[i] {
			t.Errorf("Expected '%s', but got '%s'", a[i], b[i])
		}
	}
}
