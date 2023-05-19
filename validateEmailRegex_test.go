package messengerutils

import (
	"testing"
)

func TestValidateEmailRegex(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{
			name:  "Test valid email",
			email: "test@example.com",
			want:  true,
		},
		{
			name:  "Test valid email with numbers",
			email: "123test@456example.com",
			want:  true,
		},
		{
			name:  "Test valid email with special characters",
			email: "test.email@sub.example.com",
			want:  true,
		},
		{
			name:  "Test valid email with domain ending number",
			email: "test@email1.com",
			want:  true,
		},
		{
			name:  "Test valid email with uppercase letters",
			email: "TEST@Example.com",
			want:  true,
		},
		{
			name:  "Test valid email with +",
			email: "TEST+google.com@Example.com",
			want:  true,
		},
		{
			name:  "Test email without @ symbol",
			email: "testexample.com",
			want:  false,
		},
		{
			name:  "Test email without domain",
			email: "test@",
			want:  false,
		},
		{
			name:  "Test email with multiple @ symbols",
			email: "test@@example.com",
			want:  false,
		},
		{
			name:  "Test email with space",
			email: "test @example.com",
			want:  false,
		},
		{
			name:  "Test empty email",
			email: "",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateEmailRegex(tt.email); got != tt.want {
				t.Errorf("validateEmailRegex() = %v, want %v", got, tt.want)
			}
		})
	}
}
