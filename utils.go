package messengerutils

import "fmt"

type MessengerUtils struct {
	Verbose bool
}

// PrintInfo prints the provided message to the console with the prefix "INFO" in bold, if the MessengerUtils's verbose flag is set to true.
// @param {string} message - The message to print to the console.
// @returns {void}
func (m *MessengerUtils) PrintInfo(message string) {
	if m.Verbose {
		fmt.Printf("\033[1m%s\033[0m: %s\n", "INFO", message)
	}
}

// PrintError formats and prints an error message to standard output.
// @param {string} message - A string representing the error message to display.
// @param {Error} err - An error object providing more information about the error.
// The function formats the error message with the provided message and error object,
// then prints it to the console with a bold "ERROR:" label using ANSI escape codes.
// @returns {void}
func PrintError(message string, err error) {
	fmt.Printf("\033[1mERROR:\033[0m %s: %v\n", message, err)
}
