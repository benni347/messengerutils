# **Messenger Utils**

This Go package provides a simple utility for printing informational and error messages to the console. It is designed to allow for easy management of logging verbosity and formatting of output.

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

```
INFO: This is an informational message.
```

Use PrintError to print error messages:

```go
err := errors.New("This is an error message.")
messengerutils.PrintError("An error occurred", err)
```

Output:

```
ERROR: An error occurred: This is an error message.
```

# API

## MessengerUtils struct

    Verbose: A boolean flag to determine if the PrintInfo method should output messages. If set to true, messages will be printed; otherwise, they will be silenced.

## Methods

    PrintInfo(message string): Prints the provided message to the console with the prefix "INFO" in bold, if the MessengerUtils's verbose flag is set to true.
    PrintError(message string, err error): Formats and prints an error message to the console with a bold "ERROR:" label.

## License

This package is released under the MIT License. See the [LICENSE](LICENSE) file for more information.
