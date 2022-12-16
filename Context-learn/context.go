package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Select_chan()
	// Context_fun()
	// Multi_Context()
	// Context_WithDeadline()
	// Context_WithTimeout()
	Context_WithValue()
}

func Select_chan() {
	exit := make(chan byte)

	go func() { // 创建一个 gorouting
		for {
			select {
			case <-exit: // 当 exit 通道有数据传入时,执行此命令
				fmt.Println("exiting...")
				return
			default:
				fmt.Println("goroutine is running...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second) // 睡眠10秒,模拟程序运行了10秒
	fmt.Println("done!")
	exit <- 'Y' // 程序执行完成,向exit传入信号以终止程序
}

func Context_fun() {
	ctx, cancel := context.WithCancel(context.Background()) // 创建一个带结束的 Context,context.Background()返回一个空Context
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("exiting...")
				return
			default:
				fmt.Println("gorouting running...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("done!")
	cancel()                // 调用取消函数,结束掉 gorouting
	time.Sleep(time.Second) // 程序结束过快,还没执行到输出语句就结束了,所以睡眠1秒
}

func Multi_Context() {
	ctx, cancel := context.WithCancel(context.Background())
	go Just_run(ctx, "one")
	go Just_run(ctx, "two")
	go Just_run(ctx, "three")
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(time.Second)
}

func Just_run(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name + " has done!")
			return
		default:
			fmt.Println(name + " is running!")
			time.Sleep(2 * time.Second)
		}
	}
}

func Context_WithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(20*time.Second)) // 创建一个可以随时取消的,20秒后自动结束的 Context
	defer cancel()                                                                            // 最后取消所有的 Context
	go Just_run(ctx, "one")
	go Just_run(ctx, "two")
	go Just_run(ctx, "three")
	time.Sleep(4 * time.Second)
	fmt.Println("done! do not run anymore!")
	cancel()
}

func Context_WithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // 和 WithDeadline 差不多,不过传入的是一个 time.Duration
	defer cancel()
	go Just_run(ctx, "one")
	go Just_run(ctx, "two")
	go Just_run(ctx, "three")
	time.Sleep(15 * time.Second)
	fmt.Println("done!")
}

var key = "name" // 定义一个键

func Context_WithValue() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2 := context.WithValue(ctx, key, "one") // 创建一个带值的 Context
	ctx3 := context.WithValue(ctx, key, "two")
	ctx4 := context.WithValue(ctx, key, "three")
	go Always_run(ctx2)
	go Always_run(ctx3)
	go Always_run(ctx4)
	time.Sleep(10 * time.Second)
	fmt.Println("all done!")
	cancel()
	time.Sleep(time.Second)
}

func Always_run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), " has done!") // 使用 ctx.Value() 来获取值
			return
		default:
			fmt.Println((ctx.Value(key)), " is running!")
			time.Sleep(2 * time.Second)
		}
	}
}
