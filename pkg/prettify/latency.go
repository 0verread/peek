package prettify

import (
	"github.com/fatih/color"
)

type LatencyFunc func(a ...interface{}) string

func Latency(latency int) LatencyFunc {
	return color.New(LatencyColor).SprintFunc()
}
