package fakeclient

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

//CreateResponse is the Response (struct) for CREATE
type CreateResponse struct {
	Code int
	Data `json:"data"`
}

//Data is Response data for Create Request
type Data struct {
	Type           string `json:"type"`
	ID             string `json:"id"`
	Version        int    `json:"version"`
	OrganisationID string `json:"organisation_id"`
	Attributes     `json:"attributes"`
	Relationships  `json:"relationships"`
}

//Attributes are data attributes
type Attributes struct {
	Country               string   `json:"country"`
	BaseCurrency          string   `json:"base_currency"`
	AccountNumber         string   `json:"account_number"`
	BankID                string   `json:"bank_id"`
	BankIDCode            string   `json:"bank_id_code"`
	Bic                   string   `json:"bic"`
	Name                  []string `json:"name"`
	PrivateIdentification `json:"private_identification"`
	Status                string `json:"status"`
}

//Relationships is ...
type Relationships struct {
	MasterAccount `json:"master_account"`
	AccountEvents `json:"account_events"`
}

//PrivateIdentification is ...
type PrivateIdentification struct {
	BirthDate      string   `json:"birth_date"`
	BirthCountry   string   `json:"birth_country"`
	Identification string   `json:"identification"`
	Address        []string `json:"address"`
	City           string   `json:"city"`
	Country        string   `json:"country"`
}

//MasterAccount is ...
type MasterAccount struct {
	Data []MasterAccountData `json:"data"`
}

//MasterAccountData is ...
type MasterAccountData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

//AccountEvents is
type AccountEvents struct {
	MasterAccountData `json:"data"`
}

//CreateAccount is to create new account
func (c *Client) CreateAccount(ctx context.Context, js string) (*CreateResponse, error) {

	var jsonStr = []byte(js)
	req, err := http.NewRequest("POST", fmt.Sprintf("%sv1/organisation/accounts", c.BaseURL), bytes.NewBuffer(jsonStr))

	if err != nil {
		return nil, err
	}

	res := CreateResponse{}

	if rescode, err := c.sendRequest(ctx, req, &res); err != nil {
		return nil, err
	} else {
		res.Code = rescode
	}

	return &res, nil
}
