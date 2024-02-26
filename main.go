package main

import (
    "flag"
    "fmt"
    "math/rand"
    "time"
    "github.com/atotto/clipboard"
)

func generateRandomPassword(length int, useLowercase bool, useUppercase bool, useNumbers bool, useSymbols bool) string {
    rand.Seed(time.Now().UnixNano())

    lowercaseChars := "abcdefghijklmnopqrstuvwxyz"
    uppercaseChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    numberChars := "0123456789"
    symbolChars := "!@#$%^&*()_+-=[]{}|;:,.<>?"

    allChars := ""
    password := ""

    if useLowercase {
        allChars += lowercaseChars
    }
    if useUppercase {
        allChars += uppercaseChars
    }
    if useNumbers {
        allChars += numberChars
    }
    if useSymbols {
        allChars += symbolChars
    }

    if allChars == "" {
        return ""
    }

    for i := 0; i < length; i++ {
        randomIndex := rand.Intn(len(allChars))
        password += string(allChars[randomIndex])
    }

    return password
}

func main() {
    // get arguments
    length := flag.Int("length", 16, "password length")
    noLowercase := flag.Bool("no-lowercase", false, "Avoid use lowercase characters")
    noUppercase := flag.Bool("no-uppercase", false, "Avoid use uppercase characters")
    noNumbers := flag.Bool("no-numbers", false, "Avoid use numbers characters")
    noSymbols := flag.Bool("no-symbols", false, "Avoid use symbols characters")
    flag.Parse()

    // generate the password
    password := generateRandomPassword(*length, !*noLowercase, !*noUppercase, !*noNumbers, !*noSymbols)

    if password != "" {
        fmt.Println("Generated password:", password)

        // write password in clipboard
        clipboard.WriteAll(password)
    } else {
	    fmt.Println("No password generated")
	}
}
