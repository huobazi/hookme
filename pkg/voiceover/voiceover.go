package voiceover

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

type VoiceOver interface {
	Say(...interface{})
	Sayf(string, ...interface{})

	WithError(error) VoiceOver
	WithField(string, interface{}) VoiceOver
	WithFields(map[string]interface{}) VoiceOver
}

var vo = newVoiceOver()

func newVoiceOver() VoiceOver {
	l := logrus.New()
	var logLevel = logrus.InfoLevel
	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   "logs/hookme.log",
		MaxSize:    50,  // megabytes
		MaxBackups: 10,  // amount
		MaxAge:     100, //days
		Level:      logLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: time.RFC822,
		},
	})

	if err != nil {
		logrus.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	l.SetOutput(os.Stdout)
	l.SetLevel(logLevel)
	l.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})

	l.AddHook(rotateFileHook)

	return &logrusVoiceOver{log: l}
}

func Say(args ...interface{}) {
	vo.Say(args...)
}

func Sayf(format string, args ...interface{}) {
	vo.Sayf(format, args...)
}

func WithError(err error) VoiceOver {
	return vo.WithError(err)
}

func WithField(key string, value interface{}) VoiceOver {
	return vo.WithField(key, value)
}

func WithFields(fields map[string]interface{}) VoiceOver {
	return vo.WithFields(fields)
}
