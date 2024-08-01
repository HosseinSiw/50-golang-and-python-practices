package main

/*
A robot moves in a plane starting from the original point (0,0). The robot can move toward
UP, DOWN, LEFT and RIGHT with a given steps. The trace of robot movement is shown as the following:
UP 5
DOWN 3
LEFT 3
RIGHT 2
¡­
The numbers after the direction are steps. Please write a program to compute the distance from current position after a sequence of movement and original point. If the distance is a float, then just print the nearest integer.
Example:
If the following tuples are given as input to the program:
UP 5
DOWN 3
LEFT 3
RIGHT 2
Then, the output of the program should be:
2

Hints:
In case of input data being supplied to the question, it should be assumed to be a console input.

Solution
*/
import (
	"fmt"
	"math"
)

type Robut struct {
	X int
	Y int
	x_move int
	y_move int
}

func (r *Robut) move(x, y int) {
	r.x_move += x
	r.y_move += y
}
func (r *Robut) calculate() float64 {
	xdir_movement := math.Pow(float64(r.x_move - r.X), float64(2))
	ydir_movement := math.Pow(float64(r.y_move - r.Y), float64(2))
	movement := math.Sqrt(xdir_movement + ydir_movement)	
	return movement
}
func (r *Robut) report() {
	msg := fmt.Sprintf("The current cordinate is equal to: (%d, %d)", r.x_move, r.y_move)
	fmt.Println(msg)
}


func main() {
	r := Robut{
		X: 0,
		Y: 0,
	}
	r.move(1, 2)
	fmt.Println(r.calculate())
	r.report()
}


