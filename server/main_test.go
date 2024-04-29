package main

import (
	"secure-messaging-app/server/config"
	"testing"
)

func TestLoadConfigFailure(t *testing.T) {
	// Mock the LoadConfig function
	config.LoadConfig = func() {
		// Simulate a failure in loading the configuration
		panic("Failed to load configuration")
	}

	// Call the main function
	main()

	// Assert that the program exits with an error message
	// Add your assertion here
}
