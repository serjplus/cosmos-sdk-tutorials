package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MinHaytPrice is Initial Starting Price for a Hayt that was never previously owned
var MinHaytPrice = sdk.Coins{sdk.NewInt64Coin("verotoken", 1)}

/*
// Whois is a struct that contains all the metadata of a name
type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}*/

// Whois is a struct that contains all the metadata of a Hayt
type Whois struct {
	Value         string         `json:"value"`
	Owner         sdk.AccAddress `json:"owner"`
	Price         sdk.Coins      `json:"price"`
	HaytOwnerName string         `json:"haytownername"`
}

/*
// NewWhois returns a new Whois with the minprice as the price
func NewWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}*/

// NewWhois returns a new Whois with the minprice as the price
func NewWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}

// implement fmt.Stringer
func (w Whois) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
HaytOwnerName: %s
Price: %s`, w.Owner, w.Value, w.HaytOwnerName, w.Price))
}
