package messengerutils

import (
	"fmt"
	"strconv"
	"strings"
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
		finalMessage := formatMessage(message...)
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

// PrintToDo formats and prints a to-do message to standard output.
//
// @param {interface{}} message - An array of any type that forms the to-do message to display.
//
//	The function checks each element in the message. If it is a time.Time type,
//	it is formatted according to RFC3339. All elements are then converted into
//	their string representations, concatenated into a single string,
//	then formatted and printed the string to the console with a bold "TODO:" label
//	using ANSI escape codes.
//
// @returns {void}
func PrintToDo(message ...interface{}) {
	finalMessage := formatMessage(message...)
	fmt.Printf("\033[1mTODO:\033[0m %s\n", finalMessage)
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

func formatMessage(message ...interface{}) string {
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
		case []uint8:
			finalMessage += string(v) // Converts the []uint8 slice to a string
		case []uint16:
			var numbers []string
			for _, val := range v {
				numbers = append(numbers, fmt.Sprintf("%d", val))
			}
			finalMessage += strings.Join(numbers, ", ")
		case []uint32:
			var numbers []string
			for _, val := range v {
				numbers = append(numbers, fmt.Sprintf("%d", val))
			}
			finalMessage += strings.Join(numbers, ", ")
		case []uint64:
			var numbers []string
			for _, val := range v {
				numbers = append(numbers, fmt.Sprintf("%d", val))
			}
			finalMessage += strings.Join(numbers, ", ")
		case uint:
			finalMessage += fmt.Sprintf("%d", v)
		case uint8:
			finalMessage += fmt.Sprintf("%d", v)
		case uint16:
			finalMessage += fmt.Sprintf("%d", v)
		case uint32:
			finalMessage += fmt.Sprintf("%d", v)
		case uint64:
			finalMessage += fmt.Sprintf("%d", v)
		default:
			finalMessage += fmt.Sprintf("Unknown type: %T", v)
		}
	}
	return finalMessage
}
