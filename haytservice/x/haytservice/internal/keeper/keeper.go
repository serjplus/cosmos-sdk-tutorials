package keeper

import (
	"github.com/serjplus/cosmos-sdk-tutorials/haytservice/x/haytservice/internal/types"
	"github.com/serjplus/cosmos-sdk/codec"
	sdk "github.com/serjplus/cosmos-sdk/types"
	"github.com/serjplus/cosmos-sdk/x/bank"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	CoinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the haytservice Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// Gets the entire Whois metadata struct for a name
func (k Keeper) GetWhois(ctx sdk.Context, hayt string) types.Whois {
	store := ctx.KVStore(k.storeKey)
	if !k.IsNamePresent(ctx, hayt) {
		return types.NewWhois()
	}
	bz := store.Get([]byte(hayt))
	var whois types.Whois
	k.cdc.MustUnmarshalBinaryBare(bz, &whois)
	return whois
}

// Sets the entire Whois metadata struct for a name
func (k Keeper) SetWhois(ctx sdk.Context, hayt string, whois types.Whois) {
	if whois.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(hayt), k.cdc.MustMarshalBinaryBare(whois))
}

// Deletes the entire Whois metadata struct for a name
func (k Keeper) DeleteWhois(ctx sdk.Context, hayt string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(hayt))
}

// ResolveName - returns the string that the hayt resolves to
func (k Keeper) ResolveName(ctx sdk.Context, hayt string) string {
	return k.GetWhois(ctx, hayt).Value
}

// SetHayt - sets the value string that a hayt resolves to
func (k Keeper) SetHayt(ctx sdk.Context, hayt string, value string) {
	whois := k.GetWhois(ctx, hayt)
	whois.Value = value
	k.SetWhois(ctx, hayt, whois)
}

// HasOwner - returns whether or not the hayt already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, hayt string) bool {
	return !k.GetWhois(ctx, hayt).Owner.Empty()
}

// GetOwner - get the current owner of a hayt
func (k Keeper) GetOwner(ctx sdk.Context, hayt string) sdk.AccAddress {
	return k.GetWhois(ctx, hayt).Owner
}

// SetOwner - sets the current owner of a hayt
func (k Keeper) SetOwner(ctx sdk.Context, hayt string, owner sdk.AccAddress) {
	whois := k.GetWhois(ctx, hayt)
	whois.Owner = owner
	k.SetWhois(ctx, hayt, whois)
}

// GetPrice - gets the current price of a hayt
func (k Keeper) GetPrice(ctx sdk.Context, hayt string) sdk.Coins {
	return k.GetWhois(ctx, hayt).Price
}

// SetPrice - sets the current price of a hayt
func (k Keeper) SetPrice(ctx sdk.Context, hayt string, price sdk.Coins) {
	whois := k.GetWhois(ctx, hayt)
	whois.Price = price
	k.SetWhois(ctx, hayt, whois)
}

// Get an iterator over all hayts in which the keys are the hayts and the values are the whois
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, hayt string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(hayt))
}
