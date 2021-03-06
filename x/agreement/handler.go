package agreement

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgNewAgreement:
			return handleMsgNewAgreement(ctx, keeper, msg)
		
		case types.MsgAmendAgreement:
			return handleMsgAmendAgreement(ctx, keeper, msg)
		
		case types.MsgRenewAgreement:
			return handleMsgRenewwAgreement(ctx, keeper, msg)

		case types.MsgTerminateAgreement:
			return handleMsgTerminateAgreement(ctx, keeper, msg)

		case types.MsgExpireAgreement:
			return handleMsgExpireAgreement(ctx, keeper, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized agreement message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgNewAgreement(ctx sdk.Context,  keeper Keeper, msg types.MsgNewAgreement) sdk.Result {
	order, err := types.NewAgreement(msg.Source, msg.Destination, msg.Owner, ctx.BlockTime(), msg.ClientOrderId)
	if err != nil {
		return err.Result()
	}


	return keeper.NewAgreementSingle(ctx, agreement)
}

func handleMsgAmendAgreement(ctx sdk.Context,  keeper Keeper, msg types.MsgAmendAgreement) sdk.Result {

	return keeper.AmendAgreement(ctx, msg.Owner, msg.AgreementId)
}

func handleMsgRenewAgreement(ctx sdk.Context, keeper Keeper, msg types.MsgRenewAgreement) sdk.Result {

	return keeper.RenewAgreement(ctx, msg.Owner, msg.AgreementId)
}

func handleMsgTerminateAgreement(ctx sdk.Context,  keeper Keeper, msg types.MsgTerminateAgreement) sdk.Result {

	return keeper.TerminateAgreement(ctx, msg.Owner, msg.AgreementId)
}

func handleMsgExpireAgreement(ctx sdk.Context,  keeper Keeper, msg types.MsgExpireAgreement) sdk.Result {

	return keeper.ExpireAgreement(ctx, msg.Owner, msg.AgreementId)
}