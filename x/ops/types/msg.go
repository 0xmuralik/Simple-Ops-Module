package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	_ sdk.Msg = &MsgCreateNameRecord{}
	_ sdk.Msg = &MsgUpdateNameRecord{}
	_ sdk.Msg = &MsgDeleteNameRecord{}
)

func NewMsgCreateRecord(signer sdk.AccAddress, name string, age uint64) MsgCreateNameRecord {
	return MsgCreateNameRecord{
		Signer: signer.String(),
		Name:   name,
		Age:    age,
	}
}

func (msg MsgCreateNameRecord) Route() string { return RouterKey }

func (msg MsgCreateNameRecord) Type() string { return "create" }

func (msg MsgCreateNameRecord) ValidateBasic() error { return nil }

func (msg MsgCreateNameRecord) GetSigners() []sdk.AccAddress {
	accAddr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{accAddr}
}

func NewMsgUpdateRecord(signer sdk.AccAddress, id string, name string, age uint64) MsgUpdateNameRecord {
	return MsgUpdateNameRecord{
		Signer: signer.String(),
		Id:     id,
		Name:   name,
		Age:    age,
	}
}

func (msg MsgUpdateNameRecord) Route() string { return RouterKey }

func (msg MsgUpdateNameRecord) Type() string { return "update" }

func (msg MsgUpdateNameRecord) ValidateBasic() error { return nil }

func (msg MsgUpdateNameRecord) GetSigners() []sdk.AccAddress {
	accAddr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{accAddr}
}

func NewMsgDeleteRecord(signer sdk.AccAddress, id string) MsgDeleteNameRecord {
	return MsgDeleteNameRecord{
		Signer: signer.String(),
		Id:     id,
	}
}

func (msg MsgDeleteNameRecord) Route() string { return RouterKey }

func (msg MsgDeleteNameRecord) Type() string { return "delete" }

func (msg MsgDeleteNameRecord) ValidateBasic() error { return nil }

func (msg MsgDeleteNameRecord) GetSigners() []sdk.AccAddress {
	accAddr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{accAddr}
}
