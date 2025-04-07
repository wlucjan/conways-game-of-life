package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/wlucjan/conways-game-of-life/gameoflife"
)

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	seed, err := gameoflife.NewPattern(gameoflife.GliderPattern)
	if err != nil {
		panic(fmt.Errorf("failed to create pattern: %v", err))
	}

	universe := gameoflife.NewUniverse(25, 25, seed)
	printer := gameoflife.NewPrinter("O ", ". ")

	// TODO: extract this into a separate generations iterator to allow testing
	// TODO: generations iterator could accept max generations as a parameter, to avoid infinite loops
	for {
		ClearScreen()
		fmt.Println("Conway's Game of Life")
		fmt.Println(printer.Print(universe))

		nextUniverse := universe.Next()

		if universe.Equals(nextUniverse) {
			fmt.Println("Universe has stabilized. Exiting.")
			break
		}

		universe = nextUniverse
		time.Sleep(150 * time.Millisecond)
	}
}
