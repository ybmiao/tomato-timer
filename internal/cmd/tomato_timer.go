package cmd

import (
	"countdown/internal/tool"
	"github.com/urfave/cli"
	"time"
	log "unknwon.dev/clog/v2"
)

func TomatoTimer(c *cli.Context) error {
	// 1. 获取时间单位
	unit := c.String("unit")
	if unit == "" {
		unit = "m"
	}

	// 2. 获取番茄时钟工作时长
	workDurationTime := c.Int("work_time")
	if workDurationTime == 0 {
		workDurationTime = 25
	}
	workDuration := calculateTime(unit, workDurationTime)

	// 3. 获取番茄时钟休息时长
	restDurationTime := c.Int("rest_time")
	if restDurationTime == 0 {
		restDurationTime = 5
	}
	restDuration := calculateTime(unit, restDurationTime)

	// 4.获取是否倒计时
	isCountdown := !(c.Bool("hide"))

	// 记录次数
	startTime := 0
	for {
		startTime = startTime + 1
		log.Info("Starting tomato timer %d times...", startTime)
		// 5. 开始工作番茄钟
		StartWorkTimer(workDuration, isCountdown)
		// 6. 开始休息番茄钟
		StartRestTimer(restDuration, isCountdown)
	}
}

func StartWorkTimer(workDuration int, isCountdown bool) {
	log.Info("Starting work tomato timer...")
	tool.MessagePop("番茄时钟", "开始工作...")
	StartTimer(workDuration, isCountdown)
	log.Info("Work tomato timer has ended...")
	tool.MessagePop("番茄时钟", "工作结束...")
}

func StartRestTimer(restDuration int, isCountdown bool) {
	log.Info("Starting rest tomato timer...")
	tool.MessagePop("番茄时钟", "开始休息...")
	StartTimer(restDuration, isCountdown)
	log.Info("Rest tomato timer has ended...")
	tool.MessagePop("番茄时钟", "休息结束...")
}

func StartTimer(duration int, countdown bool) {
	// 计算结束时间
	start := time.Now()
	endTime := start.Unix() + int64(duration)
	log.Info("Start time: %s", start.Format("2006-01-02 15:04:05"))
	log.Info("End time: %s", time.Unix(endTime, 0).Format("2006-01-02 15:04:05"))

	for {
		if time.Now().Unix() > endTime {
			break
		}

		if countdown {
			// 计算剩余时间
			remainingTime := endTime - time.Now().Unix()
			printTimeRemaining(remainingTime)
			switch {
			case remainingTime > 60*60:
				time.Sleep(time.Hour)
			case remainingTime > 60:
				time.Sleep(time.Minute)
			default:
				time.Sleep(time.Second)
			}
		}
	}
}

func calculateTime(unit string, time int) int {
	switch unit {
	case "h":
		return time * 60 * 60
	case "m":
		return time * 60
	case "s":
		return time
	default:
		log.Fatal("Invalid time unit")
		return 0
	}
}

func printTimeRemaining(remainingTime int64) {
	hours := remainingTime / 3600
	minutes := (remainingTime % 3600) / 60
	seconds := remainingTime % 60
	log.Info("Time remaining: %02d:%02d:%02d", hours, minutes, seconds)
}
