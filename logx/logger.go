package logx

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/bwoil/erase-game/utils/syncx"
)

type Logger struct {
	o        Option
	output   *os.File
	sig      *syncx.Signal
	lastTime time.Time
	counter  int
	logC     chan string
}

func New(opts ...OptionFunc) *Logger {
	o := defaultOption()
	for _, opt := range opts {
		opt(&o)
	}

	var output *os.File
	if o.Prefix != "" {
		// 确保日志目录存在
		err := os.MkdirAll(o.OutPath, 0755)
		if err != nil {
			panic(fmt.Errorf("unable to create log directory %s: %v", o.OutPath, err))
		}

		filename := filepath.Join(o.OutPath, o.Prefix+".log")
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			panic(fmt.Errorf("unable to open log file %s: %v", filename, err))
		}
		output = file
	}

	logger := &Logger{
		o:        o,
		output:   output,
		sig:      syncx.NewSignal(),
		lastTime: timeTruncate(o.ScrollTime),
		logC:     make(chan string, 1024),
	}

	go logger.start()

	return logger
}

func (p *Logger) Debug(format string, v ...any) { p.log(LevelDebug, format, v...) }

func (p *Logger) Info(format string, v ...any) { p.log(LevelInfo, format, v...) }

func (p *Logger) Warn(format string, v ...any) { p.log(LevelWarn, format, v...) }

func (p *Logger) Error(format string, v ...any) { p.log(LevelError, format, v...) }

func (p *Logger) Fatal(format string, v ...any) {
	p.log(LevelFatal, format, v...)
	os.Exit(1)
}

func (p *Logger) log(level Level, format string, v ...any) {
	if level < p.o.Level {
		return
	}

	caller := "unknown"
	_, file, line, ok := runtime.Caller(2 + p.o.CallerDepth)
	if ok {
		caller = fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}

	msg := fmt.Sprintf(format, v...)
	msg = fmt.Sprintf(
		"%s [%s] %s : %s\n",
		time.Now().Format(time.DateTime),
		levelString(level),
		caller,
		msg,
	)

	p.logC <- msg
}

func (p *Logger) start() {
	defer func() {
		if p.output != nil {
			p.output.Close()
		}
	}()

	timer := time.NewTimer(time.Until(p.lastTime.Add(time.Duration(p.o.ScrollTime))))
	defer timer.Stop()

	defer close(p.logC)

loop:
	for {
		select {
		case msg := <-p.logC:
			p.write(msg)
		case <-timer.C:
			p.switchFile(true)
			timer.Reset(time.Duration(p.o.ScrollTime))
		case <-p.sig.Channel():
			break loop
		}
	}

	// write remain message
	for {
		select {
		case msg, ok := <-p.logC:
			if !ok {
				return
			}
			p.write(msg)
		case <-time.After(3 * time.Second):
			return
		}
	}

}

func (p *Logger) write(msg string) {
	if p.o.ConsoleOutput {
		os.Stdout.WriteString(msg)
	}

	if p.output != nil {
		p.output.WriteString(msg)
		p.output.Sync()
		if p.o.ScrollSize != ScrollSize_None {
			fi, _ := p.output.Stat()
			if fi.Size() > int64(p.o.ScrollSize) {
				p.switchFile(false)
			}
		}
	}

}

func (p *Logger) switchFile(isTimeout bool) error {
	if isTimeout {
		p.counter = 0
		p.lastTime = timeTruncate(p.o.ScrollTime)
	} else {
		p.counter++
	}

	filename := fmt.Sprintf("%s-%s.log", p.o.Prefix, p.lastTime.Format("20060102.150405"))
	if p.counter > 0 {
		filename += "." + strconv.Itoa(p.counter)
	}
	filename = filepath.Join(p.o.OutPath, filename)

	if p.output != nil {
		oldFilename := p.output.Name()
		p.output.Close()
		os.Rename(oldFilename, filename)
	}

	filename = filepath.Join(p.o.OutPath, p.o.Prefix+".log")
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	p.output = file
	return nil
}
