# Multiple Logger Wrapper for Self Use

This repository provides a simple logger wrapper for personal use in Go projects. The logger wrapper supports multiple logging libraries, including Zerolog, Slog, Logrus, and ZapLog, allowing you to easily switch between different loggers based on your requirements.

## Features
- Easy to set up and use.
- Supports multiple logging libraries: Zerolog, Slog, Logrus, and ZapLog.
- Provides a consistent interface for logging across different libraries.

## Usage
To use the logger wrapper:

```
 go get github.com/Chris-Basson-Grundling/logger
```

The logger wrapper uses the following interface for consistency across different logging libraries:

```go
    type Logger interface {
        Debug(msg string, fields map[string]interface{})
        Info(msg string, fields map[string]interface{})
        Warn(msg string, fields map[string]interface{})
        Error(msg string, fields map[string]interface{})
        Fatal(msg string, fields map[string]interface{})
    }

```

Example of how to use the logger wrapper in your Go project:

```go
package main

import (
	"context"
	
	"github.com/Chris-Basson-Grundling/logger"
)

func main() {
	commonFields := map[string]interface{}{
		"userId":    "123456789",
		"ipAddress": "172.0.0.1",
	}
	ctx := context.WithValue(context.Background(), "commonFields", commonFields)

	logruslogger := logger.NewLoggerWrapper("logrus", ctx)
	logruslogger.Debug("This is an debug log message.", commonFields)

	sloglogger := logger.NewLoggerWrapper("slog", ctx)
	sloglogger.Info("This is an info log message.", commonFields)

	zaplogger := logger.NewLoggerWrapper("zap", ctx)
	zaplogger.Warn("This is an warn log message.", commonFields)

	zerologger := logger.NewLoggerWrapper("zerolog", ctx)
	zerologger.Error("This is an error log message.", commonFields)
}

```







