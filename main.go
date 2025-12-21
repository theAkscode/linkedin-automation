package main

import (
	"fmt"

	"linkedin-automation/internal/browser"
)

func main() {
	fmt.Println("Starting the LinkedIn Automation")

	br, err := browser.StartBrowser()
	if err != nil {
		fmt.Println("Failed to Start Browser:", err)
		return
	}

	if br == nil {
		fmt.Println("Browser is nil - connection failed")
		return
	}

	fmt.Println("Browser Started Successfully!")
	fmt.Println("Browser is running. Press Ctrl+C to exit.")

	// Keep the program running
	select {}
}
