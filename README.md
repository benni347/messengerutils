# **Messenger Utils**

[![Lint Code Base](https://github.com/benni347/messengerutils/actions/workflows/super-linter.yml/badge.svg?branch=main)](https://github.com/benni347/messengerutils/actions/workflows/super-linter.yml)

This Go package provides a simple utility for printing informational and error messages to the console. It is designed to allow for easy management of logging verbosity and formatting of output. Additionally, it includes a simple event system that allows multiple listeners to subscribe and receive notifications when an event is emitted.

## **Installation**

To install this package, simply run:

```sh
go get github.com/benni347/messengerutils
```

## **Usage**

First, import the package:

```go
import (
"github.com/benni347/messengerutils"
)
```

Create a MessengerUtils instance with the desired verbosity:

```go
messenger := &messengerutils.MessengerUtils{
Verbose: true,
}
```

Use PrintInfo to print informational messages when the verbose flag is set to true:

```go
messenger.PrintInfo("This is an informational message.")
```

Output:

```sh
INFO: This is an informational message.
```

Use PrintError to print error messages:

```go
err := errors.New("This is an error message.")
messengerutils.PrintError("An error occurred", err)
```

Output:

```sh
ERROR: An error occurred: This is an error message.
```

Create an Event instance:

```go
event := &messengerutils.Event{}
```

Subscribe to the event with a listener function:

```go
event.Subscribe(func(data interface{}) {
    fmt.Printf("Received data: %v\n", data)
})
```

Emit an event with data:

```go
event.Emit("Some data")
```

Output:

```sh
Received data: Some data
```

## API

### MessengerUtils struct

```sh
    Verbose: A boolean flag to determine if the PrintInfo method should output messages. If set to true, messages will be printed; otherwise, they will be silenced.
```

#### Methods

```go
    PrintError(message string, err error): Formats and prints an error message to the console with a bold "ERROR:" label.
```

### Event struct

```sh
    listeners: An array of listener functions to be called when the event is emitted.
    lock: A mutex used for ensuring thread-safety when modifying the listeners array or emitting events.
```

#### Event Methods

```go
    Subscribe(listener func(interface{})): Adds a listener function to the Event object.
    Emit(data interface{}): Triggers the event, notifying all subscribed listeners with the provided data.
```

## License

This package is released under the MIT License. See the [LICENSE](LICENSE) file for more information.
