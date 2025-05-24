package client

import (
	"context"
)

var adminModule = "admin"

func (ec *Client) CommitValidatorsKeys(ctx context.Context, pubkeys []string, privates []string) error {
	err := ec.c.CallContext(ctx, nil, adminModule+"_commitValidatorsKeys", pubkeys, privates)
	return err
}
