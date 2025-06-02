package client

import (
	"context"
)

var adminModule = "admin"

func (ec *Client) CommitValidatorsKeys(ctx context.Context, pubkeys []string, privates []string) error {
	err := ec.c.CallContext(ctx, nil, adminModule+"_commitValidatorsKeys", pubkeys, privates)
	return err
}

func (ec *Client) CommitReceivedAttestation(ctx context.Context, siginedAttestDataBase64 string) error {
	err := ec.c.CallContext(ctx, nil, adminModule+"_commitReceivedAttestation", siginedAttestDataBase64)
	return err
}
