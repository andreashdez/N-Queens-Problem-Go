package main

import (
	"fmt"
	"strings"
)

func DrawBoard(positions []int, conflicts []int) {
	size := len(positions)
	drawRowTop(size)
	for y := 0; y < size; y++ {
		fmt.Print("║ ")
		for x := 0; x < size; x++ {
			yPosition := positions[x]
			if yPosition == y {
				currentConflicts := conflicts[x]
				fmt.Printf("%02d", currentConflicts)
			} else {
				fmt.Print("  ")
			}
			if x < size-1 {
				fmt.Print(" │ ")
			} else {
				fmt.Println(" ║")
			}
		}
		if y < size-1 {
			drawRowMiddle(size)
		}
	}
	drawRowBottom(size)
}

func drawRowTop(size int) {
	s := "╔══"
	s += strings.Repeat("══╤══", size-1)
	s += "══╗"
	fmt.Println(s)
}

func drawRowMiddle(size int) {
	s := "╟──"
	s += strings.Repeat("──┼──", size-1)
	s += "──╢"
	fmt.Println(s)
}

func drawRowBottom(size int) {
	s := "╚══"
	s += strings.Repeat("══╧══", size-1)
	s += "══╝"
	fmt.Println(s)
}
