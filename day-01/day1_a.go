package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
func main(){
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    number, previous_number, count := 0, 0, 0
    var chaine string
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    fmt.Sscanf(scanner.Text(), "%d %s", &previous_number, &chaine)
    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%d %s", &number, &chaine)
        if number > previous_number {
            count++
        }
        previous_number = number
    }

    fmt.Println(count)
}