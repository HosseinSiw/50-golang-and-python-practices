package main
import (
	"testing"
	// "fmt"
)

func TestRobut(t *testing.T) {
	r := Robut{
		X: 5, Y: 5,
	}
	r.move(10, 20)
	movement := int(r.calculate())
	wanted := 15
	if movement != wanted{
		t.Errorf("Wanted: %d, but given %d", wanted, movement)
	}
}