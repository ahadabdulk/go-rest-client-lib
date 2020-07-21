package fakeclient

import (
	"context"
	"fmt"
	"net/http"
)

//DeleteResponse is the Respose (struct) on DELETE operation (which is NIL)
type DeleteResponse struct {
	Code int
}

//DeleteAccount method deletes an account based on the provided id. version is mandatory here
func (c *Client) DeleteAccount(ctx context.Context, id string, ver string) (*DeleteResponse, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%sv1/organisation/accounts/%s?version=%s", c.BaseURL, id, ver), nil)

	if err != nil {
		return nil, err
	}

	res := DeleteResponse{}

	if rescode, err := c.sendRequest(ctx, req, &res); err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		res.Code = rescode
	}

	return &res, nil
}
