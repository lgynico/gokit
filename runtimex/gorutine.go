package runtimex

import (
	"runtime"
	"strconv"
	"strings"
)

func GetGoroutineId() int {
	buf := make([]byte, 20) // goroutine xxxx
	runtime.Stack(buf[:], false)

	strs := strings.Split(string(buf), " ")

	if len(strs) > 2 {
		if id, err := strconv.Atoi(strs[1]); err == nil {
			return id
		}
	}

	return -1
}
