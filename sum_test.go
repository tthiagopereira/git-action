package main

import "testing"

func TestSoma(t *testing.T) {
	total := Sum(10, 12)

	if total != 22 {
		t.Errorf("Error")
	}
}
