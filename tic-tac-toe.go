package main

import (
	"fmt"
	"math/rand"
	"time"
)

var x int

func main() {
	list := make([][]string, 3)
	for i := 0; i < 3; i++ {
		list[i] = make([]string, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			list[i][j] = " "
		}
	}
	var user string
	var auto string
Again:
	fmt.Print("Choose your nick (X || O):")
	fmt.Scanln(&user)
	if user == "X" || user == "x" {
		user = "X"
		auto = "O"
	} else if user == "O" || user == "o" {
		user = "O"
		auto = "X"
	} else {
		fmt.Println("Enter X || O")
		goto Again
	}
	fmt.Println(`   |   |
--- --- ---
   |   |
--- --- ---
   |   |   `)
	var row, column int
	counter := 1
	for {
		counter++
		check(list)
		if x == 1 {
			break
		}
	enter:
		fmt.Print("row: ")
		fmt.Scan(&row)
		fmt.Print("column: ")
		fmt.Scan(&column)
		if row > 3 || column > 3 || row < 1 || column < 1 {
			fmt.Println("You should enter from 1 to 3 in both row and column!!!")
			goto enter
		}
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if i+1 == row && j+1 == column {
					list[i][j] = user
					if counter == 6 {
						goto new
					}
				random:
					randomIndexI := rand.Intn(3)
					randomIndexJ := rand.Intn(3)
					if list[randomIndexI][randomIndexJ] == " " {
						list[randomIndexI][randomIndexJ] = auto
					} else {
						goto random
					}
				}
			}
		}
	new:
		out(list)

	}
}
func out(ls [][]string) {
	fmt.Printf(" %s | %s | %s \n--- --- ---\n %s | %s | %s \n--- --- ---\n %s | %s | %s \n", ls[0][0], ls[0][1], ls[0][2], ls[1][0], ls[1][1], ls[1][2], ls[2][0], ls[2][1], ls[2][2])
}

func check(list [][]string) {
	if list[0][0] == list[0][1] && list[0][1] == list[0][2] && list[0][0] != " " {
		fmt.Println(list[0][0], "wins!")
		x = 1
	} else if list[0][0] == list[1][0] && list[1][0] == list[2][0] && list[0][0] != " " {
		fmt.Println(list[0][0], "wins!")
		x = 1
	} else if list[1][0] == list[1][1] && list[1][1] == list[1][2] && list[1][0] != " " {
		fmt.Println(list[1][0], "wins!")
		x = 1
	} else if list[0][1] == list[1][1] && list[1][1] == list[2][1] && list[0][1] != " " {
		fmt.Println(list[0][1], "wins!")
		x = 1
	} else if list[2][0] == list[2][1] && list[2][1] == list[2][2] && list[2][0] != " " {
		fmt.Println(list[2][0], "wins!")
		x = 1
	} else if list[0][2] == list[1][2] && list[1][2] == list[2][2] && list[0][2] != " " {
		fmt.Println(list[0][2], "wins!")
		x = 1
	} else if list[0][0] == list[1][1] && list[1][1] == list[2][2] && list[0][0] != " " {
		fmt.Println(list[0][0], "wins!")
		x = 1
	} else if list[0][2] == list[1][1] && list[1][1] == list[2][0] && list[2][0] != " " {
		fmt.Println(list[0][2], "wins!")
		x = 1
	} else {
		count := 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if list[i][j] != " " {
					count++
				}
			}
		}
		if count == 9 {
			fmt.Println("Draw!!")
			x = 1
		}
	}
}
