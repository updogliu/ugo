package main

import (
	"fmt"
	"time"

	"github.com/updogliu/ugo/utime"
)

func main() {
	rt := utime.NewReadyTicker(3 * time.Second)

	go func() {
		time.Sleep(10 * time.Second)
		rt.Stop()
	}()

	fmt.Println("Started at", time.Now())
	for t := range rt.C {
		fmt.Println("Tick ", t)
	}
}
