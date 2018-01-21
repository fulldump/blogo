package kip

import (
	"strconv"
	"time"
)

func random_name(prefix string) string {
	return prefix + "_" + strconv.FormatInt(time.Now().UnixNano(), 10)
}
