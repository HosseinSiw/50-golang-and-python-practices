package main

import "testing"

func TestEx4(t *testing.T) {
	msg := "Hossein Siw Github"
	wanted := map[string]int{
		"Lower":        13,
		"Upper":        3,
		"White spaces": 2,
	}
	got := Ex4(msg)
	if got["Lower"] != wanted["Lower"] || got["Upper"] != wanted["Upper"] || wanted["White spaces"] != got["White spaces"] {
		t.Errorf("An error is occured. ")
	}

}
