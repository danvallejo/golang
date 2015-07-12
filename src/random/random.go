package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now)

    seed := now.Nanosecond()

    fmt.Printf("%T / %d\n", seed, seed)

    rand.Seed(time.Now().UnixNano())

    fmt.Println("My favorite number is", rand.Intn(100))
}