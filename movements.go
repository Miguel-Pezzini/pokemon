package main

import (
	"fmt"
	"math/rand"
)

func setCursorPosition(x, y int) {
	// ANSI: ESC[row;colH (1-indexed). Original called with (col, row).
	fmt.Printf("\033[%d;%dH", y+1, x+1)
}

func clearPosition(x, y int) {
	setCursorPosition(x, y)
	fmt.Print(" ")
}

func putBush(x, y int) {
	setCursorPosition(x, y)
	fmt.Print("\033[0;42m")
	fmt.Print(" ")
	fmt.Print("\033[0m")
}

func drawPosition(x, y int) {
	setCursorPosition(x, y)
	fmt.Print("@")
}

func clearAndDrawn(mat [][]int, x, y *int, newX, newY int) {
	clearPosition(*y, *x)
	mat[*x][*y] = 0
	*x = newX
	*y = newY
	mat[*x][*y] = 2
	drawPosition(*y, *x)
}

func putBushAndDrawn(mat [][]int, x, y *int, newX, newY int) {
	putBush(*y, *x)
	mat[*x][*y] = 7
	*x = newX
	*y = newY
	mat[*x][*y] = 2
	drawPosition(*y, *x)
}

func movement(mat [][]int, key byte, xPosicao, yPosicao *int, character Character) int {
	if key == '/' {
		menuInGame(character)
		return 2
	}
	newX := *xPosicao
	newY := *yPosicao
	switch key {
	case 'w':
		newX--
	case 's':
		newX++
	case 'd':
		newY++
	case 'a':
		newY--
	default:
		return 0
	}

	if newX >= 0 && newX < len(mat) && newY >= 0 && newY < len(mat[0]) {
		if mat[newX][newY] != 1 && mat[newX][newY] != 4 {
			if mat[newX][newY] == 3 {
				return 3
			}
			if mat[newX][newY] == 5 {
				return 5
			}
			if mat[newX][newY] == 6 {
				return 6
			}
			if mat[newX][newY] == 7 {
				random := rand.Intn(10)
				if newX < *xPosicao || newX > *xPosicao {
					if mat[*xPosicao][*yPosicao+1] == 0 || mat[*xPosicao][*yPosicao-1] == 0 {
						if mat[*xPosicao][*yPosicao-1] == 7 && mat[*xPosicao-1][*yPosicao] == 7 {
						} else if mat[*xPosicao][*yPosicao+1] == 7 && mat[*xPosicao-1][*yPosicao] == 7 {
						} else if mat[*xPosicao+1][*yPosicao] == 7 && mat[*xPosicao][*yPosicao+1] == 7 {
						} else if mat[*xPosicao+1][*yPosicao] == 7 && mat[*xPosicao][*yPosicao-1] == 7 {
						} else {
							clearAndDrawn(mat, xPosicao, yPosicao, newX, newY)
						}
					}
				}
				if newY > *yPosicao || newY < *yPosicao {
					if mat[*xPosicao+1][*yPosicao] == 0 || mat[*xPosicao-1][*yPosicao] == 0 {
						if mat[*xPosicao-1][*yPosicao] == 7 && mat[*xPosicao][*yPosicao-1] == 7 {
						} else if mat[*xPosicao-1][*yPosicao] == 7 && mat[*xPosicao][*yPosicao+1] == 7 {
						} else if mat[*xPosicao+1][*yPosicao] == 7 && mat[*xPosicao][*yPosicao-1] == 7 {
						} else if mat[*xPosicao+1][*yPosicao] == 7 && mat[*xPosicao][*yPosicao+1] == 7 {
						} else {
							clearAndDrawn(mat, xPosicao, yPosicao, newX, newY)
						}
					}
				}
				putBushAndDrawn(mat, xPosicao, yPosicao, newX, newY)
				if random == 7 {
					return 7
				}
				return 0
			}
			if mat[newX][newY] == 0 {
				if newX < *xPosicao || newX > *xPosicao {
					if mat[*xPosicao][*yPosicao+1] == 7 && mat[*xPosicao][*yPosicao-1] == 7 {
						putBushAndDrawn(mat, xPosicao, yPosicao, newX, newY)
					}
					if mat[*xPosicao-1][*yPosicao] == 7 && mat[*xPosicao][*yPosicao-1] == 7 {
						putBushAndDrawn(mat, xPosicao, yPosicao, newX, newY)
					}
					if mat[*xPosicao-1][*yPosicao] == 7 && mat[*xPosicao][*yPosicao+1] == 7 {
						putBushAndDrawn(mat, xPosicao, yPosicao, newX, newY)
					}
					if mat[*xPosicao+1][*yPosicao] == 7 && mat[*xPosicao][*yPosicao-1] == 7 {
						putBushAndDrawn(mat, xPosicao, yPosicao, newX, newY)
					}
					if mat[*xPosicao+1][*yPosicao] == 7 && mat[*xPosicao][*yPosicao+1] == 7 {
						putBushAndDrawn(mat, xPosicao, yPosicao, newX, newY)
					}
				}
				if newY > *yPosicao || newY < *yPosicao {
					if mat[*xPosicao+1][*yPosicao] == 7 && mat[*xPosicao-1][*yPosicao] == 7 {
						putBushAndDrawn(mat, xPosicao, yPosicao, newX, newY)
					}
					if mat[*xPosicao-1][*yPosicao] == 7 && mat[*xPosicao][*yPosicao-1] == 7 {
						putBushAndDrawn(mat, xPosicao, yPosicao, newX, newY)
					}
					if mat[*xPosicao-1][*yPosicao] == 7 && mat[*xPosicao][*yPosicao+1] == 7 {
						putBushAndDrawn(mat, xPosicao, yPosicao, newX, newY)
					}
					if mat[*xPosicao+1][*yPosicao] == 7 && mat[*xPosicao][*yPosicao-1] == 7 {
						putBushAndDrawn(mat, xPosicao, yPosicao, newX, newY)
					}
					if mat[*xPosicao+1][*yPosicao] == 7 && mat[*xPosicao][*yPosicao+1] == 7 {
						putBushAndDrawn(mat, xPosicao, yPosicao, newX, newY)
					}
				}
			}

			clearPosition(*yPosicao, *xPosicao)
			mat[*xPosicao][*yPosicao] = 0
			*xPosicao = newX
			*yPosicao = newY
			mat[*xPosicao][*yPosicao] = 2
			drawPosition(*yPosicao, *xPosicao)
		}
	}

	return 0
}
