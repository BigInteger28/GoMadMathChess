package main

import (
	"fmt"
	"strconv"
	"strings"
)

var GRID_WIDTH int
var GRID_EDGE int
var GRID_SIZE int
var piecesPos = [12]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
var grid []int
var PIECES string = "PNBRQKpnbrqk"

func to_sq(y int, x int) int {
	return x*GRID_WIDTH + y
}

func whiteWon() bool {
	for i := 6; i < 12; i++ {
		if piecesPos[i] != -1 {
			return false
		}
	}
	return true
}

func blackWon() bool {
	for i := 0; i < 6; i++ {
		if piecesPos[i] != -1 {
			return false
		}
	}
	return true
}

func drawBoard() {
	for y := GRID_EDGE; y >= 0; y-- {
		for x := 0; x <= GRID_EDGE; x++ {
			pos := to_sq(y, x)
			piece := false
			for i := 0; i < 12; i++ {
				if pos == piecesPos[i] {
					piece = true
					fmt.Printf("%c \t", PIECES[i])
				}
			}
			if !piece {
				fmt.Printf("%d \t", grid[pos])
			}
		}
		fmt.Println("\n")
	}
	fmt.Println("\n")
}

func initialize(grid_width int, boardsetup string) {
	GRID_WIDTH = grid_width
	GRID_EDGE = GRID_WIDTH - 1
	GRID_SIZE = GRID_WIDTH * GRID_WIDTH
	grid = make([]int, GRID_SIZE, GRID_SIZE)
	results := strings.Split(boardsetup, ",")

	for i := 0; i < GRID_SIZE; i++ {
		trigger := false
		for j := 0; j < 12; j++ {
			if results[i] == string(PIECES[j]) {
				piecesPos[j] = i
				grid[i] = 0
				trigger = true
			}
		}
		if trigger == false {
			grid[i], _ = strconv.Atoi(results[i])
		}
	}
}
