package client

import (
	"context"
)

var blockModule = "block"

func (ec *Client) BlockGetNewParentRoot(ctx context.Context, slot uint64, pubkey string, parentRoot string) (AttackerResponse, error) {
	var result AttackerResponse
	err := ec.c.CallContext(ctx, &result, blockModule+"_getNewParentRoot", slot, pubkey, parentRoot)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (ec *Client) DelayForReceiveBlock(ctx context.Context, slot uint64) (AttackerResponse, error) {
	var result AttackerResponse
	err := ec.c.CallContext(ctx, &result, blockModule+"_delayForReceiveBlock", slot)
	if err != nil {
		return result, err
	}
	return result, nil

}

func (ec *Client) BlockBeforeBroadCast(ctx context.Context, slot uint64) (AttackerResponse, error) {
	var result AttackerResponse
	err := ec.c.CallContext(ctx, &result, blockModule+"_beforeBroadCast", slot)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (ec *Client) BlockAfterBroadCast(ctx context.Context, slot uint64) (AttackerResponse, error) {
	var result AttackerResponse
	err := ec.c.CallContext(ctx, &result, blockModule+"_afterBroadCast", slot)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (ec *Client) BlockBeforeSign(ctx context.Context, slot uint64, pubkey string, blockDataBase64 string) (AttackerResponse, error) {
	var result AttackerResponse
	err := ec.c.CallContext(ctx, &result, blockModule+"_beforeSign", slot, pubkey, blockDataBase64)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (ec *Client) BlockAfterSign(ctx context.Context, slot uint64, pubkey string, siginedBlockDataBase64 string) (AttackerResponse, error) {
	var result AttackerResponse
	err := ec.c.CallContext(ctx, &result, blockModule+"_afterSign", slot, pubkey, siginedBlockDataBase64)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (ec *Client) BlockBeforePropose(ctx context.Context, slot uint64, pubkey string, siginedBlockDataBase64 string) (AttackerResponse, error) {
	var result AttackerResponse
	err := ec.c.CallContext(ctx, &result, blockModule+"_beforePropose", slot, pubkey, siginedBlockDataBase64)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (ec *Client) BlockAfterPropose(ctx context.Context, slot uint64, pubkey string, siginedBlockDataBase64 string) (AttackerResponse, error) {
	var result AttackerResponse
	err := ec.c.CallContext(ctx, &result, blockModule+"_afterPropose", slot, pubkey, siginedBlockDataBase64)
	if err != nil {
		return result, err
	}
	return result, nil
}
