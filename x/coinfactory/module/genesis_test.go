package coinfactory_test

import (
	"testing"

	keepertest "dex/testutil/keeper"
	"dex/testutil/nullify"
	coinfactory "dex/x/coinfactory/module"
	"dex/x/coinfactory/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CoinfactoryKeeper(t)
	coinfactory.InitGenesis(ctx, k, genesisState)
	got := coinfactory.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
