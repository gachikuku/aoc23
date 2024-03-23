package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
    "strconv"
)

func minima(conudrum string) int {

    var red_max int
    var green_max int
    var blue_max int

    colorRegex := regexp.MustCompile(`(\d+) (\w+)`)
    colorMatches := colorRegex.FindAllStringSubmatch(conudrum, -1)

    for _, match := range colorMatches {
        colorNumber := match[1]
        number, err := strconv.Atoi(colorNumber)
        if err != nil {
            panic(err)
        }
        colorName := match[2]
        switch colorName {
        case "red":
            if number > red_max { red_max = number }
        case "green":
            if number > green_max { green_max = number }
        case "blue":
            if number > blue_max { blue_max = number }
        }
    }

    return red_max * green_max * blue_max
}


func main() {

    var ans int

    readFile, err := os.Open("input.txt")

    if err != nil {
        panic(err)
    }

    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
            ans += minima(fileScanner.Text())
        }

    readFile.Close()
    fmt.Println(ans)
}



