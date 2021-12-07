package main

import (
	"bufio"
	"fmt"
	"strconv"
	"log"
	"os"
)
func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var slice [][]int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		slice = append(slice, numbers(scanner.Text()))
    }


	slice_gamma, slice_epsilon := []int{}, []int{}

	for j := 0 ; j < len(slice[0]) ; j++ {
		count_one := 0
		count_zero := 0
		for i := 0 ; i < len(slice) ; i++ {
			if slice[i][j] == 1 {
				count_one++
			} else {
				count_zero++
			}
		}
		if count_one > count_zero {
			slice_gamma = append(slice_gamma, 1)
			slice_epsilon = append(slice_epsilon, 0)
		} else {
			slice_gamma = append(slice_gamma, 0)
			slice_epsilon = append(slice_epsilon, 1)
		}
	}

	fmt.Println(convertBinary(slice_gamma) * convertBinary(slice_epsilon))
}



func numbers (s string) []int {
	var slice []int
	for _, char := range s {
		str, err := strconv.Atoi(string(char))
		if err != nil {
			log.Fatal(err)
		}
		slice = append(slice, str)
	}
	return slice
}


func convertBinary (slice []int) int {
	sum := 0
	resized_len := len(slice) - 1
	for i, v := range(slice) {
		sum += v * IntPow(2, resized_len - i)
	}
	return sum
}

func IntPow(n, m int) int {
    if m == 0 {
        return 1
    }
    result := n
    for i := 2; i <= m; i++ {
        result *= n
    }
    return result
}