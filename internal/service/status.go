package service

import (
	"os/exec"
	"strings"
	"syscall"
)

func GetServiceStatus(name string) string {
	cmd := exec.Command("sc", "query", name)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	if strings.Contains(string(output), "RUNNING") {
		return "online"
	}
	return "offline"
}
