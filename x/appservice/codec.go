package appservice

import (
  "github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
  cdc.RegisterConcrete(MsgBuyName{}, "appservice/BuyName", nil)
  cdc.RegisterConcrete(MsgSetName{}, "appservice/SetName", nil)
}
