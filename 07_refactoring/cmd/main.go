// * added test program
package main

import (
	"fmt"
	"go_course_thinknetika/07_refactoring/pkg/geom"
)

func main() {
	g := geom.Coords{}

	g.X1, g.Y1, g.X2, g.Y2 = 1, 1, 4, 5
	calc(g)

	g.X1, g.Y1, g.X2, g.Y2 = -1, 1, 4, 5
	calc(g)

	g.X1, g.Y1, g.X2, g.Y2 = 0, 0, 0, 0
	calc(g)
}

func calc(g geom.Coords) {
	// * the task condition was moved from the package to the executable program
	if g.X1 < 0 || g.X2 < 0 || g.Y1 < 0 || g.Y2 < 0 {
		fmt.Println("Coordinates can't be less than 0")
		return
	}

	fmt.Printf("Distance between (%v, %v) and (%v, %v) is %v.\n", g.X1, g.Y1, g.X2, g.Y2, g.Distance())
}
