package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/updogliu/ugo/utime"
	"github.com/updogliu/ugo/utask"
)

func main() {
	counter := 0

	ctx := utime.CtxTimeoutMs(12e3)
	utask.Retry(ctx, "ShowCounter", 1*time.Second, func() error {
		counter++
		fmt.Println(counter)

		if counter < 10 {
			return errors.New("Fake error")
		}
		return nil
	})
}
