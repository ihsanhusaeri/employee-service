package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func openLogFile(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {
		return nil, err
	}
	return file, nil
}

//MakeLogEntry to make log format
func MakeLogEntry(e echo.Context) *log.Entry {
	// Custom Formater
	customFormater := new(log.TextFormatter)
	customFormater.DisableTimestamp = false
	customFormater.TimestampFormat = "[02-01-2006][15:04:05]"
	customFormater.ForceColors = true

	// Write log into a file, and change filename everyday
	filename := fmt.Sprintf("logs/%s_log_employee-service.txt", time.Now().Format("02-01-2006"))
	logFile, err := openLogFile(filename)
	if err != nil {
		return log.WithFields(log.Fields{
			"WARNING": err,
		})
	}

	// Write multiple writer into printout and to a file
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Set all the misc
	log.SetFormatter(customFormater)
	log.SetOutput(multiWriter)

	if e == nil {
		return log.WithFields(log.Fields{
			"Time": time.Now().Format("[02-01-2006][15:04:05]"),
		})
	}

	return log.WithFields(log.Fields{
		"Method":   e.Request().Method,
		"URL":      e.Request().URL.String(),
		"IP":       e.RealIP(),
		"Lenght":   e.Request().ContentLength,
		"Protocol": e.Request().Proto,
		"Type":     e.Request().Header.Get("Content-Type"),
		// "Time":   time.Now().Format("[02-01-2006][15:04:05]"),
	})
}
