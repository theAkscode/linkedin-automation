package main

import (
	"os"

	"linkedin-automation/internal/automation"
	"linkedin-automation/internal/browser"
	"linkedin-automation/internal/logger"

	"github.com/joho/godotenv"
)

// main orchestrates the LinkedIn automation workflow:
// 1. Loads environment variables
// 2. Initializes a browser instance
// 3. Opens the LinkedIn login page
// 4. Extracts credentials from environment
// 5. Performs the login action
// 6. Applies browser fingerprint masking for stealth
func main() {
	// Log the start of the automation process
	logger.Info("Starting LinkedIn Automation")

	// Step 1: Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
		return
	}

	// Step 2: Start the browser instance for automation
	br, err := browser.StartBrowser()
	if err != nil {
		logger.Error("Failed to start Browser: " + err.Error())
		return
	}
	// Ensure browser is properly closed when the function exits
	defer br.Close()

	// Step 3: Open the LinkedIn login page in the browser
	page, err := browser.OpenPage(br, "https://www.linkedin.com/login")
	if err != nil {
		logger.Error("Failed to open LinkedIn Page: " + err.Error())
		return
	}

	// Step 4: Read LinkedIn credentials from environment variables
	// Make sure LINKEDIN_EMAIL and LINKEDIN_PASSWORD are set in your .env file
	email := os.Getenv("LINKEDIN_EMAIL")
	password := os.Getenv("LINKEDIN_PASSWORD")

	// Step 5: Perform the login action with credentials
	err = automation.LoginLinkedln(page, email, password)
	if err != nil {
		logger.Error("Login Failed: " + err.Error())
		return
	}
	logger.Info("Login Successful")

	// Step 6: Apply browser fingerprint masking to avoid detection
	// This applies stealth techniques like mouse movements, scrolling, and delays
	logger.Info("Applying fingerprint masking...")
	browser.ApplyFingerprintMasking(br)

	// Step 7: Perform stealth actions on the LinkedIn home page
	// This simulates human-like mouse movements and scrolling behavior
	logger.Info("Starting human-like behavior simulation...")
	browser.PerformStealthActions(page)

	// Keep the browser open to see results before closing
	select {}
}
