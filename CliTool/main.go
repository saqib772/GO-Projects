package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    // Define flags
    name := flag.String("name", "World", "a string")
    age := flag.Int("age", 0, "an int")
    married := flag.Bool("married", false, "a bool")

    // Parse the flags
    flag.Parse()

    // Print the values of the flags
    fmt.Printf("Hello %s!\n", *name)
    fmt.Printf("You are %d years old.\n", *age)

    if *married {
        fmt.Println("You are married.")
    } else {
        fmt.Println("You are not married.")
    }

    os.Exit(0)
}
