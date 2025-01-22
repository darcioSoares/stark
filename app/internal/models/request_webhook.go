package models

var RequestWebhook struct {
	Event struct {
		Created      string `json:"created"`
		ID           string `json:"id"`
		Subscription string `json:"subscription"`
		WorkspaceID  string `json:"workspaceId"`
		Log          struct {
			Type    string `json:"type"`
			Created string `json:"created"`
			Invoice struct {
				ID         string `json:"id"`
				Status     string `json:"status"`
				Amount     int    `json:"amount"`
				Name       string `json:"name"`
				TaxID      string `json:"taxId"`
				Created    string `json:"created"`
				Nominal    int    `json:"nominalAmount"`
				Link       string `json:"link"`
				PDF        string `json:"pdf"`
				Expiration int    `json:"expiration"`
			} `json:"invoice"`
		} `json:"log"`
	} `json:"event"`
}