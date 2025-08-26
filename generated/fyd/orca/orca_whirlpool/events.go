package orca_whirlpool

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type LiquidityDecreased struct {
	Whirlpool         solanago.PublicKey
	Position          solanago.PublicKey
	TickLowerIndex    int32
	TickUpperIndex    int32
	Liquidity         binary.Uint128
	TokenAAmount      uint64
	TokenBAmount      uint64
	TokenATransferFee uint64
	TokenBTransferFee uint64
}

func (obj LiquidityDecreased) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_LiquidityDecreased[:], false)
	if err != nil {
		return err
	}
	// Serialize `Whirlpool`:
	if err = encoder.Encode(obj.Whirlpool); err != nil {
		return fmt.Errorf("error while marshaling Whirlpool:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `TickLowerIndex`:
	if err = encoder.Encode(obj.TickLowerIndex); err != nil {
		return fmt.Errorf("error while marshaling TickLowerIndex:%w", err)
	}
	// Serialize `TickUpperIndex`:
	if err = encoder.Encode(obj.TickUpperIndex); err != nil {
		return fmt.Errorf("error while marshaling TickUpperIndex:%w", err)
	}
	// Serialize `Liquidity`:
	if err = encoder.Encode(obj.Liquidity); err != nil {
		return fmt.Errorf("error while marshaling Liquidity:%w", err)
	}
	// Serialize `TokenAAmount`:
	if err = encoder.Encode(obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenAAmount:%w", err)
	}
	// Serialize `TokenBAmount`:
	if err = encoder.Encode(obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenBAmount:%w", err)
	}
	// Serialize `TokenATransferFee`:
	if err = encoder.Encode(obj.TokenATransferFee); err != nil {
		return fmt.Errorf("error while marshaling TokenATransferFee:%w", err)
	}
	// Serialize `TokenBTransferFee`:
	if err = encoder.Encode(obj.TokenBTransferFee); err != nil {
		return fmt.Errorf("error while marshaling TokenBTransferFee:%w", err)
	}
	return nil
}

func (obj LiquidityDecreased) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding LiquidityDecreased: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *LiquidityDecreased) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_LiquidityDecreased[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_LiquidityDecreased[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Whirlpool`:
	if err = decoder.Decode(&obj.Whirlpool); err != nil {
		return fmt.Errorf("error while unmarshaling Whirlpool:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `TickLowerIndex`:
	if err = decoder.Decode(&obj.TickLowerIndex); err != nil {
		return fmt.Errorf("error while unmarshaling TickLowerIndex:%w", err)
	}
	// Deserialize `TickUpperIndex`:
	if err = decoder.Decode(&obj.TickUpperIndex); err != nil {
		return fmt.Errorf("error while unmarshaling TickUpperIndex:%w", err)
	}
	// Deserialize `Liquidity`:
	if err = decoder.Decode(&obj.Liquidity); err != nil {
		return fmt.Errorf("error while unmarshaling Liquidity:%w", err)
	}
	// Deserialize `TokenAAmount`:
	if err = decoder.Decode(&obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAAmount:%w", err)
	}
	// Deserialize `TokenBAmount`:
	if err = decoder.Decode(&obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBAmount:%w", err)
	}
	// Deserialize `TokenATransferFee`:
	if err = decoder.Decode(&obj.TokenATransferFee); err != nil {
		return fmt.Errorf("error while unmarshaling TokenATransferFee:%w", err)
	}
	// Deserialize `TokenBTransferFee`:
	if err = decoder.Decode(&obj.TokenBTransferFee); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBTransferFee:%w", err)
	}
	return nil
}

func (obj *LiquidityDecreased) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling LiquidityDecreased: %w", err)
	}
	return nil
}

func UnmarshalLiquidityDecreased(buf []byte) (*LiquidityDecreased, error) {
	obj := new(LiquidityDecreased)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type LiquidityIncreased struct {
	Whirlpool         solanago.PublicKey
	Position          solanago.PublicKey
	TickLowerIndex    int32
	TickUpperIndex    int32
	Liquidity         binary.Uint128
	TokenAAmount      uint64
	TokenBAmount      uint64
	TokenATransferFee uint64
	TokenBTransferFee uint64
}

func (obj LiquidityIncreased) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_LiquidityIncreased[:], false)
	if err != nil {
		return err
	}
	// Serialize `Whirlpool`:
	if err = encoder.Encode(obj.Whirlpool); err != nil {
		return fmt.Errorf("error while marshaling Whirlpool:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `TickLowerIndex`:
	if err = encoder.Encode(obj.TickLowerIndex); err != nil {
		return fmt.Errorf("error while marshaling TickLowerIndex:%w", err)
	}
	// Serialize `TickUpperIndex`:
	if err = encoder.Encode(obj.TickUpperIndex); err != nil {
		return fmt.Errorf("error while marshaling TickUpperIndex:%w", err)
	}
	// Serialize `Liquidity`:
	if err = encoder.Encode(obj.Liquidity); err != nil {
		return fmt.Errorf("error while marshaling Liquidity:%w", err)
	}
	// Serialize `TokenAAmount`:
	if err = encoder.Encode(obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenAAmount:%w", err)
	}
	// Serialize `TokenBAmount`:
	if err = encoder.Encode(obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenBAmount:%w", err)
	}
	// Serialize `TokenATransferFee`:
	if err = encoder.Encode(obj.TokenATransferFee); err != nil {
		return fmt.Errorf("error while marshaling TokenATransferFee:%w", err)
	}
	// Serialize `TokenBTransferFee`:
	if err = encoder.Encode(obj.TokenBTransferFee); err != nil {
		return fmt.Errorf("error while marshaling TokenBTransferFee:%w", err)
	}
	return nil
}

func (obj LiquidityIncreased) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding LiquidityIncreased: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *LiquidityIncreased) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_LiquidityIncreased[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_LiquidityIncreased[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Whirlpool`:
	if err = decoder.Decode(&obj.Whirlpool); err != nil {
		return fmt.Errorf("error while unmarshaling Whirlpool:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `TickLowerIndex`:
	if err = decoder.Decode(&obj.TickLowerIndex); err != nil {
		return fmt.Errorf("error while unmarshaling TickLowerIndex:%w", err)
	}
	// Deserialize `TickUpperIndex`:
	if err = decoder.Decode(&obj.TickUpperIndex); err != nil {
		return fmt.Errorf("error while unmarshaling TickUpperIndex:%w", err)
	}
	// Deserialize `Liquidity`:
	if err = decoder.Decode(&obj.Liquidity); err != nil {
		return fmt.Errorf("error while unmarshaling Liquidity:%w", err)
	}
	// Deserialize `TokenAAmount`:
	if err = decoder.Decode(&obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAAmount:%w", err)
	}
	// Deserialize `TokenBAmount`:
	if err = decoder.Decode(&obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBAmount:%w", err)
	}
	// Deserialize `TokenATransferFee`:
	if err = decoder.Decode(&obj.TokenATransferFee); err != nil {
		return fmt.Errorf("error while unmarshaling TokenATransferFee:%w", err)
	}
	// Deserialize `TokenBTransferFee`:
	if err = decoder.Decode(&obj.TokenBTransferFee); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBTransferFee:%w", err)
	}
	return nil
}

func (obj *LiquidityIncreased) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling LiquidityIncreased: %w", err)
	}
	return nil
}

func UnmarshalLiquidityIncreased(buf []byte) (*LiquidityIncreased, error) {
	obj := new(LiquidityIncreased)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type PoolInitialized struct {
	Whirlpool        solanago.PublicKey
	WhirlpoolsConfig solanago.PublicKey
	TokenMintA       solanago.PublicKey
	TokenMintB       solanago.PublicKey
	TickSpacing      uint16
	TokenProgramA    solanago.PublicKey
	TokenProgramB    solanago.PublicKey
	DecimalsA        uint8
	DecimalsB        uint8
	InitialSqrtPrice binary.Uint128
}

func (obj PoolInitialized) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_PoolInitialized[:], false)
	if err != nil {
		return err
	}
	// Serialize `Whirlpool`:
	if err = encoder.Encode(obj.Whirlpool); err != nil {
		return fmt.Errorf("error while marshaling Whirlpool:%w", err)
	}
	// Serialize `WhirlpoolsConfig`:
	if err = encoder.Encode(obj.WhirlpoolsConfig); err != nil {
		return fmt.Errorf("error while marshaling WhirlpoolsConfig:%w", err)
	}
	// Serialize `TokenMintA`:
	if err = encoder.Encode(obj.TokenMintA); err != nil {
		return fmt.Errorf("error while marshaling TokenMintA:%w", err)
	}
	// Serialize `TokenMintB`:
	if err = encoder.Encode(obj.TokenMintB); err != nil {
		return fmt.Errorf("error while marshaling TokenMintB:%w", err)
	}
	// Serialize `TickSpacing`:
	if err = encoder.Encode(obj.TickSpacing); err != nil {
		return fmt.Errorf("error while marshaling TickSpacing:%w", err)
	}
	// Serialize `TokenProgramA`:
	if err = encoder.Encode(obj.TokenProgramA); err != nil {
		return fmt.Errorf("error while marshaling TokenProgramA:%w", err)
	}
	// Serialize `TokenProgramB`:
	if err = encoder.Encode(obj.TokenProgramB); err != nil {
		return fmt.Errorf("error while marshaling TokenProgramB:%w", err)
	}
	// Serialize `DecimalsA`:
	if err = encoder.Encode(obj.DecimalsA); err != nil {
		return fmt.Errorf("error while marshaling DecimalsA:%w", err)
	}
	// Serialize `DecimalsB`:
	if err = encoder.Encode(obj.DecimalsB); err != nil {
		return fmt.Errorf("error while marshaling DecimalsB:%w", err)
	}
	// Serialize `InitialSqrtPrice`:
	if err = encoder.Encode(obj.InitialSqrtPrice); err != nil {
		return fmt.Errorf("error while marshaling InitialSqrtPrice:%w", err)
	}
	return nil
}

func (obj PoolInitialized) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding PoolInitialized: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *PoolInitialized) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_PoolInitialized[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_PoolInitialized[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Whirlpool`:
	if err = decoder.Decode(&obj.Whirlpool); err != nil {
		return fmt.Errorf("error while unmarshaling Whirlpool:%w", err)
	}
	// Deserialize `WhirlpoolsConfig`:
	if err = decoder.Decode(&obj.WhirlpoolsConfig); err != nil {
		return fmt.Errorf("error while unmarshaling WhirlpoolsConfig:%w", err)
	}
	// Deserialize `TokenMintA`:
	if err = decoder.Decode(&obj.TokenMintA); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMintA:%w", err)
	}
	// Deserialize `TokenMintB`:
	if err = decoder.Decode(&obj.TokenMintB); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMintB:%w", err)
	}
	// Deserialize `TickSpacing`:
	if err = decoder.Decode(&obj.TickSpacing); err != nil {
		return fmt.Errorf("error while unmarshaling TickSpacing:%w", err)
	}
	// Deserialize `TokenProgramA`:
	if err = decoder.Decode(&obj.TokenProgramA); err != nil {
		return fmt.Errorf("error while unmarshaling TokenProgramA:%w", err)
	}
	// Deserialize `TokenProgramB`:
	if err = decoder.Decode(&obj.TokenProgramB); err != nil {
		return fmt.Errorf("error while unmarshaling TokenProgramB:%w", err)
	}
	// Deserialize `DecimalsA`:
	if err = decoder.Decode(&obj.DecimalsA); err != nil {
		return fmt.Errorf("error while unmarshaling DecimalsA:%w", err)
	}
	// Deserialize `DecimalsB`:
	if err = decoder.Decode(&obj.DecimalsB); err != nil {
		return fmt.Errorf("error while unmarshaling DecimalsB:%w", err)
	}
	// Deserialize `InitialSqrtPrice`:
	if err = decoder.Decode(&obj.InitialSqrtPrice); err != nil {
		return fmt.Errorf("error while unmarshaling InitialSqrtPrice:%w", err)
	}
	return nil
}

func (obj *PoolInitialized) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling PoolInitialized: %w", err)
	}
	return nil
}

func UnmarshalPoolInitialized(buf []byte) (*PoolInitialized, error) {
	obj := new(PoolInitialized)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type Traded struct {
	Whirlpool         solanago.PublicKey
	AToB              bool
	PreSqrtPrice      binary.Uint128
	PostSqrtPrice     binary.Uint128
	InputAmount       uint64
	OutputAmount      uint64
	InputTransferFee  uint64
	OutputTransferFee uint64
	LpFee             uint64
	ProtocolFee       uint64
}

func (obj Traded) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_Traded[:], false)
	if err != nil {
		return err
	}
	// Serialize `Whirlpool`:
	if err = encoder.Encode(obj.Whirlpool); err != nil {
		return fmt.Errorf("error while marshaling Whirlpool:%w", err)
	}
	// Serialize `AToB`:
	if err = encoder.Encode(obj.AToB); err != nil {
		return fmt.Errorf("error while marshaling AToB:%w", err)
	}
	// Serialize `PreSqrtPrice`:
	if err = encoder.Encode(obj.PreSqrtPrice); err != nil {
		return fmt.Errorf("error while marshaling PreSqrtPrice:%w", err)
	}
	// Serialize `PostSqrtPrice`:
	if err = encoder.Encode(obj.PostSqrtPrice); err != nil {
		return fmt.Errorf("error while marshaling PostSqrtPrice:%w", err)
	}
	// Serialize `InputAmount`:
	if err = encoder.Encode(obj.InputAmount); err != nil {
		return fmt.Errorf("error while marshaling InputAmount:%w", err)
	}
	// Serialize `OutputAmount`:
	if err = encoder.Encode(obj.OutputAmount); err != nil {
		return fmt.Errorf("error while marshaling OutputAmount:%w", err)
	}
	// Serialize `InputTransferFee`:
	if err = encoder.Encode(obj.InputTransferFee); err != nil {
		return fmt.Errorf("error while marshaling InputTransferFee:%w", err)
	}
	// Serialize `OutputTransferFee`:
	if err = encoder.Encode(obj.OutputTransferFee); err != nil {
		return fmt.Errorf("error while marshaling OutputTransferFee:%w", err)
	}
	// Serialize `LpFee`:
	if err = encoder.Encode(obj.LpFee); err != nil {
		return fmt.Errorf("error while marshaling LpFee:%w", err)
	}
	// Serialize `ProtocolFee`:
	if err = encoder.Encode(obj.ProtocolFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFee:%w", err)
	}
	return nil
}

func (obj Traded) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Traded: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Traded) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_Traded[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_Traded[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Whirlpool`:
	if err = decoder.Decode(&obj.Whirlpool); err != nil {
		return fmt.Errorf("error while unmarshaling Whirlpool:%w", err)
	}
	// Deserialize `AToB`:
	if err = decoder.Decode(&obj.AToB); err != nil {
		return fmt.Errorf("error while unmarshaling AToB:%w", err)
	}
	// Deserialize `PreSqrtPrice`:
	if err = decoder.Decode(&obj.PreSqrtPrice); err != nil {
		return fmt.Errorf("error while unmarshaling PreSqrtPrice:%w", err)
	}
	// Deserialize `PostSqrtPrice`:
	if err = decoder.Decode(&obj.PostSqrtPrice); err != nil {
		return fmt.Errorf("error while unmarshaling PostSqrtPrice:%w", err)
	}
	// Deserialize `InputAmount`:
	if err = decoder.Decode(&obj.InputAmount); err != nil {
		return fmt.Errorf("error while unmarshaling InputAmount:%w", err)
	}
	// Deserialize `OutputAmount`:
	if err = decoder.Decode(&obj.OutputAmount); err != nil {
		return fmt.Errorf("error while unmarshaling OutputAmount:%w", err)
	}
	// Deserialize `InputTransferFee`:
	if err = decoder.Decode(&obj.InputTransferFee); err != nil {
		return fmt.Errorf("error while unmarshaling InputTransferFee:%w", err)
	}
	// Deserialize `OutputTransferFee`:
	if err = decoder.Decode(&obj.OutputTransferFee); err != nil {
		return fmt.Errorf("error while unmarshaling OutputTransferFee:%w", err)
	}
	// Deserialize `LpFee`:
	if err = decoder.Decode(&obj.LpFee); err != nil {
		return fmt.Errorf("error while unmarshaling LpFee:%w", err)
	}
	// Deserialize `ProtocolFee`:
	if err = decoder.Decode(&obj.ProtocolFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFee:%w", err)
	}
	return nil
}

func (obj *Traded) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Traded: %w", err)
	}
	return nil
}

func UnmarshalTraded(buf []byte) (*Traded, error) {
	obj := new(Traded)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
