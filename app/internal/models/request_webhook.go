package models

type RequestWebhook struct {
	ID           string `json:"id"`
	IsDelivered  bool   `json:"isDelivered"`
	Subscription string `json:"subscription"`
	Created      string `json:"created"`
	Log          struct {
		ID       string   `json:"id"`
		Errors   []string `json:"errors"`
		Type     string   `json:"type"`
		Created  string   `json:"created"`
		Transfer struct {
			ID             string   `json:"id"`
			Status         string   `json:"status"`
			Amount         int      `json:"amount"`
			Name           string   `json:"name"`
			BankCode       string   `json:"bankCode"`
			BranchCode     string   `json:"branchCode"`
			AccountNumber  string   `json:"accountNumber"`
			TaxID          string   `json:"taxId"`
			Tags           []string `json:"tags"`
			Created        string   `json:"created"`
			Updated        string   `json:"updated"`
			TransactionIds []string `json:"transactionIds"`
			Fee            int      `json:"fee"`
		} `json:"transfer"`
	} `json:"log"`
}

type RequestWebhooktPayload struct {
	Event RequestWebhook `json:"event"`
}