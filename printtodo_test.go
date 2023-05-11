package messengerutils

import (
	"bytes"
	"os"
	"testing"
	"time"
)

func TestPrintToDo(t *testing.T) {
	testCases := []struct {
		name           string
		message        []interface{}
		expectedOutput string
	}{
		{
			name:           "Test with a single string message",
			message:        []interface{}{"Finish the report!"},
			expectedOutput: "\033[1mTODO:\033[0m Finish the report!\n",
		},
		{
			name:           "Test with multiple string messages",
			message:        []interface{}{"Remember,", "buy milk!"},
			expectedOutput: "\033[1mTODO:\033[0m Remember, buy milk!\n",
		},
		{
			name:           "Test with a single integer message",
			message:        []interface{}{42},
			expectedOutput: "\033[1mTODO:\033[0m 42\n",
		},
		{
			name:           "Test with a single time.Time message",
			message:        []interface{}{time.Date(2023, 5, 5, 12, 0, 0, 0, time.UTC)},
			expectedOutput: "\033[1mTODO:\033[0m 2023-05-05T12:00:00Z\n",
		},
		{
			name: "Test with a mix of message types",
			message: []interface{}{
				"Complete task",
				42,
				"by",
				time.Date(2023, 5, 5, 12, 0, 0, 0, time.UTC),
			},
			expectedOutput: "\033[1mTODO:\033[0m Complete task 42 by 2023-05-05T12:00:00Z\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			PrintToDo(tc.message...)

			w.Close()
			os.Stdout = oldStdout
			var buf bytes.Buffer
			buf.ReadFrom(r)
			got := buf.String()

			if got != tc.expectedOutput {
				t.Errorf("expected: %s, got: %s", tc.expectedOutput, buf.String())
			}
		})
	}
}
