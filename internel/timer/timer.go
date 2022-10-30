package timer

import (
	"time"
)

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

func GetCalculateTime(currentTimer time.Time, durationStr string) (time.Time, error) {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return time.Time{}, err
	}

	return currentTimer.Add(duration), nil
}
