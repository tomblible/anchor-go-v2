package raydium_cpmm

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

// Holds the current owner of the factory
type AmmConfig struct {
	// Bump to identify PDA
	Bump uint8

	// Status to control if new pool can be create
	DisableCreatePool bool

	// Config index
	Index uint16

	// The trade fee, denominated in hundredths of a bip (10^-6)
	TradeFeeRate uint64

	// The protocol fee
	ProtocolFeeRate uint64

	// The fund fee, denominated in hundredths of a bip (10^-6)
	FundFeeRate uint64

	// Fee for create a new pool
	CreatePoolFee uint64

	// Address of the protocol fee owner
	ProtocolOwner solanago.PublicKey

	// Address of the fund fee owner
	FundOwner solanago.PublicKey

	// padding
	Padding [16]uint64
}

func (obj AmmConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_AmmConfig[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	// Serialize `DisableCreatePool`:
	if err = encoder.Encode(obj.DisableCreatePool); err != nil {
		return fmt.Errorf("error while marshaling DisableCreatePool:%w", err)
	}
	// Serialize `Index`:
	if err = encoder.Encode(obj.Index); err != nil {
		return fmt.Errorf("error while marshaling Index:%w", err)
	}
	// Serialize `TradeFeeRate`:
	if err = encoder.Encode(obj.TradeFeeRate); err != nil {
		return fmt.Errorf("error while marshaling TradeFeeRate:%w", err)
	}
	// Serialize `ProtocolFeeRate`:
	if err = encoder.Encode(obj.ProtocolFeeRate); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeRate:%w", err)
	}
	// Serialize `FundFeeRate`:
	if err = encoder.Encode(obj.FundFeeRate); err != nil {
		return fmt.Errorf("error while marshaling FundFeeRate:%w", err)
	}
	// Serialize `CreatePoolFee`:
	if err = encoder.Encode(obj.CreatePoolFee); err != nil {
		return fmt.Errorf("error while marshaling CreatePoolFee:%w", err)
	}
	// Serialize `ProtocolOwner`:
	if err = encoder.Encode(obj.ProtocolOwner); err != nil {
		return fmt.Errorf("error while marshaling ProtocolOwner:%w", err)
	}
	// Serialize `FundOwner`:
	if err = encoder.Encode(obj.FundOwner); err != nil {
		return fmt.Errorf("error while marshaling FundOwner:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj AmmConfig) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding AmmConfig: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *AmmConfig) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_AmmConfig[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_AmmConfig[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	// Deserialize `DisableCreatePool`:
	if err = decoder.Decode(&obj.DisableCreatePool); err != nil {
		return fmt.Errorf("error while unmarshaling DisableCreatePool:%w", err)
	}
	// Deserialize `Index`:
	if err = decoder.Decode(&obj.Index); err != nil {
		return fmt.Errorf("error while unmarshaling Index:%w", err)
	}
	// Deserialize `TradeFeeRate`:
	if err = decoder.Decode(&obj.TradeFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling TradeFeeRate:%w", err)
	}
	// Deserialize `ProtocolFeeRate`:
	if err = decoder.Decode(&obj.ProtocolFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeRate:%w", err)
	}
	// Deserialize `FundFeeRate`:
	if err = decoder.Decode(&obj.FundFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling FundFeeRate:%w", err)
	}
	// Deserialize `CreatePoolFee`:
	if err = decoder.Decode(&obj.CreatePoolFee); err != nil {
		return fmt.Errorf("error while unmarshaling CreatePoolFee:%w", err)
	}
	// Deserialize `ProtocolOwner`:
	if err = decoder.Decode(&obj.ProtocolOwner); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolOwner:%w", err)
	}
	// Deserialize `FundOwner`:
	if err = decoder.Decode(&obj.FundOwner); err != nil {
		return fmt.Errorf("error while unmarshaling FundOwner:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *AmmConfig) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling AmmConfig: %w", err)
	}
	return nil
}

func UnmarshalAmmConfig(buf []byte) (*AmmConfig, error) {
	obj := new(AmmConfig)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type ObservationState struct {
	// Whether the ObservationState is initialized
	Initialized bool

	// the most-recently updated index of the observations array
	ObservationIndex uint16
	PoolId           solanago.PublicKey

	// observation array
	Observations [100]Observation

	// padding for feature update
	Padding [4]uint64
}

func (obj ObservationState) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_ObservationState[:], false)
	if err != nil {
		return err
	}
	// Serialize `Initialized`:
	if err = encoder.Encode(obj.Initialized); err != nil {
		return fmt.Errorf("error while marshaling Initialized:%w", err)
	}
	// Serialize `ObservationIndex`:
	if err = encoder.Encode(obj.ObservationIndex); err != nil {
		return fmt.Errorf("error while marshaling ObservationIndex:%w", err)
	}
	// Serialize `PoolId`:
	if err = encoder.Encode(obj.PoolId); err != nil {
		return fmt.Errorf("error while marshaling PoolId:%w", err)
	}
	// Serialize `Observations`:
	if err = encoder.Encode(obj.Observations); err != nil {
		return fmt.Errorf("error while marshaling Observations:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj ObservationState) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding ObservationState: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *ObservationState) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_ObservationState[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_ObservationState[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Initialized`:
	if err = decoder.Decode(&obj.Initialized); err != nil {
		return fmt.Errorf("error while unmarshaling Initialized:%w", err)
	}
	// Deserialize `ObservationIndex`:
	if err = decoder.Decode(&obj.ObservationIndex); err != nil {
		return fmt.Errorf("error while unmarshaling ObservationIndex:%w", err)
	}
	// Deserialize `PoolId`:
	if err = decoder.Decode(&obj.PoolId); err != nil {
		return fmt.Errorf("error while unmarshaling PoolId:%w", err)
	}
	// Deserialize `Observations`:
	if err = decoder.Decode(&obj.Observations); err != nil {
		return fmt.Errorf("error while unmarshaling Observations:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *ObservationState) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling ObservationState: %w", err)
	}
	return nil
}

func UnmarshalObservationState(buf []byte) (*ObservationState, error) {
	obj := new(ObservationState)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type PoolState struct {
	// Which config the pool belongs
	AmmConfig solanago.PublicKey

	// pool creator
	PoolCreator solanago.PublicKey

	// Token A
	Token0Vault solanago.PublicKey

	// Token B
	Token1Vault solanago.PublicKey

	// Pool tokens are issued when A or B tokens are deposited.
	// Pool tokens can be withdrawn back to the original A or B token.
	LpMint solanago.PublicKey

	// Mint information for token A
	Token0Mint solanago.PublicKey

	// Mint information for token B
	Token1Mint solanago.PublicKey

	// token_0 program
	Token0Program solanago.PublicKey

	// token_1 program
	Token1Program solanago.PublicKey

	// observation account to store oracle data
	ObservationKey solanago.PublicKey
	AuthBump       uint8

	// Bitwise representation of the state of the pool
	// bit0, 1: disable deposit(value is 1), 0: normal
	// bit1, 1: disable withdraw(value is 2), 0: normal
	// bit2, 1: disable swap(value is 4), 0: normal
	Status         uint8
	LpMintDecimals uint8

	// mint0 and mint1 decimals
	Mint0Decimals uint8
	Mint1Decimals uint8

	// True circulating supply without burns and lock ups
	LpSupply uint64

	// The amounts of token_0 and token_1 that are owed to the liquidity provider.
	ProtocolFeesToken0 uint64
	ProtocolFeesToken1 uint64
	FundFeesToken0     uint64
	FundFeesToken1     uint64

	// The timestamp allowed for swap in the pool.
	OpenTime uint64

	// recent epoch
	RecentEpoch uint64

	// padding for future updates
	Padding [31]uint64
}

func (obj PoolState) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_PoolState[:], false)
	if err != nil {
		return err
	}
	// Serialize `AmmConfig`:
	if err = encoder.Encode(obj.AmmConfig); err != nil {
		return fmt.Errorf("error while marshaling AmmConfig:%w", err)
	}
	// Serialize `PoolCreator`:
	if err = encoder.Encode(obj.PoolCreator); err != nil {
		return fmt.Errorf("error while marshaling PoolCreator:%w", err)
	}
	// Serialize `Token0Vault`:
	if err = encoder.Encode(obj.Token0Vault); err != nil {
		return fmt.Errorf("error while marshaling Token0Vault:%w", err)
	}
	// Serialize `Token1Vault`:
	if err = encoder.Encode(obj.Token1Vault); err != nil {
		return fmt.Errorf("error while marshaling Token1Vault:%w", err)
	}
	// Serialize `LpMint`:
	if err = encoder.Encode(obj.LpMint); err != nil {
		return fmt.Errorf("error while marshaling LpMint:%w", err)
	}
	// Serialize `Token0Mint`:
	if err = encoder.Encode(obj.Token0Mint); err != nil {
		return fmt.Errorf("error while marshaling Token0Mint:%w", err)
	}
	// Serialize `Token1Mint`:
	if err = encoder.Encode(obj.Token1Mint); err != nil {
		return fmt.Errorf("error while marshaling Token1Mint:%w", err)
	}
	// Serialize `Token0Program`:
	if err = encoder.Encode(obj.Token0Program); err != nil {
		return fmt.Errorf("error while marshaling Token0Program:%w", err)
	}
	// Serialize `Token1Program`:
	if err = encoder.Encode(obj.Token1Program); err != nil {
		return fmt.Errorf("error while marshaling Token1Program:%w", err)
	}
	// Serialize `ObservationKey`:
	if err = encoder.Encode(obj.ObservationKey); err != nil {
		return fmt.Errorf("error while marshaling ObservationKey:%w", err)
	}
	// Serialize `AuthBump`:
	if err = encoder.Encode(obj.AuthBump); err != nil {
		return fmt.Errorf("error while marshaling AuthBump:%w", err)
	}
	// Serialize `Status`:
	if err = encoder.Encode(obj.Status); err != nil {
		return fmt.Errorf("error while marshaling Status:%w", err)
	}
	// Serialize `LpMintDecimals`:
	if err = encoder.Encode(obj.LpMintDecimals); err != nil {
		return fmt.Errorf("error while marshaling LpMintDecimals:%w", err)
	}
	// Serialize `Mint0Decimals`:
	if err = encoder.Encode(obj.Mint0Decimals); err != nil {
		return fmt.Errorf("error while marshaling Mint0Decimals:%w", err)
	}
	// Serialize `Mint1Decimals`:
	if err = encoder.Encode(obj.Mint1Decimals); err != nil {
		return fmt.Errorf("error while marshaling Mint1Decimals:%w", err)
	}
	// Serialize `LpSupply`:
	if err = encoder.Encode(obj.LpSupply); err != nil {
		return fmt.Errorf("error while marshaling LpSupply:%w", err)
	}
	// Serialize `ProtocolFeesToken0`:
	if err = encoder.Encode(obj.ProtocolFeesToken0); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeesToken0:%w", err)
	}
	// Serialize `ProtocolFeesToken1`:
	if err = encoder.Encode(obj.ProtocolFeesToken1); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeesToken1:%w", err)
	}
	// Serialize `FundFeesToken0`:
	if err = encoder.Encode(obj.FundFeesToken0); err != nil {
		return fmt.Errorf("error while marshaling FundFeesToken0:%w", err)
	}
	// Serialize `FundFeesToken1`:
	if err = encoder.Encode(obj.FundFeesToken1); err != nil {
		return fmt.Errorf("error while marshaling FundFeesToken1:%w", err)
	}
	// Serialize `OpenTime`:
	if err = encoder.Encode(obj.OpenTime); err != nil {
		return fmt.Errorf("error while marshaling OpenTime:%w", err)
	}
	// Serialize `RecentEpoch`:
	if err = encoder.Encode(obj.RecentEpoch); err != nil {
		return fmt.Errorf("error while marshaling RecentEpoch:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj PoolState) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding PoolState: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *PoolState) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_PoolState[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_PoolState[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `AmmConfig`:
	if err = decoder.Decode(&obj.AmmConfig); err != nil {
		return fmt.Errorf("error while unmarshaling AmmConfig:%w", err)
	}
	// Deserialize `PoolCreator`:
	if err = decoder.Decode(&obj.PoolCreator); err != nil {
		return fmt.Errorf("error while unmarshaling PoolCreator:%w", err)
	}
	// Deserialize `Token0Vault`:
	if err = decoder.Decode(&obj.Token0Vault); err != nil {
		return fmt.Errorf("error while unmarshaling Token0Vault:%w", err)
	}
	// Deserialize `Token1Vault`:
	if err = decoder.Decode(&obj.Token1Vault); err != nil {
		return fmt.Errorf("error while unmarshaling Token1Vault:%w", err)
	}
	// Deserialize `LpMint`:
	if err = decoder.Decode(&obj.LpMint); err != nil {
		return fmt.Errorf("error while unmarshaling LpMint:%w", err)
	}
	// Deserialize `Token0Mint`:
	if err = decoder.Decode(&obj.Token0Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Token0Mint:%w", err)
	}
	// Deserialize `Token1Mint`:
	if err = decoder.Decode(&obj.Token1Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Token1Mint:%w", err)
	}
	// Deserialize `Token0Program`:
	if err = decoder.Decode(&obj.Token0Program); err != nil {
		return fmt.Errorf("error while unmarshaling Token0Program:%w", err)
	}
	// Deserialize `Token1Program`:
	if err = decoder.Decode(&obj.Token1Program); err != nil {
		return fmt.Errorf("error while unmarshaling Token1Program:%w", err)
	}
	// Deserialize `ObservationKey`:
	if err = decoder.Decode(&obj.ObservationKey); err != nil {
		return fmt.Errorf("error while unmarshaling ObservationKey:%w", err)
	}
	// Deserialize `AuthBump`:
	if err = decoder.Decode(&obj.AuthBump); err != nil {
		return fmt.Errorf("error while unmarshaling AuthBump:%w", err)
	}
	// Deserialize `Status`:
	if err = decoder.Decode(&obj.Status); err != nil {
		return fmt.Errorf("error while unmarshaling Status:%w", err)
	}
	// Deserialize `LpMintDecimals`:
	if err = decoder.Decode(&obj.LpMintDecimals); err != nil {
		return fmt.Errorf("error while unmarshaling LpMintDecimals:%w", err)
	}
	// Deserialize `Mint0Decimals`:
	if err = decoder.Decode(&obj.Mint0Decimals); err != nil {
		return fmt.Errorf("error while unmarshaling Mint0Decimals:%w", err)
	}
	// Deserialize `Mint1Decimals`:
	if err = decoder.Decode(&obj.Mint1Decimals); err != nil {
		return fmt.Errorf("error while unmarshaling Mint1Decimals:%w", err)
	}
	// Deserialize `LpSupply`:
	if err = decoder.Decode(&obj.LpSupply); err != nil {
		return fmt.Errorf("error while unmarshaling LpSupply:%w", err)
	}
	// Deserialize `ProtocolFeesToken0`:
	if err = decoder.Decode(&obj.ProtocolFeesToken0); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeesToken0:%w", err)
	}
	// Deserialize `ProtocolFeesToken1`:
	if err = decoder.Decode(&obj.ProtocolFeesToken1); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeesToken1:%w", err)
	}
	// Deserialize `FundFeesToken0`:
	if err = decoder.Decode(&obj.FundFeesToken0); err != nil {
		return fmt.Errorf("error while unmarshaling FundFeesToken0:%w", err)
	}
	// Deserialize `FundFeesToken1`:
	if err = decoder.Decode(&obj.FundFeesToken1); err != nil {
		return fmt.Errorf("error while unmarshaling FundFeesToken1:%w", err)
	}
	// Deserialize `OpenTime`:
	if err = decoder.Decode(&obj.OpenTime); err != nil {
		return fmt.Errorf("error while unmarshaling OpenTime:%w", err)
	}
	// Deserialize `RecentEpoch`:
	if err = decoder.Decode(&obj.RecentEpoch); err != nil {
		return fmt.Errorf("error while unmarshaling RecentEpoch:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *PoolState) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling PoolState: %w", err)
	}
	return nil
}

func UnmarshalPoolState(buf []byte) (*PoolState, error) {
	obj := new(PoolState)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
