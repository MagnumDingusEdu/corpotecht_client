package linux

import (
	"corpotecht_client/config"
	"corpotecht_client/utils"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
)

func IsServiceInstalled() bool {
	currentUser, _ := user.Current()
	blobLoc := currentUser.HomeDir + config.LinuxBlobLocation + config.LinuxBlobName
	_, err := os.Open(blobLoc)
	return !os.IsNotExist(err)
}

func AutoStart() {
	currentUser, _ := user.Current()
	err := os.MkdirAll(currentUser.HomeDir+config.LinuxBlobLocation, os.ModePerm)
	utils.HandleError(err)
	blobLoc := currentUser.HomeDir + config.LinuxBlobLocation + config.LinuxBlobName
	currentBinaryLocation, err := os.Executable()
	utils.HandleError(err)

	_ = os.Remove(blobLoc)

	copyBinaryToDestination(currentBinaryLocation, blobLoc)
	createSystemdService(*currentUser, blobLoc)
	utils.DebugLog("Installed systemd service successfully!")
}

var systemdTemplate = `
[Unit]
Description=Corpotecht Client

[Service]
ExecStart=%s

[Install]
WantedBy=default.target
`

func createSystemdService(userDetails user.User, blobLoc string) {

	serviceName := config.LinuxSystemdService

	userServicesLoc := userDetails.HomeDir + "/.config/systemd/user"
	err := os.MkdirAll(userServicesLoc, os.ModePerm)
	utils.HandleError(err)

	// Write service file
	systemdFileLoc := userServicesLoc + "/" + serviceName
	systemdFileContent := fmt.Sprintf(systemdTemplate, blobLoc)
	err = ioutil.WriteFile(systemdFileLoc, []byte(systemdFileContent), 0644)
	utils.HandleError(err)

	// Enable and Start Service
	_, err = exec.Command("systemctl", "--user", "daemon-reload").CombinedOutput()
	utils.HandleError(err)
	_, err = exec.Command("systemctl", "--user", "enable", serviceName).CombinedOutput()
	utils.HandleError(err)
	_, err = exec.Command("systemctl", "--user", "start", serviceName).CombinedOutput()
	utils.HandleError(err)
	out, err := exec.Command("systemctl", "--user", "status", serviceName).CombinedOutput()
	utils.DebugLog(string(out[:]))

}

func copyBinaryToDestination(source string, destination string) {

	// Open source and destination files
	sourceFile, err := os.Open(source)
	utils.HandleError(err)
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	utils.HandleError(err)
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	utils.HandleError(err)

	err = os.Chmod(destination, 0755)
	utils.HandleError(err)
}
