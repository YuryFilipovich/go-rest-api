package main

import "fmt"

// Responsible for the instantiation of go app
func Run() error {
	fmt.Println("Starting up Go app")
	return nil
}

func main() {
	fmt.Println("asd")
	err := Run()
	if err != nil {
		fmt.Println(err)
	}
}
