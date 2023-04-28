package main

import (
    "flag"
    "fmt"
    "os"

    "golang.org/x/crypto/bcrypt"
)

func main() {
    // Define flags
    name := flag.String("name", "World", "a string")
    age := flag.Int("age", 0, "an int")
    married := flag.Bool("married", false, "a bool")
    password := flag.String("password", "", "a string")

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

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println("Error hashing password:", err)
        os.Exit(1)
    }
    fmt.Printf("Your hashed password is: %s\n", string(hashedPassword))

    os.Exit(0)
}
// Run The Following Command 
// ./mycli -name=John -age=30 -married=true -password=secret 
// or 
// go run main.go -name=John -age=30 -married=true -password=secret 
