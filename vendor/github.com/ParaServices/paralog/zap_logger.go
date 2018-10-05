// Package paralog ...
package paralog

import (
	"errors"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// AppEnv ...
type AppEnv interface {
	String() string
	Development() bool
}

type appEnv string

// String ...
func (a appEnv) String() string {
	return string(a)
}

func (a appEnv) Development() bool {

	if string(a) == "development" {
		return true

	}
	return false
}

const (
	// DevelopmentEnv ...
	DevelopmentEnv appEnv = "development"
	// TestEnv ...
	TestEnv appEnv = "test"
	// StagingEnv ...
	StagingEnv appEnv = "staging"
	// ProductionEnv ...
	ProductionEnv appEnv = "production"
)

type logger struct {
	zl *zap.Logger
}

// Logger ...
type Logger interface {
	Debug(string, ...zapcore.Field)
	Info(string, ...zapcore.Field)
	Warn(string, ...zapcore.Field)
	Error(string, ...zapcore.Field)
	Fatal(string, ...zapcore.Field)
}

// ZapConfig ...
type ZapConfig struct {
	AppName        string
	AppEnv         AppEnv
	AppRevision    string
	AppRevisionTag string
	Hostname       string
	PID            int
	Config         *zap.Config
}

// NewZapLogger ...
func NewZapLogger(zc *ZapConfig) (Logger, error) {
	if zc.AppEnv == nil {
		return nil, errors.New("AppEnv must be set to use this logger")
	}

	if zc.Hostname == "" {
		var err error
		zc.Hostname, err = os.Hostname()
		if err != nil {
			return nil, err
		}
	}

	if zc.PID == 0 {
		zc.PID = os.Getpid()
	}

	context := []zapcore.Field{
		zap.String("app_name", zc.AppName),
		zap.String("hostname", zc.Hostname),
		zap.String("app_revision", zc.AppRevision),
		zap.String("app_revision_tag", zc.AppRevisionTag),
		zap.String("env", zc.AppEnv.String()),
		zap.Int("PID", zc.PID),
	}

	if zc.Config == nil {
		zc.Config = &zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
			Development: zc.AppEnv.Development(),
			Sampling: &zap.SamplingConfig{
				Initial:    0,
				Thereafter: 1,
			},
			Encoding:         "json",
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:     "message",
				LevelKey:       "level",
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeDuration: zapcore.NanosDurationEncoder,
			},
		}
	}

	var zl *zap.Logger
	zl, zErr := zc.Config.Build(zap.Fields(context...))
	if zErr != nil {
		return nil, zErr
	}

	return &logger{zl: zl}, nil
}

// Debug logs the given message that ends with an os.Exit(1). If a sentry
// client was set it will report the given message prior to logging and
// subsequently calling os.Exit(1)
func (l *logger) Debug(message string, fields ...zapcore.Field) {
	l.zl.Debug(message, fields...)
}

// Info logs the given message and fields with a log level of Warn and sends
// the given message and fields to sentry if a sentry client is set.
func (l *logger) Info(message string, fields ...zapcore.Field) {
	l.zl.Info(message, fields...)
}

// Warn logs the given message and fields with a log level of Warn and sends
// the given message and fields to sentry if a sentry client is set.
func (l *logger) Warn(message string, fields ...zapcore.Field) {
	l.zl.Warn(message, fields...)
}

// Error logs the given message that ends with an os.Exit(1). If a sentry
// client was set it will report the given message prior to logging and
// subsequently calling os.Exit(1)
func (l *logger) Error(message string, fields ...zapcore.Field) {
	l.zl.Error(message, fields...)
}

// Fatal logs the given message that ends with an os.Exit(1). If a sentry
// client was set it will report the given message prior to logging and
// subsequently calling os.Exit(1)
func (l *logger) Fatal(message string, fields ...zapcore.Field) {
	l.zl.Fatal(message, fields...)
}
