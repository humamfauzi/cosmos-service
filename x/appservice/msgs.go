package appservice

import (
  "encoding/json"

  sdk "github.com/cosmos/cosmos-sdk/types"
)
/*

general interface of msgs
  Route()         string
  Type()          string
  ValidateBasic() bool
  GetSignBytes()  []byte
  GetSigners()    []sdk.AccAdress

*/

// Handler for SetName Action
type MsgSetName struct {
  NameID string
  Value  string
  Owner  sdk.AccAdress
}

func NewMsgSetName(name string, value string, owner sdk.AccAdress) MsgSetName {
  return MsgSetName{
    NameID: name,
    Value:  value,
    Owner:  owner,
  }
}

func (msg MsgSetName) Route() string { return "serviceapp"}

func (msg MsgSetName) Type() string { return "set_name"}

func (msg MsgSetName) ValidateBasic() sdk.Error {
  if msg.Owner.Empty() {
    return sdk.ErrInvalidAddress(msg.Owner.String())
  }

  if len(msg.NameID) == 0 || len(msg.Value) == 0{
    return sdk.ErrUnkownRequest("Name or Value is missing")
  }

  return nil
}

func (msg MsgSetName) GetSignBytes() []byte {
  b, err := json.Marshal(msg)
  if err != nil {
    panic(err)
  }
  return sdk.MustSortJSON(b)
}

func (msg MsgSetName) GetSigners() []sdk.AccAdress {
  return []sdk.Address{msg.Owner}
}

// Handler for BuyName Action
type MsgBuyName struct {
  NameID string
  Bid sdk.Coins
  Buyer sdk.AccAdress
}

func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAdress) MsgBuyName {
  return MsgBuyName{
    NameID: name,
    Bid: bid,
    Buyer: buyer,
  }
}

func (msg MsgBuyName) Route() string { return "serviceapp" }

func (msg MsgBuyName) Type() string { return "buy_name"}

func (msg MsgBuyName) ValidateBasic() sdk.Error {
  if msg.Buyer.Empty() {
    return sdk.ErrInvalidAddress(msg.Buyer.String())
  }

  if len(msg.NameID) == 0 {
    return sdk.ErrUnkownRequest("Name cannot empty")
  }

  if !msg.Bid.IsPositive() {
    return sdk.ErrInsufficientCoins("Bid must be postive")
  }

  return nil
}

func (msg MsgBuyName) GetSignBytes() []byte {
  b, err := json.Marshal(msg)
  if err != nil {
    panic(err)
  }

  return sdk.MustSortJSON(b)
}

func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{msg.Buyer}
}
