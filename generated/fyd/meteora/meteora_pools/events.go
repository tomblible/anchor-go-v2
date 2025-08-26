package meteora_pools

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type EvtAddLiquidity struct {
	LpMintAmount uint64
	TokenAAmount uint64
	TokenBAmount uint64
}

func (obj EvtAddLiquidity) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_AddLiquidity[:], false)
	if err != nil {
		return err
	}
	// Serialize `LpMintAmount`:
	if err = encoder.Encode(obj.LpMintAmount); err != nil {
		return fmt.Errorf("error while marshaling LpMintAmount:%w", err)
	}
	// Serialize `TokenAAmount`:
	if err = encoder.Encode(obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenAAmount:%w", err)
	}
	// Serialize `TokenBAmount`:
	if err = encoder.Encode(obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenBAmount:%w", err)
	}
	return nil
}

func (obj EvtAddLiquidity) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtAddLiquidity: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtAddLiquidity) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_AddLiquidity[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_AddLiquidity[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LpMintAmount`:
	if err = decoder.Decode(&obj.LpMintAmount); err != nil {
		return fmt.Errorf("error while unmarshaling LpMintAmount:%w", err)
	}
	// Deserialize `TokenAAmount`:
	if err = decoder.Decode(&obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAAmount:%w", err)
	}
	// Deserialize `TokenBAmount`:
	if err = decoder.Decode(&obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBAmount:%w", err)
	}
	return nil
}

func (obj *EvtAddLiquidity) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtAddLiquidity: %w", err)
	}
	return nil
}

func UnmarshalEvtAddLiquidity(buf []byte) (*EvtAddLiquidity, error) {
	obj := new(EvtAddLiquidity)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtRemoveLiquidity struct {
	LpUnmintAmount  uint64
	TokenAOutAmount uint64
	TokenBOutAmount uint64
}

func (obj EvtRemoveLiquidity) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_RemoveLiquidity[:], false)
	if err != nil {
		return err
	}
	// Serialize `LpUnmintAmount`:
	if err = encoder.Encode(obj.LpUnmintAmount); err != nil {
		return fmt.Errorf("error while marshaling LpUnmintAmount:%w", err)
	}
	// Serialize `TokenAOutAmount`:
	if err = encoder.Encode(obj.TokenAOutAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenAOutAmount:%w", err)
	}
	// Serialize `TokenBOutAmount`:
	if err = encoder.Encode(obj.TokenBOutAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenBOutAmount:%w", err)
	}
	return nil
}

func (obj EvtRemoveLiquidity) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtRemoveLiquidity: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtRemoveLiquidity) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_RemoveLiquidity[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_RemoveLiquidity[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LpUnmintAmount`:
	if err = decoder.Decode(&obj.LpUnmintAmount); err != nil {
		return fmt.Errorf("error while unmarshaling LpUnmintAmount:%w", err)
	}
	// Deserialize `TokenAOutAmount`:
	if err = decoder.Decode(&obj.TokenAOutAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAOutAmount:%w", err)
	}
	// Deserialize `TokenBOutAmount`:
	if err = decoder.Decode(&obj.TokenBOutAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBOutAmount:%w", err)
	}
	return nil
}

func (obj *EvtRemoveLiquidity) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtRemoveLiquidity: %w", err)
	}
	return nil
}

func UnmarshalEvtRemoveLiquidity(buf []byte) (*EvtRemoveLiquidity, error) {
	obj := new(EvtRemoveLiquidity)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtBootstrapLiquidity struct {
	LpMintAmount uint64
	TokenAAmount uint64
	TokenBAmount uint64
	Pool         solanago.PublicKey
}

func (obj EvtBootstrapLiquidity) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_BootstrapLiquidity[:], false)
	if err != nil {
		return err
	}
	// Serialize `LpMintAmount`:
	if err = encoder.Encode(obj.LpMintAmount); err != nil {
		return fmt.Errorf("error while marshaling LpMintAmount:%w", err)
	}
	// Serialize `TokenAAmount`:
	if err = encoder.Encode(obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenAAmount:%w", err)
	}
	// Serialize `TokenBAmount`:
	if err = encoder.Encode(obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenBAmount:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	return nil
}

func (obj EvtBootstrapLiquidity) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtBootstrapLiquidity: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtBootstrapLiquidity) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_BootstrapLiquidity[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_BootstrapLiquidity[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LpMintAmount`:
	if err = decoder.Decode(&obj.LpMintAmount); err != nil {
		return fmt.Errorf("error while unmarshaling LpMintAmount:%w", err)
	}
	// Deserialize `TokenAAmount`:
	if err = decoder.Decode(&obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAAmount:%w", err)
	}
	// Deserialize `TokenBAmount`:
	if err = decoder.Decode(&obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBAmount:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	return nil
}

func (obj *EvtBootstrapLiquidity) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtBootstrapLiquidity: %w", err)
	}
	return nil
}

func UnmarshalEvtBootstrapLiquidity(buf []byte) (*EvtBootstrapLiquidity, error) {
	obj := new(EvtBootstrapLiquidity)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtSwap struct {
	InAmount    uint64
	OutAmount   uint64
	TradeFee    uint64
	ProtocolFee uint64
	HostFee     uint64
}

func (obj EvtSwap) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_Swap[:], false)
	if err != nil {
		return err
	}
	// Serialize `InAmount`:
	if err = encoder.Encode(obj.InAmount); err != nil {
		return fmt.Errorf("error while marshaling InAmount:%w", err)
	}
	// Serialize `OutAmount`:
	if err = encoder.Encode(obj.OutAmount); err != nil {
		return fmt.Errorf("error while marshaling OutAmount:%w", err)
	}
	// Serialize `TradeFee`:
	if err = encoder.Encode(obj.TradeFee); err != nil {
		return fmt.Errorf("error while marshaling TradeFee:%w", err)
	}
	// Serialize `ProtocolFee`:
	if err = encoder.Encode(obj.ProtocolFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFee:%w", err)
	}
	// Serialize `HostFee`:
	if err = encoder.Encode(obj.HostFee); err != nil {
		return fmt.Errorf("error while marshaling HostFee:%w", err)
	}
	return nil
}

func (obj EvtSwap) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtSwap: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtSwap) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_Swap[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_Swap[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `InAmount`:
	if err = decoder.Decode(&obj.InAmount); err != nil {
		return fmt.Errorf("error while unmarshaling InAmount:%w", err)
	}
	// Deserialize `OutAmount`:
	if err = decoder.Decode(&obj.OutAmount); err != nil {
		return fmt.Errorf("error while unmarshaling OutAmount:%w", err)
	}
	// Deserialize `TradeFee`:
	if err = decoder.Decode(&obj.TradeFee); err != nil {
		return fmt.Errorf("error while unmarshaling TradeFee:%w", err)
	}
	// Deserialize `ProtocolFee`:
	if err = decoder.Decode(&obj.ProtocolFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFee:%w", err)
	}
	// Deserialize `HostFee`:
	if err = decoder.Decode(&obj.HostFee); err != nil {
		return fmt.Errorf("error while unmarshaling HostFee:%w", err)
	}
	return nil
}

func (obj *EvtSwap) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtSwap: %w", err)
	}
	return nil
}

func UnmarshalEvtSwap(buf []byte) (*EvtSwap, error) {
	obj := new(EvtSwap)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtSetPoolFees struct {
	TradeFeeNumerator           uint64
	TradeFeeDenominator         uint64
	ProtocolTradeFeeNumerator   uint64
	ProtocolTradeFeeDenominator uint64
	Pool                        solanago.PublicKey
}

func (obj EvtSetPoolFees) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_SetPoolFees[:], false)
	if err != nil {
		return err
	}
	// Serialize `TradeFeeNumerator`:
	if err = encoder.Encode(obj.TradeFeeNumerator); err != nil {
		return fmt.Errorf("error while marshaling TradeFeeNumerator:%w", err)
	}
	// Serialize `TradeFeeDenominator`:
	if err = encoder.Encode(obj.TradeFeeDenominator); err != nil {
		return fmt.Errorf("error while marshaling TradeFeeDenominator:%w", err)
	}
	// Serialize `ProtocolTradeFeeNumerator`:
	if err = encoder.Encode(obj.ProtocolTradeFeeNumerator); err != nil {
		return fmt.Errorf("error while marshaling ProtocolTradeFeeNumerator:%w", err)
	}
	// Serialize `ProtocolTradeFeeDenominator`:
	if err = encoder.Encode(obj.ProtocolTradeFeeDenominator); err != nil {
		return fmt.Errorf("error while marshaling ProtocolTradeFeeDenominator:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	return nil
}

func (obj EvtSetPoolFees) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtSetPoolFees: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtSetPoolFees) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_SetPoolFees[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_SetPoolFees[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TradeFeeNumerator`:
	if err = decoder.Decode(&obj.TradeFeeNumerator); err != nil {
		return fmt.Errorf("error while unmarshaling TradeFeeNumerator:%w", err)
	}
	// Deserialize `TradeFeeDenominator`:
	if err = decoder.Decode(&obj.TradeFeeDenominator); err != nil {
		return fmt.Errorf("error while unmarshaling TradeFeeDenominator:%w", err)
	}
	// Deserialize `ProtocolTradeFeeNumerator`:
	if err = decoder.Decode(&obj.ProtocolTradeFeeNumerator); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolTradeFeeNumerator:%w", err)
	}
	// Deserialize `ProtocolTradeFeeDenominator`:
	if err = decoder.Decode(&obj.ProtocolTradeFeeDenominator); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolTradeFeeDenominator:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	return nil
}

func (obj *EvtSetPoolFees) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtSetPoolFees: %w", err)
	}
	return nil
}

func UnmarshalEvtSetPoolFees(buf []byte) (*EvtSetPoolFees, error) {
	obj := new(EvtSetPoolFees)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtPoolInfo struct {
	TokenAAmount     uint64
	TokenBAmount     uint64
	VirtualPrice     float64
	CurrentTimestamp uint64
}

func (obj EvtPoolInfo) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_PoolInfo[:], false)
	if err != nil {
		return err
	}
	// Serialize `TokenAAmount`:
	if err = encoder.Encode(obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenAAmount:%w", err)
	}
	// Serialize `TokenBAmount`:
	if err = encoder.Encode(obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenBAmount:%w", err)
	}
	// Serialize `VirtualPrice`:
	if err = encoder.Encode(obj.VirtualPrice); err != nil {
		return fmt.Errorf("error while marshaling VirtualPrice:%w", err)
	}
	// Serialize `CurrentTimestamp`:
	if err = encoder.Encode(obj.CurrentTimestamp); err != nil {
		return fmt.Errorf("error while marshaling CurrentTimestamp:%w", err)
	}
	return nil
}

func (obj EvtPoolInfo) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtPoolInfo: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtPoolInfo) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_PoolInfo[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_PoolInfo[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TokenAAmount`:
	if err = decoder.Decode(&obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAAmount:%w", err)
	}
	// Deserialize `TokenBAmount`:
	if err = decoder.Decode(&obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBAmount:%w", err)
	}
	// Deserialize `VirtualPrice`:
	if err = decoder.Decode(&obj.VirtualPrice); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualPrice:%w", err)
	}
	// Deserialize `CurrentTimestamp`:
	if err = decoder.Decode(&obj.CurrentTimestamp); err != nil {
		return fmt.Errorf("error while unmarshaling CurrentTimestamp:%w", err)
	}
	return nil
}

func (obj *EvtPoolInfo) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtPoolInfo: %w", err)
	}
	return nil
}

func UnmarshalEvtPoolInfo(buf []byte) (*EvtPoolInfo, error) {
	obj := new(EvtPoolInfo)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtTransferAdmin struct {
	Admin    solanago.PublicKey
	NewAdmin solanago.PublicKey
	Pool     solanago.PublicKey
}

func (obj EvtTransferAdmin) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_TransferAdmin[:], false)
	if err != nil {
		return err
	}
	// Serialize `Admin`:
	if err = encoder.Encode(obj.Admin); err != nil {
		return fmt.Errorf("error while marshaling Admin:%w", err)
	}
	// Serialize `NewAdmin`:
	if err = encoder.Encode(obj.NewAdmin); err != nil {
		return fmt.Errorf("error while marshaling NewAdmin:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	return nil
}

func (obj EvtTransferAdmin) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtTransferAdmin: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtTransferAdmin) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_TransferAdmin[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_TransferAdmin[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Admin`:
	if err = decoder.Decode(&obj.Admin); err != nil {
		return fmt.Errorf("error while unmarshaling Admin:%w", err)
	}
	// Deserialize `NewAdmin`:
	if err = decoder.Decode(&obj.NewAdmin); err != nil {
		return fmt.Errorf("error while unmarshaling NewAdmin:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	return nil
}

func (obj *EvtTransferAdmin) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtTransferAdmin: %w", err)
	}
	return nil
}

func UnmarshalEvtTransferAdmin(buf []byte) (*EvtTransferAdmin, error) {
	obj := new(EvtTransferAdmin)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtOverrideCurveParam struct {
	NewAmp           uint64
	UpdatedTimestamp uint64
	Pool             solanago.PublicKey
}

func (obj EvtOverrideCurveParam) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_OverrideCurveParam[:], false)
	if err != nil {
		return err
	}
	// Serialize `NewAmp`:
	if err = encoder.Encode(obj.NewAmp); err != nil {
		return fmt.Errorf("error while marshaling NewAmp:%w", err)
	}
	// Serialize `UpdatedTimestamp`:
	if err = encoder.Encode(obj.UpdatedTimestamp); err != nil {
		return fmt.Errorf("error while marshaling UpdatedTimestamp:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	return nil
}

func (obj EvtOverrideCurveParam) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtOverrideCurveParam: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtOverrideCurveParam) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_OverrideCurveParam[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_OverrideCurveParam[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `NewAmp`:
	if err = decoder.Decode(&obj.NewAmp); err != nil {
		return fmt.Errorf("error while unmarshaling NewAmp:%w", err)
	}
	// Deserialize `UpdatedTimestamp`:
	if err = decoder.Decode(&obj.UpdatedTimestamp); err != nil {
		return fmt.Errorf("error while unmarshaling UpdatedTimestamp:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	return nil
}

func (obj *EvtOverrideCurveParam) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtOverrideCurveParam: %w", err)
	}
	return nil
}

func UnmarshalEvtOverrideCurveParam(buf []byte) (*EvtOverrideCurveParam, error) {
	obj := new(EvtOverrideCurveParam)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtPoolCreated struct {
	LpMint     solanago.PublicKey
	TokenAMint solanago.PublicKey
	TokenBMint solanago.PublicKey
	PoolType   PoolType
	Pool       solanago.PublicKey
}

func (obj EvtPoolCreated) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_PoolCreated[:], false)
	if err != nil {
		return err
	}
	// Serialize `LpMint`:
	if err = encoder.Encode(obj.LpMint); err != nil {
		return fmt.Errorf("error while marshaling LpMint:%w", err)
	}
	// Serialize `TokenAMint`:
	if err = encoder.Encode(obj.TokenAMint); err != nil {
		return fmt.Errorf("error while marshaling TokenAMint:%w", err)
	}
	// Serialize `TokenBMint`:
	if err = encoder.Encode(obj.TokenBMint); err != nil {
		return fmt.Errorf("error while marshaling TokenBMint:%w", err)
	}
	// Serialize `PoolType`:
	if err = encoder.Encode(obj.PoolType); err != nil {
		return fmt.Errorf("error while marshaling PoolType:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	return nil
}

func (obj EvtPoolCreated) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtPoolCreated: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtPoolCreated) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_PoolCreated[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_PoolCreated[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LpMint`:
	if err = decoder.Decode(&obj.LpMint); err != nil {
		return fmt.Errorf("error while unmarshaling LpMint:%w", err)
	}
	// Deserialize `TokenAMint`:
	if err = decoder.Decode(&obj.TokenAMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAMint:%w", err)
	}
	// Deserialize `TokenBMint`:
	if err = decoder.Decode(&obj.TokenBMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBMint:%w", err)
	}
	// Deserialize `PoolType`:
	if err = decoder.Decode(&obj.PoolType); err != nil {
		return fmt.Errorf("error while unmarshaling PoolType:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	return nil
}

func (obj *EvtPoolCreated) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtPoolCreated: %w", err)
	}
	return nil
}

func UnmarshalEvtPoolCreated(buf []byte) (*EvtPoolCreated, error) {
	obj := new(EvtPoolCreated)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtPoolEnabled struct {
	Pool    solanago.PublicKey
	Enabled bool
}

func (obj EvtPoolEnabled) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_PoolEnabled[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Enabled`:
	if err = encoder.Encode(obj.Enabled); err != nil {
		return fmt.Errorf("error while marshaling Enabled:%w", err)
	}
	return nil
}

func (obj EvtPoolEnabled) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtPoolEnabled: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtPoolEnabled) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_PoolEnabled[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_PoolEnabled[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Enabled`:
	if err = decoder.Decode(&obj.Enabled); err != nil {
		return fmt.Errorf("error while unmarshaling Enabled:%w", err)
	}
	return nil
}

func (obj *EvtPoolEnabled) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtPoolEnabled: %w", err)
	}
	return nil
}

func UnmarshalEvtPoolEnabled(buf []byte) (*EvtPoolEnabled, error) {
	obj := new(EvtPoolEnabled)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtMigrateFeeAccount struct {
	Pool              solanago.PublicKey
	NewAdminTokenAFee solanago.PublicKey
	NewAdminTokenBFee solanago.PublicKey
	TokenAAmount      uint64
	TokenBAmount      uint64
}

func (obj EvtMigrateFeeAccount) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_MigrateFeeAccount[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `NewAdminTokenAFee`:
	if err = encoder.Encode(obj.NewAdminTokenAFee); err != nil {
		return fmt.Errorf("error while marshaling NewAdminTokenAFee:%w", err)
	}
	// Serialize `NewAdminTokenBFee`:
	if err = encoder.Encode(obj.NewAdminTokenBFee); err != nil {
		return fmt.Errorf("error while marshaling NewAdminTokenBFee:%w", err)
	}
	// Serialize `TokenAAmount`:
	if err = encoder.Encode(obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenAAmount:%w", err)
	}
	// Serialize `TokenBAmount`:
	if err = encoder.Encode(obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenBAmount:%w", err)
	}
	return nil
}

func (obj EvtMigrateFeeAccount) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtMigrateFeeAccount: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtMigrateFeeAccount) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_MigrateFeeAccount[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_MigrateFeeAccount[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `NewAdminTokenAFee`:
	if err = decoder.Decode(&obj.NewAdminTokenAFee); err != nil {
		return fmt.Errorf("error while unmarshaling NewAdminTokenAFee:%w", err)
	}
	// Deserialize `NewAdminTokenBFee`:
	if err = decoder.Decode(&obj.NewAdminTokenBFee); err != nil {
		return fmt.Errorf("error while unmarshaling NewAdminTokenBFee:%w", err)
	}
	// Deserialize `TokenAAmount`:
	if err = decoder.Decode(&obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAAmount:%w", err)
	}
	// Deserialize `TokenBAmount`:
	if err = decoder.Decode(&obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBAmount:%w", err)
	}
	return nil
}

func (obj *EvtMigrateFeeAccount) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtMigrateFeeAccount: %w", err)
	}
	return nil
}

func UnmarshalEvtMigrateFeeAccount(buf []byte) (*EvtMigrateFeeAccount, error) {
	obj := new(EvtMigrateFeeAccount)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtCreateLockEscrow struct {
	Pool  solanago.PublicKey
	Owner solanago.PublicKey
}

func (obj EvtCreateLockEscrow) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CreateLockEscrow[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	return nil
}

func (obj EvtCreateLockEscrow) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtCreateLockEscrow: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtCreateLockEscrow) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CreateLockEscrow[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CreateLockEscrow[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	return nil
}

func (obj *EvtCreateLockEscrow) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtCreateLockEscrow: %w", err)
	}
	return nil
}

func UnmarshalEvtCreateLockEscrow(buf []byte) (*EvtCreateLockEscrow, error) {
	obj := new(EvtCreateLockEscrow)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtLock struct {
	Pool   solanago.PublicKey
	Owner  solanago.PublicKey
	Amount uint64
}

func (obj EvtLock) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_Lock[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `Amount`:
	if err = encoder.Encode(obj.Amount); err != nil {
		return fmt.Errorf("error while marshaling Amount:%w", err)
	}
	return nil
}

func (obj EvtLock) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtLock: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtLock) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_Lock[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_Lock[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `Amount`:
	if err = decoder.Decode(&obj.Amount); err != nil {
		return fmt.Errorf("error while unmarshaling Amount:%w", err)
	}
	return nil
}

func (obj *EvtLock) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtLock: %w", err)
	}
	return nil
}

func UnmarshalEvtLock(buf []byte) (*EvtLock, error) {
	obj := new(EvtLock)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtClaimFee struct {
	Pool   solanago.PublicKey
	Owner  solanago.PublicKey
	Amount uint64
	AFee   uint64
	BFee   uint64
}

func (obj EvtClaimFee) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_ClaimFee[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `Amount`:
	if err = encoder.Encode(obj.Amount); err != nil {
		return fmt.Errorf("error while marshaling Amount:%w", err)
	}
	// Serialize `AFee`:
	if err = encoder.Encode(obj.AFee); err != nil {
		return fmt.Errorf("error while marshaling AFee:%w", err)
	}
	// Serialize `BFee`:
	if err = encoder.Encode(obj.BFee); err != nil {
		return fmt.Errorf("error while marshaling BFee:%w", err)
	}
	return nil
}

func (obj EvtClaimFee) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtClaimFee: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtClaimFee) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_ClaimFee[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_ClaimFee[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `Amount`:
	if err = decoder.Decode(&obj.Amount); err != nil {
		return fmt.Errorf("error while unmarshaling Amount:%w", err)
	}
	// Deserialize `AFee`:
	if err = decoder.Decode(&obj.AFee); err != nil {
		return fmt.Errorf("error while unmarshaling AFee:%w", err)
	}
	// Deserialize `BFee`:
	if err = decoder.Decode(&obj.BFee); err != nil {
		return fmt.Errorf("error while unmarshaling BFee:%w", err)
	}
	return nil
}

func (obj *EvtClaimFee) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtClaimFee: %w", err)
	}
	return nil
}

func UnmarshalEvtClaimFee(buf []byte) (*EvtClaimFee, error) {
	obj := new(EvtClaimFee)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtCreateConfig struct {
	TradeFeeNumerator         uint64
	ProtocolTradeFeeNumerator uint64
	Config                    solanago.PublicKey
}

func (obj EvtCreateConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CreateConfig[:], false)
	if err != nil {
		return err
	}
	// Serialize `TradeFeeNumerator`:
	if err = encoder.Encode(obj.TradeFeeNumerator); err != nil {
		return fmt.Errorf("error while marshaling TradeFeeNumerator:%w", err)
	}
	// Serialize `ProtocolTradeFeeNumerator`:
	if err = encoder.Encode(obj.ProtocolTradeFeeNumerator); err != nil {
		return fmt.Errorf("error while marshaling ProtocolTradeFeeNumerator:%w", err)
	}
	// Serialize `Config`:
	if err = encoder.Encode(obj.Config); err != nil {
		return fmt.Errorf("error while marshaling Config:%w", err)
	}
	return nil
}

func (obj EvtCreateConfig) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtCreateConfig: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtCreateConfig) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CreateConfig[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CreateConfig[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TradeFeeNumerator`:
	if err = decoder.Decode(&obj.TradeFeeNumerator); err != nil {
		return fmt.Errorf("error while unmarshaling TradeFeeNumerator:%w", err)
	}
	// Deserialize `ProtocolTradeFeeNumerator`:
	if err = decoder.Decode(&obj.ProtocolTradeFeeNumerator); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolTradeFeeNumerator:%w", err)
	}
	// Deserialize `Config`:
	if err = decoder.Decode(&obj.Config); err != nil {
		return fmt.Errorf("error while unmarshaling Config:%w", err)
	}
	return nil
}

func (obj *EvtCreateConfig) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtCreateConfig: %w", err)
	}
	return nil
}

func UnmarshalEvtCreateConfig(buf []byte) (*EvtCreateConfig, error) {
	obj := new(EvtCreateConfig)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtCloseConfig struct {
	Config solanago.PublicKey
}

func (obj EvtCloseConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CloseConfig[:], false)
	if err != nil {
		return err
	}
	// Serialize `Config`:
	if err = encoder.Encode(obj.Config); err != nil {
		return fmt.Errorf("error while marshaling Config:%w", err)
	}
	return nil
}

func (obj EvtCloseConfig) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtCloseConfig: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtCloseConfig) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CloseConfig[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CloseConfig[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Config`:
	if err = decoder.Decode(&obj.Config); err != nil {
		return fmt.Errorf("error while unmarshaling Config:%w", err)
	}
	return nil
}

func (obj *EvtCloseConfig) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtCloseConfig: %w", err)
	}
	return nil
}

func UnmarshalEvtCloseConfig(buf []byte) (*EvtCloseConfig, error) {
	obj := new(EvtCloseConfig)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtWithdrawProtocolFees struct {
	Pool              solanago.PublicKey
	ProtocolAFee      uint64
	ProtocolBFee      uint64
	ProtocolAFeeOwner solanago.PublicKey
	ProtocolBFeeOwner solanago.PublicKey
}

func (obj EvtWithdrawProtocolFees) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_WithdrawProtocolFees[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `ProtocolAFee`:
	if err = encoder.Encode(obj.ProtocolAFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolAFee:%w", err)
	}
	// Serialize `ProtocolBFee`:
	if err = encoder.Encode(obj.ProtocolBFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolBFee:%w", err)
	}
	// Serialize `ProtocolAFeeOwner`:
	if err = encoder.Encode(obj.ProtocolAFeeOwner); err != nil {
		return fmt.Errorf("error while marshaling ProtocolAFeeOwner:%w", err)
	}
	// Serialize `ProtocolBFeeOwner`:
	if err = encoder.Encode(obj.ProtocolBFeeOwner); err != nil {
		return fmt.Errorf("error while marshaling ProtocolBFeeOwner:%w", err)
	}
	return nil
}

func (obj EvtWithdrawProtocolFees) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtWithdrawProtocolFees: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtWithdrawProtocolFees) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_WithdrawProtocolFees[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_WithdrawProtocolFees[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `ProtocolAFee`:
	if err = decoder.Decode(&obj.ProtocolAFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolAFee:%w", err)
	}
	// Deserialize `ProtocolBFee`:
	if err = decoder.Decode(&obj.ProtocolBFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolBFee:%w", err)
	}
	// Deserialize `ProtocolAFeeOwner`:
	if err = decoder.Decode(&obj.ProtocolAFeeOwner); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolAFeeOwner:%w", err)
	}
	// Deserialize `ProtocolBFeeOwner`:
	if err = decoder.Decode(&obj.ProtocolBFeeOwner); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolBFeeOwner:%w", err)
	}
	return nil
}

func (obj *EvtWithdrawProtocolFees) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtWithdrawProtocolFees: %w", err)
	}
	return nil
}

func UnmarshalEvtWithdrawProtocolFees(buf []byte) (*EvtWithdrawProtocolFees, error) {
	obj := new(EvtWithdrawProtocolFees)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtPartnerClaimFees struct {
	Pool    solanago.PublicKey
	FeeA    uint64
	FeeB    uint64
	Partner solanago.PublicKey
}

func (obj EvtPartnerClaimFees) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_PartnerClaimFees[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `FeeA`:
	if err = encoder.Encode(obj.FeeA); err != nil {
		return fmt.Errorf("error while marshaling FeeA:%w", err)
	}
	// Serialize `FeeB`:
	if err = encoder.Encode(obj.FeeB); err != nil {
		return fmt.Errorf("error while marshaling FeeB:%w", err)
	}
	// Serialize `Partner`:
	if err = encoder.Encode(obj.Partner); err != nil {
		return fmt.Errorf("error while marshaling Partner:%w", err)
	}
	return nil
}

func (obj EvtPartnerClaimFees) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtPartnerClaimFees: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtPartnerClaimFees) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_PartnerClaimFees[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_PartnerClaimFees[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `FeeA`:
	if err = decoder.Decode(&obj.FeeA); err != nil {
		return fmt.Errorf("error while unmarshaling FeeA:%w", err)
	}
	// Deserialize `FeeB`:
	if err = decoder.Decode(&obj.FeeB); err != nil {
		return fmt.Errorf("error while unmarshaling FeeB:%w", err)
	}
	// Deserialize `Partner`:
	if err = decoder.Decode(&obj.Partner); err != nil {
		return fmt.Errorf("error while unmarshaling Partner:%w", err)
	}
	return nil
}

func (obj *EvtPartnerClaimFees) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtPartnerClaimFees: %w", err)
	}
	return nil
}

func UnmarshalEvtPartnerClaimFees(buf []byte) (*EvtPartnerClaimFees, error) {
	obj := new(EvtPartnerClaimFees)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
