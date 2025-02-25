package logx

type (
	Option struct {
		Level         Level
		ScrollSize    ScrollSize
		ScrollTime    ScrollTime
		ConsoleOutput bool
		OutPath       string
		Prefix        string
		CallerDepth   int // 调用栈深度: 解决 Logger 被封装时, 没有正确获取调用者信息
	}

	OptionFunc func(*Option)
)

func defaultOption() Option {
	return Option{
		Level:         LevelInfo,
		ScrollSize:    ScrollSize_None,
		ScrollTime:    ScrollTime_Default,
		ConsoleOutput: true,
		OutPath:       "./logs",
		Prefix:        "log",
		CallerDepth:   0,
	}
}

func WithLevel(level Level) OptionFunc {
	return func(o *Option) {
		o.Level = level
	}
}

func WithScrollSize(size ScrollSize) OptionFunc {
	return func(o *Option) {
		o.ScrollSize = size
	}
}

func WithScrollTime(time ScrollTime) OptionFunc {
	return func(o *Option) {
		o.ScrollTime = time
	}
}

func WithConsoleOutput(output bool) OptionFunc {
	return func(o *Option) {
		o.ConsoleOutput = output
	}
}

func WithOutPath(path string) OptionFunc {
	return func(o *Option) {
		o.OutPath = path
	}
}

func WithPrefix(prefix string) OptionFunc {
	return func(o *Option) {
		o.Prefix = prefix
	}
}

func WithCallerDepth(depth int) OptionFunc {
	return func(o *Option) {
		o.CallerDepth = depth
	}
}

func WithFileOutputDisable() OptionFunc {
	return func(o *Option) {
		o.Prefix = ""
	}
}
