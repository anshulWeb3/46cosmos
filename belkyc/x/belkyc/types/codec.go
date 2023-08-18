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
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
