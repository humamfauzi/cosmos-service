package appservice

import (
  "fmt"

  sdk "github.com/cosmos/cosmos-sdk/types"
)

// Function for interacting with chain
// Handler Construction
func NewHandler(keeper Keeper) sdk.Handler {
  return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
    switch msg := msg.(type) {
    case MsgSetName:
      return handleMsgSetName(ctx, keeper, msg)
    case MsgBuyName:
      return handleMsgBuyName(ctx, keeper, msg)
    default:
      return sdk.ErrUnkownRequest(errMsg).Result()
    }
  }
}

func handleMsgSetName(ctx sdk.Context, keeper Keeper, msg MsgSetName) sdk.Result {
  if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.NameID)) {
    return sdk.ErrUnauthorized("Incorrect Owner").Result()
  }

  keeper.SetName(ctx, msg.NameID, msg.Value)
  return sdk.Result{}
}

func handleMsgBuyName(ctx sdk.Context, keeper Keeper, msg MsgBuyName) sdk.Result {
  if keeper.GetPrice(ctx, msg.NameID).IsGTE(msg.Bid) {
    return sdk.ErrInsufficientCoins("Bid is too low").Result()
  }

  if keeper.HasOwner(ctx, msg.NameID) {
    _, err := keeper.coinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.NameID), msg.Bid)
    if err != nil {
      return sdk.ErrInsufficientCoins("Buyer is too poor")
    }

  } else {
    _, _, err := keeper.coinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid)
    if err != nil {
      return sdk.ErrInsufficientCoins("Buyer is too poor")
    }
  }
  keeper.SetOwner(ctx, msg.NameID, msg.Buyer)
  keeper.SetPrice(ctx, msg.NameID, msg.Bid)
  return sdk.Result{}
}
