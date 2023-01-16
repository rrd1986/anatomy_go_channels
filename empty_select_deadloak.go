package main

import "fmt"

func service33() {
	fmt.Println("Hello from service!")
}

func Empty_Select() {
	fmt.Println("main() started")

	go service33()

	select {}

	fmt.Println("main() stopped")
}
