package storage

import (
	"encoding/json"
	"os"
	"time"
)

// AppState represents the current state of the LinkedIn automation application.
// It tracks whether a login was attempted and the timestamp of the last run.
type AppState struct {
	// LoginAttempted indicates whether a login attempt was made
	LoginAttempted bool `json:"login_attempted"`
	// LastRun stores the timestamp of when the automation last ran
	LastRun time.Time `json:"last_run"`
}

// SaveState saves the current application state to a JSON file.
// It creates or overwrites the data/state.json file with the current timestamp and login status.
// Returns an error if file creation or encoding fails.
func SaveState() error {
	// Create an AppState struct with current timestamp and login status
	state := AppState{
		LoginAttempted: true,
		LastRun:        time.Now(),
	}

	// Create or truncate the state file at the specified path
	file, err := os.Create("data/state.json")
	if err != nil {
		return err
	}

	// Ensure the file is closed when the function returns
	defer file.Close()

	// Encode the state struct to JSON and write it to the file
	encoder := json.NewEncoder(file)
	return encoder.Encode(state)
}
