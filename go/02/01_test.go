package main

import (
	"testing"
)

func TestEx5(t *testing.T) {
	wanted := 780
	result := Ex5(5)
	if wanted != result {
		t.Errorf("Expected: %d\nThe programm: %d", wanted, result)
	}	
}