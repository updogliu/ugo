package main

import (
	"fmt"
	"time"

	"github.com/updogliu/ugo/utime"
)

func main() {
	rt := utime.NewReadyTicker(5 * time.Second)
	fmt.Println("Started at", time.Now())
	for t := range rt.C {
		fmt.Println("Tick ", t)
	}
}
