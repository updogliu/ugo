package main

import (
	"errors"
	"fmt"

	"github.com/updogliu/ugo/utask"
	"github.com/updogliu/ugo/utime"
)

func main() {
	counter := 0

	ctx := utime.CtxTimeoutMs(12e3)
	utask.Retry(ctx, "ShowCounter", 1e3, func() error {
		counter++
		fmt.Println(counter)

		if counter < 10 {
			return errors.New("Fake error")
		}
		return nil
	})
}
