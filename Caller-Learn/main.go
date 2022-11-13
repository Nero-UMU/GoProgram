package main

import (
	"fmt"
	"runtime"
)

func main() {
	trace2()
}

func hello() {
	fmt.Printf("我是 %s, %s 在调用我\n", printName(), printCallerName())
	world()
}

func world() {
	fmt.Printf("我是 %s, %s 在调用我\n", printName(), printCallerName())
	trace()
}

func printName() string {
	pc, _, _, _ := runtime.Caller(1) // 本函数向上推一层即可获取调用者的 program counter
	return runtime.FuncForPC(pc).Name()
}

func printCallerName() string {
	{
		pc, _, _, _ := runtime.Caller(2)
		return runtime.FuncForPC(pc).Name()
	}
}

func trace() {
	pc := make([]uintptr, 10)   // 生成一个指向10个空间的uintptr指针
	n := runtime.Callers(0, pc) // 栈的program counter,放到pc这个指针中
	for i := 0; i < n; i++ {
		f := runtime.FuncForPC(pc[i])   // 把程序计数器地址对应的函数的信息获取出来
		file, line := f.FileLine(pc[i]) // 获取调用函数名称和行号
		fmt.Printf("%s:%d %s\n", file, line, f.Name())
	}
}

func trace2() {
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc[:n]) // 直接获取整个栈信息,放到pc中,返回一个pc切片指针
	for {
		frame, more := frames.Next()
		// 此处解析名字时就不需要使用 FuncForPC,直接调用
		fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
}
