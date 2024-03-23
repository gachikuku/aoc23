package main

import (
        "fmt"
        "regexp"
)

func main() {
        str := "Game 5: 10 blue, 1 red, 9 green; 1 blue, 8 red, 1 green; 1 red, 6 green"

        // Extract the number after "Game"
        gameRegex := regexp.MustCompile(`Game (\d+)`)
        gameMatch := gameRegex.FindStringSubmatch(str)
        gameNumber := gameMatch[1]
        fmt.Println("Game number:", gameNumber)

        // Extract the color names and numbers
        colorRegex := regexp.MustCompile(`(\d+) (\w+)`)
        colorMatches := colorRegex.FindAllStringSubmatch(str, -1)
        for _, match := range colorMatches {
                colorNumber := match[1]
                colorName := match[2]
                fmt.Printf("Color: %s, Number: %s\n", colorName, colorNumber)
        }
}


