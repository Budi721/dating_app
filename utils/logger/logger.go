package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var (
	Log *zerolog.Logger
)

func NewLogger() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	Log = &logger
}
