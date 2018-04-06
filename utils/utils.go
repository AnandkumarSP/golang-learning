package utils

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"../logger"
)

// SingleChannelConfig defines the config of a single channel
type SingleChannelConfig struct {
	StepsSequence []struct {
		StepName    string
		InputConfig []struct {
			Name          string
			DataType      string
			Input         string
			Default       string
			CanUseDefault bool
		}
	}
}

// configuration defines the uploader config.
type configuration struct {
	Channels             map[string]SingleChannelConfig
	FirstStepInputConfig []struct {
		Name          string
		DataType      string
		Default       string
		CanUseDefault bool
	}
}

// Config global variable to hold the uploader configuration.
var Config configuration

// pluginsConfiguration defines the pluginConfig.
type pluginsConfiguration struct {
	CanBeFirst        bool
	ShouldBeFirstOnly bool
	Desc              string
	InputConfig       []struct {
		Name          string
		DataType      string
		Default       string
		CanUseDefault bool
	}
	OutputConfig []struct {
		Name     string
		DataType string
	}
}

// PluginsConfig global variable to hold the plugins configuration.
var PluginsConfig map[string]pluginsConfiguration

// GenericErrorHandler is for handling all sort of errors.
// NOTE: For specific error handling create one.
func GenericErrorHandler(err error) {
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	if r := recover(); r != nil {
		logger.Error(r)
		panic(r)
	}
}

// SafeErrorHandler is for handling all sort of errors.
// NOTE: For specific error handling create one.
func SafeErrorHandler(err error) {
	if err != nil {
		logger.Error(err)
	}

	if r := recover(); r != nil {
		logger.Error(r)
	}
}

// NewCustomError returns an error that formats as the given text.
func NewCustomError(text string, code int) error {
	return &CustomError{text, code}
}

// CustomError is a trivial implementation of error.
type CustomError struct {
	msg  string
	code int
}

// Error return the message.
func (e *CustomError) Error() string {
	return e.msg
}

// ErrorCode return the HTTP status code.
func (e *CustomError) ErrorCode() int {
	return e.code
}

func UpdateConfig(obj interface{}) (err error) {
	defer GenericErrorHandler(nil)
	jsonString, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile("config.json", jsonString, 0664)
	if err != nil {
		return
	}
	RefreshUploaderConfig()
	return
}

// RefreshUploaderConfig returns the uploader configuration.
func RefreshUploaderConfig() {
	defer GenericErrorHandler(nil)
	data, err := ioutil.ReadFile("config.json")
	GenericErrorHandler(err)

	Config = configuration{}
	json.Unmarshal(data, &Config)
	return
}

// RefreshPluginsConfig returns the uploader configuration.
func RefreshPluginsConfig() {
	defer GenericErrorHandler(nil)
	data, err := ioutil.ReadFile("pluginsConfig.json")
	GenericErrorHandler(err)

	PluginsConfig = make(map[string]pluginsConfiguration, 0)
	json.Unmarshal(data, &PluginsConfig)
	return
}

// ConvertStrToDatatype converts to any string value given dataType format
func ConvertStrToDatatype(val string, dataType string) (interface{}, error) {
	switch dataType {
	case "int":
		return StrToInt(val)
	case "string":
		return val, nil
	default:
		panic("Unsupported datatype conversion")
	}
}

// StrToInt converts string to int
func StrToInt(val string) (int, error) {
	return strconv.Atoi(val)
}
