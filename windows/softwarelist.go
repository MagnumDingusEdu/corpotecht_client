//go:build windows
// +build windows

package windows

import (
	"corpotecht_client/crossplatform"
	"net/url"
)

func getInstalledSoftware() {
	command := "Get-Package -Provider Programs -IncludeWindowsInstaller *"

	output := executeStringInShell(command)
	payload := url.Values{
		"command_output": {output},
	}
	payload.Add("Directive", "listsoftware")
	crossplatform.SendDirectiveResponse(payload)

}
