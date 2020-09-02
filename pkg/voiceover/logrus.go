package voiceover

import (
	"github.com/sirupsen/logrus"
)

type logrusVoiceOver struct{ log *logrus.Logger }

func (l *logrusVoiceOver) Say(args ...interface{}) {
	l.log.Println(args...)
}

func (l *logrusVoiceOver) Sayf(format string, args ...interface{}) {
	l.log.Printf(format, args...)
}

func (l *logrusVoiceOver) WithError(err error) VoiceOver {
	return (*logrusEntry)(l.log.WithError(err))
}

func (l *logrusVoiceOver) WithField(key string, value interface{}) VoiceOver {
	return (*logrusEntry)(l.log.WithField(key, value))
}

func (l *logrusVoiceOver) WithFields(fields map[string]interface{}) VoiceOver {
	return (*logrusEntry)(l.log.WithFields(fields))
}

type logrusEntry logrus.Entry

func (e *logrusEntry) Say(args ...interface{}) {
	(*logrus.Entry)(e).Println(args...)
}

func (e *logrusEntry) Sayf(format string, args ...interface{}) {
	(*logrus.Entry)(e).Printf(format, args...)
}

func (e *logrusEntry) WithError(err error) VoiceOver {
	return (*logrusEntry)((*logrus.Entry)(e).WithError(err))
}

func (e *logrusEntry) WithField(key string, value interface{}) VoiceOver {
	return (*logrusEntry)((*logrus.Entry)(e).WithField(key, value))
}

func (e *logrusEntry) WithFields(fields map[string]interface{}) VoiceOver {
	return (*logrusEntry)((*logrus.Entry)(e).WithFields(fields))
}
