package belkyc_test

import (
	"testing"

	keepertest "belkyc/testutil/keeper"
	"belkyc/testutil/nullify"
	"belkyc/x/belkyc"
	"belkyc/x/belkyc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		KycList: []types.Kyc{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		ValidatorKYCList: []types.ValidatorKYC{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BelkycKeeper(t)
	belkyc.InitGenesis(ctx, *k, genesisState)
	got := belkyc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.KycList, got.KycList)
	require.ElementsMatch(t, genesisState.ValidatorKYCList, got.ValidatorKYCList)
	// this line is used by starport scaffolding # genesis/test/assert
}
