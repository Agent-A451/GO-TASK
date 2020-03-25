package main

import "fmt"

func main() {
	var number int
	for {
		fmt.Println("Please input your number: ")
		fmt.Scan(&number)
		fmt.Printf("number = %d, in binary format = %b\n", number, number)
	}
}
