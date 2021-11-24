/*
	The logger package extends the logrus.Logger with helpful reuseable functions
*/
package logs

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/sirupsen/logrus"
)

// Internal variable to store all loggers
var loggers = make(map[string]*logrus.Logger)

// GetLogger always gets the default logger, otherwise, it initializes and adds the logger
func GetLogger() *logrus.Logger {
	// return logger if it exists
	if logger, ok := GetLoggerByName(defaultLoggerName); ok {
		return logger
	}

	// create the logger if it doesn't exist (panic on error)
	c := NewConfiguration(defaultLoggerName, "", "")
	logger, err := InitLoggerWithFileOutput(c)
	if err != nil {
		panic(fmt.Errorf("Unable to initialize logger while trying to get the logger: %+v", err))
	}
	return logger
}

// GetLoggerByName tries to get a logger by a given name
func GetLoggerByName(name string) (*logrus.Logger, bool) {
	logger, ok := loggers[name]
	return logger, ok
}

// InitLoggerWithFileOutput initializes a logger for a given configuration
func InitLoggerWithFileOutput(c *Configuration) (*logrus.Logger, error) {
	var logger *logrus.Logger
	var ok bool
	// check if the logger already exists, ignore the initialization if it already exists
	if logger, ok = GetLoggerByName(c.Name); ok {
		return logger, nil
	}

	// create a new logger
	logger = logrus.New()

	// set formatting for logger
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	// setup the file output
	fileout, err := os.OpenFile(c.Path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, fs.FileMode(c.Permissions))
	if err != nil {
		return nil, err
	}
	logger.SetOutput(fileout)

	// set logging level
	level, err := logrus.ParseLevel(c.Level)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(level)

	// add logger to the map
	loggers[c.Name] = logger

	// return the logger
	return logger, nil
}
