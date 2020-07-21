// +build integration

package fakeclient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {

	json_body := `{
		"data": {
		  "type": "accounts",
		  "id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		  "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		  "attributes": {
			"country": "GB",
			"base_currency": "GBP",
			"bank_id": "400300",
			"bank_id_code": "GBDSC",
			"bic": "NWBKGB22"
		  }
		}
	  }`

	ctx := context.Background()

	c := NewClient()
	res, err := c.CreateAccount(ctx, json_body)
	assert.Nil(t, err, "Expected nil error")
	assert.NotNil(t, res, "Expected non-nil response")
	assert.EqualValues(t, 201, res.Code, "Expected 201, for resource created successfully")
}

func TestFetchAccount(t *testing.T) {

	id := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	ctx := context.Background()

	c := NewClient()
	res, err := c.FetchAccount(ctx, id)
	assert.Nil(t, err, "Expected nil error")
	assert.NotNil(t, res, "Expected non-nil response")
}

func TestListAccounts(t *testing.T) {

	c := NewClient()
	ctx := context.Background()

	res, err := c.ListAccounts(ctx, nil)

	assert.Nil(t, err, "Expected nil error")
	assert.NotNil(t, res, "Expected non-nil response")

	if res != nil {
		assert.Equal(t, 1, len(res.Data), "1 account expected")
	}
}

func TestPaging(t *testing.T) {

	//adding second entry
	json_body := `{
		"data": {
		  "type": "accounts",
		  "id": "ad27e265-9605-5b5b-a0e5-3003ea9cc5ef",
		  "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		  "attributes": {
			"country": "IN",
			"base_currency": "INR",
			"bank_id": "600301",
			"bank_id_code": "HDFCI",
			"bic": "NWBINH22"
		  }
		}
	  }`

	c := NewClient()
	ctx := context.Background()
	res, err := c.CreateAccount(ctx, json_body)

	assert.Nil(t, err, "Expected nil error")
	assert.NotNil(t, res, "Expected non-nil response")

	op := AccountListOptions{PageNumber: 1, PageSize: 1}

	res2, err2 := c.ListAccounts(ctx, &op)

	assert.Nil(t, err2, "Expected nil error")
	assert.NotNil(t, res2, "Expected non-nil response")

	if res != nil {
		assert.Equal(t, 1, len(res2.Data), "1 account out of 2 expected")
	}

}

func TestDeleteAccount(t *testing.T) {

	id := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	ver := "0"
	ctx := context.Background()

	c := NewClient()
	res, _ := c.DeleteAccount(ctx, id, ver)

	assert.NotNil(t, res, "nil response expected for delete")
	assert.EqualValues(t, 204, res.Code, "Expected 204 response code, resource deleted successfully")

	res2, _ := c.FetchAccount(ctx, id)
	if res2 != nil {
		assert.EqualValues(t, 404, res2.Code, "Expected 404 response code, resource does not exist")
	}

}

func TestErrorResponse(t *testing.T) {
	//Accssing an already deleted account wii give 404
	id := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	ctx := context.Background()

	c := NewClient()
	_, err := c.FetchAccount(ctx, id)

	if err.Error() != "record ad27e265-9605-4b4b-a0e5-3003ea9cc4dc does not exist" {
		t.Errorf("Expected error: record ad27e265-9605-5b5b-a0e5-3003ea9cc5ef does not exist, but got %v", err.Error())
	}
}

func TestClearAllTestData(t *testing.T) {
	//Deleting if any remaining account
	id := "ad27e265-9605-5b5b-a0e5-3003ea9cc5ef"
	ver := "0"
	ctx := context.Background()

	c := NewClient()
	res, _ := c.DeleteAccount(ctx, id, ver)
	assert.NotNil(t, res, "nil response expected for delete")

	res2, _ := c.ListAccounts(ctx, nil)

	//Check if the list is empty
	if res2 != nil {
		assert.Equal(t, 0, len(res2.Data), "0 account out of 2 expected")
	}
}
