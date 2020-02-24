package types

import (
	sdkerrors "github.com/serjplus/cosmos-sdk/types/errors"
)

var (
	ErrNameDoesNotExist = sdkerrors.Register(ModuleName, 1, "name does not exist")
)
