package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("введите имя аргументом, например: go run main.go макан")
		return
	}

	name := os.Args[1]

	if name == "макан" {
		for {
			fmt.Println("хуесос")
		}
	} else {
		fmt.Println("не тот ввод, брат")
	}
}
