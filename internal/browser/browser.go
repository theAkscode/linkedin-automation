package browser

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"linkedin-automation/internal/logger"
	"linkedin-automation/internal/stealth"
)

// StartBrowser launches and returns a Rod Browser instances
func StartBrowser() (*rod.Browser, error) {
	fmt.Println("Launching browser...")

	u := launcher.New().
		Delete("leakless").
		Headless(false).
		MustLaunch()

	fmt.Println("Browser launched, connecting...")

	browser := rod.New().ControlURL(u)

	err := browser.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to browser: %w", err)
	}

	fmt.Println("Browser connected successfully!")

	// Navigate to a test page so you can see the browser
	page := browser.MustPage("https://www.google.com")
	_ = page

	return browser, nil
}

// PerformStealthActions executes human-like behavior on the page (mouse movements and scrolling)
// to avoid detection by anti-bot systems
func PerformStealthActions(page *rod.Page) {
	logger.Info("Performing stealth actions - simulating human-like behavior")

	// Perform random mouse movements
	logger.Info("Executing random mouse movements...")
	stealth.MoveMouseRandomly(page)

	// Perform random scrolling
	logger.Info("Executing random page scrolling...")
	stealth.RandomScroll(page)

	logger.Info("Stealth actions completed")
}

// OpenPage opens a new page and navigates to the specified URL
func OpenPage(browser *rod.Browser, url string) (*rod.Page, error) {
	page := browser.MustPage("about:blank")

	err := page.Navigate(url)
	if err != nil {
		return nil, fmt.Errorf("failed to navigate to %s: %w", url, err)
	}

	return page, nil
}
