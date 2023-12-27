package osenv // import "github.com/NehanSikder/osenv"

Package osenv follows the GO flag package pattern of letting users set a default
value when getting values of environment variables It provides helper methods
to convert the env variable values to int, string or bool It also an interface
for users to provide custom logic to convert the env variable value If the env
variable is not set, is empty or the conversion code fails for any reason,
then the functions return the default value

FUNCTIONS

func Get(envKey string, c Converter) any
    Get takes an string env variable key and a Converter interface, extracts the
    value using the env variable key, converts it using the provided Converter
    implementation and returns it If the extracted value is empty, has not
    been set or the Converter fails to convert, then it returns the Converters
    default value Get has a return type of any and the Converter implementation
    determines what type the env variable value is converted and returned as

func GetBool(envKey string, defaultValue bool) bool
    GetBool takes an string env variable key and a bool default value It gets
    the env value associated with the provided env variable key and returns the
    value as a bool It considers the following strings as true: 1, t, T, TRUE,
    true, True It considers the following strings as false: 0, f, F, FALSE,
    false, False It returns the default value for any other string If the env
    variable is an empty string, has not been set, or the conversion fails then
    it returns the provided int default value

func GetInt(envKey string, defaultValue int) int
    GetInt takes an string env variable key and an int default value It gets
    the env value associated with the provided env variable key and returns the
    value as an int If the env variable is an empty string, has not been set,
    or the conversion fails then it returns the provided int default value

func GetString(envKey string, defaultValue string) string
    GetString takes an string env variable key and a string default value It
    gets the env value associated with the provided env variable key and returns
    the value as a string If the env variable is an empty string or has not been
    set then returns the provided string default value


TYPES

type BoolConverter struct {
	DefaultValue bool
}
    BoolConverter converts extracted env variable to a bool if possible and
    returns the default bool value otherwise

func (b BoolConverter) Convert(s string) (interface{}, error)
    Convert converts the provided string parameter to a bool and returns the
    bool value If the conversion fails, then it returns the default bool
    value that is set when the BoolConverter is instantiated It considers
    the following strings as true: 1, t, T, TRUE, true, True It considers the
    following strings as false: 0, f, F, FALSE, false, False It returns the
    default value for any other string

func (b BoolConverter) GetDefaultValue() interface{}
    GetDefaultValue returns the default bool value that is set when the
    BoolConverter is instantiated

type Converter interface {
	GetDefaultValue() interface{}
	Convert(string) (interface{}, error)
}
    A Converter converts the extract env variable value to a specific return
    type Users can implement their own custom Converters to convert the env
    variables to their needs similar to whats available in the GO flag package
    See the details of the IntConverter or the BoolConverter for details on how
    to implement your own custom Converter

type IntConverter struct {
	DefaultValue int
}
    IntConverter converts extracted env variable to an int if possible and
    returns the default int value otherwise

func (i IntConverter) Convert(s string) (interface{}, error)
    Convert converts the provided string parameter to an int and returns the int
    value If the conversion fails, then it returns the default int value that is
    set when the IntConverter is instantiated

func (i IntConverter) GetDefaultValue() interface{}
    GetDefaultValue returns the default int value that is set when the
    IntConverter is instantiated

