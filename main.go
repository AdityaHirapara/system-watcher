package main

import (
	"fmt"
	"time"

	systray "github.com/getlantern/systray"
	load "github.com/shirou/gopsutil/load"
)

var memUsage float64 = 0
var loadavg = 0.0

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("💹")
	systray.SetTooltip("System Watcher")
	titleItem := systray.AddMenuItem("System Watcher", "")

	ramStatusItem := systray.AddMenuItem(fmt.Sprintf("Memory usage: %.2f%%", memUsage), "")
	loadavgItem := systray.AddMenuItem(fmt.Sprintf("Load avg: %.2f", loadavg), "")

	systray.AddSeparator()
	quitItem := systray.AddMenuItem("Quit", "")

	titleItem.Disable()
	ramStatusItem.Disable()
	loadavgItem.Disable()

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

			loadavgObj, _ := load.Avg()
			loadavg = loadavgObj.Load5
			loadavgItem.SetTitle((fmt.Sprintf("Load avg: %.2f", loadavg)))

			time.Sleep(1 * time.Second)
		}
	}()
}

func onExit() {
	fmt.Println("Exited")
}
