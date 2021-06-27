package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	GuessingGame()
}

func GuessingGame() {
	option := ""
	fmt.Print("Welcome to guessing game, ")
	for option != "c" {
		fmt.Println("\nChoose a button below: ")
		fmt.Println("a. Start game\nb. Read manual\nc. Exit")
		fmt.Scan(&option)
		option = strings.ToLower(option)

		switch option {
		case "a":
			startGame()
		case "b":
			loadManual()
		case "c":
			fmt.Println("See you, bye!")
		default:
			fmt.Println("Error, no such an option. Please, try again.")
		}
	}
}

func startGame() {
	var n_lives int = 7
	unk_number := rand.Intn(100)
	// fmt.Println(unk_number)
	printStars()
	fmt.Println("\nAlright let's get started.")
	printStars()
	var g_num int

	for n_lives > 0 {
		// fmt.Println("Number of lives: ", n_lives)
		fmt.Println("\nGuess a number in interval (0, 100):")
		fmt.Scan(&g_num)
		if g_num == unk_number {
			fmt.Println("Yeaaaa! You get it right!")
			return
		} else if g_num > unk_number {
			fmt.Println("Almost there, your number is slighlty greater, try again")
			n_lives--
		} else if g_num < unk_number {
			fmt.Println("Almost there, your number is less than guessed one, try again")
			n_lives--
		}
	}
	printStars()
	fmt.Println("\nOuch, no more lives left, you lose.")
	printStars()
	return
}

func loadManual() {
	printStars()
	fmt.Println("\nSimple rules: You gotta guess a number in a given interval (0, 100).")
	fmt.Println("Initially you have only 7 lives, for incorrect guessing you lose 1 live.")
	fmt.Println("When the number of your lives would be equal to 0, the game overs.")
	printStars()
}

func printStars() {
	for i := 50; i > 0; i-- {
		fmt.Print("*")
	}
}
