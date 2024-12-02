package logging

import (
	"log/slog"
	"os"

	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/logging/lf"
)

// TODO: inject config

func NewSLogger() *SLogger {
	return &SLogger{
		logger: slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			})),
	}
}

type SLogger struct {
	logger *slog.Logger
}

func (s *SLogger) WithFields(fields ...lf.Field) Logger {
	return &SLogger{
		logger: s.logger.With(s.parseFields(fields)...),
	}
}

func (s *SLogger) parseFields(fields []lf.Field) []any {
	ret := make([]any, 0, len(fields)*2)
	for _, f := range fields {
		ret = append(ret, f.Key(), f.Value())
	}
	return ret
}

func (s *SLogger) Debug(message string, fields ...lf.Field) {
	s.logger.Debug(message, s.parseFields(fields)...)
}

func (s *SLogger) Info(message string, fields ...lf.Field) {
	s.logger.Info(message, s.parseFields(fields)...)
}

func (s *SLogger) Warning(message string, fields ...lf.Field) {
	s.logger.Warn(message, s.parseFields(fields)...)
}

func (s *SLogger) Error(message string, fields ...lf.Field) {
	s.logger.Error(message, s.parseFields(fields)...)
}
