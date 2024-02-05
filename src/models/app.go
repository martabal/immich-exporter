package models

type TypeAppConfig struct {
	Port     int
	Error    bool
	LogLevel string
}

var AppConfig TypeAppConfig

func SetApp(setport int, seterror bool, setloglevel string) {
	AppConfig = TypeAppConfig{
		Port:     setport,
		Error:    seterror,
		LogLevel: setloglevel,
	}
}

func GetPort() int {
	return AppConfig.Port
}

func SetPromptError(prompt bool) {
	AppConfig.Error = prompt
}

func GetPromptError() bool {
	return AppConfig.Error
}

func GetLogLevel() string {
	return AppConfig.LogLevel
}
