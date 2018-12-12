package appservice

import (
  "github.com/cosmos/cosmos-sdk/codec"
  "github.com/cosmos/cosmos-sdk/x/bank"

  sdk "github.com/cosmos/cosmos-sdk/types"
)
/*

  Keeper
  L [names] value <string>
  L [names] owners <types.AccAdress>
  L [names] price <types.Coins>

  > cosmos-sdk/types to binary should use cdc.MustMarshalBinary and vice versa
*/
type Keeper struct {
  coinKeeper bank.Keeper

  namesStoreKey  sdk.StoreKey
  ownersStoreKey sdk.StoreKey
  pricesStoreKey sdk.StoreKey

  cdc *codec.Codec
}

func NewKeeper(coinKeeper bank.Keeper, namesStoreKey sdk.StoreKey, ownersStoreKey sdk.StoreKey, pricesStoreKey sdk.StoreKey cdc *codec.Codec) Keeper {
  return Keeper{
    coinKeeper:     coinKeeper,
    namesStoreKey:  namesStoreKey,
    ownersStoreKey: ownersStoreKey,
    pricesStoreKey: pricesStoreKey,
    cdc:            cdc,
  }
}

func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
  // Returns the string where the names belong to
  store := ctx.KVStore(k.namesStoreKey)
  bz := store.Get([]byte(name))
  return string(bz)
}

func (k Keeper) SetName(ctx sdk.Context, name string, value string) {
  // Set name and value pair
  store := ctx.KVStore(k.namesStoreKey)
  store.Set([]byte(name), []byte(value))
}

func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
  // Exploring whether a certain name has owner
  store := ctx.KVStore(k.ownersStoreKey)
  bz := store.Get([]byte(name))
  return bz != nil
}

func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAdress {
  store := ctx.KVStore(k.ownersStoreKey)
  bz := store.Get([]byte(name))
  return bz
}

func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAdress) {
  store := ctx.KVStore(k.ownersStoreKey)
  store.Set([]byte(name), owner)
}

func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
  if !k.HasOwner(ctx, name) {
    return sdk.Coins{sdk.NewInt64Coin("mycoin", 1)}
  }
  store := ctx.KVStore(k.pricesStoreKey)
  bz := store.Get([]byte(name))
  var price sdk.Coins
  k.cdc.MustUnmarshalBinary(bz, &price)
  return price
}

func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
  store := ctx.KVStore([]byte(name))
  store.Set([]byte(name), k.cdc.MustMarshalBinary(price))
}
