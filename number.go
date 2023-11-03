package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	secretNumber := rand.Intn(100) + 1

	fmt.Println("Welcome to the Guess the Number game!")
	fmt.Println("I'm thinking of a number between 1 and 100. Can you guess it?")

	var guess int
	attempts := 0

	for {
		attempts++
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		if guess < secretNumber {
			fmt.Println("Too low. Try again.")
		} else if guess > secretNumber {
			fmt.Println("Too high. Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed the number %d in %d attempts!\n", secretNumber, attempts)
			break
		}
	}
}
