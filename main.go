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

var secretNumber int
var guess int
var win bool

func createRandomNumber() int {
	min, max := 1, 10
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(max-min) + min
	// fmt.Println("The secret number is: ", secretNumber)
	return secretNumber
}

func getUserInput() int {
	guessMessage := "Guess a number between 1 and 10"

	fmt.Println(guessMessage)

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

	if guess < secretNumber {
		fmt.Println("Your guess is less than the secret number. Try again")
		win = false
	} else if guess > secretNumber {
		fmt.Println("Your guess is greater than the secret number. Try again")
		win = false
	} else {
		fmt.Println("Correct, you Legend!")
		win = true
	}
	return win
}

func main() {
	attempts := 0

	secretNumber = createRandomNumber()
	for {
		attempts++
		guess = getUserInput()
		win = compareGuess(guess, secretNumber)
		if win == true {
			break
		}
		if attempts == 3 && win == false {
			fmt.Println("So Sorry you failed to guess the number, it was: ", secretNumber)
			break
		}
	}
}
