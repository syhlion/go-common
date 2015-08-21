package common

import (
	"strconv"
	"time"
)

func TimeToString() string {
	return strconv.FormatInt(Time(), 10)
}

func Time() (t int64) {
	t = time.Now().UnixNano() / int64(time.Millisecond)
	return
}
