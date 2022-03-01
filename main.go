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
	min          int
	max          int
	secretNumber int
	guess        int
	attempts     int
	win          bool
)

func createRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(max-min) + min
	fmt.Println("The secret number is: ", randomNumber)
	return randomNumber
}

func getUserInput(min, max int) int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	// validates the input is a valid integer
	guess, err := strconv.Atoi(input)
	if err != nil {
		invalidMessage := "Invalid input. Please enter an integer value"
		fmt.Println(invalidMessage)
	}

	//validate the number is between the min and max:
	if guess < min || guess > max {
		fmt.Printf("\nYour guess must be between %v and %v.\nTry Again: ", min, max)
	} else {
		fmt.Println("\nYour Guess is: ", guess)
	}

	return guess
}

func compareGuess(guess, secretNumber int) bool {
	win := false
	if attempts <= 3 {
		if guess < secretNumber {
			fmt.Printf("Your guess is less than the secret number.")
			win = false
		} else if guess > secretNumber {
			fmt.Printf("Your guess is greater than the secret number.")
			win = false
		} else {
			fmt.Println("Correct, you Legend! You guessed right after", attempts, "attempts")
			win = true
		}
		if attempts < 3 && win == false {
			fmt.Printf("You have %v more tries.\nTry again: ", 3-attempts)
		}
	}
	return win
}

func main() {
	attempts = 0
	min := 1
	max := 10

	secretNumber = createRandomNumber(min, max)
	fmt.Printf("You have 3 chances to guess a number between %v and %v.\nMake your guess: ", min, max)
	for {
		guess = getUserInput(min, max)
		if guess >= min && guess <= max {
			attempts++
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
}
