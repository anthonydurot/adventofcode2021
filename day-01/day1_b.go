package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)
func main(){
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	prevPrevPrevNumber, prevPrevNumber, prevNumber, number, count := 0, 0, 0, 0, 0 
    scanner := bufio.NewScanner(file)

    scanner.Scan()
	prevPrevPrevNumber, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	prevPrevNumber, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	prevNumber, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

    for scanner.Scan() {
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if number > prevPrevPrevNumber {
			count++
		}

		prevPrevPrevNumber = prevPrevNumber
		prevPrevNumber = prevNumber
		prevNumber = number
    }

    fmt.Println(count)
}