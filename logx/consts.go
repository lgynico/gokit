package logx

import "time"

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func levelString(l Level) string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO "
	case LevelWarn:
		return "WARN "
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "?????"
	}
}

func levelColor(l Level) string {
	switch l {
	case LevelDebug:
		return "\033[0;34m%s\033[0m"
	case LevelInfo:
		return "\033[0;32m%s\033[0m"
	case LevelWarn:
		return "\033[0;33m%s\033[0m"
	case LevelError:
		return "\033[0;31m%s\033[0m"
	case LevelFatal:
		return "\033[0;31m%s\033[0m"
	default:
		return "%s"
	}
}

type ScrollSize int

const (
	ScrollSize_None  ScrollSize = 0
	ScrollSize_1MB   ScrollSize = 1 * 1024 * 1024
	ScrollSize_10MB  ScrollSize = 10 * 1024 * 1024
	ScrollSize_50MB  ScrollSize = 50 * 1024 * 1024
	ScrollSize_100MB ScrollSize = 100 * 1024 * 1024
	ScrollSize_500MB ScrollSize = 500 * 1024 * 1024
	ScrollSize_1GB   ScrollSize = 1 * 1024 * 1024 * 1024
)

type ScrollTime time.Duration

const (
	ScrollTime_Default            = ScrollTime_1H
	ScrollTime_1H      ScrollTime = ScrollTime(1 * time.Hour)
	ScrollTime_2H      ScrollTime = ScrollTime(2 * time.Hour)
	ScrollTime_4H      ScrollTime = ScrollTime(4 * time.Hour)
	ScrollTime_8H      ScrollTime = ScrollTime(8 * time.Hour)
	ScrollTime_12H     ScrollTime = ScrollTime(12 * time.Hour)
	ScrollTime_1D      ScrollTime = ScrollTime(24 * time.Hour)
	scrollTime_Test    ScrollTime = ScrollTime(time.Minute)
)

func timeTruncate(scrollTime ScrollTime) time.Time {
	if scrollTime < ScrollTime_1H {
		return time.Now().Truncate(time.Minute)
	} else if scrollTime < ScrollTime_1D {
		return time.Now().Truncate(time.Hour)
	} else {
		return time.Now().Truncate(time.Hour * 24)
	}
}
