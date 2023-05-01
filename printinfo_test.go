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
		m := &MessengerUtils{verbose: verbose}

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
