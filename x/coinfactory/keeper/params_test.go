package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "dex/testutil/keeper"
	"dex/x/coinfactory/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CoinfactoryKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
