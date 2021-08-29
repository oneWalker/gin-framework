package schedule

import (
	"github.com/sirupsen/logrus"
	"time"
)

func InitTimer() {
	StartMonthlyTimer(testFuncDemo)
	StartMinuteTimer(1, testFuncDemo)
}

//注意事项，如果定义有参数则在定时任务函数上也有参数
func testFuncDemo() {
	logrus.Info("test is running")
	return
}

//set timer monthly
func StartMonthlyTimer(f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下个月零点
			year := now.Year()
			month := now.Month()
			if month == time.December {
				year += 1
				month = time.January
			} else {
				month = time.Month(int(month) + 1)
			}
			next := time.Date(year, month, 1, 0, 0, 0, 0, time.Now().Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

//set timer monthly
func StartMinuteTimer(intMin uint, f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			next := now.Add(time.Duration(intMin) * time.Minute)
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

//set timer secondly by tricker
func StartSecondTimer(seconds int, f func()) {
	ticker := time.NewTicker(time.Second * time.Duration(seconds))
	go func() {
		for range ticker.C {
			f()
		}
	}()
}
