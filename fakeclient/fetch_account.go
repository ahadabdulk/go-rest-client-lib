package fakeclient

import (
	"context"
	"fmt"
	"net/http"
)

//FetchResponse is Response (struct) for FETCH (GET)
type FetchResponse struct {
	Code int         //Status Code
	Data AccountData `json:"data"`
}

//AccountData is ...
type AccountData struct {
	Type           string               `json:"type"`
	ID             string               `json:"id"`
	OrganisationID string               `json:"organisation_id"`
	Version        int                  `json:"version"`
	Attributes     AccountAttributes    `json:"attributes"`
	Relationships  AccountRelationships `json:"relationships"`
}

//AccountAttributes is
type AccountAttributes struct {
	Country       string `json:"country"`
	BaseCurrency  string `json:"base_currency"`
	AccountNumber string `json:"account_number"`
	BankID        string `json:"bank_id"`
	BankIDCode    string `json:"bank_id_code"`
	Bic           string `json:"bic"`
	Iban          string `json:"iban"`
	Status        string `json:"status"`
}

//AccountRelationships is
type AccountRelationships struct {
	AccountEvents AccountEventsInfo `json:"account_events"`
}

//AccountEventsInfo is
type AccountEventsInfo struct {
	Data []AccountEventData `json:"data"`
}

//AccountEventData is
type AccountEventData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

//FetchAccount is to fetch the details of a specific account with given id
func (c *Client) FetchAccount(ctx context.Context, id string) (*FetchResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%sv1/organisation/accounts/%s", c.BaseURL, id), nil)

	if err != nil {
		return nil, err
	}

	res := FetchResponse{}
	/*if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}*/

	if rescode, err := c.sendRequest(ctx, req, &res); err != nil {
		return nil, err
	} else {
		res.Code = rescode
	}

	return &res, nil
}
