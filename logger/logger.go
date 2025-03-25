package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var log zerolog.Logger

// Init khởi tạo logger toàn cục
func Init(isDebug bool) {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		NoColor:    false,
	}

	// Định dạng outputs
	output.FormatLevel = func(i interface{}) string {
		return fmt.Sprintf("%-6s", i)
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("| %s |", i)
	}

	// Thiết lập log level
	level := zerolog.InfoLevel
	if isDebug {
		level = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(level)

	// Tạo logger với timestamp
	log = zerolog.New(output).
		With().
		Timestamp().
		Logger()
}

// Debug logs debug message
func Debug(msg string, args ...interface{}) {
	if len(args) > 0 {
		log.Debug().Fields(argsToMap(args...)).Msg(msg)
	} else {
		log.Debug().Msg(msg)
	}
}

// Info logs info message
func Info(msg string, args ...interface{}) {
	if len(args) > 0 {
		log.Info().Fields(argsToMap(args...)).Msg(msg)
	} else {
		log.Info().Msg(msg)
	}
}

// Warn logs warning message
func Warn(msg string, args ...interface{}) {
	if len(args) > 0 {
		log.Warn().Fields(argsToMap(args...)).Msg(msg)
	} else {
		log.Warn().Msg(msg)
	}
}

// Error logs error message
func Error(msg string, err error, args ...interface{}) {
	if err != nil {
		if len(args) > 0 {
			log.Error().Err(err).Fields(argsToMap(args...)).Msg(msg)
		} else {
			log.Error().Err(err).Msg(msg)
		}
	} else {
		if len(args) > 0 {
			log.Error().Fields(argsToMap(args...)).Msg(msg)
		} else {
			log.Error().Msg(msg)
		}
	}
}

// WithAccount tạo logger mới với context tài khoản
func WithAccount(username string) zerolog.Logger {
	return log.With().Str("account", username).Logger()
}

// AccountLogger tạo logger với context tài khoản
type AccountLogger struct {
	logger zerolog.Logger
}

// NewAccountLogger tạo account logger mới
func NewAccountLogger(username string) *AccountLogger {
	return &AccountLogger{
		logger: log.With().Str("account", username).Logger(),
	}
}

// Debug logs debug message with account context
func (l *AccountLogger) Debug(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.logger.Debug().Fields(argsToMap(args...)).Msg(msg)
	} else {
		l.logger.Debug().Msg(msg)
	}
}

// Info logs info message with account context
func (l *AccountLogger) Info(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.logger.Info().Fields(argsToMap(args...)).Msg(msg)
	} else {
		l.logger.Info().Msg(msg)
	}
}

// Warn logs warning message with account context
func (l *AccountLogger) Warn(msg string, args ...interface{}) {
	if len(args) > 0 {
		l.logger.Warn().Fields(argsToMap(args...)).Msg(msg)
	} else {
		l.logger.Warn().Msg(msg)
	}
}

// Error logs error message with account context
func (l *AccountLogger) Error(msg string, err error, args ...interface{}) {
	if err != nil {
		if len(args) > 0 {
			l.logger.Error().Err(err).Fields(argsToMap(args...)).Msg(msg)
		} else {
			l.logger.Error().Err(err).Msg(msg)
		}
	} else {
		if len(args) > 0 {
			l.logger.Error().Fields(argsToMap(args...)).Msg(msg)
		} else {
			l.logger.Error().Msg(msg)
		}
	}
}

// argsToMap chuyển đổi các tham số thành map
func argsToMap(args ...interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			m[args[i].(string)] = args[i+1]
		}
	}
	return m
}