package main

import (
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

func main() {
	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigc)

		<-sigc
		log.Println("Got interupt, shutting down")
		for i := 10; i > 0; i-- {
			<-sigc
			if i > 1 {
				log.Println("already shutting down")
			}
		}
		debug.SetTraceback("all")
		panic("boom")
	}()
	time.Sleep(30 * time.Second)

}
