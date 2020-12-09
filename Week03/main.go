package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func businessLogic(ctx context.Context) error {
	return nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%+v", err)
		}
	}()

	eg, ctx := errgroup.WithContext(context.Background())
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)
	serverErr := make(chan error, 1)
	sigC := make(chan os.Signal, 1)

	// business logic
	eg.Go(func() error {
		go businessLogic(ctx)
		select {
		case err := <-serverErr:
			log.Print("优雅关闭。。。")
			close(sigC)
			close(serverErr)
			ctx.Done()
			return err
		}
	})
	server := &http.Server{Addr: ":8001"}

	eg.Go(func() error {
		go func() {
			log.Print("httpServer start...")
			serverErr <- server.ListenAndServe()
		}()
		select {
		case err := <-serverErr:
			log.Print("优雅关闭。。。")
			close(sigC)
			close(serverErr)
			return err
		}
	})

	eg.Go(func() error {
		signal.Notify(sigC,
			syscall.SIGINT|syscall.SIGTERM|syscall.SIGKILL)
		<-sigC
		log.Print("中断运行")
		return server.Shutdown(ctx)
	})

	if err := eg.Wait(); err != nil {
		log.Printf("eg errors is %+v", err)
	}
	log.Println("game over!")

}
