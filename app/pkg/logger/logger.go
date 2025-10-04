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
	l.Debug("app: starting component lifecycle hook", "component", e.CallerName, "hook", e.FunctionName)
}

func (l *Logger) logOnStartExecuted(e *fxevent.OnStartExecuted) {
	if e.Err != nil {
		l.Error("app: component failed to start", "component", e.CallerName, "hook", e.FunctionName, "error", e.Err)
	} else {
		l.Debug("app: component started successfully", "component", e.CallerName, "duration", e.Runtime.String())
	}
}

func (l *Logger) logOnStopExecuting(e *fxevent.OnStopExecuting) {
	l.Debug("app: stopping component", "component", e.CallerName, "hook", e.FunctionName)
}

func (l *Logger) logOnStopExecuted(e *fxevent.OnStopExecuted) {
	if e.Err != nil {
		l.Error("app: component failed to stop gracefully", "component", e.CallerName, "hook", e.FunctionName, "error", e.Err)
	} else {
		l.Debug("app: component stopped successfully", "component", e.CallerName, "duration", e.Runtime.String())
	}
}

func (l *Logger) logSupplied(e *fxevent.Supplied) {
	if e.Err != nil {
		l.Error("app: failed to supply dependency", "type", e.TypeName, "error", e.Err)
	} else {
		l.Debug("app: dependency supplied to container", "type", e.TypeName)
	}
}

func (l *Logger) logProvided(e *fxevent.Provided) {
	if e.Err != nil {
		l.Error("app: failed to provide dependency", "constructor", e.ConstructorName, "error", e.Err)
	} else {
		l.Debug("app: dependencies registered", "constructor", e.ConstructorName, "provides", e.OutputTypeNames)
	}
}

func (l *Logger) logDecorated(e *fxevent.Decorated) {
	if e.Err != nil {
		l.Error("app: failed to decorate dependency", "decorator", e.DecoratorName, "error", e.Err)
	} else {
		l.Debug("app: dependency decorated", "decorator", e.DecoratorName, "outputs", e.OutputTypeNames)
	}
}

func (l *Logger) logInvoking(e *fxevent.Invoking) {
	l.Debug("app: invoking function", "function", e.FunctionName)
}

func (l *Logger) logInvoked(e *fxevent.Invoked) {
	if e.Err != nil {
		l.Error("app: function invocation failed", "function", e.FunctionName, "error", e.Err)
	} else {
		l.Debug("app: function invoked successfully", "function", e.FunctionName)
	}
}

func (l *Logger) logStopping(e *fxevent.Stopping) {
	l.Info("app: shutting down application", "signal", e.Signal.String())
}

func (l *Logger) logStopped(e *fxevent.Stopped) {
	if e.Err != nil {
		l.Error("app: stopped with errors", "error", e.Err)
	}
}

func (l *Logger) logRollingBack(e *fxevent.RollingBack) {
	l.Error("app: rolling back application startup", "error", e.StartErr)
}

func (l *Logger) logRolledBack(e *fxevent.RolledBack) {
	if e.Err != nil {
		l.Error("app: rollback failed", "error", e.Err)
	}
}

func (l *Logger) logStarted(e *fxevent.Started) {
	if e.Err != nil {
		l.Error("app: failed to start", "error", e.Err)
	} else {
		l.Info("app: started successfully")
	}
}

func (l *Logger) logLoggerInitialized(e *fxevent.LoggerInitialized) {
	if e.Err != nil {
		l.Error("app: logger initialization failed", "error", e.Err)
	} else {
		l.Debug("app: logger initialized", "constructor", e.ConstructorName)
	}
}
