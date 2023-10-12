package main

import (
	"fmt"
	"os"
)

func pause_exit_when_press_enter_or_press_any_key_then_enter() {
	fmt.Println("--------------------------------------")
	fmt.Printf("[info] Press Enter or Press any key then Press Enter to exit...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
	fmt.Println("[info] server closed")
}

func pause_exit_when_press_any_key_then_enter() {
	fmt.Println("--------------------------------------")
	fmt.Printf("[info] Press any key then Press Enter to exit...")
	var s string
	fmt.Scan(&s)
	fmt.Println("[info] server closed")
	// not work for enter!
}

func pause_exit_when_typing_exit() {
	fmt.Println("--------------------------------------")
	var s string
	fmt.Println("[info] Press exit to exit...")
	fmt.Scan(&s)
	if s == "exit" {
		// not work!but exit with ctrl + c
		fmt.Println("[info] server closed")
	} else {
		pause_exit_when_typing_exit()
	}
}
