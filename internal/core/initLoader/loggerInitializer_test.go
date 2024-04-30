package initLoader

import (
	"github.com/glieske/kube-nanny/internal/core/helpers"
	"github.com/glieske/kube-nanny/internal/dto"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestLoadLoggerInitHandler_execute(t *testing.T) {
	type args struct {
		cfg              dto.MainConfig
		expectedLogLevel log.Level
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "logger_debug", args: args{cfg: dto.MainConfig{LoggingLevel: helpers.StringPtr("debug")}, expectedLogLevel: log.DebugLevel}},
		{name: "logger_info", args: args{cfg: dto.MainConfig{LoggingLevel: helpers.StringPtr("info")}, expectedLogLevel: log.InfoLevel}},
		{name: "logger_warn", args: args{cfg: dto.MainConfig{LoggingLevel: helpers.StringPtr("warn")}, expectedLogLevel: log.WarnLevel}},
		{name: "logger_error", args: args{cfg: dto.MainConfig{LoggingLevel: helpers.StringPtr("error")}, expectedLogLevel: log.ErrorLevel}},
		{name: "logger_other", args: args{cfg: dto.MainConfig{LoggingLevel: helpers.StringPtr("other")}, expectedLogLevel: log.InfoLevel}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := LoadLoggerInitHandler{
				next: nil,
			}
			s.execute(&tt.args.cfg)
			if log.GetLevel() != tt.args.expectedLogLevel {
				t.Errorf("Expected log level to be %v, got %v", tt.args.expectedLogLevel, log.GetLevel())
			}
		})
	}
}

func TestLoadLoggerInitHandler_setNext(t *testing.T) {
	handler := LoadLoggerInitHandler{
		next: nil,
	}

	nextHandler := LoadLoggerInitHandler{
		next: nil,
	}

	handler.setNext(&nextHandler)
	if handler.next != &nextHandler {
		t.Errorf("Expected next handler to be set, got %v", handler.next)
	}
}

func TestLoadLoggingConfiguration(t *testing.T) {
	cfg := &dto.MainConfig{
		LoggingLevel: new(string),
	}

	*cfg.LoggingLevel = "debug"
	loadLoggingConfiguration(cfg)
	if log.GetLevel() != log.DebugLevel {
		t.Errorf("Expected log level to be Debug, got %v", log.GetLevel())
	}
}
