package logger

import (
	"log/slog"
	"os"

	"go.uber.org/fx/fxevent"
)

type Logger struct {
	*slog.Logger
}

func New() *Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(log)

	return &Logger{log}
}

func (l *Logger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.logOnStartExecuting(e)
	case *fxevent.OnStartExecuted:
		l.logOnStartExecuted(e)
	case *fxevent.OnStopExecuting:
		l.logOnStopExecuting(e)
	case *fxevent.OnStopExecuted:
		l.logOnStopExecuted(e)
	case *fxevent.Supplied:
		l.logSupplied(e)
	case *fxevent.Provided:
		l.logProvided(e)
	case *fxevent.Decorated:
		l.logDecorated(e)
	case *fxevent.Invoking:
		l.logInvoking(e)
	case *fxevent.Invoked:
		l.logInvoked(e)
	case *fxevent.Stopping:
		l.logStopping(e)
	case *fxevent.Stopped:
		l.logStopped(e)
	case *fxevent.RollingBack:
		l.logRollingBack(e)
	case *fxevent.RolledBack:
		l.logRolledBack(e)
	case *fxevent.Started:
		l.logStarted(e)
	case *fxevent.LoggerInitialized:
		l.logLoggerInitialized(e)
	}
}

func (l *Logger) logOnStartExecuting(e *fxevent.OnStartExecuting) {
	l.Debug("fx: hook OnStart executing", "callee", e.FunctionName, "caller", e.CallerName)
}

func (l *Logger) logOnStartExecuted(e *fxevent.OnStartExecuted) {
	if e.Err != nil {
		l.Error("fx: hook OnStart failed", "callee", e.FunctionName, "caller", e.CallerName, "error", e.Err)
	} else {
		l.Debug("fx: hook OnStart executed", "callee", e.FunctionName, "caller", e.CallerName, "runtime", e.Runtime.String())
	}
}

func (l *Logger) logOnStopExecuting(e *fxevent.OnStopExecuting) {
	l.Debug("fx: hook OnStop executing", "callee", e.FunctionName, "caller", e.CallerName)
}

func (l *Logger) logOnStopExecuted(e *fxevent.OnStopExecuted) {
	if e.Err != nil {
		l.Error("fx: hook OnStop failed", "callee", e.FunctionName, "caller", e.CallerName, "error", e.Err)
	} else {
		l.Debug("fx: hook OnStop executed", "callee", e.FunctionName, "caller", e.CallerName, "runtime", e.Runtime.String())
	}
}

func (l *Logger) logSupplied(e *fxevent.Supplied) {
	if e.Err != nil {
		l.Error("fx: supplied error", "type", e.TypeName, "error", e.Err)
	} else {
		l.Debug("fx: supplied", "type", e.TypeName)
	}
}

func (l *Logger) logProvided(e *fxevent.Provided) {
	if e.Err != nil {
		l.Error("fx: provided error", "constructor", e.ConstructorName, "error", e.Err)
	} else {
		l.Debug("fx: provided", "constructor", e.ConstructorName, "outputs", e.OutputTypeNames)
	}
}

func (l *Logger) logDecorated(e *fxevent.Decorated) {
	if e.Err != nil {
		l.Error("fx: decorated error", "decorator", e.DecoratorName, "error", e.Err)
	} else {
		l.Debug("fx: decorated", "decorator", e.DecoratorName, "outputs", e.OutputTypeNames)
	}
}

func (l *Logger) logInvoking(e *fxevent.Invoking) {
	l.Debug("fx: invoking", "function", e.FunctionName)
}

func (l *Logger) logInvoked(e *fxevent.Invoked) {
	if e.Err != nil {
		l.Error("fx: invoked error", "function", e.FunctionName, "error", e.Err)
	} else {
		l.Debug("fx: invoked", "function", e.FunctionName)
	}
}

func (l *Logger) logStopping(e *fxevent.Stopping) {
	l.Debug("fx: stopping", "signal", e.Signal.String())
}

func (l *Logger) logStopped(e *fxevent.Stopped) {
	if e.Err != nil {
		l.Error("fx: stopped with error", "error", e.Err)
	}
}

func (l *Logger) logRollingBack(e *fxevent.RollingBack) {
	l.Error("fx: rolling back", "error", e.StartErr)
}

func (l *Logger) logRolledBack(e *fxevent.RolledBack) {
	if e.Err != nil {
		l.Error("fx: rollback error", "error", e.Err)
	}
}

func (l *Logger) logStarted(e *fxevent.Started) {
	if e.Err != nil {
		l.Error("fx: start failed", "error", e.Err)
	} else {
		l.Debug("fx: started")
	}
}

func (l *Logger) logLoggerInitialized(e *fxevent.LoggerInitialized) {
	if e.Err != nil {
		l.Error("fx: logger failed to initialize", "error", e.Err)
	} else {
		l.Debug("fx: logger initialized", "constructor", e.ConstructorName)
	}
}
