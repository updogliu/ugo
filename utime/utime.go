package utime

import (
	"context"
	"fmt"
	"time"
)

func NowSec() int64 {
	return time.Now().Unix()
}

func NowMs() int64 {
	return time.Now().UnixNano() / 1e6
}

func TimeToSec(t time.Time) int64 {
	return t.UnixNano() / 1e9
}

func TimeToMs(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func DurToSec(d time.Duration) int64 {
	return int64(d / time.Second)
}

func DurToMs(d time.Duration) int64 {
	return int64(d / time.Millisecond)
}

func SecToTime(timeSec int64) time.Time {
	return time.Unix(timeSec, 0)
}

func MsToTime(timeMs int64) time.Time {
	return time.Unix(0, timeMs*1e6)
}

// Abs(t1 - t2)
func Gap(t1, t2 time.Time) time.Duration {
	if t1.Before(t2) {
		return t2.Sub(t1)
	}
	return t1.Sub(t2)
}

func GapSec(t1, t2 time.Time) int64 {
	return DurToSec(Gap(t1, t2))
}

func GapMs(t1, t2 time.Time) int64 {
	return DurToMs(Gap(t1, t2))
}

// Returns whether a timestamp (in second) is in 2017/12/01 ~ 2027/12/01
func IsRealTimeSec(timeSec int64) bool {
	return IsRealTimeMs(timeSec * 1000)
}

// Returns whether a timestamp (in millisecond) is in 2017/12/01 ~ 2027/12/01
func IsRealTimeMs(timeMs int64) bool {
	return 1512086400000 <= timeMs && timeMs <= 1827705599999
}

// Assert a timestamp (in second) is in 2017/12/01 ~ 2027/12/01
func AssertRealTimeSec(timeSec int64) {
	if !IsRealTimeSec(timeSec) {
		panic(fmt.Sprintf("timeSec %v is not in 2017/12/01 ~ 2027/12/01 UTC", timeSec))
	}
}

// Assert a timestamp (in millisecond) is in 2017/12/01 ~ 2027/12/01 UTC.
func AssertRealTimeMs(timeMs int64) {
	if !IsRealTimeMs(timeMs) {
		panic(fmt.Sprintf("timeMs %v is not in 2017/12/01 ~ 2027/12/01 UTC", timeMs))
	}
}

func CtxTimeoutMs(timeoutMs int64) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeoutMs)*time.Millisecond)
	return ctx
}

// DEPRECATED - use `CtxTimeoutMs`
// Precondition: `timeoutMs` >= 0
func CtxWithTimeoutMs(timeoutMs int64) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeoutMs)*time.Millisecond)
	return ctx
}

func SleepMs(durationMs int64) {
	time.Sleep(time.Duration(durationMs) * time.Millisecond)
}
