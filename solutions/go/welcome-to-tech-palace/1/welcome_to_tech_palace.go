package techpalace

import (
    "fmt"
    "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    upperCust := strings.ToUpper(customer)
	return fmt.Sprintf("Welcome to the Tech Palace, %v", upperCust)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := fmt.Sprint(strings.Repeat("*", numStarsPerLine))
    message :=  border + "\n" + welcomeMsg + "\n" + border
    return message
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	removeStar := strings.ReplaceAll(oldMsg, "*", "")
    return strings.TrimSpace(removeStar)
}
