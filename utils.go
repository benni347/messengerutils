package messengerutils

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type MessengerUtils struct {
	Verbose bool
}

// PrintInfo prints the provided message to the console with the prefix "INFO" in bold, if the MessengerUtils's verbose flag is set to true. The function accepts a variadic string parameter and concatenates all strings before printing.
// @param {string} message - One or more strings to print as the message.
// @returns {void}
func (m *MessengerUtils) PrintInfo(message ...interface{}) {
	if m.Verbose {
		finalMessage := ""
		for i, word := range message {
			if i > 0 {
				finalMessage += " "
			}
			switch v := word.(type) {
			case string:
				finalMessage += v
			case int:
				finalMessage += strconv.Itoa(v)
			case time.Time:
				finalMessage += v.Format(time.RFC3339)
			default:
				finalMessage += fmt.Sprintf("Unknown type: %T", v)
			}
		}
		fmt.Printf("\033[1m%s\033[0m: %s\n", "INFO", finalMessage)
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

// Event is a simple event system that allows multiple listeners
// to subscribe and receive notifications when an event is emitted.
// @typedef {Object} Event
// @property {Array<Function>} listeners - An array of listener functions to be called when the event is emitted.
// @property {sync.Mutex} lock - A mutex used for ensuring thread-safety when modifying the listeners array or emitting events.
type Event struct {
	listeners []func(interface{})
	lock      sync.Mutex
}

// Subscribe adds a listener function to the Event object.
// @method
// @memberof Event
// @param {Function} listener - A function to be called when the event is emitted. It takes a single argument, which is the data passed when emitting the event.
// @returns {void}
func (e *Event) Subscribe(listener func(interface{})) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.listeners = append(e.listeners, listener)
}

// Emit triggers the event, notifying all subscribed listeners with the provided data.
// @method
// @memberof Event
// @param {interface{}} data - The data to be passed to each listener when the event is emitted.
// @returns {void}
func (e *Event) Emit(data interface{}) {
	e.lock.Lock()
	defer e.lock.Unlock()
	for _, listener := range e.listeners {
		listener(data)
	}
}
