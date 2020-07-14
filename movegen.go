package main

func goNorth(square int, value int) int {
	return (square + value) % GRID_SIZE
}

func goEast(square int, value int) int {
	dest := square
	for i := 0; i < value; i++ {
		if square < GRID_WIDTH*GRID_EDGE {
			dest = square + GRID_WIDTH
		} else {
			dest = square + 1 - (GRID_WIDTH * GRID_EDGE)
		}
	}
	return dest
}

func goSouth(square int, value int) int {
	return (square - value) % GRID_SIZE
}

func goWest(square int, value int) int {
	dest := square
	for i := 0; i < value; i++ {
		if square > GRID_WIDTH {
			dest = square - GRID_WIDTH
		} else {
			dest = square - 1 + (GRID_WIDTH * GRID_EDGE)
		}
	}
	return dest
}

func goNorthEast(square int, value int) int {
	dest := square
	for i := 0; i < value; i++ {
		dest = goNorth(square, 1)
		dest = goEast(square, 1)
	}
	return dest
}

func goSouthEast(square int, value int) int {
	dest := square
	for i := 0; i < value; i++ {
		dest = goSouth(square, 1)
		dest = goEast(square, 1)
	}
	return dest
}

func goSouthWest(square int, value int) int {
	dest := square
	for i := 0; i < value; i++ {
		dest = goSouth(square, 1)
		dest = goWest(square, 1)
	}
	return dest
}

func goNorthWest(square int, value int) int {
	dest := square
	for i := 0; i < value; i++ {
		dest = goNorth(square, 1)
		dest = goWest(square, 1)
	}
	return dest
}

func takePiece(piecesPos []int, own int, newSquare int) {
	for i := 0; i < 12; i++ {
		if i != own && piecesPos[i] == newSquare {
			piecesPos[i] = -1
		}
	}
}

func doWhitePawn(piecesPos []int, grid []int) {
	square := piecesPos[0]
	value := grid[square]
	var newSquare int
	if value == 0 {
		newSquare = goNorth(square, 1)
	} else {
		newSquare = goNorth(square, value)
	}
	piecesPos[0] = newSquare
	takePiece(piecesPos, 0, newSquare)
}

func doWhiteKnight(piecesPos []int, grid []int) {
	square := piecesPos[1]
	value := grid[square]
	var newSquare int
	if value == 0 {
		newSquare = goEast(square, 2)
		newSquare = goNorth(square, 1)
	} else {
		for i := 0; i < value; i++ {
			newSquare = goEast(square, 2)
			newSquare = goNorth(square, 1)
		}
	}
	piecesPos[1] = newSquare
	takePiece(piecesPos, 1, newSquare)
}

func doWhiteBishop(piecesPos []int, grid []int) {
	square := piecesPos[2]
	value := grid[square]
	var newSquare int
	if value == 0 {
		newSquare = goNorthEast(square, 1)
	} else {
		newSquare = goNorthEast(square, value)
	}
	piecesPos[2] = newSquare
	takePiece(piecesPos, 2, newSquare)
}

func doWhiteRook(piecesPos []int, grid []int) {
	square := piecesPos[3]
	value := grid[square]
	var newSquare int
	if value == 0 {
		newSquare = goEast(square, 1)
	} else {
		newSquare = goEast(square, value)
	}
	piecesPos[3] = newSquare
	takePiece(piecesPos, 3, newSquare)
}

func doWhiteQueen(piecesPos []int, grid []int) {
	square := piecesPos[4]
	value := grid[square]
	var newSquare int
	if value == 0 {
		newSquare = goNorthWest(square, 1)
	} else {
		newSquare = goNorthWest(square, value)
	}
	piecesPos[4] = newSquare
	takePiece(piecesPos, 4, newSquare)
}

func doWhiteKing(piecesPos []int, grid []int) {
	square := piecesPos[5]
	value := grid[square]
	var newSquare int
	if value == 0 {
		newSquare = goWest(square, 1)
	} else {
		newSquare = goWest(square, value)
	}
	piecesPos[5] = newSquare
	takePiece(piecesPos, 5, newSquare)
}

func doBlackPawn(piecesPos []int, grid []int) {
	square := piecesPos[6]
	value := grid[square]
	var newSquare int
	if value == 0 {
		newSquare = goSouth(square, 1)
	} else {
		newSquare = goSouth(square, value)
	}
	piecesPos[6] = newSquare
	takePiece(piecesPos, 6, newSquare)
}

func doBlackKnight(piecesPos []int, grid []int) {
	square := piecesPos[7]
	value := grid[square]
	var newSquare int
	if value == 0 {
		newSquare = goWest(square, 2)
		newSquare = goSouth(square, 1)
	} else {
		for i := 0; i < value; i++ {
			newSquare = goWest(square, 2)
			newSquare = goSouth(square, 1)
		}
	}
	piecesPos[7] = newSquare
	takePiece(piecesPos, 7, newSquare)
}

func doBlackBishop(piecesPos []int, grid []int) {
	square := piecesPos[8]
	value := grid[square]
	var newSquare int
	if value == 0 {
		newSquare = goSouthWest(square, 1)
	} else {
		newSquare = goSouthWest(square, value)
	}
	piecesPos[8] = newSquare
	takePiece(piecesPos, 8, newSquare)
}

func doBlackRook(piecesPos []int, grid []int) {
	square := piecesPos[9]
	value := grid[square]
	var newSquare int
	if value == 0 {
		newSquare = goWest(square, 1)
	} else {
		newSquare = goWest(square, value)
	}
	piecesPos[9] = newSquare
	takePiece(piecesPos, 9, newSquare)
}

func doBlackQueen(piecesPos []int, grid []int) {
	square := piecesPos[10]
	value := grid[square]
	var newSquare int
	if value == 0 {
		newSquare = goSouthEast(square, 1)
	} else {
		newSquare = goSouthEast(square, value)
	}
	piecesPos[10] = newSquare
	takePiece(piecesPos, 10, newSquare)
}

func doBlackKing(piecesPos []int, grid []int) {
	square := piecesPos[11]
	value := grid[square]
	var newSquare int
	if value == 0 {
		newSquare = goEast(square, 1)
	} else {
		newSquare = goEast(square, value)
	}
	piecesPos[11] = newSquare
	takePiece(piecesPos, 11, newSquare)
}
