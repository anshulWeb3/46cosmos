package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateKyc{}, "belkyc/CreateKyc", nil)
	cdc.RegisterConcrete(&MsgUpdateKyc{}, "belkyc/UpdateKyc", nil)
	cdc.RegisterConcrete(&MsgDeleteKyc{}, "belkyc/DeleteKyc", nil)
	cdc.RegisterConcrete(&MsgChangeAdmin{}, "belkyc/ChangeAdmin", nil)
	cdc.RegisterConcrete(&MsgCreateValidatorKYC{}, "belkyc/CreateValidatorKYC", nil)
	cdc.RegisterConcrete(&MsgUpdateValidatorKYC{}, "belkyc/UpdateValidatorKYC", nil)
	cdc.RegisterConcrete(&MsgDeleteValidatorKYC{}, "belkyc/DeleteValidatorKYC", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateKyc{},
		&MsgUpdateKyc{},
		&MsgDeleteKyc{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgChangeAdmin{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateValidatorKYC{},
		&MsgUpdateValidatorKYC{},
		&MsgDeleteValidatorKYC{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
