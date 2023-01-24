package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	endPart := strings.ToUpper(customer)
	result := fmt.Sprintf("Welcome to the Tech Palace, %s", endPart)
	return result
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starLine := strings.Repeat("*", numStarsPerLine)
	result := fmt.Sprintf("%s\n%s\n%s", starLine, welcomeMsg, starLine)
	return result
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	replaceStar := strings.ReplaceAll(oldMsg, "*", "")
	result := strings.TrimSpace(replaceStar)
	return result
}
