package browser

import (
	"context"
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// StartBrowser launches and returns a Rod Browser instances
func StartBrowser() (*rod.Browser, error) {
	fmt.Println("Launching browser...")

	u := launcher.New().
		Delete("leakless").
		Headless(false).
		MustLaunch()

	fmt.Println("Browser launched, connecting...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	browser := rod.New().ControlURL(u)
	browser = browser.Context(ctx)

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
