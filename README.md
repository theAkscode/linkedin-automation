# LinkedIn Automation

A Go application for automating LinkedIn interactions using the Rod browser automation library.

## Overview

This project uses [Rod](https://go-rod.github.io/) to automate browser interactions with LinkedIn. It launches a headless Chrome browser and provides a foundation for building LinkedIn automation workflows.

## Features

- Automated browser launch with Chrome/Chromium
- Headless and headed browser modes
- Easy browser control through Rod library
- Configured for Windows, macOS, and Linux

## Prerequisites

- **Go 1.24.5** or higher
- **Chrome/Chromium browser** (Rod will auto-download if not present)
- **Windows Defender exclusion** (see troubleshooting)

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

3. Build the project:
```bash
go build
```

## Usage

Run the application:
```bash
go run main.go
```

The program will:
1. Launch a Chrome browser window
2. Navigate to Google.com
3. Keep the browser open until you press `Ctrl+C`

## Project Structure

```
linkedin-automation/
├── main.go                 # Entry point
├── go.mod                  # Go module definition
├── go.sum                  # Dependency checksums
├── README.md               # This file
└── internal/
    └── browser/
        └── browser.go      # Browser initialization and control
```

## Configuration

### Browser Settings

Edit [internal/browser/browser.go](internal/browser/browser.go) to customize:

- **Headless mode**: Change `Headless(false)` to `Headless(true)` for headless mode
- **Browser binary**: Set custom Chrome path with `.Bin("/path/to/chrome")`
- **Additional flags**: Add launcher flags as needed

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

## Dependencies

- `github.com/go-rod/rod` - Browser automation library
- `github.com/ysmood/leakless` - Process management
- `github.com/ysmood/gson` - JSON handling

See [go.mod](go.mod) for complete dependencies.

## Development

### Running with Debug Output

Add logging to see what's happening:
```go
browser.Logger(os.Stdout)
```

### Adding New Features

Create new modules in the `internal/` directory for additional automation tasks.

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
