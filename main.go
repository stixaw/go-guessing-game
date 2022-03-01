package main

import (
	"bufio"
	"fmt"
	"log"
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

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Println("An error occurred while reading input. Please try again", err)
	}

	input = strings.TrimSuffix(input, "\n")
	return input
}

func convertInputToInt(input string) int {
	guess, err := strconv.Atoi(input)
	if err != nil {
		invalidMessage := "Invalid input. Please enter an integer value.\nTry Again: "
		fmt.Println(invalidMessage)
	}
	return guess
}

func validateUserGuess(guess, min, max int) bool {
	// if the input is a valid guess we return the guess
	isValid := false
	if guess < min || guess > max {
		fmt.Printf("\nYour guess must be between %v and %v.\nTry Again: ", min, max)
	} else {
		isValid = true
		fmt.Println("\nYour Guess is: ", guess)
	}
	return isValid
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
	fmt.Printf("Your goal is to guess a random number between %v and %v.\nMake your guess: ", min, max)
	for {
		input := getUserInput()
		guess = convertInputToInt(input)
		isValid := validateUserGuess(guess, min, max)
		if isValid {
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
