package messengerutils

import (
	"testing"
	"time"
)

func TestFormatMessage(t *testing.T) {
	cases := []struct {
		in   []interface{}
		want string
	}{
		{[]interface{}{"Hello", "world"}, "Hello world"},
		{[]interface{}{"The", 42, "answer is"}, "The 42 answer is"},
		{[]interface{}{3.14}, "Unknown type: float64"},
		{[]interface{}{"Numbers:", 1, 2, 3, 4, 5}, "Numbers: 1 2 3 4 5"},
	}

	for _, c := range cases {
		got := formatMessage(c.in...)
		if got != c.want {
			t.Errorf("formatMessage(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFormatMessageWithTime(t *testing.T) {
	now := time.Now()
	in := []interface{}{"Current time:", now}
	want := "Current time: " + now.Format(time.RFC3339)

	got := formatMessage(in...)
	if got != want {
		t.Errorf("formatMessage(%q) == %q, want %q", in, got, want)
	}
}
