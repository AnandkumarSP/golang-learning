package logger

import (
	"encoding/json"
	"fmt"
	"log"
)

// Info prints message with [INFO] prefix.
func Info(args ...interface{}) {
	log.Printf("[INFO] %s", fmt.Sprint(args...))
}

// Infof formats message prints it with [INFO] prefix.
func Infof(format string, args ...interface{}) {
	log.Printf("[INFO] %s", fmt.Sprintf(format, args...))
}

// Warn prints message with [WARN] prefix.
func Warn(args ...interface{}) {
	log.Printf("[WARN] %s", fmt.Sprint(args...))
}

// Warnf formats message prints it with [WARN] prefix.
func Warnf(format string, args ...interface{}) {
	log.Printf("[WARN] %s", fmt.Sprintf(format, args...))
}

// Error prints message with [ERR] prefix.
func Error(args ...interface{}) {
	log.Printf("[ERR] %s", fmt.Sprint(args...))
}

// Errorf formats message prints it with [ERR] prefix.
func Errorf(format string, args ...interface{}) {
	log.Printf("[ERR] %s", fmt.Sprintf(format, args...))
}

// LogObj pretty prints the Obj.
func LogObj(obj interface{}) {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		Error(err)
	}
	Info(string(b))
}
