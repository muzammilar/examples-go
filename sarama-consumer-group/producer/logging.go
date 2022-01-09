// The module producer contains the code for a sample data producer

package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

// InitLoggerWithFileOutput initializes a logger for a given configuration
func InitLoggerWithStdOut() *logrus.Logger {

	// create a new logger
	var logger *logrus.Logger = logrus.New()

	// set formatting for logger
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	// add caller function name (might be somewhat expensive)
	logger.SetReportCaller(true)

	// Output to stdout instead of the default stderr
	logger.SetOutput(os.Stdout)

	// set logging level
	level, err := logrus.ParseLevel(loglevel)
	if err != nil {
		panic(err)
	}
	logger.SetLevel(level)

	// return the logger
	return logger
}
