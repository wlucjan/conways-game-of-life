package main

import (
	"fmt"

	"github.com/wlucjan/conways-game-of-life/gameoflife"
)

func main() {
	universe := gameoflife.NewUniverse(25, 25, nil)
	printer := gameoflife.NewPrinter("O ", ". ")

	fmt.Println(printer.Print(universe))
}
