package haytservice

import (
	"github.com/serjplus/cosmos-sdk-tutorials/haytservice/x/haytservice/internal/keeper"
	"github.com/serjplus/cosmos-sdk-tutorials/haytservice/x/haytservice/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

/*
var (
	NewKeeper        = keeper.NewKeeper
	NewQuerier       = keeper.NewQuerier
	NewMsgBuyName    = types.NewMsgBuyName
	NewMsgSetName    = types.NewMsgSetName
	NewMsgDeleteName = types.NewMsgDeleteName
	NewWhois         = types.NewWhois
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)*/

var (
	NewKeeper        = keeper.NewKeeper
	NewQuerier       = keeper.NewQuerier
	NewMsgBuyHayt    = types.NewMsgBuyHayt
	NewMsgSetHayt    = types.NewMsgSetHayt
	NewMsgDeleteHayt = types.NewMsgDeleteHayt
	NewWhois         = types.NewWhois
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)

/*
type (
	Keeper          = keeper.Keeper
	MsgSetName      = types.MsgSetName
	MsgBuyName      = types.MsgBuyName
	MsgDeleteName   = types.MsgDeleteName
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Whois           = types.Whois
*/

type (
	Keeper          = keeper.Keeper
	MsgSetHayt      = types.MsgSetHayt
	MsgBuyHayt      = types.MsgBuyHayt
	MsgDeleteHayt   = types.MsgDeleteHayt
	QueryResResolve = types.QueryResResolve
	QueryResHayts   = types.QueryResHayts
	Whois           = types.Whois
)
