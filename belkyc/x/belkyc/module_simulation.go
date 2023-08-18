package belkyc

import (
	"math/rand"

	"belkyc/testutil/sample"
	belkycsimulation "belkyc/x/belkyc/simulation"
	"belkyc/x/belkyc/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = belkycsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateKyc = "op_weight_msg_kyc"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateKyc int = 100

	opWeightMsgUpdateKyc = "op_weight_msg_kyc"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateKyc int = 100

	opWeightMsgDeleteKyc = "op_weight_msg_kyc"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteKyc int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	belkycGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		KycList: []types.Kyc{
			{
				Creator: sample.AccAddress(),
				Address: "0",
			},
			{
				Creator: sample.AccAddress(),
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&belkycGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateKyc int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateKyc, &weightMsgCreateKyc, nil,
		func(_ *rand.Rand) {
			weightMsgCreateKyc = defaultWeightMsgCreateKyc
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateKyc,
		belkycsimulation.SimulateMsgCreateKyc(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateKyc int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateKyc, &weightMsgUpdateKyc, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateKyc = defaultWeightMsgUpdateKyc
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateKyc,
		belkycsimulation.SimulateMsgUpdateKyc(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteKyc int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteKyc, &weightMsgDeleteKyc, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteKyc = defaultWeightMsgDeleteKyc
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteKyc,
		belkycsimulation.SimulateMsgDeleteKyc(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
