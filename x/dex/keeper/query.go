package keeper

import (
	"dex/x/dex/types"
)

var _ types.QueryServer = Keeper{}
