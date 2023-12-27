package osenv

import (
	"os"
	"testing"
)

func TestGetInt(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		defaultValue   int
		expectedOutput int
	}{
		{
			name:           "empty",
			input:          "",
			defaultValue:   0,
			expectedOutput: 0,
		}, {
			name:           "env_var_has_value",
			input:          "10",
			defaultValue:   1,
			expectedOutput: 10,
		}, {
			name:           "env_var_has_value",
			input:          "hello world",
			defaultValue:   1,
			expectedOutput: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := "test"
			os.Setenv(key, tt.input)
			actualOutput := GetInt(key, tt.defaultValue)
			os.Unsetenv(key)
			if actualOutput != tt.expectedOutput {
				t.Errorf("want %d; got %d", tt.expectedOutput, actualOutput)
			}
		})
	}
}

func TestGetString(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		defaultValue   string
		expectedOutput string
	}{
		{
			name:           "empty",
			input:          "",
			defaultValue:   "hello",
			expectedOutput: "hello",
		},
		{
			name:           "env_var_has_value",
			input:          "hello",
			defaultValue:   "world",
			expectedOutput: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := "test"
			os.Setenv(key, tt.input)
			actualOutput := GetString(key, tt.defaultValue)
			os.Unsetenv(key)
			if actualOutput != tt.expectedOutput {
				t.Errorf("want %s; got %s", tt.expectedOutput, actualOutput)
			}
		})
	}
}

func TestGetBool(t *testing.T) {

	// DEFINE TEST DATA

	tests := []struct {
		name           string
		input          string
		defaultValue   bool
		expectedOutput bool
	}{
		{
			name:           "true",
			input:          "true",
			defaultValue:   false,
			expectedOutput: true,
		},
		{
			name:           "True",
			input:          "True",
			defaultValue:   false,
			expectedOutput: true,
		},
		{
			name:           "false",
			input:          "false",
			defaultValue:   true,
			expectedOutput: false,
		},
		{
			name:           "False",
			input:          "False",
			defaultValue:   true,
			expectedOutput: false,
		},
		{
			name:           "hello",
			input:          "hello",
			defaultValue:   true,
			expectedOutput: true,
		},
		{
			name:           "1",
			input:          "1",
			defaultValue:   false,
			expectedOutput: true,
		},
		{
			name:           "0",
			input:          "0",
			defaultValue:   true,
			expectedOutput: false,
		},
		{
			name:           "Yes",
			input:          "Yes",
			defaultValue:   false,
			expectedOutput: false,
		},
		{
			name:           "No",
			input:          "No",
			defaultValue:   true,
			expectedOutput: true,
		},
		{
			name:           "1.0",
			input:          "1.0",
			defaultValue:   false,
			expectedOutput: false,
		},
		{
			name:           "0.0",
			input:          "0.0",
			defaultValue:   true,
			expectedOutput: true,
		},
		{
			name:           "empty",
			input:          "",
			defaultValue:   false,
			expectedOutput: false,
		},
	}

	// RUN TESTS USING TEST DATA
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := "test"
			os.Setenv(key, tt.input)
			actualOutput := GetBool(key, tt.defaultValue)
			os.Unsetenv(key)
			if actualOutput != tt.expectedOutput {
				t.Errorf("want %t; got %t", tt.expectedOutput, actualOutput)
			}
		})
	}
}
