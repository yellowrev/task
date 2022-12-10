package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Print("\tFizzBuzz \n")
		} else if i%5 == 0 {
			fmt.Printf("\tBuzz ")
		} else if i%3 == 0 {
			fmt.Print("\tFizz ")
		} else {
			fmt.Print("\t", i, " ")
		}
	}
}
