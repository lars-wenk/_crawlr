package model

type ComdirectAuth struct {
	OAuthURL      string `json:"oauthurl" `
	Client_id     string `json:"client_id"`
	Secret_id     string `json:"secret_id"`
	Zugangsnummer string `json:"zugangsnummer"`
	Pin           string `json:"pin"`
}

type ComdirectRequestInformation struct {
	ClientRequestID struct {
		SessionID string `json:"sessionId"`
		RequestID string `json:"requestId"`
	} `json:"clientRequestId"`
}

type ComdirectAuthResponse struct {
	AccessToken       string `json:"access_token"`
	TokenType         string `json:"token_type"`
	RefreshToken      string `json:"refresh_token"`
	ExpiresIn         int    `json:"expires_in"`
	Scope             string `json:"scope"`
	Kdnr              string `json:"kdnr"`
	Bpid              int    `json:"bpid"`
	KontaktID         int64  `json:"kontaktId"`
	Error             string `json:"error"`
	Error_description string `json:"error_description"`
}

type ComdirectSessions []struct {
	Identifier       string `json:"identifier"`
	SessionTanActive bool   `json:"sessionTanActive"`
	Activated2FA     bool   `json:"activated2FA"`
}

type ComdirectSession struct {
	Identifier       string `json:"identifier"`
	SessionTanActive bool   `json:"sessionTanActive"`
	Activated2FA     bool   `json:"activated2FA"`
}

type ComdirectBalance struct {
	Paging struct {
		Index   int `json:"index"`
		Matches int `json:"matches"`
	} `json:"paging"`
	Values []struct {
		Account struct {
			AccountID        string `json:"accountId"`
			AccountDisplayID string `json:"accountDisplayId"`
			Currency         string `json:"currency"`
			ClientID         string `json:"clientId"`
			AccountType      struct {
				Key  string `json:"key"`
				Text string `json:"text"`
			} `json:"accountType"`
			Iban        string `json:"iban"`
			CreditLimit struct {
				Value string `json:"value"`
				Unit  string `json:"unit"`
			} `json:"creditLimit"`
		} `json:"account"`
		AccountID string `json:"accountId"`
		Balance   struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"balance"`
		BalanceEUR struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"balanceEUR"`
		AvailableCashAmount struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"availableCashAmount"`
		AvailableCashAmountEUR struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"availableCashAmountEUR"`
	} `json:"values"`
}

type ComdirectDepot struct {
	Paging struct {
		Index   int `json:"index"`
		Matches int `json:"matches"`
	} `json:"paging"`
	Aggregated struct {
		Depot struct {
			DepotID                    string        `json:"depotId"`
			DepotDisplayID             string        `json:"depotDisplayId"`
			ClientID                   string        `json:"clientId"`
			DefaultSettlementAccountID string        `json:"defaultSettlementAccountId"`
			SettlementAccountIds       []interface{} `json:"settlementAccountIds"`
		} `json:"depot"`
		PrevDayValue struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"prevDayValue"`
		CurrentValue struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"currentValue"`
		PurchaseValue struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"purchaseValue"`
		ProfitLossPurchaseAbs struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"profitLossPurchaseAbs"`
		ProfitLossPurchaseRel string `json:"profitLossPurchaseRel"`
		ProfitLossPrevDayAbs  struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"profitLossPrevDayAbs"`
		ProfitLossPrevDayRel string `json:"profitLossPrevDayRel"`
	} `json:"aggregated"`
	Values []struct {
		DepotID     string `json:"depotId"`
		PositionID  string `json:"positionId"`
		Wkn         string `json:"wkn"`
		CustodyType string `json:"custodyType"`
		Quantity    struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"quantity"`
		AvailableQuantity struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"availableQuantity"`
		CurrentPrice struct {
			Price struct {
				Value string `json:"value"`
				Unit  string `json:"unit"`
			} `json:"price"`
			PriceDateTime string `json:"priceDateTime"`
		} `json:"currentPrice"`
		PurchasePrice struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"purchasePrice"`
		PrevDayPrice struct {
			Price struct {
				Value string `json:"value"`
				Unit  string `json:"unit"`
			} `json:"price"`
			PriceDateTime string `json:"priceDateTime"`
		} `json:"prevDayPrice"`
		CurrentValue struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"currentValue"`
		PurchaseValue struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"purchaseValue"`
		ProfitLossPurchaseAbs struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"profitLossPurchaseAbs"`
		ProfitLossPurchaseRel string `json:"profitLossPurchaseRel"`
		ProfitLossPrevDayAbs  struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"profitLossPrevDayAbs"`
		ProfitLossPrevDayRel     string      `json:"profitLossPrevDayRel"`
		Version                  interface{} `json:"version"`
		Hedgeability             string      `json:"hedgeability"`
		AvailableQuantityToHedge struct {
			Value string `json:"value"`
			Unit  string `json:"unit"`
		} `json:"availableQuantityToHedge"`
	} `json:"values"`
}
