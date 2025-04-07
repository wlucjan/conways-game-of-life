package gameoflife

import "strings"

type Printer struct {
	alive string
	dead  string
}

func NewPrinter(alive, dead string) Printer {
	return Printer{
		alive: alive,
		dead:  dead,
	}
}

func (p Printer) Print(u Universe) string {
	var lines []string

	for _, row := range u {
		var line string
		for _, cell := range row {
			if cell {
				line += p.alive
			} else {
				line += p.dead
			}
		}
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
