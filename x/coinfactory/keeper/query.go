package keeper

import (
	"dex/x/coinfactory/types"
)

var _ types.QueryServer = Keeper{}
