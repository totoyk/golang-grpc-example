package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func NewLogger(logLevel string) (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
	logger.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, fmt.Errorf("failed to parse log level: %w", err)
	}
	logger.SetLevel(level)

	return logger, nil
}
