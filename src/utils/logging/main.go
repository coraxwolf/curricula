/*
package logging provides a structured logging utility for the application and to track API Requests.
It uses Go's slog package to log messages in JSON format, making it easy to parse and analyze logs.
*/
package logging

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"
)

// AppLogger handles general application logging.
type AppLogger struct {
	logger *slog.Logger
	path   string
	writer *os.File
}

// APILogger handles logging API interactions.
// It tracks the requests made to the API and includes the Rate Limit Cost, Duration, and other relevant details.
type APILogger struct {
	logger *slog.Logger
	path   string
	writer *os.File
}

func NewAppLogger(path string) (*AppLogger, error) {
	var writers []io.Writer
	writers = append(writers, os.Stdout)
	// check if path exists
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// create the file
		err := os.MkdirAll(path, os.ModeAppend)
		if err != nil {
			return nil, err
		}
		slog.Default().Info("Log Directory Created", slog.String("path", path))
	}
	file, err := os.OpenFile(path+"/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	writers = append(writers, file)
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(writers...), &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true}))
	return &AppLogger{logger: logger, path: path, writer: file}, nil
}

func NewAPILogger(path string, console bool) (*APILogger, error) {
	var writers []io.Writer
	if console {
		slog.Default().Info("API Logger Console Output Enabled")
		writers = append(writers, os.Stdout)
	}
	// check if path exists
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// create the file
		err := os.MkdirAll(path, os.ModeAppend)
		if err != nil {
			return nil, err
		}
		slog.Default().Info("Log Directory Created", slog.String("path", path))
	}
	filename := fmt.Sprintf("%s-%s-%s-%s-%s-api.log", time.Now().Format("2006"), time.Now().Format("01"), time.Now().Format("02"), time.Now().Format("15"), time.Now().Format("04"))
	file, err := os.OpenFile(path+"/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	writers = append(writers, file)
	slog.Default().Info("API Log File Created", slog.String("path", path+"/"+filename))
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(writers...), &slog.HandlerOptions{Level: slog.LevelInfo, AddSource: true}))
	return &APILogger{logger: logger, path: path, writer: file}, nil
}

// App Logging Methods

// Info logs an informational message with optional details.
func (al *AppLogger) Info(msg string, details map[string]any) {
	al.logger.Info(msg, slog.Any("details", details))
}

// Warn logs a warning message with optional details.
func (al *AppLogger) Warn(msg string, details map[string]any) {
	al.logger.Warn(msg, slog.Any("details", details))
}

// Error logs an error message with the error and optional details.
func (al *AppLogger) Error(msg string, err error, details map[string]any) {
	al.logger.Error(msg, slog.String("error", err.Error()), slog.Any("details", details))
}

// Debug logs a debug message with optional details.
func (al *AppLogger) Debug(msg string, details map[string]any) {
	al.logger.Debug(msg, slog.Any("details", details))
}

// API Logging Methods

// Info logs the general details about an API reuests and the response.
func (al *APILogger) Info(url, method, status, msg string, rateCost, duration float64, details map[string]any) {
	al.logger.Info(msg,
		slog.String("url", url),
		slog.String("method", method),
		slog.String("status", status),
		slog.Float64("rateCost", rateCost),
		slog.Float64("duration", duration),
		slog.Any("details", details),
	)
}

// Warn logs any warnings generated from the API or Concerns that might be of importance to the user.
func (al *APILogger) Warn(url, method, status, msg string, rateCost, duration float64, details map[string]any) {
	al.logger.Warn(msg,
		slog.String("url", url),
		slog.String("method", method),
		slog.String("status", status),
		slog.Float64("rateCost", rateCost),
		slog.Float64("duration", duration),
		slog.Any("details", details),
	)
}

// Error logs any failures from the API.
func (al *APILogger) Error(url, method, status, msg string, rateCost, duration float64, error error) {
	al.logger.Error(msg,
		slog.String("url", url),
		slog.String("method", method),
		slog.String("status", status),
		slog.Float64("rateCost", rateCost),
		slog.Float64("duration", duration),
		slog.String("error", error.Error()),
	)
}
