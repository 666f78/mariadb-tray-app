package ui

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/666f78/mariadb-tray-app/internal/service"

	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
)

//go:embed assets/icon_offline.ico
var iconOffline []byte

//go:embed assets/icon_online.ico
var iconOnline []byte

func OnReady() {
	updateIconStatus()
	systray.SetTitle("MariaDB Tray App")
	systray.SetTooltip("Control MariaDB Service")

	mEnable := systray.AddMenuItem("Enable MariaDB", "Start MariaDB service")
	mDisable := systray.AddMenuItem("Disable MariaDB", "Stop MariaDB service")
	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	go func() {
		for {
			select {
			case <-mEnable.ClickedCh:
				err := service.StartService("MariaDB")
				if err != nil {
					beeep.Alert("Error", fmt.Sprintf("Failed to start MariaDB service: %v", err), "")
				} else {
					beeep.Notify("MariaDB Service", "MariaDB has been enabled", "")
					waitForServiceStatus("RUNNING")
					updateIconStatus()
				}
			case <-mDisable.ClickedCh:
				err := service.StopService("MariaDB")
				if err != nil {
					beeep.Alert("Error", fmt.Sprintf("Failed to stop MariaDB service: %v", err), "")
				} else {
					beeep.Notify("MariaDB Service", "MariaDB has been disabled", "")
					waitForServiceStatus("STOPPED")
					updateIconStatus()
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func OnExit() {
	// Cleanup tasks here
}

func updateIconStatus() {
    status := service.GetServiceStatus("MariaDB")
    var iconData []byte
    if status == "online" {
        iconData = iconOnline
    } else {
        iconData = iconOffline
    }
    systray.SetIcon(iconData)
}

func waitForServiceStatus(targetStatus string) {
	for {
		status := service.GetServiceStatus("MariaDB")
		if (targetStatus == "RUNNING" && status == "online") || (targetStatus == "STOPPED" && status == "offline") {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
