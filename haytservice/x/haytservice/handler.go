package haytservice

import (
	"fmt"

	"github.com/serjplus/cosmos-sdk-tutorials/haytservice/x/haytservice/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//###############################################################################################
/*
// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case MsgBuyName:
			return handleMsgBuyName(ctx, keeper, msg)
		case MsgDeleteName:
			return handleMsgDeleteName(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
*/

// NewHandler returns a handler for "haytservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetHayt:
			return handleMsgSetHayt(ctx, keeper, msg)
		case MsgBuyHayt:
			return handleMsgBuyHayt(ctx, keeper, msg)
		case MsgDeleteHayt:
			return handleMsgDeleteHayt(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized haytservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

//##########################################################################################################
/*
// Handle a message to set name
func handleMsgSetName(ctx sdk.Context, keeper Keeper, msg MsgSetName) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}
	keeper.SetName(ctx, msg.Name, msg.Value) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                      // return
}*/

// Handle a message to set hayt
func handleMsgSetHayt(ctx sdk.Context, keeper Keeper, msg MsgSetHayt) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Hayt)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}
	keeper.SetHayt(ctx, msg.Hayt, msg.Value) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                      // return
}

//###########################################################################################################
/*
// Handle a message to buy name
func handleMsgBuyName(ctx sdk.Context, keeper Keeper, msg MsgBuyName) sdk.Result {
	// Checks if the the bid price is greater than the price paid by the current owner
	if keeper.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) {
		return sdk.ErrInsufficientCoins("Bid not high enough").Result() // If not, throw an error
	}
	if keeper.HasOwner(ctx, msg.Name) {
		err := keeper.CoinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	} else {
		_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid) // If so, deduct the Bid amount from the sender
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	}
	keeper.SetOwner(ctx, msg.Name, msg.Buyer)
	keeper.SetPrice(ctx, msg.Name, msg.Bid)
	return sdk.Result{}
}*/

// Handle a message to buy hayt
func handleMsgBuyHayt(ctx sdk.Context, keeper Keeper, msg MsgBuyHayt) sdk.Result {
	// Checks if the the bid price is greater than the price paid by the current owner
	if keeper.GetPrice(ctx, msg.Hayt).IsAllGT(msg.Bid) {
		return sdk.ErrInsufficientCoins("Bid not high enough").Result() // If not, throw an error
	}
	if keeper.HasOwner(ctx, msg.Hayt) {
		err := keeper.CoinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Hayt), msg.Bid)
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	} else {
		_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid) // If so, deduct the Bid amount from the sender
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	}
	keeper.SetOwner(ctx, msg.Hayt, msg.Buyer)
	keeper.SetPrice(ctx, msg.Hayt, msg.Bid)
	return sdk.Result{}
}

//##################################################################################################################
/*
// Handle a message to delete name
func handleMsgDeleteName(ctx sdk.Context, keeper Keeper, msg MsgDeleteName) sdk.Result {
	if !keeper.IsNamePresent(ctx, msg.Name) {
		return types.ErrNameDoesNotExist(types.DefaultCodespace).Result()
	}
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	keeper.DeleteWhois(ctx, msg.Name)
	return sdk.Result{}
}*/

// Handle a message to delete Hayt
func handleMsgDeleteHayt(ctx sdk.Context, keeper Keeper, msg MsgDeleteHayt) sdk.Result {
	if !keeper.IsNamePresent(ctx, msg.Hayt) {
		return types.ErrNameDoesNotExist(types.DefaultCodespace).Result()
	}
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Hayt)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	keeper.DeleteWhois(ctx, msg.Hayt)
	return sdk.Result{}
}
