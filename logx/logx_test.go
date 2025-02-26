package logx

import (
	"fmt"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	logger := New(
		WithScrollSize(ScrollSize_1MB),
		WithConsoleOutput(true),
		WithScrollTime(scrollTime_Test),
		WithLevel(LevelDebug),
	)
	msg := "adsfasdfasdfasdiofjaosidfjoasjdfojoisdjfoijsdfjoisajdofjsiodfj"
	for {
		logger.Info(msg)
		logger.Debug(msg)
		logger.Error(msg)
		logger.Warn(msg)

		time.Sleep(time.Second)
	}

}

func TestTruncate(t *testing.T) {
	fmt.Println(time.Date(2025, 2, 9, 0, 0, 1, 0, time.UTC).Truncate(time.Hour * 4).Format(time.DateTime))
}
