package client

import (
	"context"
)

var attestModule = "attest"

func (ec *Client) AttestBeforeBroadCast(ctx context.Context, slot uint64) (AttackerResponse, error) {
	var result AttackerResponse
	err := ec.c.CallContext(ctx, &result, attestModule+"_beforeBroadCast", slot)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (ec *Client) AttestAfterBroadCast(ctx context.Context, slot uint64) (AttackerResponse, error) {
	var result AttackerResponse
	err := ec.c.CallContext(ctx, &result, attestModule+"_afterBroadCast", slot)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (ec *Client) AttestBeforeSign(ctx context.Context, slot uint64, pubkey string, attestDataBase64 string) (AttackerResponse, error) {
	var result AttackerResponse
	err := ec.c.CallContext(ctx, &result, attestModule+"_beforeSign", slot, pubkey, attestDataBase64)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (ec *Client) AttestAfterSign(ctx context.Context, slot uint64, pubkey string, siginedAttestDataBase64 string) (AttackerResponse, error) {
	var result AttackerResponse
	err := ec.c.CallContext(ctx, &result, attestModule+"_afterSign", slot, pubkey, siginedAttestDataBase64)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (ec *Client) AttestBeforePropose(ctx context.Context, slot uint64, pubkey string, siginedAttestDataBase64 string) (AttackerResponse, error) {
	var result AttackerResponse
	err := ec.c.CallContext(ctx, &result, attestModule+"_beforePropose", slot, pubkey, siginedAttestDataBase64)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (ec *Client) AttestAfterPropose(ctx context.Context, slot uint64, pubkey string, siginedAttestDataBase64 string) (AttackerResponse, error) {
	var result AttackerResponse
	err := ec.c.CallContext(ctx, &result, attestModule+"_afterPropose", slot, pubkey, siginedAttestDataBase64)
	if err != nil {
		return result, err
	}
	return result, nil
}
