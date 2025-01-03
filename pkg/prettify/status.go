package prettify

import (
	"github.com/fatih/color"
)

type ColorFunc func(a ...interface{}) string

func statusColor(status int) ColorFunc {
	if status >= 200 && status < 300 {
		return color.New(SuccessColor).SprintFunc()
	} else if status >= 300 && status < 400 {
		return color.New(RedirectColor).SprintFunc()
	} else if status >= 400 && status < 500 {
		return color.New(ClientErrorColor).SprintFunc()
	} else {
		return color.New(ServerErrorColor).SprintFunc()
	}
}

func Status(status int) ColorFunc {
	return statusColor(status)
}
