package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"
)

type Level int8

// 共有字段
type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

// 日志的结构体
type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

// 新建日志
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger {
	n1 := *l
	return &n1
}

// 设置共有字段
func (l *Logger) WithFields(f Fields) *Logger {
	l1 := l.clone()
	if l1.fields == nil {
		l1.fields = make(Fields)
	}
	for k, v := range f {
		l1.fields[k] = v
	}
	return l1
}

// 设置上下文
func (l *Logger) WithContext(ctx context.Context) *Logger {
	l1 := l.clone()
	l1.ctx = ctx
	return l1
}

// 设置当前某一层调用栈的信息
func (l *Logger) WithCaller(skip int) *Logger {
	l1 := l.clone()

	// 获取当前执行函数的指针,函数所在文件名目录和行号,skip为0时表示当前函数,1为上一层的调用者,依次类推
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		l1.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return l1
}

// 设置当前的整个调用栈信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)

	// 返回调用栈的程序深度,0表示计算到Callers这个函数本身,1代表当前函数,依次类推
	depth := runtime.Callers(minCallerDepth, pcs)

	// 获得整个栈的信息
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}
	l1 := l.clone()
	l1.callers = callers
	return l1
}
