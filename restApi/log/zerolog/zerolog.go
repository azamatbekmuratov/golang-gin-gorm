package zerolog

import (
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

const (
	serviceName = "serviceName"
)

type Zerolog struct {
	log         zerolog.Logger
	serviceName string
}

func NewLogger(serviceName string) *Zerolog {
	log := zerolog.New(customConsole()).
		With().
		Caller().
		Timestamp().
		Logger()

	zerolog.CallerMarshalFunc = customMarshalFunc()
	zerolog.CallerSkipFrameCount = 3
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	return &Zerolog{
		log:         log,
		serviceName: serviceName,
	}

}

func (z *Zerolog) Info(msg string) {
	z.log.Info().Str(serviceName, z.serviceName).Msg(msg)
}

func (z *Zerolog) Error(msg string, err error) {
	z.log.Error().Err(err).Str(serviceName, z.serviceName).Msg(msg)
}

func (z *Zerolog) Debug(msg string) {
	z.log.Debug().Str(serviceName, z.serviceName).Msg(msg)
}

func (z *Zerolog) Warn(msg string) {
	z.log.Warn().Str(serviceName, z.serviceName).Msg(msg)
}

func customConsole() zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		PartsOrder: customOrders(),
	}
}

func customOrders() []string {
	return []string{
		zerolog.TimestampFieldName,
		zerolog.LevelFieldName,
		zerolog.MessageFieldName,
		zerolog.CallerFieldName,
	}
}

func customMarshalFunc() func(file string, line int) string {
	return func(file string, line int) string {
		short := file

		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]

				break
			}
		}

		file = short

		return file + ":" + strconv.Itoa(line)
	}
}
