# LinkedIn Automation

A sophisticated Go application for automating LinkedIn login and interactions with human-like behavior simulation using the Rod browser automation library.

## Overview

This project uses [Rod](https://go-rod.github.io/) to automate browser interactions with LinkedIn. It features advanced stealth techniques including random mouse movements, natural scrolling patterns, and realistic typing speeds to avoid detection by anti-bot systems.

## Features

- **Automated LinkedIn Login** - Logs into LinkedIn with email/password credentials
- **Human-like Behavior Simulation**:
  - Random mouse movements (3-5 movements with natural pauses)
  - Random page scrolling (3-5 scrolls with reading time)
  - Realistic typing speeds
  - Random delays between actions (simulating human reaction time)
- **Stealth Fingerprint Masking** - Masks automation indicators to avoid detection
- **Comprehensive Logging** - Detailed logs of all actions performed
- **Browser Automation** - Full control over browser interactions using Rod library
- **Environment-based Credentials** - Secure credential management via .env file

## Prerequisites

- **Go 1.24.5** or higher
- **Chrome/Chromium browser** (Rod will auto-download if not present)
- **Windows Defender exclusion** (optional, see troubleshooting)
- **.env file** with LinkedIn credentials

## Installation

1. Clone the repository:
```bash
git clone https://github.com/YOUR_USERNAME/linkedin-automation.git
cd linkedin-automation
```

2. Install dependencies:
```bash
go mod tidy
```

3. Create a `.env` file in the project root:
```env
LINKEDIN_EMAIL=your_email@example.com
LINKEDIN_PASSWORD=your_password
```

4. Build the project (optional):
```bash
go build
```

## Usage

Run the application:
```bash
go run main.go
```

### What the Application Does

1. **Loads credentials** from the `.env` file
2. **Launches a Chrome browser** in non-headless mode (visible)
3. **Navigates to LinkedIn login page**
4. **Performs automated login** with human-like delays
5. **Verifies successful login** by checking page URL
6. **Applies fingerprint masking** to avoid detection
7. **Simulates human behavior**:
   - Random mouse movements across the page
   - Random scrolling with natural pauses
   - Realistic timing between all actions
8. **Keeps browser open** for 30+ seconds showing the results

### Expected Output

```
[2025-12-24 19:08:28] INFO: Starting LinkedIn Automation
[2025-12-24 19:08:29] INFO: Opening Linkedin Login page
[2025-12-24 19:08:30] INFO: Locating email input field
[2025-12-24 19:08:31] INFO: Locating password input field
[2025-12-24 19:08:32] INFO: Locating and clicking on sign in button
[2025-12-24 19:08:33] INFO: Checking if login was successful
[2025-12-24 19:08:34] INFO: Login Successful - Redirected to home page
[2025-12-24 19:08:35] INFO: Applying fingerprint masking...
[2025-12-24 19:08:36] INFO: Starting human-like behavior simulation...
[2025-12-24 19:08:37] INFO: Performing stealth actions - simulating human-like behavior
[2025-12-24 19:08:38] INFO: Executing random mouse movements...
[2025-12-24 19:08:39] INFO: Executing random page scrolling...
[2025-12-24 19:08:40] INFO: Stealth actions completed
```

## Project Structure

```
linkedin-automation/
├── main.go                    # Entry point - orchestrates the automation workflow
├── go.mod                     # Go module definition
├── go.sum                     # Dependency checksums
├── .env                       # Environment variables (credentials)
├── README.md                  # This file
│
├── internal/
│   ├── automation/
│   │   └── login.go           # LinkedIn login automation logic
│   │
│   ├── browser/
│   │   ├── browser.go         # Browser initialization, page opening, stealth actions
│   │   └── fingerprint.go     # Anti-detection fingerprint masking
│   │
│   ├── logger/
│   │   └── logger.go          # Centralized logging utility
│   │
│   ├── stealth/
│   │   ├── delay.go           # Random delay generation for human-like timing
│   │   ├── mouse.go           # Random mouse movement simulation (3-5 movements)
│   │   ├── scroll.go          # Random page scrolling simulation (3-5 scrolls)
│   │   └── typing.go          # Human-like typing speed simulation
│   │
│   ├── storage/
│   │   └── state.go           # Application state persistence
│   │
│   ├── config/                # Configuration management
│   └── models/                # Data models
│
└── pkg/
    ├── models/                # Public data models
    └── utils/                 # Public utility functions

tests/                         # Test files
```

## Configuration

### .env File Requirements

Create a `.env` file in the project root with your LinkedIn credentials:
```env
LINKEDIN_EMAIL=your_email@example.com
LINKEDIN_PASSWORD=your_password
```

⚠️ **Security Warning**: Never commit `.env` files to version control. Add it to `.gitignore`.

### Browser Settings

Edit [internal/browser/browser.go](internal/browser/browser.go) to customize:

- **Headless mode**: Change `Headless(false)` to `Headless(true)` for invisible browsing
- **Browser binary**: Set custom Chrome path with `.Bin("/path/to/chrome")`
- **Additional flags**: Add launcher flags as needed

### Stealth Settings

Customize behavior simulation in:
- [internal/stealth/delay.go](internal/stealth/delay.go) - Adjust delay ranges
- [internal/stealth/mouse.go](internal/stealth/mouse.go) - Adjust movement ranges (lines 7-8)
- [internal/stealth/scroll.go](internal/stealth/scroll.go) - Adjust scroll ranges (lines 7-8)
- [internal/stealth/typing.go](internal/stealth/typing.go) - Typing speed parameters

## How Stealth Works

### Mouse Movement Simulation
- Performs 3-5 random mouse movements
- Coordinates range: X(100-800px), Y(100-600px)
- Pause between movements: 300-800ms
- Avoids corners to appear natural

### Scrolling Simulation
- Performs 3-5 random scroll actions
- Scroll distance per action: 200-600 pixels
- Pause between scrolls: 800-1500ms
- Mimics human reading behavior

### Typing Simulation
- Realistic keystroke delays between characters
- Variable speed based on character type
- Occasional typos and corrections (optional)

### Timing & Delays
- Login field detection: 800-1500ms delay
- Form submission: 1000-2000ms delay
- Page load verification: 3000-5000ms delay
- All delays are randomized to appear natural

## Troubleshooting

### Windows Defender Blocking Rod

If you see: `Operation did not complete successfully because the file contains a virus or potentially unwanted software`

Add an exclusion to Windows Defender:
```powershell
Add-MpPreference -ExclusionPath "$env:TEMP\leakless-*" -Force
```

Or manually:
1. Open Windows Security
2. Go to Virus & threat protection settings
3. Add exclusion for `%TEMP%\leakless-*`

### Browser Not Opening

- Ensure Chrome/Chromium is installed or let Rod download it
- Check that `Headless(false)` is set in browser.go
- Verify no firewall is blocking the browser process

### Login Fails

- Verify credentials in `.env` file are correct
- Check that LinkedIn is accessible from your network
- Ensure the LinkedIn login page HTML selectors haven't changed
- Check console logs for specific error messages

### No Mouse Movements or Scrolling Visible

- Ensure `Headless(false)` is set (browser must be visible)
- Check that `PerformStealthActions(page)` is being called in main.go
- Verify logger output shows stealth action messages
- Increase pause durations in stealth files if movements are too fast to see

## Dependencies

- `github.com/go-rod/rod` - Browser automation library
- `github.com/joho/godotenv` - .env file loading
- `github.com/ysmood/leakless` - Process management (included with rod)

See [go.mod](go.mod) for complete dependencies.

## Development

### Running with Debug Output

Add logging to see detailed execution:
```go
logger.Info("Your debug message")
```

### Adding New Features

1. Create new modules in the `internal/` directory
2. Follow the existing logging patterns
3. Use stealth functions to maintain natural behavior
4. Test thoroughly before deployment

### Extending Stealth Techniques

Add new stealth methods in `internal/stealth/`:
```go
// Example: New stealth function
func RandomHover(page *rod.Page) {
    // Implementation
}
```

Then call from `browser.PerformStealthActions()`.

## Important Notes

⚠️ **LinkedIn's Terms of Service**: Ensure your automation complies with LinkedIn's ToS. Unauthorized automation may result in account suspension.

✅ **Best Practices**:
- Use delays between actions
- Don't perform bulk operations rapidly
- Respect rate limits
- Log all activities for compliance
- Test with a dedicated account first

## Contributing

Feel free to submit issues and enhancement requests!

## License

MIT License - see LICENSE file for details (if added)

## Author

Created for LinkedIn automation tasks.

## Resources

- [Rod Documentation](https://go-rod.github.io/)
- [Go Language Docs](https://golang.org/doc/)
- [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/)
- [LinkedIn Developer](https://developer.linkedin.com/)

## Changelog

### v1.0.0 (2025-12-24)
- Initial release
- LinkedIn login automation
- Mouse movement simulation
- Page scrolling simulation
- Fingerprint masking
- Comprehensive logging

