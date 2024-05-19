package app

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func Shutdown(wg *sync.WaitGroup, done chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-c:
		fmt.Println("Shutting down gracefully...")
		close(done)
		wg.Wait()
		return
	case <-done:
		return
	}
}
