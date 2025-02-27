package keeper

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/umee-network/umee/v4/util/sdkutil"
	"github.com/umee-network/umee/v4/x/uibc"
)

var _ uibc.MsgServer = msgServer{}

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of MsgServer for the x/uibc
// module.
func NewMsgServerImpl(keeper Keeper) uibc.MsgServer {
	return &msgServer{keeper: keeper}
}

// GovUpdateQuota implements types.MsgServer
func (m msgServer) GovUpdateQuota(goCtx context.Context, msg *uibc.MsgGovUpdateQuota) (
	*uibc.MsgGovUpdateQuotaResponse, error,
) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// validate the msg
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := m.keeper.UpdateQuotaParams(ctx, msg.Total, msg.PerDenom, msg.QuotaDuration); err != nil {
		return nil, err
	}
	return &uibc.MsgGovUpdateQuotaResponse{}, nil
}

// GovSetIBCStatus implements types.MsgServer
func (m msgServer) GovSetIBCStatus(
	goCtx context.Context, msg *uibc.MsgGovSetIBCStatus,
) (*uibc.MsgGovSetIBCStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := m.keeper.SetIBCStatus(ctx, msg.IbcStatus); err != nil {
		return &uibc.MsgGovSetIBCStatusResponse{}, err
	}
	sdkutil.Emit(&ctx, &uibc.EventIBCTransferStatus{
		Status: msg.IbcStatus,
	})

	return &uibc.MsgGovSetIBCStatusResponse{}, nil
}
