package client

import (
	"context"
	"net/url"

	"go.sia.tech/renterd/api"
	rhpv3 "go.thebigfile.com/core/rhp/v3"
	"go.thebigfile.com/core/types"
)

// Accounts returns all accounts.
func (c *Client) Accounts(ctx context.Context, owner string) (accounts []api.Account, err error) {
	values := url.Values{}
	values.Set("owner", owner)
	err = c.c.WithContext(ctx).GET("/accounts?"+values.Encode(), &accounts)
	return
}

func (c *Client) FundAccount(ctx context.Context, account rhpv3.Account, fcid types.FileContractID, amount types.Currency) (types.Currency, error) {
	var resp api.AccountsFundResponse
	err := c.c.WithContext(ctx).POST("/accounts/fund", api.AccountsFundRequest{
		AccountID:  account,
		Amount:     amount,
		ContractID: fcid,
	}, &resp)
	return resp.Deposit, err
}

// UpdateAccounts saves all accounts.
func (c *Client) UpdateAccounts(ctx context.Context, accounts []api.Account) (err error) {
	err = c.c.WithContext(ctx).POST("/accounts", api.AccountsSaveRequest{
		Accounts: accounts,
	}, nil)
	return
}
