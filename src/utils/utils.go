package utils

import "log"

func LogError(message string, err error) {
	log.Printf("%s: %v", message, err)
}
