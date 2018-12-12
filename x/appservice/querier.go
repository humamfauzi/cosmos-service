package appservice

import (
  "github.com/cosmos/cosmos-sdk/codec"

  sdk "github.com/cosmos/cosmos-sdk/types"
  abci "github.com/tendermint/tendermint/abci/types"
)

const (
  // list of query task
  QueryResolve = "resolve"
  QueryWhois   = "whois"
)

// Query constructor for every task
func NewQuerier(keeper Keeper) sdk.Querirer {
  return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
    switch path[0] {
    case QueryResolve:
      return queryResolve(ctx, path[1:], req, keeper), nil
    case QueryWhois:
      return queryWhois(ctx, path[1:], req, keeper), nil
    default:
      return nil, sdk.ErrUnkownRequest("unknown query")
    }
  }
}

func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
  name := path[0]
  value := keeper.ResolveName(ctx, name)

  if value == "" {
    return []byte{}, sdk.ErrUnkownRequest("could not resolve name")
  }

  return []byte(value), nil
}

type Whois struct {
  Value string        `json:"value"`
  Owner sdk.AccAdress `json:"owner"`
  Price sdk.Coins     `json:"price"`
}

func queryWhois(ctx sdk.Context, path []string, req abci.RequstQuery, keeper Keeper) (res []byte, err sdk.Error) {
  name := path[0]
  whois := Whois{
    Value: keeper.ResolveName(ctx, name),
    Owner: keeper.GetOwner(ctx, name),
    Price: keeper.GetPrice(ctx, name),
  }

  bz, err2 := codec.MarshatJSONIndent(keeper.cdc, whois)
  if err2 != nil {
    panic("Could not marshal to JSON")
  }
  
  return bz, nil
}
