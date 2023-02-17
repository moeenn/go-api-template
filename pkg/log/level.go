package log

type Level uint8

func (l Level) String() string {
	switch l {
	case INFO:
		return "INFO"

	case WARNING:
		return "WARNING"

	case ERROR:
		return "ERROR"
	}

	return "UNKNOWN"
}

const (
	INFO Level = iota
	WARNING
	ERROR
)
