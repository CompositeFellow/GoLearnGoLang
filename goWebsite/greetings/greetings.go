package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Ehhhh, Whats up %v", name)
	return message
}