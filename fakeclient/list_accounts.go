package fakeclient

import (
	"context"
	"fmt"
	"net/http"
)

//ListResponse is Response (struct) for LIST
type ListResponse struct {
	Code int
	Data []AccountsList `json:"data"`
}

//AccountsList is data in the list of accounts
type AccountsList struct {
	Type           string `json:"type"`
	ID             string `json:"id"`
	OrganisationID string `json:"organisation_id"`
	Version        int    `json:"version"`
	Attributes2    `json:"attributes"`
}

//Attributes2 are data attributes
type Attributes2 struct {
	Country               string `json:"country"`
	BaseCurrency          string `json:"base_currency"`
	AccountNumber         string `json:"account_number"`
	BankID                string `json:"bank_id"`
	BankIDCode            string `json:"bank_id_code"`
	Bic                   string `json:"bic"`
	Iban                  string `json:"iban"`
	AccountClassification string `json:"account_classification"`
	JointAccount          bool   `json:"joint_account"`
	Switched              bool   `json:"switched"`
	AccountMatchingOptOut bool   `json:"account_matching_opt_out"`
	Status                string `json:"status"`
}

//AccountListOptions is for optional fields
type AccountListOptions struct {
	PageNumber int
	PageSize   int
	*FilterOption
}

//FilterOption is used to get account details based on the filter applied
type FilterOption struct {
	Attribute string
	Value     string
}

//ListAccounts is to list the details of all accounts
func (c *Client) ListAccounts(ctx context.Context, options *AccountListOptions) (*ListResponse, error) {

	pageNumber := 0
	pageSize := 100
	filterAttribute := ""
	filterValue := ""
	url := fmt.Sprintf("%sv1/organisation/accounts", c.BaseURL)

	if options != nil {
		pageNumber = options.PageNumber
		pageSize = options.PageSize
		url = fmt.Sprintf("%s?page[number]=%d&page[size]=%d", url, pageNumber, pageSize)

		if options.FilterOption != nil {
			filterAttribute = options.FilterOption.Attribute
			filterValue = options.FilterOption.Value
			url = fmt.Sprintf("%s&filter[%s]=%s", url, filterAttribute, filterValue)
		}
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	res := []AccountsList{}
	res2 := ListResponse{}

	if rescode, err := c.sendRequest(ctx, req, &res); err != nil {
		return nil, err
	} else {
		res2 = ListResponse{
			Code: rescode,
			Data: res,
		}
	}

	return &res2, nil
}
