package utils

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadFile(link string, file *os.File) {
	err := file.Chmod(0770)
	HandleError(err)

	DebugLog("Downloading ....")
	// Get the response bytes from the url
	response, err := http.Get(link)
	HandleError(err)
	defer response.Body.Close()

	// Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	HandleError(err)
	DebugLog("Finished downloading.")
}

func SaveStringAsFile(contents string, file *os.File) {
	err := file.Chmod(0770)
	HandleError(err)

	// Write the bytes to the file
	_, err = io.Copy(file, strings.NewReader(contents))
	HandleError(err)
	DebugLog("Finished writing to file.")
}
