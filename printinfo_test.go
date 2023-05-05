package messengerutils

import (
	"bytes"
	"os"
	"testing"
	"time"
)

func FuzzPrintInfo(f *testing.F) {
	f.Add(true, "Test message")
	f.Add(true, "-1230")
	f.Add(true, "1")
	f.Add(true, "123")
	f.Add(true, "123message")
	f.Add(true, "#fsdaf320")
	f.Add(false, "Test message")
	f.Add(false, "-1230")
	f.Add(false, "1")
	f.Add(false, "123")
	f.Add(false, "123message")
	f.Add(false, "#fsdaf320")

	f.Fuzz(func(t *testing.T, verbose bool, message string) {
		m := &MessengerUtils{Verbose: verbose}

		// Capture output using os.Pipe()
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		m.PrintInfo(message)

		// Restore os.Stdout and read the captured output
		w.Close()
		os.Stdout = oldStdout
		var buf bytes.Buffer
		buf.ReadFrom(r)
		got := buf.String()

		if verbose && got != "\033[1mINFO\033[0m: "+message+"\n" {
			t.Errorf("PrintInfo() got = %q, want = %q", got, "\033[1mINFO\033[0m: "+message+"\n")
		} else if !verbose && got != "" {
			t.Errorf("PrintInfo() got = %q, want empty string", got)
		}
	})
}

func TestPrintInfo(t *testing.T) {
	testCases := []struct {
		name           string
		verbose        bool
		message        []interface{}
		expectedOutput string
	}{
		{
			name:           "Test with a single string message",
			verbose:        true,
			message:        []interface{}{"Hello, world!"},
			expectedOutput: "\033[1mINFO\033[0m: Hello, world!\n",
		},
		{
			name:           "Test with multiple string messages",
			verbose:        true,
			message:        []interface{}{"Hello,", "world!"},
			expectedOutput: "\033[1mINFO\033[0m: Hello, world!\n",
		},
		{
			name:           "Test with a single integer message",
			verbose:        true,
			message:        []interface{}{42},
			expectedOutput: "\033[1mINFO\033[0m: 42\n",
		},
		{
			name:           "Test with a single time.Time message",
			verbose:        true,
			message:        []interface{}{time.Date(2023, 5, 5, 12, 0, 0, 0, time.UTC)},
			expectedOutput: "\033[1mINFO\033[0m: 2023-05-05T12:00:00Z\n",
		},
		{
			name:    "Test with a mix of message types",
			verbose: true,
			message: []interface{}{
				"The answer is",
				42,
				"at",
				time.Date(2023, 5, 5, 12, 0, 0, 0, time.UTC),
			},
			expectedOutput: "\033[1mINFO\033[0m: The answer is 42 at 2023-05-05T12:00:00Z\n",
		},
		{
			name:           "Test with verbose flag set to false",
			verbose:        false,
			message:        []interface{}{"This message should not be printed"},
			expectedOutput: "",
		},
		{
			name:           "Test with verbose flag set to true and message is empty",
			verbose:        true,
			message:        []interface{}{},
			expectedOutput: "\033[1mINFO\033[0m: \n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			m := &MessengerUtils{Verbose: tc.verbose}
			m.PrintInfo(tc.message...)

			w.Close()
			os.Stdout = oldStdout
			var buf bytes.Buffer
			buf.ReadFrom(r)
			got := buf.String()

			if got != tc.expectedOutput {
				t.Errorf("PrintInfo() = %q, want %q", got, tc.expectedOutput)
			}
		})
	}
}
