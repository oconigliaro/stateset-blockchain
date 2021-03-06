package invoice

import "github.com/cosmos/cosmos-sdk/codec"

// RegisterCodec registers all the necessary types and interfaces for the module
func RegisterCodec(c *codec.Codec) {
	c.RegisterConcrete(MsgCreateInvoice{}, "stateset/MsgCreateInvoice", nil)
	c.RegisterConcrete(MsgEditInvoice{}, "stateset/MsgEditInvoice", nil)
	c.RegisterConcrete(MsgDeleteInvoice{}, "stateset/MsgDeleteInvoice", nil)
	c.RegisterConcrete(MsgFactorInvoice{}, "stateset/MsgFactorInvoice", nil)
	c.RegisterConcrete(MsgPayInvoice{}, "stateset/MsgPayInvoice", nil)
	c.RegisterConcrete(MsgAddAdmin{}, "invoice/MsgAddAdmin", nil)
	c.RegisterConcrete(MsgRemoveAdmin{}, "invoice/MsgRemoveAdmin", nil)
	c.RegisterConcrete(MsgUpdateParams{}, "invoice/MsgUpdateParams", nil)

	c.RegisterConcrete(Invoice{}, "stateset/Invoice", nil)
}

// ModuleCodec encodes module codec
var ModuleCodec *codec.Codec

func init() {
	ModuleCodec = codec.New()
	RegisterCodec(ModuleCodec)
	codec.RegisterCrypto(ModuleCodec)
	ModuleCodec.Seal()
}