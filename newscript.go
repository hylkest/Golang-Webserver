package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Seed the random number generator with the current time
    rand.Seed(time.Now().UnixNano())

    // Generate a random number between 1 and 100
    randomNumber := rand.Intn(100) + 1

    fmt.Printf("Generated random number: %d\n", randomNumber)

    if randomNumber%2 == 0 {
        fmt.Println("It's an even number!")
    } else {
        fmt.Println("It's an odd number!")
    }
}
