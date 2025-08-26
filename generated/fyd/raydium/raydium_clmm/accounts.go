package raydium_clmm

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

// Holds the current owner of the factory
type AmmConfig struct {
	// Bump to identify PDA
	Bump  uint8
	Index uint16

	// Address of the protocol owner
	Owner solanago.PublicKey

	// The protocol fee
	ProtocolFeeRate uint32

	// The trade fee, denominated in hundredths of a bip (10^-6)
	TradeFeeRate uint32

	// The tick spacing
	TickSpacing uint16

	// The fund fee, denominated in hundredths of a bip (10^-6)
	FundFeeRate uint32
	PaddingU32  uint32
	FundOwner   solanago.PublicKey
	Padding     [3]uint64
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
	// Serialize `Index`:
	if err = encoder.Encode(obj.Index); err != nil {
		return fmt.Errorf("error while marshaling Index:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `ProtocolFeeRate`:
	if err = encoder.Encode(obj.ProtocolFeeRate); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeRate:%w", err)
	}
	// Serialize `TradeFeeRate`:
	if err = encoder.Encode(obj.TradeFeeRate); err != nil {
		return fmt.Errorf("error while marshaling TradeFeeRate:%w", err)
	}
	// Serialize `TickSpacing`:
	if err = encoder.Encode(obj.TickSpacing); err != nil {
		return fmt.Errorf("error while marshaling TickSpacing:%w", err)
	}
	// Serialize `FundFeeRate`:
	if err = encoder.Encode(obj.FundFeeRate); err != nil {
		return fmt.Errorf("error while marshaling FundFeeRate:%w", err)
	}
	// Serialize `PaddingU32`:
	if err = encoder.Encode(obj.PaddingU32); err != nil {
		return fmt.Errorf("error while marshaling PaddingU32:%w", err)
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
	// Deserialize `Index`:
	if err = decoder.Decode(&obj.Index); err != nil {
		return fmt.Errorf("error while unmarshaling Index:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `ProtocolFeeRate`:
	if err = decoder.Decode(&obj.ProtocolFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeRate:%w", err)
	}
	// Deserialize `TradeFeeRate`:
	if err = decoder.Decode(&obj.TradeFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling TradeFeeRate:%w", err)
	}
	// Deserialize `TickSpacing`:
	if err = decoder.Decode(&obj.TickSpacing); err != nil {
		return fmt.Errorf("error while unmarshaling TickSpacing:%w", err)
	}
	// Deserialize `FundFeeRate`:
	if err = decoder.Decode(&obj.FundFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling FundFeeRate:%w", err)
	}
	// Deserialize `PaddingU32`:
	if err = decoder.Decode(&obj.PaddingU32); err != nil {
		return fmt.Errorf("error while unmarshaling PaddingU32:%w", err)
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

	// recent update epoch
	RecentEpoch uint64

	// the most-recently updated index of the observations array
	ObservationIndex uint16

	// belongs to which pool
	PoolId solanago.PublicKey

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
	// Serialize `RecentEpoch`:
	if err = encoder.Encode(obj.RecentEpoch); err != nil {
		return fmt.Errorf("error while marshaling RecentEpoch:%w", err)
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
	// Deserialize `RecentEpoch`:
	if err = decoder.Decode(&obj.RecentEpoch); err != nil {
		return fmt.Errorf("error while unmarshaling RecentEpoch:%w", err)
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

// Holds the current owner of the factory
type OperationState struct {
	// Bump to identify PDA
	Bump uint8

	// Address of the operation owner
	OperationOwners [10]solanago.PublicKey

	// The mint address of whitelist to emit reward
	WhitelistMints [100]solanago.PublicKey
}

func (obj OperationState) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_OperationState[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	// Serialize `OperationOwners`:
	if err = encoder.Encode(obj.OperationOwners); err != nil {
		return fmt.Errorf("error while marshaling OperationOwners:%w", err)
	}
	// Serialize `WhitelistMints`:
	if err = encoder.Encode(obj.WhitelistMints); err != nil {
		return fmt.Errorf("error while marshaling WhitelistMints:%w", err)
	}
	return nil
}

func (obj OperationState) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding OperationState: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *OperationState) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_OperationState[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_OperationState[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	// Deserialize `OperationOwners`:
	if err = decoder.Decode(&obj.OperationOwners); err != nil {
		return fmt.Errorf("error while unmarshaling OperationOwners:%w", err)
	}
	// Deserialize `WhitelistMints`:
	if err = decoder.Decode(&obj.WhitelistMints); err != nil {
		return fmt.Errorf("error while unmarshaling WhitelistMints:%w", err)
	}
	return nil
}

func (obj *OperationState) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling OperationState: %w", err)
	}
	return nil
}

func UnmarshalOperationState(buf []byte) (*OperationState, error) {
	obj := new(OperationState)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type PersonalPositionState struct {
	// Bump to identify PDA
	Bump [1]uint8

	// Mint address of the tokenized position
	NftMint solanago.PublicKey

	// The ID of the pool with which this token is connected
	PoolId solanago.PublicKey

	// The lower bound tick of the position
	TickLowerIndex int32

	// The upper bound tick of the position
	TickUpperIndex int32

	// The amount of liquidity owned by this position
	Liquidity binary.Uint128

	// The token_0 fee growth of the aggregate position as of the last action on the individual position
	FeeGrowthInside0LastX64 binary.Uint128

	// The token_1 fee growth of the aggregate position as of the last action on the individual position
	FeeGrowthInside1LastX64 binary.Uint128

	// The fees owed to the position owner in token_0, as of the last computation
	TokenFeesOwed0 uint64

	// The fees owed to the position owner in token_1, as of the last computation
	TokenFeesOwed1 uint64
	RewardInfos    [3]PositionRewardInfo
	RecentEpoch    uint64
	Padding        [7]uint64
}

func (obj PersonalPositionState) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_PersonalPositionState[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	// Serialize `NftMint`:
	if err = encoder.Encode(obj.NftMint); err != nil {
		return fmt.Errorf("error while marshaling NftMint:%w", err)
	}
	// Serialize `PoolId`:
	if err = encoder.Encode(obj.PoolId); err != nil {
		return fmt.Errorf("error while marshaling PoolId:%w", err)
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
	// Serialize `FeeGrowthInside0LastX64`:
	if err = encoder.Encode(obj.FeeGrowthInside0LastX64); err != nil {
		return fmt.Errorf("error while marshaling FeeGrowthInside0LastX64:%w", err)
	}
	// Serialize `FeeGrowthInside1LastX64`:
	if err = encoder.Encode(obj.FeeGrowthInside1LastX64); err != nil {
		return fmt.Errorf("error while marshaling FeeGrowthInside1LastX64:%w", err)
	}
	// Serialize `TokenFeesOwed0`:
	if err = encoder.Encode(obj.TokenFeesOwed0); err != nil {
		return fmt.Errorf("error while marshaling TokenFeesOwed0:%w", err)
	}
	// Serialize `TokenFeesOwed1`:
	if err = encoder.Encode(obj.TokenFeesOwed1); err != nil {
		return fmt.Errorf("error while marshaling TokenFeesOwed1:%w", err)
	}
	// Serialize `RewardInfos`:
	if err = encoder.Encode(obj.RewardInfos); err != nil {
		return fmt.Errorf("error while marshaling RewardInfos:%w", err)
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

func (obj PersonalPositionState) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding PersonalPositionState: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *PersonalPositionState) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_PersonalPositionState[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_PersonalPositionState[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	// Deserialize `NftMint`:
	if err = decoder.Decode(&obj.NftMint); err != nil {
		return fmt.Errorf("error while unmarshaling NftMint:%w", err)
	}
	// Deserialize `PoolId`:
	if err = decoder.Decode(&obj.PoolId); err != nil {
		return fmt.Errorf("error while unmarshaling PoolId:%w", err)
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
	// Deserialize `FeeGrowthInside0LastX64`:
	if err = decoder.Decode(&obj.FeeGrowthInside0LastX64); err != nil {
		return fmt.Errorf("error while unmarshaling FeeGrowthInside0LastX64:%w", err)
	}
	// Deserialize `FeeGrowthInside1LastX64`:
	if err = decoder.Decode(&obj.FeeGrowthInside1LastX64); err != nil {
		return fmt.Errorf("error while unmarshaling FeeGrowthInside1LastX64:%w", err)
	}
	// Deserialize `TokenFeesOwed0`:
	if err = decoder.Decode(&obj.TokenFeesOwed0); err != nil {
		return fmt.Errorf("error while unmarshaling TokenFeesOwed0:%w", err)
	}
	// Deserialize `TokenFeesOwed1`:
	if err = decoder.Decode(&obj.TokenFeesOwed1); err != nil {
		return fmt.Errorf("error while unmarshaling TokenFeesOwed1:%w", err)
	}
	// Deserialize `RewardInfos`:
	if err = decoder.Decode(&obj.RewardInfos); err != nil {
		return fmt.Errorf("error while unmarshaling RewardInfos:%w", err)
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

func (obj *PersonalPositionState) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling PersonalPositionState: %w", err)
	}
	return nil
}

func UnmarshalPersonalPositionState(buf []byte) (*PersonalPositionState, error) {
	obj := new(PersonalPositionState)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// The pool state
//
// PDA of `[POOL_SEED, config, token_mint_0, token_mint_1]`
type PoolState struct {
	// Bump to identify PDA
	Bump      [1]uint8
	AmmConfig solanago.PublicKey
	Owner     solanago.PublicKey

	// Token pair of the pool, where token_mint_0 address < token_mint_1 address
	TokenMint0 solanago.PublicKey
	TokenMint1 solanago.PublicKey

	// Token pair vault
	TokenVault0 solanago.PublicKey
	TokenVault1 solanago.PublicKey

	// observation account key
	ObservationKey solanago.PublicKey

	// mint0 and mint1 decimals
	MintDecimals0 uint8
	MintDecimals1 uint8

	// The minimum number of ticks between initialized ticks
	TickSpacing uint16

	// The currently in range liquidity available to the pool.
	Liquidity binary.Uint128

	// The current price of the pool as a sqrt(token_1/token_0) Q64.64 value
	SqrtPriceX64 binary.Uint128

	// The current tick of the pool, i.e. according to the last tick transition that was run.
	TickCurrent int32
	Padding3    uint16
	Padding4    uint16

	// The fee growth as a Q64.64 number, i.e. fees of token_0 and token_1 collected per
	// unit of liquidity for the entire life of the pool.
	FeeGrowthGlobal0X64 binary.Uint128
	FeeGrowthGlobal1X64 binary.Uint128

	// The amounts of token_0 and token_1 that are owed to the protocol.
	ProtocolFeesToken0 uint64
	ProtocolFeesToken1 uint64

	// The amounts in and out of swap token_0 and token_1
	SwapInAmountToken0  binary.Uint128
	SwapOutAmountToken1 binary.Uint128
	SwapInAmountToken1  binary.Uint128
	SwapOutAmountToken0 binary.Uint128

	// Bitwise representation of the state of the pool
	// bit0, 1: disable open position and increase liquidity, 0: normal
	// bit1, 1: disable decrease liquidity, 0: normal
	// bit2, 1: disable collect fee, 0: normal
	// bit3, 1: disable collect reward, 0: normal
	// bit4, 1: disable swap, 0: normal
	Status uint8

	// Leave blank for future use
	Padding     [7]uint8
	RewardInfos [3]RewardInfo

	// Packed initialized tick array state
	TickArrayBitmap [16]uint64

	// except protocol_fee and fund_fee
	TotalFeesToken0 uint64

	// except protocol_fee and fund_fee
	TotalFeesClaimedToken0 uint64
	TotalFeesToken1        uint64
	TotalFeesClaimedToken1 uint64
	FundFeesToken0         uint64
	FundFeesToken1         uint64
	OpenTime               uint64
	RecentEpoch            uint64
	Padding1               [24]uint64
	Padding2               [32]uint64
}

func (obj PoolState) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_PoolState[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	// Serialize `AmmConfig`:
	if err = encoder.Encode(obj.AmmConfig); err != nil {
		return fmt.Errorf("error while marshaling AmmConfig:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `TokenMint0`:
	if err = encoder.Encode(obj.TokenMint0); err != nil {
		return fmt.Errorf("error while marshaling TokenMint0:%w", err)
	}
	// Serialize `TokenMint1`:
	if err = encoder.Encode(obj.TokenMint1); err != nil {
		return fmt.Errorf("error while marshaling TokenMint1:%w", err)
	}
	// Serialize `TokenVault0`:
	if err = encoder.Encode(obj.TokenVault0); err != nil {
		return fmt.Errorf("error while marshaling TokenVault0:%w", err)
	}
	// Serialize `TokenVault1`:
	if err = encoder.Encode(obj.TokenVault1); err != nil {
		return fmt.Errorf("error while marshaling TokenVault1:%w", err)
	}
	// Serialize `ObservationKey`:
	if err = encoder.Encode(obj.ObservationKey); err != nil {
		return fmt.Errorf("error while marshaling ObservationKey:%w", err)
	}
	// Serialize `MintDecimals0`:
	if err = encoder.Encode(obj.MintDecimals0); err != nil {
		return fmt.Errorf("error while marshaling MintDecimals0:%w", err)
	}
	// Serialize `MintDecimals1`:
	if err = encoder.Encode(obj.MintDecimals1); err != nil {
		return fmt.Errorf("error while marshaling MintDecimals1:%w", err)
	}
	// Serialize `TickSpacing`:
	if err = encoder.Encode(obj.TickSpacing); err != nil {
		return fmt.Errorf("error while marshaling TickSpacing:%w", err)
	}
	// Serialize `Liquidity`:
	if err = encoder.Encode(obj.Liquidity); err != nil {
		return fmt.Errorf("error while marshaling Liquidity:%w", err)
	}
	// Serialize `SqrtPriceX64`:
	if err = encoder.Encode(obj.SqrtPriceX64); err != nil {
		return fmt.Errorf("error while marshaling SqrtPriceX64:%w", err)
	}
	// Serialize `TickCurrent`:
	if err = encoder.Encode(obj.TickCurrent); err != nil {
		return fmt.Errorf("error while marshaling TickCurrent:%w", err)
	}
	// Serialize `Padding3`:
	if err = encoder.Encode(obj.Padding3); err != nil {
		return fmt.Errorf("error while marshaling Padding3:%w", err)
	}
	// Serialize `Padding4`:
	if err = encoder.Encode(obj.Padding4); err != nil {
		return fmt.Errorf("error while marshaling Padding4:%w", err)
	}
	// Serialize `FeeGrowthGlobal0X64`:
	if err = encoder.Encode(obj.FeeGrowthGlobal0X64); err != nil {
		return fmt.Errorf("error while marshaling FeeGrowthGlobal0X64:%w", err)
	}
	// Serialize `FeeGrowthGlobal1X64`:
	if err = encoder.Encode(obj.FeeGrowthGlobal1X64); err != nil {
		return fmt.Errorf("error while marshaling FeeGrowthGlobal1X64:%w", err)
	}
	// Serialize `ProtocolFeesToken0`:
	if err = encoder.Encode(obj.ProtocolFeesToken0); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeesToken0:%w", err)
	}
	// Serialize `ProtocolFeesToken1`:
	if err = encoder.Encode(obj.ProtocolFeesToken1); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeesToken1:%w", err)
	}
	// Serialize `SwapInAmountToken0`:
	if err = encoder.Encode(obj.SwapInAmountToken0); err != nil {
		return fmt.Errorf("error while marshaling SwapInAmountToken0:%w", err)
	}
	// Serialize `SwapOutAmountToken1`:
	if err = encoder.Encode(obj.SwapOutAmountToken1); err != nil {
		return fmt.Errorf("error while marshaling SwapOutAmountToken1:%w", err)
	}
	// Serialize `SwapInAmountToken1`:
	if err = encoder.Encode(obj.SwapInAmountToken1); err != nil {
		return fmt.Errorf("error while marshaling SwapInAmountToken1:%w", err)
	}
	// Serialize `SwapOutAmountToken0`:
	if err = encoder.Encode(obj.SwapOutAmountToken0); err != nil {
		return fmt.Errorf("error while marshaling SwapOutAmountToken0:%w", err)
	}
	// Serialize `Status`:
	if err = encoder.Encode(obj.Status); err != nil {
		return fmt.Errorf("error while marshaling Status:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	// Serialize `RewardInfos`:
	if err = encoder.Encode(obj.RewardInfos); err != nil {
		return fmt.Errorf("error while marshaling RewardInfos:%w", err)
	}
	// Serialize `TickArrayBitmap`:
	if err = encoder.Encode(obj.TickArrayBitmap); err != nil {
		return fmt.Errorf("error while marshaling TickArrayBitmap:%w", err)
	}
	// Serialize `TotalFeesToken0`:
	if err = encoder.Encode(obj.TotalFeesToken0); err != nil {
		return fmt.Errorf("error while marshaling TotalFeesToken0:%w", err)
	}
	// Serialize `TotalFeesClaimedToken0`:
	if err = encoder.Encode(obj.TotalFeesClaimedToken0); err != nil {
		return fmt.Errorf("error while marshaling TotalFeesClaimedToken0:%w", err)
	}
	// Serialize `TotalFeesToken1`:
	if err = encoder.Encode(obj.TotalFeesToken1); err != nil {
		return fmt.Errorf("error while marshaling TotalFeesToken1:%w", err)
	}
	// Serialize `TotalFeesClaimedToken1`:
	if err = encoder.Encode(obj.TotalFeesClaimedToken1); err != nil {
		return fmt.Errorf("error while marshaling TotalFeesClaimedToken1:%w", err)
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
	// Serialize `Padding1`:
	if err = encoder.Encode(obj.Padding1); err != nil {
		return fmt.Errorf("error while marshaling Padding1:%w", err)
	}
	// Serialize `Padding2`:
	if err = encoder.Encode(obj.Padding2); err != nil {
		return fmt.Errorf("error while marshaling Padding2:%w", err)
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
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	// Deserialize `AmmConfig`:
	if err = decoder.Decode(&obj.AmmConfig); err != nil {
		return fmt.Errorf("error while unmarshaling AmmConfig:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `TokenMint0`:
	if err = decoder.Decode(&obj.TokenMint0); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMint0:%w", err)
	}
	// Deserialize `TokenMint1`:
	if err = decoder.Decode(&obj.TokenMint1); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMint1:%w", err)
	}
	// Deserialize `TokenVault0`:
	if err = decoder.Decode(&obj.TokenVault0); err != nil {
		return fmt.Errorf("error while unmarshaling TokenVault0:%w", err)
	}
	// Deserialize `TokenVault1`:
	if err = decoder.Decode(&obj.TokenVault1); err != nil {
		return fmt.Errorf("error while unmarshaling TokenVault1:%w", err)
	}
	// Deserialize `ObservationKey`:
	if err = decoder.Decode(&obj.ObservationKey); err != nil {
		return fmt.Errorf("error while unmarshaling ObservationKey:%w", err)
	}
	// Deserialize `MintDecimals0`:
	if err = decoder.Decode(&obj.MintDecimals0); err != nil {
		return fmt.Errorf("error while unmarshaling MintDecimals0:%w", err)
	}
	// Deserialize `MintDecimals1`:
	if err = decoder.Decode(&obj.MintDecimals1); err != nil {
		return fmt.Errorf("error while unmarshaling MintDecimals1:%w", err)
	}
	// Deserialize `TickSpacing`:
	if err = decoder.Decode(&obj.TickSpacing); err != nil {
		return fmt.Errorf("error while unmarshaling TickSpacing:%w", err)
	}
	// Deserialize `Liquidity`:
	if err = decoder.Decode(&obj.Liquidity); err != nil {
		return fmt.Errorf("error while unmarshaling Liquidity:%w", err)
	}
	// Deserialize `SqrtPriceX64`:
	if err = decoder.Decode(&obj.SqrtPriceX64); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtPriceX64:%w", err)
	}
	// Deserialize `TickCurrent`:
	if err = decoder.Decode(&obj.TickCurrent); err != nil {
		return fmt.Errorf("error while unmarshaling TickCurrent:%w", err)
	}
	// Deserialize `Padding3`:
	if err = decoder.Decode(&obj.Padding3); err != nil {
		return fmt.Errorf("error while unmarshaling Padding3:%w", err)
	}
	// Deserialize `Padding4`:
	if err = decoder.Decode(&obj.Padding4); err != nil {
		return fmt.Errorf("error while unmarshaling Padding4:%w", err)
	}
	// Deserialize `FeeGrowthGlobal0X64`:
	if err = decoder.Decode(&obj.FeeGrowthGlobal0X64); err != nil {
		return fmt.Errorf("error while unmarshaling FeeGrowthGlobal0X64:%w", err)
	}
	// Deserialize `FeeGrowthGlobal1X64`:
	if err = decoder.Decode(&obj.FeeGrowthGlobal1X64); err != nil {
		return fmt.Errorf("error while unmarshaling FeeGrowthGlobal1X64:%w", err)
	}
	// Deserialize `ProtocolFeesToken0`:
	if err = decoder.Decode(&obj.ProtocolFeesToken0); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeesToken0:%w", err)
	}
	// Deserialize `ProtocolFeesToken1`:
	if err = decoder.Decode(&obj.ProtocolFeesToken1); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeesToken1:%w", err)
	}
	// Deserialize `SwapInAmountToken0`:
	if err = decoder.Decode(&obj.SwapInAmountToken0); err != nil {
		return fmt.Errorf("error while unmarshaling SwapInAmountToken0:%w", err)
	}
	// Deserialize `SwapOutAmountToken1`:
	if err = decoder.Decode(&obj.SwapOutAmountToken1); err != nil {
		return fmt.Errorf("error while unmarshaling SwapOutAmountToken1:%w", err)
	}
	// Deserialize `SwapInAmountToken1`:
	if err = decoder.Decode(&obj.SwapInAmountToken1); err != nil {
		return fmt.Errorf("error while unmarshaling SwapInAmountToken1:%w", err)
	}
	// Deserialize `SwapOutAmountToken0`:
	if err = decoder.Decode(&obj.SwapOutAmountToken0); err != nil {
		return fmt.Errorf("error while unmarshaling SwapOutAmountToken0:%w", err)
	}
	// Deserialize `Status`:
	if err = decoder.Decode(&obj.Status); err != nil {
		return fmt.Errorf("error while unmarshaling Status:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	// Deserialize `RewardInfos`:
	if err = decoder.Decode(&obj.RewardInfos); err != nil {
		return fmt.Errorf("error while unmarshaling RewardInfos:%w", err)
	}
	// Deserialize `TickArrayBitmap`:
	if err = decoder.Decode(&obj.TickArrayBitmap); err != nil {
		return fmt.Errorf("error while unmarshaling TickArrayBitmap:%w", err)
	}
	// Deserialize `TotalFeesToken0`:
	if err = decoder.Decode(&obj.TotalFeesToken0); err != nil {
		return fmt.Errorf("error while unmarshaling TotalFeesToken0:%w", err)
	}
	// Deserialize `TotalFeesClaimedToken0`:
	if err = decoder.Decode(&obj.TotalFeesClaimedToken0); err != nil {
		return fmt.Errorf("error while unmarshaling TotalFeesClaimedToken0:%w", err)
	}
	// Deserialize `TotalFeesToken1`:
	if err = decoder.Decode(&obj.TotalFeesToken1); err != nil {
		return fmt.Errorf("error while unmarshaling TotalFeesToken1:%w", err)
	}
	// Deserialize `TotalFeesClaimedToken1`:
	if err = decoder.Decode(&obj.TotalFeesClaimedToken1); err != nil {
		return fmt.Errorf("error while unmarshaling TotalFeesClaimedToken1:%w", err)
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
	// Deserialize `Padding1`:
	if err = decoder.Decode(&obj.Padding1); err != nil {
		return fmt.Errorf("error while unmarshaling Padding1:%w", err)
	}
	// Deserialize `Padding2`:
	if err = decoder.Decode(&obj.Padding2); err != nil {
		return fmt.Errorf("error while unmarshaling Padding2:%w", err)
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

// Info stored for each user's position
type ProtocolPositionState struct {
	// Bump to identify PDA
	Bump uint8

	// The ID of the pool with which this token is connected
	PoolId solanago.PublicKey

	// The lower bound tick of the position
	TickLowerIndex int32

	// The upper bound tick of the position
	TickUpperIndex int32

	// The amount of liquidity owned by this position
	Liquidity binary.Uint128

	// The token_0 fee growth per unit of liquidity as of the last update to liquidity or fees owed
	FeeGrowthInside0LastX64 binary.Uint128

	// The token_1 fee growth per unit of liquidity as of the last update to liquidity or fees owed
	FeeGrowthInside1LastX64 binary.Uint128

	// The fees owed to the position owner in token_0
	TokenFeesOwed0 uint64

	// The fees owed to the position owner in token_1
	TokenFeesOwed1 uint64

	// The reward growth per unit of liquidity as of the last update to liquidity
	RewardGrowthInside [3]binary.Uint128
	RecentEpoch        uint64
	Padding            [7]uint64
}

func (obj ProtocolPositionState) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_ProtocolPositionState[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	// Serialize `PoolId`:
	if err = encoder.Encode(obj.PoolId); err != nil {
		return fmt.Errorf("error while marshaling PoolId:%w", err)
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
	// Serialize `FeeGrowthInside0LastX64`:
	if err = encoder.Encode(obj.FeeGrowthInside0LastX64); err != nil {
		return fmt.Errorf("error while marshaling FeeGrowthInside0LastX64:%w", err)
	}
	// Serialize `FeeGrowthInside1LastX64`:
	if err = encoder.Encode(obj.FeeGrowthInside1LastX64); err != nil {
		return fmt.Errorf("error while marshaling FeeGrowthInside1LastX64:%w", err)
	}
	// Serialize `TokenFeesOwed0`:
	if err = encoder.Encode(obj.TokenFeesOwed0); err != nil {
		return fmt.Errorf("error while marshaling TokenFeesOwed0:%w", err)
	}
	// Serialize `TokenFeesOwed1`:
	if err = encoder.Encode(obj.TokenFeesOwed1); err != nil {
		return fmt.Errorf("error while marshaling TokenFeesOwed1:%w", err)
	}
	// Serialize `RewardGrowthInside`:
	if err = encoder.Encode(obj.RewardGrowthInside); err != nil {
		return fmt.Errorf("error while marshaling RewardGrowthInside:%w", err)
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

func (obj ProtocolPositionState) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding ProtocolPositionState: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *ProtocolPositionState) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_ProtocolPositionState[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_ProtocolPositionState[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	// Deserialize `PoolId`:
	if err = decoder.Decode(&obj.PoolId); err != nil {
		return fmt.Errorf("error while unmarshaling PoolId:%w", err)
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
	// Deserialize `FeeGrowthInside0LastX64`:
	if err = decoder.Decode(&obj.FeeGrowthInside0LastX64); err != nil {
		return fmt.Errorf("error while unmarshaling FeeGrowthInside0LastX64:%w", err)
	}
	// Deserialize `FeeGrowthInside1LastX64`:
	if err = decoder.Decode(&obj.FeeGrowthInside1LastX64); err != nil {
		return fmt.Errorf("error while unmarshaling FeeGrowthInside1LastX64:%w", err)
	}
	// Deserialize `TokenFeesOwed0`:
	if err = decoder.Decode(&obj.TokenFeesOwed0); err != nil {
		return fmt.Errorf("error while unmarshaling TokenFeesOwed0:%w", err)
	}
	// Deserialize `TokenFeesOwed1`:
	if err = decoder.Decode(&obj.TokenFeesOwed1); err != nil {
		return fmt.Errorf("error while unmarshaling TokenFeesOwed1:%w", err)
	}
	// Deserialize `RewardGrowthInside`:
	if err = decoder.Decode(&obj.RewardGrowthInside); err != nil {
		return fmt.Errorf("error while unmarshaling RewardGrowthInside:%w", err)
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

func (obj *ProtocolPositionState) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolPositionState: %w", err)
	}
	return nil
}

func UnmarshalProtocolPositionState(buf []byte) (*ProtocolPositionState, error) {
	obj := new(ProtocolPositionState)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Holds the current owner of the factory
type SupportMintAssociated struct {
	// Bump to identify PDA
	Bump uint8

	// Address of the supported token22 mint
	Mint    solanago.PublicKey
	Padding [8]uint64
}

func (obj SupportMintAssociated) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_SupportMintAssociated[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj SupportMintAssociated) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding SupportMintAssociated: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *SupportMintAssociated) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_SupportMintAssociated[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_SupportMintAssociated[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *SupportMintAssociated) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling SupportMintAssociated: %w", err)
	}
	return nil
}

func UnmarshalSupportMintAssociated(buf []byte) (*SupportMintAssociated, error) {
	obj := new(SupportMintAssociated)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type TickArrayBitmapExtension struct {
	PoolId solanago.PublicKey

	// Packed initialized tick array state for start_tick_index is positive
	PositiveTickArrayBitmap [14][8]uint64

	// Packed initialized tick array state for start_tick_index is negitive
	NegativeTickArrayBitmap [14][8]uint64
}

func (obj TickArrayBitmapExtension) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_TickArrayBitmapExtension[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolId`:
	if err = encoder.Encode(obj.PoolId); err != nil {
		return fmt.Errorf("error while marshaling PoolId:%w", err)
	}
	// Serialize `PositiveTickArrayBitmap`:
	if err = encoder.Encode(obj.PositiveTickArrayBitmap); err != nil {
		return fmt.Errorf("error while marshaling PositiveTickArrayBitmap:%w", err)
	}
	// Serialize `NegativeTickArrayBitmap`:
	if err = encoder.Encode(obj.NegativeTickArrayBitmap); err != nil {
		return fmt.Errorf("error while marshaling NegativeTickArrayBitmap:%w", err)
	}
	return nil
}

func (obj TickArrayBitmapExtension) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding TickArrayBitmapExtension: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *TickArrayBitmapExtension) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_TickArrayBitmapExtension[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_TickArrayBitmapExtension[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolId`:
	if err = decoder.Decode(&obj.PoolId); err != nil {
		return fmt.Errorf("error while unmarshaling PoolId:%w", err)
	}
	// Deserialize `PositiveTickArrayBitmap`:
	if err = decoder.Decode(&obj.PositiveTickArrayBitmap); err != nil {
		return fmt.Errorf("error while unmarshaling PositiveTickArrayBitmap:%w", err)
	}
	// Deserialize `NegativeTickArrayBitmap`:
	if err = decoder.Decode(&obj.NegativeTickArrayBitmap); err != nil {
		return fmt.Errorf("error while unmarshaling NegativeTickArrayBitmap:%w", err)
	}
	return nil
}

func (obj *TickArrayBitmapExtension) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling TickArrayBitmapExtension: %w", err)
	}
	return nil
}

func UnmarshalTickArrayBitmapExtension(buf []byte) (*TickArrayBitmapExtension, error) {
	obj := new(TickArrayBitmapExtension)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type TickArrayState struct {
	PoolId               solanago.PublicKey
	StartTickIndex       int32
	Ticks                [60]TickState
	InitializedTickCount uint8
	RecentEpoch          uint64
	Padding              [107]uint8
}

func (obj TickArrayState) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_TickArrayState[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolId`:
	if err = encoder.Encode(obj.PoolId); err != nil {
		return fmt.Errorf("error while marshaling PoolId:%w", err)
	}
	// Serialize `StartTickIndex`:
	if err = encoder.Encode(obj.StartTickIndex); err != nil {
		return fmt.Errorf("error while marshaling StartTickIndex:%w", err)
	}
	// Serialize `Ticks`:
	if err = encoder.Encode(obj.Ticks); err != nil {
		return fmt.Errorf("error while marshaling Ticks:%w", err)
	}
	// Serialize `InitializedTickCount`:
	if err = encoder.Encode(obj.InitializedTickCount); err != nil {
		return fmt.Errorf("error while marshaling InitializedTickCount:%w", err)
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

func (obj TickArrayState) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding TickArrayState: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *TickArrayState) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_TickArrayState[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_TickArrayState[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolId`:
	if err = decoder.Decode(&obj.PoolId); err != nil {
		return fmt.Errorf("error while unmarshaling PoolId:%w", err)
	}
	// Deserialize `StartTickIndex`:
	if err = decoder.Decode(&obj.StartTickIndex); err != nil {
		return fmt.Errorf("error while unmarshaling StartTickIndex:%w", err)
	}
	// Deserialize `Ticks`:
	if err = decoder.Decode(&obj.Ticks); err != nil {
		return fmt.Errorf("error while unmarshaling Ticks:%w", err)
	}
	// Deserialize `InitializedTickCount`:
	if err = decoder.Decode(&obj.InitializedTickCount); err != nil {
		return fmt.Errorf("error while unmarshaling InitializedTickCount:%w", err)
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

func (obj *TickArrayState) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling TickArrayState: %w", err)
	}
	return nil
}

func UnmarshalTickArrayState(buf []byte) (*TickArrayState, error) {
	obj := new(TickArrayState)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
