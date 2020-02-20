package types

import (
	sdk "github.com/serjplus/cosmos-sdk/types"
)

// RouterKey is the module name router key
const RouterKey = ModuleName // this was defined in your key.go file

// MsgSetHayt defines a SetHayt message
type MsgSetHayt struct {
	Hayt          string         `json:"hayt"`
	Value         string         `json:"value"`
	Owner         sdk.AccAddress `json:"owner"`
	HaytOwnerName string         `json:"haytownername"`
}

/*
// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}*/

// NewMsgSetHayt is a constructor function for MsgSetHayt
func NewMsgSetHayt(hayt string, value string, owner sdk.AccAddress, haytownername string) MsgSetHayt {
	return MsgSetHayt{
		Hayt:          hayt,
		Value:         value,
		Owner:         owner,
		HaytOwnerName: haytownername,
	}
}

/*
// Route should return the name of the module
func (msg MsgSetName) Route() string { return RouterKey }
*/
// Route should return the name of the module
func (msg MsgSetHayt) Route() string { return RouterKey }

/*
// Type should return the action
func (msg MsgSetName) Type() string { return "set_name" }
*/

// Type should return the action
func (msg MsgSetHayt) Type() string { return "set_hayt" }

/*
// ValidateBasic runs stateless checks on the message
func (msg MsgSetName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}
*/

// ValidateBasic runs stateless checks on the message
func (msg MsgSetHayt) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Hayt) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Hayt and/or Value cannot be empty")
	}
	return nil
}

/*
// GetSignBytes encodes the message for signing
func (msg MsgSetName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}*/

// GetSignBytes encodes the message for signing
func (msg MsgSetHayt) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

/*
// GetSigners defines whose signature is required
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}*/

// GetSigners defines whose signature is required
func (msg MsgSetHayt) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

//#############################################################################################################################################
/*
// MsgBuyName defines the BuyName message
type MsgBuyName struct {
	Name  string         `json:"name"`
	Bid   sdk.Coins      `json:"bid"`
	Buyer sdk.AccAddress `json:"buyer"`
}
*/

// MsgBuyHayt defines the BuyHayt message
type MsgBuyHayt struct {
	Hayt          string         `json:"hayt"`
	Bid           sdk.Coins      `json:"bid"`
	Buyer         sdk.AccAddress `json:"buyer"`
	HaytOwnerName string         `json:"haytownername"`
}

/*
// NewMsgBuyName is the constructor function for MsgBuyName
func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyName {
	return MsgBuyName{
		Name:  name,
		Bid:   bid,
		Buyer: buyer,
	}
}*/

// NewMsgBuyHayt is the constructor function for MsgBuyHayt
func NewMsgBuyHayt(hayt string, bid sdk.Coins, buyer sdk.AccAddress, haytownername string) MsgBuyHayt {
	return MsgBuyHayt{
		Hayt:          hayt,
		Bid:           bid,
		Buyer:         buyer,
		HaytOwnerName: haytownername,
	}
}

/*
// Route should return the name of the module
func (msg MsgBuyName) Route() string { return RouterKey }
*/

// Route should return the name of the module
func (msg MsgBuyHayt) Route() string { return RouterKey }

/*
// Type should return the action
func (msg MsgBuyName) Type() string { return "buy_name" }
*/

// Type should return the action
func (msg MsgBuyHayt) Type() string { return "buy_hayt" }

/*
// ValidateBasic runs stateless checks on the message
func (msg MsgBuyName) ValidateBasic() sdk.Error {
	if msg.Buyer.Empty() {
		return sdk.ErrInvalidAddress(msg.Buyer.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
	return nil
}*/

// ValidateBasic runs stateless checks on the message
func (msg MsgBuyHayt) ValidateBasic() sdk.Error {
	if msg.Buyer.Empty() {
		return sdk.ErrInvalidAddress(msg.Buyer.String())
	}
	if len(msg.Hayt) == 0 {
		return sdk.ErrUnknownRequest("Hayt cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
	if len(msg.HaytOwnerName) == 0 {
		return sdk.ErrUnknownRequest("HaytOwnerName cannot be empty")
	}
	return nil
}

/*
// GetSignBytes encodes the message for signing
func (msg MsgBuyName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}*/

// GetSignBytes encodes the message for signing
func (msg MsgBuyHayt) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

/*
// GetSigners defines whose signature is required
func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}*/

// GetSigners defines whose signature is required
func (msg MsgBuyHayt) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

//##################################################################################################################################################
/*
// MsgDeleteName defines a DeleteName message
type MsgDeleteName struct {
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}*/

// MsgDeleteHayt defines a DeleteHayt message
type MsgDeleteHayt struct {
	Hayt          string         `json:"hayt"`
	Owner         sdk.AccAddress `json:"owner"`
	HaytOwnerName string         `json:"haytownername"`
}

/*
func NewMsgDeleteName(name string, owner sdk.AccAddress) MsgDeleteName {
	return MsgDeleteName{
		Name:  name,
		Owner: owner,
	}
}*/

// NewMsgDeleteHayt is a constructor function for MsgDeleteHayt
func NewMsgDeleteHayt(hayt string, owner sdk.AccAddress, haytownername string) MsgDeleteHayt {
	return MsgDeleteHayt{
		Hayt:          hayt,
		Owner:         owner,
		HaytOwnerName: haytownername,
	}
}

/*
// Route should return the name of the module
func (msg MsgDeleteName) Route() string { return RouterKey }
*/

// Route should return the name of the module
func (msg MsgDeleteHayt) Route() string { return RouterKey }

/*
// Type should return the action
func (msg MsgDeleteName) Type() string { return "delete_name" }
*/

// Type should return the action
func (msg MsgDeleteHayt) Type() string { return "delete_hayt" }

/*
// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	return nil
}*/

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteHayt) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.HaytOwnerName) == 0 {
		return sdk.ErrUnknownRequest("HaytOwnerName cannot be empty")
	}
	if len(msg.Hayt) == 0 {
		return sdk.ErrUnknownRequest("Hayt cannot be empty")
	}
	return nil
}

/*
// GetSignBytes encodes the message for signing
func (msg MsgDeleteName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}*/

// GetSignBytes encodes the message for signing
func (msg MsgDeleteHayt) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

/*
// GetSigners defines whose signature is required
func (msg MsgDeleteName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}*/

// GetSigners defines whose signature is required
func (msg MsgDeleteHayt) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
