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

    number, horizontal, vertical := 0, 0, 0
    var verb string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%s %d", &verb, &number)
        switch verb {
			case "forward":
				horizontal = horizontal + number
			case "down":
				vertical = vertical + number
			case "up":
				vertical = vertical - number
		}
    }

	fmt.Println(horizontal*vertical)
	//fmt.Println(vertical)

    //fmt.Println(count)
}