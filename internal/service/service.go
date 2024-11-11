package service

import (
	"os/exec"
	"syscall"
)

func StartService(name string) error {
	cmd := exec.Command("sc", "start", name)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd.Run()
}

func StopService(name string) error {
	cmd := exec.Command("sc", "stop", name)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd.Run()
}
