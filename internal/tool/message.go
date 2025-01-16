package tool

import (
	"fmt"
	"os/exec"
	"runtime"
	log "unknwon.dev/clog/v2"
)

func MessagePop(title, message string) {
	var cmd *exec.Cmd

	// 判断操作系统类型
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("msg", "*", fmt.Sprintf("%s\n%s", title, message))
	case "darwin":
		// 在 macOS 上使用 AppleScript 实现弹窗通知
		cmd = exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "%s"`, message, title))
	case "linux":
		// 在 Linux 上使用 notify-send 命令实现弹窗通知
		cmd = exec.Command("notify-send", title, message)
	default:
		log.Error("unsupported operating system: %s", runtime.GOOS)
		return
	}

	// 执行命令并处理错误
	if err := cmd.Run(); err != nil {
		log.Error("message pop failed: %v", err)
	}

}
