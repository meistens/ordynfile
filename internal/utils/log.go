package utils

import "fmt"

func Info(msg string, args ...any) {
	fmt.Printf("[INFO] "+msg+"\n", args...)
}

func Error(msg string, args ...any) {
	fmt.Printf("[ERROR] "+msg+"\n", args...)
}
