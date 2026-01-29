package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	i := 0
	Run(func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(1 * time.Second):
				fmt.Printf("running %d\n", i)
				i++
			}

		}
	})
	stopped := i
	i = 0
	for {
		i++
		time.Sleep(time.Second)
		fmt.Printf("stopping from %d to %d\n", stopped, i)
	}
}

func RunWithContext(ctx context.Context, callback func(context.Context) error) error {
	return callback(ContextWithSigterm(ctx))
}

func Run(callback func(context.Context) error) error {
	return RunWithContext(context.Background(), callback)
}

func Context() context.Context {
	return ContextWithSigterm(context.Background())
}

func ContextWithSigterm(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt,
			os.Interrupt,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)
		<-interrupt
		cancel()
	}()
	return ctx
}
