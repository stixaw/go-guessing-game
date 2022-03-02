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
	maxAttempts  int
	win          bool
)

func createRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(max-min) + min
	// fmt.Println("The secret number is: ", randomNumber)
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
	// if the input is a valid guess we return the true
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
	if attempts <= maxAttempts {
		if guess < secretNumber {
			fmt.Printf("Your guess is less than the number, let the force guide you.")
			win = false
		} else if guess > secretNumber {
			fmt.Printf("Your guess is greater than the number, let the force guide you.")
			win = false
		} else {
			fmt.Println("Correct, Padawan!\nYou are strong in the force.\nYou guessed right after", attempts, "attempts")
			win = true
		}
		if attempts < maxAttempts && win == false {
			fmt.Printf("You have %v more tries.\nTry again: ", maxAttempts-attempts)
		}
	}
	return win
}

func main() {

	//to do make these submittable by commandline
	maxAttempts = 5
	attempts = 0
	min := 1
	max := 20

	// the program
	fmt.Println("Padwan, its time to test your use the force.\nPrepare to be tested!")

	secretNumber = createRandomNumber(min, max)
	fmt.Printf("To see how strong you are in the force you must guess what number I am thinking of between %v and %v.\nMake your guess: ", min, max)
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
			if attempts == maxAttempts && win == false {
				fmt.Printf("\nYou are not very strong in the force Padawan for the number was: %v.", secretNumber)
				break
			}
		}
	}
}
