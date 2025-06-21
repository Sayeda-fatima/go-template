package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type MyLogger struct {
	zerolog.Logger
}

var Logger MyLogger

func NewLogger(env string) MyLogger {
	var output zerolog.LevelWriter

	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Failed to get current directory: %v", err))
	}

	// Create log file path in current directory
	logDir := filepath.Join(currentDir, "logs")
	logFilePath := filepath.Join(logDir, "app.log")

	// Ensure the log directory exists
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create log directory: %v", err))
	}

	// Open or create the log file
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %v", err))
	}

	// Use ConsoleWriter in dev for better readability
	if env == "dev" {
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
		consoleWriter.FormatLevel = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("|%6s|", i))
		}
		consoleWriter.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s:", i)
		}
		consoleWriter.FormatFieldValue = func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		}
		output = zerolog.MultiLevelWriter(consoleWriter, logFile)
	} else {
		output = zerolog.MultiLevelWriter(os.Stdout, logFile)
	}

	// Create the logger
	logger := zerolog.New(output).
		With().
		Caller().
		Timestamp().
		Logger()

	// Set log level based on environment
	switch env {
	case "prod":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "stage":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default: // dev
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	Logger = MyLogger{logger}
	return Logger
}

func (l *MyLogger) LogInfo() *zerolog.Event {
	return l.Logger.Info()
}
func (l *MyLogger) LogError() *zerolog.Event {
	return l.Logger.Error()
}

func (l *MyLogger) LogDebug() *zerolog.Event {
	return l.Logger.Debug()
}

func (l *MyLogger) LogWarn() *zerolog.Event {
	return l.Logger.Warn()
}

func (l *MyLogger) LogFatal() *zerolog.Event {
	return l.Logger.Fatal()
}
