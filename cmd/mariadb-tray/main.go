package main

import (
	"github.com/666f78/mariadb-tray-app/internal/ui"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(ui.OnReady, ui.OnExit)
}