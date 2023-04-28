package messengerutils

import "fmt"

type MessengerUtils struct {
	verbose bool
}

// PrintInfo prints the provided message to the console with the prefix "INFO" in bold, if the MessengerUtils's verbose flag is set to true.
// @param {string} message - The message to print to the console.
// @returns {void}
func (m *MessengerUtils) PrintInfo(message string) {
	if m.verbose {
		fmt.Printf("\033[1m%s\033[0m: %s\n", "INFO", message)
	}
}
