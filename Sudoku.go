package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var a int = 0

func field() [][]int {
	rand.Seed(time.Now().UnixNano())
	count := 0
empty:
	sudoku := make([][]int, 9)
	for i := 0; i < 9; i++ {
		sudoku[i] = make([]int, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			count = 0
		again:
			count++
			if count > 50 {
				goto empty
			}
			number := rand.Intn(9) + 1
			if i < 3 && j < 3 { // 1
				for squareI := 0; squareI < 3; squareI++ {
					for squareJ := 0; squareJ < 3; squareJ++ {
						if sudoku[squareI][squareJ] == number {
							goto again
						}
					}
				}
			} else if i < 3 && j > 2 && j < 6 { // 2
				for squareI := 0; squareI < 3; squareI++ {
					for squareJ := 3; squareJ < 6; squareJ++ {
						if sudoku[squareI][squareJ] == number {
							goto again
						}
					}
				}
			} else if i < 3 && j > 5 { // 3
				for squareI := 0; squareI < 3; squareI++ {
					for squareJ := 6; squareJ < 9; squareJ++ {
						if sudoku[squareI][squareJ] == number {
							goto again
						}
					}
				}
			} else if i > 2 && i < 6 && j < 3 { // 4
				for squareI := 3; squareI < 6; squareI++ {
					for squareJ := 0; squareJ < 3; squareJ++ {
						if sudoku[squareI][squareJ] == number {
							goto again
						}
					}
				}
			} else if i > 2 && i < 6 && j > 2 && j < 6 { // 5
				for squareI := 3; squareI < 6; squareI++ {
					for squareJ := 3; squareJ < 6; squareJ++ {
						if sudoku[squareI][squareJ] == number {
							goto again
						}
					}
				}
			} else if i > 2 && i < 6 && j > 5 { // 6
				for squareI := 3; squareI < 6; squareI++ {
					for squareJ := 6; squareJ < 9; squareJ++ {
						if sudoku[squareI][squareJ] == number {
							goto again
						}
					}
				}
			} else if i > 5 && j < 3 { // 7
				for squareI := 6; squareI < 9; squareI++ {
					for squareJ := 0; squareJ < 3; squareJ++ {
						if sudoku[squareI][squareJ] == number {
							goto again
						}
					}
				}
			} else if i > 5 && j > 2 && j < 6 { // 8
				for squareI := 6; squareI < 9; squareI++ {
					for squareJ := 3; squareJ < 6; squareJ++ {
						if sudoku[squareI][squareJ] == number {
							goto again
						}
					}
				}
			} else if i > 5 && j > 5 { // 9
				for squareI := 6; squareI < 9; squareI++ {
					for squareJ := 6; squareJ < 9; squareJ++ {
						if sudoku[squareI][squareJ] == number {
							goto again
						}
					}
				}
			}
			for lineJ := 0; lineJ < 9; lineJ++ {
				if sudoku[i][lineJ] == number {
					goto again
				}
			}
			for lineI := 0; lineI < 9; lineI++ {
				if sudoku[lineI][j] == number {
					goto again
				}
			}
			sudoku[i][j] = number
		}
	}
	return sudoku
}

func RightInput() ([]byte, error) {
	inputReader := bufio.NewReader(os.Stdin)
	sentence, errSentence := inputReader.ReadBytes('\n')
	return sentence, errSentence
}

func RandomlyShow(FieldedSudoku [][]int) {
Choise:
	fmt.Print(`
	1 - Easy
	2 - Medium
	3 - Hard
(If you want to use ctrl + Z function, just enter 11 anytime and anywhere)
`)
	sentence, errSentence := RightInput()
	if errSentence != nil {
		fmt.Println("Invalid user input!")
		goto Choise
	}
	sentence = sentence[:len(sentence)-1]
	choise, erChoise := strconv.Atoi(string(sentence))
	if erChoise != nil {
		fmt.Println("Invalid user input!")
		goto Choise
	}
	if choise > 3 || choise < 1 {
		fmt.Println("Choose 1 of them")
		goto Choise
	}
	switch choise {
	case 1:
		InCase(FieldedSudoku, choise)
	case 2:
		InCase(FieldedSudoku, choise)
	case 3:
		InCase(FieldedSudoku, choise)
	}
}

func InCase(FieldedSudoku [][]int, choise int) {
	showRandomSudoku := SelectedMode(FieldedSudoku, choise)
	printShowRandomSudoku(showRandomSudoku, 0)
	UserEnter(showRandomSudoku)
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func UserEnter(showRandomSudoku [][]int) {
	bigSlice := make([][]int, 0)
	for {
		littleSlice := make([]int, 2)
		x := 0
		y := 0
		num := 0
		f := 0
		go func() {
			for {
				a += 1
				time.Sleep(time.Second)
				if a%5 == 0 {
					printShowRandomSudoku(showRandomSudoku, a)
					if x != 0 {
						fmt.Println("\nx:", x)
					} else {
						fmt.Print("\nx: ")
					}
					if y != 0 {
						fmt.Println("y:", y)
					} else if x != 0 {
						fmt.Print("y: ")
					}
					if x != 0 && y != 0 {
						fmt.Print("number: ")
					}
				}
				if f != 0 {
					break
				}
			}
		}()
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if showRandomSudoku[i][j] == 0 {
					break
				}
			}
		}
	xgo:
		fmt.Print("\nx: ")
		sentence, errx := RightInput()
		if errx != nil {
			fmt.Print("Invalid user input!")
			goto xgo
		}
		sentence = sentence[:len(sentence)-1]
		_, erx := strconv.Atoi(string(sentence))
		if erx != nil {
			fmt.Print("Invalid user input!")
			goto xgo
		}
		x, _ = strconv.Atoi(string(sentence))
		if x == 11 {
			bigSlice, showRandomSudoku = BackGo(bigSlice, showRandomSudoku)
			x = 0
			goto xgo
		} else if x > 9 || x < 1 {
			fmt.Println("You should enter from 1 to 9!")
			goto xgo
		}
	ygo:
		fmt.Print("y: ")
		sentence, erry := RightInput()
		if erry != nil {
			fmt.Println("Invalid user input!")
			goto ygo
		}
		sentence = sentence[:len(sentence)-1]
		_, ery := strconv.Atoi(string(sentence))
		if ery != nil {
			fmt.Println("Invalid user input!")
			goto ygo
		}
		y, _ = strconv.Atoi(string(sentence))
		if y == 11 {
			bigSlice, showRandomSudoku = BackGo(bigSlice, showRandomSudoku)
			x = 0
			y = 0
			goto xgo
		} else if y > 9 || y < 1 {
			fmt.Println("You should enter from 1 to 9!")
			goto ygo
		}
		if showRandomSudoku[x-1][y-1] != 0 {
			fmt.Println("There is not empty!")
			goto xgo
		}
	Num:
		fmt.Print("number: ")
		sentence, errNum := RightInput()
		if errNum != nil {
			fmt.Println("Invalid user input!")
			num = 0
			goto Num
		}
		sentence = sentence[:len(sentence)-1]
		_, err := strconv.Atoi(string(sentence))
		if err != nil {
			fmt.Println("Invalid user input!")
			num = 0
			goto Num
		}
		num, _ = strconv.Atoi(string(sentence))
		if num == 11 {
			bigSlice, showRandomSudoku = BackGo(bigSlice, showRandomSudoku)
			num = 0
			x = 0
			y = 0
			goto xgo
		} else if !(num > 0 && num < 10) {
			fmt.Println("Enter a number between 1 and 9!")
			num = 0
			goto Num
		}
		if CheckPlace(showRandomSudoku, x-1, y-1, num) {
			showRandomSudoku[x-1][y-1] = num
			littleSlice = append(littleSlice, x-1, y-1)
			bigSlice = append(bigSlice, littleSlice)
			f = 1
		} else {
			num = 0
			fmt.Println("You can not put there this number!")
			goto xgo
		}
		printShowRandomSudoku(showRandomSudoku, a)
		if CheckAnswer(showRandomSudoku) {
			break
		}
	}
	fmt.Println("You won!")
}

func BackGo(bigSlice, showRandomSudoku [][]int) ([][]int, [][]int) {
	if len(bigSlice) != 0 {
		if len(bigSlice) > 1 {
			showRandomSudoku[bigSlice[len(bigSlice)-1][0]][bigSlice[len(bigSlice)-1][1]] = 0
			bigSlice = bigSlice[:len(bigSlice)-1]
		} else {
			showRandomSudoku[bigSlice[0][0]][bigSlice[0][1]] = 0
			bigSlice = make([][]int, 0)
		}
		return bigSlice, showRandomSudoku
	}
	return nil, showRandomSudoku
}

func CheckAnswer(showRandomSudoku [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if showRandomSudoku[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func CheckPlace(showRandomSudoku [][]int, i, j, number int) bool {
	if i < 3 && j < 3 { // 1
		for squareI := 0; squareI < 3; squareI++ {
			for squareJ := 0; squareJ < 3; squareJ++ {
				if showRandomSudoku[squareI][squareJ] == number {
					return false
				}
			}
		}
	} else if i < 3 && j > 2 && j < 6 { // 2
		for squareI := 0; squareI < 3; squareI++ {
			for squareJ := 3; squareJ < 6; squareJ++ {
				if showRandomSudoku[squareI][squareJ] == number {
					return false
				}
			}
		}
	} else if i < 3 && j > 5 { // 3
		for squareI := 0; squareI < 3; squareI++ {
			for squareJ := 6; squareJ < 9; squareJ++ {
				if showRandomSudoku[squareI][squareJ] == number {
					return false
				}
			}
		}
	} else if i > 2 && i < 6 && j < 3 { // 4
		for squareI := 3; squareI < 6; squareI++ {
			for squareJ := 0; squareJ < 3; squareJ++ {
				if showRandomSudoku[squareI][squareJ] == number {
					return false
				}
			}
		}
	} else if i > 2 && i < 6 && j > 2 && j < 6 { // 5
		for squareI := 3; squareI < 6; squareI++ {
			for squareJ := 3; squareJ < 6; squareJ++ {
				if showRandomSudoku[squareI][squareJ] == number {
					return false
				}
			}
		}
	} else if i > 2 && i < 6 && j > 5 { // 6
		for squareI := 3; squareI < 6; squareI++ {
			for squareJ := 6; squareJ < 9; squareJ++ {
				if showRandomSudoku[squareI][squareJ] == number {
					return false
				}
			}
		}
	} else if i > 5 && j < 3 { // 7
		for squareI := 6; squareI < 9; squareI++ {
			for squareJ := 0; squareJ < 3; squareJ++ {
				if showRandomSudoku[squareI][squareJ] == number {
					return false
				}
			}
		}
	} else if i > 5 && j > 2 && j < 6 { // 8
		for squareI := 6; squareI < 9; squareI++ {
			for squareJ := 3; squareJ < 6; squareJ++ {
				if showRandomSudoku[squareI][squareJ] == number {
					return false
				}
			}
		}
	} else if i > 5 && j > 5 { // 9
		for squareI := 6; squareI < 9; squareI++ {
			for squareJ := 6; squareJ < 9; squareJ++ {
				if showRandomSudoku[squareI][squareJ] == number {
					return false
				}
			}
		}
	}
	for lineJ := 0; lineJ < 9; lineJ++ {
		if showRandomSudoku[i][lineJ] == number {
			return false
		}
	}
	for lineI := 0; lineI < 9; lineI++ {
		if showRandomSudoku[lineI][j] == number {
			return false
		}
	}
	return true
}

func printShowRandomSudoku(showRandomSudoku [][]int, n int) {
	colorGreen := "\033[34m"
	colorReset := "\033[0m"
	clear()
	fmt.Println("    1  2  3   4  5  6   7  8  9       ", n/60, ":", n%60, "\n")
	for i := 0; i < 9; i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("    ------- + ------- + -------")
		}
		fmt.Print(i+1, " ")
		for j := 0; j < 9; j++ {
			if showRandomSudoku[i][j] != 0 {
				if j%3 == 0 && j != 0 {
					fmt.Print(" | ", string(colorGreen), showRandomSudoku[i][j], string(colorReset))
				} else {
					fmt.Print(string(colorGreen), "  ", showRandomSudoku[i][j], string(colorReset))
				}
			} else {
				if j%3 == 0 && j != 0 {
					fmt.Print(" | .")
				} else {
					fmt.Print("  .")
				}
			}
		}
		fmt.Println()
	}
}

func SelectedMode(FieldedSudoku [][]int, choice int) [][]int {
	fmt.Println()
	rand.Seed(time.Now().UnixNano())
	showRandomSudoku := make([][]int, 9)
	for i := 0; i < 9; i++ {
		showRandomSudoku[i] = make([]int, 9)
	}
	if choice == 1 {
		for d := 0; d < 65; d++ {
			a := rand.Intn(9) + 1
			b := rand.Intn(9) + 1
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					if i+1 == a && j+1 == b {
						showRandomSudoku[i][j] = FieldedSudoku[i][j]
					}
				}
			}
		}
		return showRandomSudoku
	} else if choice == 2 {
		for d := 0; d < 40; d++ {
			a := rand.Intn(9) + 1
			b := rand.Intn(9) + 1
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					if i+1 == a && j+1 == b {
						showRandomSudoku[i][j] = FieldedSudoku[i][j]
					}
				}
			}
		}
		return showRandomSudoku
	}
	for d := 0; d < 30; d++ {
		a := rand.Intn(9) + 1
		b := rand.Intn(9) + 1
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if i+1 == a && j+1 == b {
					showRandomSudoku[i][j] = FieldedSudoku[i][j]
				}
			}
		}
	}
	return showRandomSudoku
}

func main() {
	FieldedSudoku := field()
	RandomlyShow(FieldedSudoku)
}
