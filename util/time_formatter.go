// pakage util utils
package util

import (
	"fmt"
	"strings"
	"time"
)

const (
	minute   = 60
	hour     = 60 * minute
	day      = 24 * hour
	threeDay = 3 * day
)

func FormatTime(Timestamp int64) string {
	var strBuilder strings.Builder
	timeDifference := time.Now().Unix() - Timestamp

	if timeDifference < minute {
		return "刚刚"
	} else if timeDifference < hour {
		strBuilder.WriteString(fmt.Sprint((timeDifference / minute)))
		strBuilder.WriteString(" 分钟前")
	} else if timeDifference < day {
		strBuilder.WriteString(fmt.Sprint((timeDifference / hour)))
		strBuilder.WriteString(" 小时前")
	} else if timeDifference < threeDay {
		strBuilder.WriteString(fmt.Sprint((timeDifference / day)))
		strBuilder.WriteString(" 天前")
	} else {
		timeStamp := time.Unix(Timestamp, 0)
		return timeStamp.Format("2006-01-02")
	}
	
	return strBuilder.String()
}
