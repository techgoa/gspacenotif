package logger

import "fmt"

func DefaultLogger(level, source, payload, err string) {
	fmt.Printf("Level: %s\nSource: %s\nPayload: %s\nError: %s\n",
		level, source, payload, err)
}
