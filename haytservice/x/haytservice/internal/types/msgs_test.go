package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var hayt = "maTurtle"

func TestMsgSetHayt(t *testing.T) {
	value := "1"
	acc := sdk.AccAddress([]byte("me"))
	var msg = NewMsgSetHayt(hayt, value, acc, haytownername)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "set_hayt")
}

func TestMsgSetHaytValidation(t *testing.T) {
	value := "1"
	acc := sdk.AccAddress([]byte("me"))
	hayt2 := "a"
	value2 := "2"
	acc2 := sdk.AccAddress([]byte("you"))

	cases := []struct {
		valid bool
		tx    MsgSetHayt
	}{
		{true, NewMsgSetHayt(hayt, value, acc, haytownername)},
		{true, NewMsgSetHayt(hayt2, value2, acc2, haytownername)},
		{true, NewMsgSetHayt(hayt2, value, acc2, haytownername)},
		{true, NewMsgSetHayt(hayt2, value2, acc)},
		{false, NewMsgSetHayt(hayt, value2, nil)},
		{false, NewMsgSetHayt("", value2, acc2)},
		{false, NewMsgSetHayt(hayt, "", acc2)},
	}

	for _, tc := range cases {
		err := tc.tx.ValidateBasic()
		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgSetNameGetSignBytes(t *testing.T) {
	value := "1"
	acc := sdk.AccAddress([]byte("me"))

	var msg = NewMsgSetHayt(hayt, value, acc, haytownername)
	res := msg.GetSignBytes()

	expected := `{"type":"haytservice/SetHayt","value":{"hayt":"maTurtle","owner":"cosmos1d4js690r9j","value":"1"}}`

	require.Equal(t, expected, string(res))
}

func TestMsgBuyHayt(t *testing.T) {
	coins := sdk.NewCoins(sdk.NewInt64Coin("atom", 10))
	acc := sdk.AccAddress([]byte("me"))
	var msg = NewMsgBuyHayt(hayt, coins, acc)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "buy_hayt")
}

func TestMsgBuyNameValidation(t *testing.T) {
	acc := sdk.AccAddress([]byte("me"))
	hayt2 := "a"
	acc2 := sdk.AccAddress([]byte("you"))
	coins := sdk.NewCoins(sdk.NewInt64Coin("atom", 10))

	cases := []struct {
		valid bool
		tx    MsgBuyHayt
	}{
		{true, NewMsgBuyHayt(hayt, coins, acc)},
		{true, NewMsgBuyHayt(hayt2, coins, acc2, haytownername)},
	}

	for _, tc := range cases {
		err := tc.tx.ValidateBasic()
		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgBuyNameGetSignBytes(t *testing.T) {
	acc := sdk.AccAddress([]byte("me"))
	coins := sdk.NewCoins(sdk.NewInt64Coin("atom", 10))
	var msg = NewMsgBuyHayt(hayt, coins, acc)
	res := msg.GetSignBytes()

	expected := `{"type":"haytservice/BuyHayt","value":{"bid":[{"amount":"10","denom":"atom"}],` +
		`"buyer":"cosmos1d4js690r9j","hayt":"maTurtle"}}`

	require.Equal(t, expected, string(res))
}

func TestMsgDeleteHayt(t *testing.T) {
	acc := sdk.AccAddress([]byte("me"))
	var msg = NewMsgDeleteHayt(hayt, acc)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "delete_hayt")
}

func TestMsgDeleteNameValidation(t *testing.T) {
	acc := sdk.AccAddress([]byte("me"))
	hayt2 := "a"
	acc2 := sdk.AccAddress([]byte("you"))

	cases := []struct {
		valid bool
		tx    MsgDeleteHayt
	}{
		{true, NewMsgDeleteHayt(hayt, acc)},
		{true, NewMsgDeleteHayt(hayt2, acc2)},
	}

	for _, tc := range cases {
		err := tc.tx.ValidateBasic()
		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgDeleteNameGetSignBytes(t *testing.T) {
	acc := sdk.AccAddress([]byte("me"))
	var msg = NewMsgDeleteHayt(hayt, acc)
	res := msg.GetSignBytes()

	expected := `{"type":"haytservice/DeleteHayt","value":{"hayt":"maTurtle","owner":"cosmos1d4js690r9j"}}`

	require.Equal(t, expected, string(res))
}
