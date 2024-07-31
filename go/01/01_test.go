package main

import "testing"

func TestExercise01(t *testing.T) {
	want := "7,14,21,28"
	given := Exercise01(0, 30)
	if want != given {
		t.Errorf("Wanted %s but we have given %s", want, given)
	}
}
