<p align="center">
<h1 align="center">fakeclient</h1>
<p align="center">Simple REST client library for Go</p>

## About me:

  * Name: Abdul Ahad
  * Golang skills: Intermediate

## Library Features

  * CREATE, FETCH, LIST, DELETE operations on fake account api
  * Support Pagination
  * Support Attributes


## Usage

#### Simple GET / Fetch
##### func (c *Client) FetchAccount(ctx context.Context, id string) (*FetchResponse, error)
```go
// Create a Client
client := NewClient()

ctx := context.Background()

// Account id to be fetched (string)
id := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

res, err := c.FetchAccount(ctx, id)
```

#### Advance GET/ List
##### func (c *Client) ListAccounts(ctx context.Context, options *AccountListOptions) (*ListResponse, error)

```go
	c := NewClient()
    ctx := context.Background()
    
    //Also support option of Paging and filtering
    op := AccountListOptions{
            PageNumber: 1, 
            PageSize: 20, 
            &FilterOptions
                { attribute:"country",
                  value: "IN"  
                }
            }

    res, err := c.ListAccounts(ctx, &op)
    
    //without options
    res, err := c.ListAccounts(ctx, nil)
```

#### POST / Create
##### func (c *Client) CreateAccount(ctx context.Context, js string) (*CreateResponse, error)

```go
    // Create a client
    c := NewClient()
    ctx := context.Background()
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

	//CreateAccount takes json body in POST request
	res, err := c.CreateAccount(ctx, json_body)
```

#### DELETE
##### func (c *Client) DeleteAccount(ctx context.Context, id string, ver string) (*DeleteResponse, error)

```go
    
    // Create a client
    c := NewClient()
    ctx := context.Background()

    //Account id to be deleted
    id := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

    //Version
    ver := "0"  
    
	res, err := c.DeleteAccount(ctx, id, ver)
```

##### Test Cases covered
```go
/*
=== RUN   TestCreateAccount
--- PASS: TestCreateAccount (0.02s)
=== RUN   TestFetchAccount
--- PASS: TestFetchAccount (0.01s)
=== RUN   TestListAccounts
--- PASS: TestListAccounts (0.01s)
=== RUN   TestPaging
--- PASS: TestPaging (0.03s)
=== RUN   TestDeleteAccount
--- PASS: TestDeleteAccount (0.03s)
=== RUN   TestErrorResponse
--- PASS: TestErrorResponse (0.02s)
=== RUN   TestClearAllTestData
--- PASS: TestClearAllTestData (0.03s)
PASS
*/
```

