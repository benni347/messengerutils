package messengerutils

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func FuzzPrintError(f *testing.F) {
	// Add seed inputs

	f.Add("Test case 1", "Error 1")
	f.Add("Test case 2", "Error 2")
	f.Add("12345", "Number error")
	f.Add("Hello, world!", "String error")

	// Define the fuzz target function
	f.Fuzz(func(t *testing.T, message string, err string) {
		err1 := errors.New(err)
		// Call the PrintError function with the fuzzed inputs
		// Capture output using os.Pipe()
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		PrintError(message, err1)

		// Restore os.Stdout and read the captured output
		w.Close()
		os.Stdout = oldStdout
		var buf bytes.Buffer
		buf.ReadFrom(r)
		got := buf.String()

		if got != "\033[1mERROR:\033[0m "+message+": "+err1.Error()+"\n" {
			t.Errorf(
				"PrintError() got = %q, want = %q",
				got,
				"\033[1mERROR:\033[0m "+message+": "+err1.Error()+"\n",
			)
		} else if got == "" {
			t.Errorf("PrintError() got = %q, want non-empty string", got)
		}
	})
}
