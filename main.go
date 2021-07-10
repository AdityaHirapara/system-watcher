package main

import (
	"fmt"
	"time"

	systray "github.com/getlantern/systray"
)

var memUsage float64 = 0

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("💹")
	systray.SetTooltip("System Watcher")
	titleItem := systray.AddMenuItem("System Watcher", "")

	ramStatusItem := systray.AddMenuItem(fmt.Sprintf("Memory usage: %.2f%%", memUsage), "")

	systray.AddSeparator()
	quitItem := systray.AddMenuItem("Quit", "")

	titleItem.Disable()
	ramStatusItem.Disable()

	go func() {
		for {
			select {
			case <-quitItem.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

	// goroutine that updates stats in every 1s cycle
	go func() {
		for {
			memUsage, _ = getMemoryUsage()
			ramStatusItem.SetTitle(fmt.Sprintf("Memory usage: %.2f%%", memUsage))

			time.Sleep(1 * time.Second)
		}
	}()
}

func onExit() {
	fmt.Println("Exited")
}
