package logx

import (
	"fmt"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	logger := New(WithScrollSize(ScrollSize_1MB), WithConsoleOutput(false), WithScrollTime(scrollTime_Test))
	msg := "adsfasdfasdfasdiofjaosidfjoasjdfojoisdjfoijsdfjoisajdofjsiodfj"
	for {
		logger.Info("%s", msg)

		time.Sleep(10 * time.Millisecond)
	}

}

func TestTruncate(t *testing.T) {
	fmt.Println(time.Date(2025, 2, 9, 0, 0, 1, 0, time.UTC).Truncate(time.Hour * 4).Format(time.DateTime))
}
