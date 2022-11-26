package tools

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

type YourLogger struct {
	log *zerolog.Logger
}

func SetupLogger() (*YourLogger, error) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	log := zerolog.New(output).With().Timestamp().Logger()
	return &YourLogger{
		log: &log,
	}, nil
}

func (yl *YourLogger) Log() *zerolog.Logger {
	return yl.log
}
