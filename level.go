package log

type Level int

const (
	AllLevel Level = iota
	TraceLevel
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

func (l Level) String() string {
	switch l {
	case AllLevel:
		return "ALL"
	case TraceLevel:
		return "TRACE"
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	case PanicLevel:
		return "PANIC"
	default:
		return ""
	}
}

type LevelSettable interface {
	SetLevel(level Level)
}

func SetLevel(level Level) Level {
	globals.level = level
	std.SetLevel(level)
	return level
}

func GetLevel() Level {
	return globals.level
}
