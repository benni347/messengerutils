package messengerutils

import "fmt"

type MessengerUtils struct {
	verbose bool
}

func (m *MessengerUtils) PrintInfo(message string) {
	if m.verbose {
		fmt.Printf("\033[1m%s\033[0m: %s\n", "INFO", message)
	}
}
