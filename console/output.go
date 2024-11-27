package console

import "io"

type OutputMode interface {
	Has(OutputMode) bool
	Append(OutputMode) OutputMode
	Remove(OutputMode) OutputMode
}

type Output interface {
	io.Writer
	io.StringWriter
	GetMode() OutputMode
	WithMode(mode OutputMode) Output
	Debug(s string)
	Debugf(format string, a ...any)
	Info(s string)
	Infof(format string, a ...any)
	Notice(s string)
	Noticef(format string, a ...any)
	Warn(s string)
	Warnf(format string, a ...any)
	Error(s string)
	Errorf(format string, a ...any)
	Success(s string)
	Successf(format string, a ...any)
	Fail(s string)
	Failf(format string, a ...any)
}
