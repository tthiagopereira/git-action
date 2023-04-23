package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(10, 12)

	if total != 22 {
		t.Errorf("Error")
	}
}
