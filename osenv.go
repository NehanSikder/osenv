// Package osenv follows the GO flag package pattern of letting users set a default value when getting values of environment variables
// It provides helper methods to convert the env variable values to int, string or bool
// It also an interface for users to provide custom logic to convert the env variable value
// If the env variable is not set, is empty or the conversion code fails for any reason, then the functions return the default value
package osenv

import (
	"errors"
	"os"
	"strconv"
)

// A Converter converts the extract env variable value to a specific return type
// Users can implement their own custom Converters to convert the env variables to their needs similar to whats available in the GO flag package
// See the details of the IntConverter or the BoolConverter for details on how to implement your own custom Converter
type Converter interface {
	GetDefaultValue() interface{}
	Convert(string) (interface{}, error)
}

// IntConverter converts extracted env variable to an int if possible and returns the default int value otherwise
type IntConverter struct {
	DefaultValue int
}

// GetDefaultValue returns the default int value that is set when the IntConverter is instantiated
func (i IntConverter) GetDefaultValue() interface{} {
	return i.DefaultValue
}

// Convert converts the provided string parameter to an int and returns the int value
// If the conversion fails, then it returns the default int value that is set when the IntConverter is instantiated
func (i IntConverter) Convert(s string) (interface{}, error) {
	value, err := strconv.Atoi(s)
	if err != nil {
		return i.GetDefaultValue(), errors.Join(err)
	}
	return value, nil
}

// BoolConverter converts extracted env variable to a bool if possible and returns the default bool value otherwise
type BoolConverter struct {
	DefaultValue bool
}

// GetDefaultValue returns the default bool value that is set when the BoolConverter is instantiated
func (b BoolConverter) GetDefaultValue() interface{} {
	return b.DefaultValue
}

// Convert converts the provided string parameter to a bool and returns the bool value
// If the conversion fails, then it returns the default bool value that is set when the BoolConverter is instantiated
// It considers the following strings as true: 1, t, T, TRUE, true, True
// It considers the following strings as false: 0, f, F, FALSE, false, False
// It returns the default value for any other string
func (b BoolConverter) Convert(s string) (interface{}, error) {
	result, err := strconv.ParseBool(s)
	if err != nil {
		return b.GetDefaultValue(), errors.Join(err)
	}
	return result, nil
}

// Get takes an string env variable key and a Converter interface, extracts the value using the env variable key, converts it using the provided Converter implementation and returns it
// If the extracted value is empty, has not been set or the Converter fails to convert, then it returns the Converters default value
// Get has a return type of any and the Converter implementation determines what type the env variable value is converted and returned as
func Get(envKey string, c Converter) any {
	defaultValue := c.GetDefaultValue()
	envValue := os.Getenv(envKey)
	// no env value associated with key
	if envValue == "" {
		return defaultValue
	}
	// try to convert
	value, _ := c.Convert(envValue)
	return value
}

// GetString takes an string env variable key and a string default value
// It gets the env value associated with the provided env variable key and returns the value as a string
// If the env variable is an empty string or has not been set then returns the provided string default value
func GetString(envKey string, defaultValue string) string {
	envValue := os.Getenv(envKey)
	if envValue == "" {
		return defaultValue
	}
	return envValue
}

// GetInt takes an string env variable key and an int default value
// It gets the env value associated with the provided env variable key and returns the value as an int
// If the env variable is an empty string, has not been set, or the conversion fails then it returns the provided int default value
func GetInt(envKey string, defaultValue int) int {
	intConverter := IntConverter{
		DefaultValue: defaultValue,
	}
	result := Get(envKey, intConverter)
	i, _ := result.(int)
	return i
}

// GetBool takes an string env variable key and a bool default value
// It gets the env value associated with the provided env variable key and returns the value as a bool
// It considers the following strings as true: 1, t, T, TRUE, true, True
// It considers the following strings as false: 0, f, F, FALSE, false, False
// It returns the default value for any other string
// If the env variable is an empty string, has not been set, or the conversion fails then it returns the provided int default value
func GetBool(envKey string, defaultValue bool) bool {
	boolConverter := BoolConverter{
		DefaultValue: defaultValue,
	}
	result := Get(envKey, boolConverter)
	b, _ := result.(bool)
	return b

}
