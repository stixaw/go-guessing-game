package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	secretNumber int
	guess        int
	attempts     int
	win          bool
)

func createRandomNumber() int {
	min, max := 1, 10
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(max-min) + min
	fmt.Println("The secret number is: ", randomNumber)
	return randomNumber
}

func getUserInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	// validates the input is a valid integer
	guess, err := strconv.Atoi(input)
	if err != nil {
		invalidMessage := "Invalid input. Please enter an integer value"
		fmt.Println(invalidMessage)
	}

	fmt.Println("Your Guess is: ", guess)
	return guess
}

func compareGuess(guess, secretNumber int) bool {
	win := false
	if attempts <= 3 {
		if guess < secretNumber {
			fmt.Printf("Your guess is %v than the secret number.", "less")
			win = false
			if attempts < 3 {
				fmt.Println("\nTry again: ")
			}
		} else if guess > secretNumber {
			fmt.Println("Your guess is greater than the secret number.\nTry again: ")
			win = false
			if attempts < 3 {
				fmt.Println("\nTry again: ")
			}
		} else {
			fmt.Println("Correct, you Legend!")
			win = true
		}
	}
	return win
}

func main() {
	attempts = 0

	secretNumber = createRandomNumber()
	guessMessage := "You have 3 chances to guess a number between 1 and 10.\nMake your guess: "
	fmt.Println(guessMessage)
	for {
		attempts++
		guess = getUserInput()
		win = compareGuess(guess, secretNumber)
		if win == true {
			break
		}
		if attempts == 3 && win == false {
			fmt.Printf("\nSo Sorry you failed to guess the number, it was: %v.", secretNumber)
			break
		}
	}
}
