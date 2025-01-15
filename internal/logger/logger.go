// Package logger provides internal logging functionality.
package logger

import "fmt"

// DefaultLogger is the default logging implementation that prints
// messages to standard output.
func DefaultLogger(level, source, payload, err string) {
	fmt.Printf("Level: %s\nSource: %s\nPayload: %s\nError: %s\n",
		level, source, payload, err)
}
