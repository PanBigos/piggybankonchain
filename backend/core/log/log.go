package log

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

var (
	root Logger
	lvl  *slog.LevelVar = &slog.LevelVar{}
	_    Logger         = (*slogWrapper)(nil)
)

type Logger interface {
	Trace(msg string, args ...any)
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Fatal(msg string, args ...any)
	With(args ...any) Logger
	WithGroup(name string) Logger
}

func Root() Logger {
	return root
}

func New(w io.Writer, source bool) *slogWrapper {
	handler := slog.NewJSONHandler(w, &slog.HandlerOptions{
		AddSource: source,
		Level:     lvl,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				source := a.Value.Any().(*slog.Source)
				if source != nil {
					source.File = filepath.Base(source.File)
					source.Function = ""
				}
			}
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				switch level.Level() {
				case LvlTrace.toSlog():
					a.Value = slog.StringValue(LvlTrace.ToString())
				case LvlFatal.toSlog():
					a.Value = slog.StringValue(LvlFatal.ToString())
				}
			}
			return a
		},
	})
	return &slogWrapper{Logger: slog.New(handler)}

}

func init() {
	ChangeLvl(LvlInfo)
	root = New(os.Stdout, true)
}

type slogWrapper struct {
	*slog.Logger
}

func (s *slogWrapper) Trace(msg string, args ...any) {
	runtime.Caller(1)
	s.Logger.Log(context.Background(), LvlTrace.toSlog(), msg, args...)
}

func (s *slogWrapper) Fatal(msg string, args ...any) {
	os.Exit(s.fatal(msg, args...))
}

func (s *slogWrapper) fatal(msg string, args ...any) int {
	runtime.Caller(2)
	s.Logger.Log(context.Background(), LvlFatal.toSlog(), msg, args...)
	return 255
}

func (s *slogWrapper) With(
	args ...any,
) Logger {
	return &slogWrapper{Logger: s.Logger.With(args...)}
}
func (s *slogWrapper) WithGroup(name string) Logger {
	return &slogWrapper{Logger: s.Logger.WithGroup(name)}
}

const (
	LvlFatal Lvl = iota
	LvlError
	LvlWarn
	LvlInfo
	LvlDebug
	LvlTrace
)

const (
	FirstLvl Lvl = LvlFatal
	LastLvl  Lvl = LvlTrace
)

type Lvl int

func (l Lvl) Lvls() []string {
	arr := make([]string, LastLvl)
	for i := 0; i < int(LastLvl); i++ {
		arr[i] = (Lvl)(i).ToString()
	}
	return arr
}

func (l *Lvl) FromString(str string) error {
	if len(str) == 0 {
		return nil
	}
	switch str {
	case "FATAL", "fatal":
		*l = LvlFatal
	case "ERROR", "error":
		*l = LvlError
	case "WARN", "warn":
		*l = LvlWarn
	case "INFO", "info":
		*l = LvlInfo
	case "DEBUG", "debug":
		*l = LvlDebug
	case "TRACE", "trace":
		*l = LvlTrace
	default:
		return fmt.Errorf("unknown lvl: %s", str)
	}
	return nil
}

func (l Lvl) ToString() string {
	switch l {
	case LvlFatal:
		return "FATAL"
	case LvlError:
		return "ERROR"
	case LvlWarn:
		return "WARN"
	case LvlInfo:
		return "INFO"
	case LvlDebug:
		return "DEBUG"
	case LvlTrace:
		return "TRACE"
	default:
		return "unknown"
	}
}

func (l Lvl) toSlog() slog.Level {
	switch l {
	case LvlFatal:
		return slog.LevelError + 4
	case LvlError:
		return slog.LevelError
	case LvlWarn:
		return slog.LevelWarn
	case LvlInfo:
		return slog.LevelInfo
	case LvlDebug:
		return slog.LevelDebug
	case LvlTrace:
		return slog.LevelDebug - 10
	default:
		panic("unsupported lvl")
	}
}

func ChangeLvl(l Lvl) {
	lvl.Set(l.toSlog())
}

func Levels() []string {
	var l Lvl
	return l.Lvls()
}
