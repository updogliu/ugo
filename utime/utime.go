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
	return time.Unix(timeSec, 0).UTC()
}

func MsToTime(timeMs int64) time.Time {
	return time.Unix(0, timeMs*1e6).UTC()
}

func MsToDur(durMs int64) time.Duration {
	return time.Duration(durMs) * time.Millisecond
}

func MsToSec(durMs int64) float64 {
	return float64(durMs) / 1e3
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

func SinceMs(timeMs int64) int64 {
	return NowMs() - timeMs
}

// Returns whether a timestamp (in second) is in 2017/12/01 ~ 2099/12/01
func IsRealTimeSec(timeSec int64) bool {
	return 1512086400 <= timeSec && timeSec <= 4099852799
}

// Returns whether a timestamp (in millisecond) is in 2017/12/01 ~ 2099/12/01
func IsRealTimeMs(timeMs int64) bool {
	return 1512086400000 <= timeMs && timeMs <= 4099852799999
}

// Assert a timestamp (in second) is in 2017/12/01 ~ 2099/12/01
func AssertRealTimeSec(timeSec int64) {
	if !IsRealTimeSec(timeSec) {
		panic(fmt.Sprintf("timeSec %v is not in 2017/12/01 ~ 2099/12/01 UTC", timeSec))
	}
}

// Assert a timestamp (in millisecond) is in 2017/12/01 ~ 2099/12/01 UTC.
func AssertRealTimeMs(timeMs int64) {
	if !IsRealTimeMs(timeMs) {
		panic(fmt.Sprintf("timeMs %v is not in 2017/12/01 ~ 2099/12/01 UTC", timeMs))
	}
}

// Precondition: `timeoutMs` >= 0
func CtxTimeoutMs(timeoutMs int64) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeoutMs)*time.Millisecond)
	return ctx
}

// Return immediately if `durationMs` is non-positive.
func SleepMs(durationMs int64) {
	time.Sleep(time.Duration(durationMs) * time.Millisecond)
}

// Return immediately if `timeMs` is a past time (earlier than `NowMs()`).
func SleepUntilMs(timeMs int64) {
	SleepMs(timeMs - NowMs())
}

const TimeLayoutSqlMs = "2006-01-02 15:04:05.000"
func FormatTimeMsAsSqlMs(timeMs int64) string {
	return MsToTime(timeMs).UTC().Format(TimeLayoutSqlMs)
}
