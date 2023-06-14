package main

import "testing"

func TestRunMain(t *testing.T) {
	main()
}

func TestSum(t *testing.T) {
	result := sum(2, 2)

	if result != 4 {
		t.Error("The result must be 5")
	}
}
