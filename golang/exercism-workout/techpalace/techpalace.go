package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.

func WelcomeMessage(customer string) string {

	var cust_name = strings.ToUpper(customer)

	return ("Welcome to the Tech Palace, " + cust_name)

}

// AddBorder adds a border to a welcome message.

func AddBorder(welcomeMsg string, numStarsPerLine int) string {

	var border = strings.Repeat("*", numStarsPerLine)

	return border + "\n" + welcomeMsg + "\n" + border

}

// CleanupMessage cleans up an old marketing message.

func CleanupMessage(oldMsg string) string {

	var clean = strings.Trim(strings.ReplaceAll(oldMsg, "*", ""), " \n")

	return clean

}
