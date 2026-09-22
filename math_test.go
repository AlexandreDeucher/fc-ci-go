package main

import "testing"

func testSum(t *testing.T) {
	total := sum(20, 20)

	if total != 40 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 40)
}