package belkyc

import (
	"belkyc/x/belkyc/keeper"
	"belkyc/x/belkyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the kyc
	for _, elem := range genState.KycList {
		k.SetKyc(ctx, elem)
	}
	// Set all the validatorKYC
	for _, elem := range genState.ValidatorKYCList {
		k.SetValidatorKYC(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.KycList = k.GetAllKyc(ctx)
	genesis.ValidatorKYCList = k.GetAllValidatorKYC(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
