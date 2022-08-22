//go:build windows
// +build windows

package windows

import (
	"corpotecht_client/crossplatform"
	"net/url"
)

func UninstallSoftware(params []string) {

	command := "get-package " + params[0] + " | uninstall-package"

	output := executeStringInShell(command)
	payload := url.Values{
		"command_output": {output},
	}
	payload.Add("Directive", "listsoftware")
	crossplatform.SendDirectiveResponse(payload)
}
