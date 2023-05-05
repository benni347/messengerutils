package messengerutils

import (
	"bytes"
	"os"
	"testing"
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
	// Test cases
	tests := []struct {
		name     string
		verbose  bool
		message  []string
		expected string
	}{
		{
			name:     "verbose_true_single_string",
			verbose:  true,
			message:  []string{"test"},
			expected: "\033[1mINFO\033[0m: test\n",
		},
		{
			name:     "verbose_true_multiple_strings",
			verbose:  true,
			message:  []string{"test", " message"},
			expected: "\033[1mINFO\033[0m: test message\n",
		},
		{
			name:     "verbose_false_single_string",
			verbose:  false,
			message:  []string{"test"},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Prepare the MessengerUtils
			messenger := &MessengerUtils{
				Verbose: test.verbose,
			}

			// Capture output
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Execute the function
			messenger.PrintInfo(test.message...)

			// Restore os.Stdout and read the captured output
			w.Close()
			os.Stdout = oldStdout
			var buf bytes.Buffer
			buf.ReadFrom(r)
			got := buf.String()

			// Check the result
			if got != test.expected {
				t.Errorf("Expected: %q, got: %q", test.expected, got)
			}
		})
	}
}
