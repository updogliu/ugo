package main

import (
	"fmt"
	"time"

	"github.com/updogliu/ugo/utime"
)

func main() {
	cd := utime.NewCooldown(3 * time.Second)

	go func() {
		for {
			if cd.Ready() {
				fmt.Println("Robot 1:", utime.NowSec())
			} else {
				utime.SleepMs(50)
			}
		}
	}()

	go func() {
		for {
			if cd.Ready() {
				fmt.Println("Robot 2:", utime.NowSec())
			} else {
				utime.SleepMs(50)
			}
		}
	}()

	utime.SleepMs(16e3)
}

/* Example Output:

Robot 1: 1536454877
Robot 1: 1536454880
Robot 1: 1536454883
Robot 2: 1536454886
Robot 1: 1536454889
Robot 2: 1536454892

*/
