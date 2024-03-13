package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Player struct {
	name string
}

func main() {

	var x uint8
	var y uint8
	var start bool = true
	var box = [3][3]uint8{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}

	var Player1 = Player{name: "Nicko"}
	var Player2 = Player{name: "Falmark"}

	var hasPlayer1_Move bool
	var current Player = Player1

	displaybox(&box)

	for start {

		if hasPlayer1_Move {
			current = Player2
		} else {
			current = Player1
		}
		fmt.Println("Enter 4 to exit the game")
		fmt.Printf("%v enter Position x: ", current.name)
		fmt.Scanln(&x)
		fmt.Printf("%v Enter Position y: ", current.name)
		fmt.Scanln(&y)

		fmt.Printf("x %v y %v \n", x, y)

		if x == 4 || y == 4 {
			break
		}

		checkX := checkAnswer(x)
		checkY := checkAnswer(y)

		if !checkX {
			fmt.Println("Invalid Position x")
		}
		if !checkY {
			fmt.Println("Invalid Position y")
		}

		if checkY && checkX {

			if !hasPlayer1_Move {
				change := &box[x-1][y-1]
				if *change == 1 {
					*change = 2
					isWinner := checkWin(box)
					if isWinner {
						displaybox(&box)
						fmt.Printf("Winner %v \n", current.name)
						break
					}

					hasPlayer1_Move = true
				} else {
					fmt.Println("The position has already have a value")
					hasPlayer1_Move = false
				}
			} else {
				change := &box[x-1][y-1]
				if *change == 1 {
					*change = 3
					isWinner := checkWin(box)
					if isWinner {
						displaybox(&box)
						fmt.Printf("Winner %v \n", current.name)

						break
					}

					hasPlayer1_Move = false
				} else {
					fmt.Println("The position has already have a value")
					hasPlayer1_Move = true
				}
			}

			displaybox(&box)
			x = 0
			y = 0
		}

		isDraw := checkDraw(box)

		if isDraw {
			fmt.Println("Game Draw")
			box = [3][3]uint8{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
			displaybox(&box)
			// break
		}
		//x is row, y is column

	}

}

func checkDraw(box [3][3]uint8) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if box[i][j] == 1 { // If any cell is empty
				return false // Game is not a draw
			}
		}
	}
	// If all cells are filled and no winner is found, it's a draw
	return true
}

func checkWin(box [3][3]uint8) bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if box[i][0] != 1 && box[i][0] == box[i][1] && box[i][1] == box[i][2] {
			return true
		}
	}

	// Check columns
	for j := 0; j < 3; j++ {
		if box[0][j] != 1 && box[0][j] == box[1][j] && box[1][j] == box[2][j] {
			return true
		}
	}

	// Check diagonals
	if box[0][0] != 1 && box[0][0] == box[1][1] && box[1][1] == box[2][2] {
		return true
	}
	if box[0][2] != 1 && box[0][2] == box[1][1] && box[1][1] == box[2][0] {
		return true
	}

	return false
}

func checkAnswer(answer uint8) bool {
	if answer >= 1 && answer <= 3 {
		return true
	}
	return false
}

func displaybox(b *[3][3]uint8) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for x := 0; x < len(b); x++ {
		if x == 0 {
			fmt.Println("   1 2 3")
		}
		for y := 0; y < len(b[x]); y++ {

			if y == 0 {
				fmt.Printf("%v ", x+1)
				fmt.Print("|")
			}
			value := b[x][y]
			if value == 1 {
				fmt.Printf("%v", "_")
			} else if value == 2 {
				fmt.Printf("%v", "o")
			} else {
				fmt.Printf("%v", "x")
			}

			fmt.Print("|")
		}

		fmt.Println()
	}
}

// func check
