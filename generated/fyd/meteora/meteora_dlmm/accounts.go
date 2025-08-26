package meteora_dlmm

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

// An account to contain a range of bin. For example: Bin 100 <-> 200.
// For example:
// BinArray index: 0 contains bin 0 <-> 599
// index: 2 contains bin 600 <-> 1199, ...
type BinArray struct {
	Index int64

	// Version of binArray
	Version uint8
	Padding [7]uint8
	LbPair  solanago.PublicKey
	Bins    [70]Bin
}

func (obj BinArray) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_BinArray[:], false)
	if err != nil {
		return err
	}
	// Serialize `Index`:
	if err = encoder.Encode(obj.Index); err != nil {
		return fmt.Errorf("error while marshaling Index:%w", err)
	}
	// Serialize `Version`:
	if err = encoder.Encode(obj.Version); err != nil {
		return fmt.Errorf("error while marshaling Version:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `Bins`:
	if err = encoder.Encode(obj.Bins); err != nil {
		return fmt.Errorf("error while marshaling Bins:%w", err)
	}
	return nil
}

func (obj BinArray) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding BinArray: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *BinArray) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_BinArray[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_BinArray[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Index`:
	if err = decoder.Decode(&obj.Index); err != nil {
		return fmt.Errorf("error while unmarshaling Index:%w", err)
	}
	// Deserialize `Version`:
	if err = decoder.Decode(&obj.Version); err != nil {
		return fmt.Errorf("error while unmarshaling Version:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `Bins`:
	if err = decoder.Decode(&obj.Bins); err != nil {
		return fmt.Errorf("error while unmarshaling Bins:%w", err)
	}
	return nil
}

func (obj *BinArray) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling BinArray: %w", err)
	}
	return nil
}

func UnmarshalBinArray(buf []byte) (*BinArray, error) {
	obj := new(BinArray)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type BinArrayBitmapExtension struct {
	LbPair solanago.PublicKey

	// Packed initialized bin array state for start_bin_index is positive
	PositiveBinArrayBitmap [12][8]uint64

	// Packed initialized bin array state for start_bin_index is negative
	NegativeBinArrayBitmap [12][8]uint64
}

func (obj BinArrayBitmapExtension) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_BinArrayBitmapExtension[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `PositiveBinArrayBitmap`:
	if err = encoder.Encode(obj.PositiveBinArrayBitmap); err != nil {
		return fmt.Errorf("error while marshaling PositiveBinArrayBitmap:%w", err)
	}
	// Serialize `NegativeBinArrayBitmap`:
	if err = encoder.Encode(obj.NegativeBinArrayBitmap); err != nil {
		return fmt.Errorf("error while marshaling NegativeBinArrayBitmap:%w", err)
	}
	return nil
}

func (obj BinArrayBitmapExtension) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding BinArrayBitmapExtension: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *BinArrayBitmapExtension) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_BinArrayBitmapExtension[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_BinArrayBitmapExtension[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `PositiveBinArrayBitmap`:
	if err = decoder.Decode(&obj.PositiveBinArrayBitmap); err != nil {
		return fmt.Errorf("error while unmarshaling PositiveBinArrayBitmap:%w", err)
	}
	// Deserialize `NegativeBinArrayBitmap`:
	if err = decoder.Decode(&obj.NegativeBinArrayBitmap); err != nil {
		return fmt.Errorf("error while unmarshaling NegativeBinArrayBitmap:%w", err)
	}
	return nil
}

func (obj *BinArrayBitmapExtension) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling BinArrayBitmapExtension: %w", err)
	}
	return nil
}

func UnmarshalBinArrayBitmapExtension(buf []byte) (*BinArrayBitmapExtension, error) {
	obj := new(BinArrayBitmapExtension)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Parameter that set by the protocol
type ClaimFeeOperator struct {
	// operator
	Operator solanago.PublicKey

	// Reserve
	Padding [128]uint8
}

func (obj ClaimFeeOperator) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_ClaimFeeOperator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Operator`:
	if err = encoder.Encode(obj.Operator); err != nil {
		return fmt.Errorf("error while marshaling Operator:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj ClaimFeeOperator) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding ClaimFeeOperator: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *ClaimFeeOperator) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_ClaimFeeOperator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_ClaimFeeOperator[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Operator`:
	if err = decoder.Decode(&obj.Operator); err != nil {
		return fmt.Errorf("error while unmarshaling Operator:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *ClaimFeeOperator) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling ClaimFeeOperator: %w", err)
	}
	return nil
}

func UnmarshalClaimFeeOperator(buf []byte) (*ClaimFeeOperator, error) {
	obj := new(ClaimFeeOperator)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type DummyZcAccount struct {
	PositionBinData PositionBinData
}

func (obj DummyZcAccount) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_DummyZcAccount[:], false)
	if err != nil {
		return err
	}
	// Serialize `PositionBinData`:
	if err = encoder.Encode(obj.PositionBinData); err != nil {
		return fmt.Errorf("error while marshaling PositionBinData:%w", err)
	}
	return nil
}

func (obj DummyZcAccount) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding DummyZcAccount: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *DummyZcAccount) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_DummyZcAccount[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_DummyZcAccount[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PositionBinData`:
	if err = decoder.Decode(&obj.PositionBinData); err != nil {
		return fmt.Errorf("error while unmarshaling PositionBinData:%w", err)
	}
	return nil
}

func (obj *DummyZcAccount) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling DummyZcAccount: %w", err)
	}
	return nil
}

func UnmarshalDummyZcAccount(buf []byte) (*DummyZcAccount, error) {
	obj := new(DummyZcAccount)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type LbPair struct {
	Parameters  StaticParameters
	VParameters VariableParameters
	BumpSeed    [1]uint8

	// Bin step signer seed
	BinStepSeed [2]uint8

	// Type of the pair
	PairType uint8

	// Active bin id
	ActiveId int32

	// Bin step. Represent the price increment / decrement.
	BinStep uint16

	// Status of the pair. Check PairStatus enum.
	Status uint8

	// Require base factor seed
	RequireBaseFactorSeed uint8

	// Base factor seed
	BaseFactorSeed [2]uint8

	// Activation type
	ActivationType uint8

	// Allow pool creator to enable/disable pool with restricted validation. Only applicable for customizable permissionless pair type.
	CreatorPoolOnOffControl uint8

	// Token X mint
	TokenXMint solanago.PublicKey

	// Token Y mint
	TokenYMint solanago.PublicKey

	// LB token X vault
	ReserveX solanago.PublicKey

	// LB token Y vault
	ReserveY solanago.PublicKey

	// Uncollected protocol fee
	ProtocolFee ProtocolFee

	// _padding_1, previous Fee owner, BE CAREFUL FOR TOMBSTONE WHEN REUSE !!
	Padding1 [32]uint8

	// Farming reward information
	RewardInfos [2]RewardInfo

	// Oracle pubkey
	Oracle solanago.PublicKey

	// Packed initialized bin array state
	BinArrayBitmap [16]uint64

	// Last time the pool fee parameter was updated
	LastUpdatedAt int64

	// _padding_2, previous whitelisted_wallet, BE CAREFUL FOR TOMBSTONE WHEN REUSE !!
	Padding2 [32]uint8

	// Address allowed to swap when the current point is greater than or equal to the pre-activation point. The pre-activation point is calculated as `activation_point - pre_activation_duration`.
	PreActivationSwapAddress solanago.PublicKey

	// Base keypair. Only required for permission pair
	BaseKey solanago.PublicKey

	// Time point to enable the pair. Only applicable for permission pair.
	ActivationPoint uint64

	// Duration before activation activation_point. Used to calculate pre-activation time point for pre_activation_swap_address
	PreActivationDuration uint64

	// _padding 3 is reclaimed free space from swap_cap_deactivate_point and swap_cap_amount before, BE CAREFUL FOR TOMBSTONE WHEN REUSE !!
	Padding3 [8]uint8

	// _padding_4, previous lock_duration, BE CAREFUL FOR TOMBSTONE WHEN REUSE !!
	Padding4 uint64

	// Pool creator
	Creator solanago.PublicKey

	// token_mint_x_program_flag
	TokenMintXProgramFlag uint8

	// token_mint_y_program_flag
	TokenMintYProgramFlag uint8

	// Reserved space for future use
	Reserved [22]uint8
}

func (obj LbPair) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_LbPair[:], false)
	if err != nil {
		return err
	}
	// Serialize `Parameters`:
	if err = encoder.Encode(obj.Parameters); err != nil {
		return fmt.Errorf("error while marshaling Parameters:%w", err)
	}
	// Serialize `VParameters`:
	if err = encoder.Encode(obj.VParameters); err != nil {
		return fmt.Errorf("error while marshaling VParameters:%w", err)
	}
	// Serialize `BumpSeed`:
	if err = encoder.Encode(obj.BumpSeed); err != nil {
		return fmt.Errorf("error while marshaling BumpSeed:%w", err)
	}
	// Serialize `BinStepSeed`:
	if err = encoder.Encode(obj.BinStepSeed); err != nil {
		return fmt.Errorf("error while marshaling BinStepSeed:%w", err)
	}
	// Serialize `PairType`:
	if err = encoder.Encode(obj.PairType); err != nil {
		return fmt.Errorf("error while marshaling PairType:%w", err)
	}
	// Serialize `ActiveId`:
	if err = encoder.Encode(obj.ActiveId); err != nil {
		return fmt.Errorf("error while marshaling ActiveId:%w", err)
	}
	// Serialize `BinStep`:
	if err = encoder.Encode(obj.BinStep); err != nil {
		return fmt.Errorf("error while marshaling BinStep:%w", err)
	}
	// Serialize `Status`:
	if err = encoder.Encode(obj.Status); err != nil {
		return fmt.Errorf("error while marshaling Status:%w", err)
	}
	// Serialize `RequireBaseFactorSeed`:
	if err = encoder.Encode(obj.RequireBaseFactorSeed); err != nil {
		return fmt.Errorf("error while marshaling RequireBaseFactorSeed:%w", err)
	}
	// Serialize `BaseFactorSeed`:
	if err = encoder.Encode(obj.BaseFactorSeed); err != nil {
		return fmt.Errorf("error while marshaling BaseFactorSeed:%w", err)
	}
	// Serialize `ActivationType`:
	if err = encoder.Encode(obj.ActivationType); err != nil {
		return fmt.Errorf("error while marshaling ActivationType:%w", err)
	}
	// Serialize `CreatorPoolOnOffControl`:
	if err = encoder.Encode(obj.CreatorPoolOnOffControl); err != nil {
		return fmt.Errorf("error while marshaling CreatorPoolOnOffControl:%w", err)
	}
	// Serialize `TokenXMint`:
	if err = encoder.Encode(obj.TokenXMint); err != nil {
		return fmt.Errorf("error while marshaling TokenXMint:%w", err)
	}
	// Serialize `TokenYMint`:
	if err = encoder.Encode(obj.TokenYMint); err != nil {
		return fmt.Errorf("error while marshaling TokenYMint:%w", err)
	}
	// Serialize `ReserveX`:
	if err = encoder.Encode(obj.ReserveX); err != nil {
		return fmt.Errorf("error while marshaling ReserveX:%w", err)
	}
	// Serialize `ReserveY`:
	if err = encoder.Encode(obj.ReserveY); err != nil {
		return fmt.Errorf("error while marshaling ReserveY:%w", err)
	}
	// Serialize `ProtocolFee`:
	if err = encoder.Encode(obj.ProtocolFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFee:%w", err)
	}
	// Serialize `Padding1`:
	if err = encoder.Encode(obj.Padding1); err != nil {
		return fmt.Errorf("error while marshaling Padding1:%w", err)
	}
	// Serialize `RewardInfos`:
	if err = encoder.Encode(obj.RewardInfos); err != nil {
		return fmt.Errorf("error while marshaling RewardInfos:%w", err)
	}
	// Serialize `Oracle`:
	if err = encoder.Encode(obj.Oracle); err != nil {
		return fmt.Errorf("error while marshaling Oracle:%w", err)
	}
	// Serialize `BinArrayBitmap`:
	if err = encoder.Encode(obj.BinArrayBitmap); err != nil {
		return fmt.Errorf("error while marshaling BinArrayBitmap:%w", err)
	}
	// Serialize `LastUpdatedAt`:
	if err = encoder.Encode(obj.LastUpdatedAt); err != nil {
		return fmt.Errorf("error while marshaling LastUpdatedAt:%w", err)
	}
	// Serialize `Padding2`:
	if err = encoder.Encode(obj.Padding2); err != nil {
		return fmt.Errorf("error while marshaling Padding2:%w", err)
	}
	// Serialize `PreActivationSwapAddress`:
	if err = encoder.Encode(obj.PreActivationSwapAddress); err != nil {
		return fmt.Errorf("error while marshaling PreActivationSwapAddress:%w", err)
	}
	// Serialize `BaseKey`:
	if err = encoder.Encode(obj.BaseKey); err != nil {
		return fmt.Errorf("error while marshaling BaseKey:%w", err)
	}
	// Serialize `ActivationPoint`:
	if err = encoder.Encode(obj.ActivationPoint); err != nil {
		return fmt.Errorf("error while marshaling ActivationPoint:%w", err)
	}
	// Serialize `PreActivationDuration`:
	if err = encoder.Encode(obj.PreActivationDuration); err != nil {
		return fmt.Errorf("error while marshaling PreActivationDuration:%w", err)
	}
	// Serialize `Padding3`:
	if err = encoder.Encode(obj.Padding3); err != nil {
		return fmt.Errorf("error while marshaling Padding3:%w", err)
	}
	// Serialize `Padding4`:
	if err = encoder.Encode(obj.Padding4); err != nil {
		return fmt.Errorf("error while marshaling Padding4:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
	}
	// Serialize `TokenMintXProgramFlag`:
	if err = encoder.Encode(obj.TokenMintXProgramFlag); err != nil {
		return fmt.Errorf("error while marshaling TokenMintXProgramFlag:%w", err)
	}
	// Serialize `TokenMintYProgramFlag`:
	if err = encoder.Encode(obj.TokenMintYProgramFlag); err != nil {
		return fmt.Errorf("error while marshaling TokenMintYProgramFlag:%w", err)
	}
	// Serialize `Reserved`:
	if err = encoder.Encode(obj.Reserved); err != nil {
		return fmt.Errorf("error while marshaling Reserved:%w", err)
	}
	return nil
}

func (obj LbPair) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding LbPair: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *LbPair) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_LbPair[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_LbPair[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Parameters`:
	if err = decoder.Decode(&obj.Parameters); err != nil {
		return fmt.Errorf("error while unmarshaling Parameters:%w", err)
	}
	// Deserialize `VParameters`:
	if err = decoder.Decode(&obj.VParameters); err != nil {
		return fmt.Errorf("error while unmarshaling VParameters:%w", err)
	}
	// Deserialize `BumpSeed`:
	if err = decoder.Decode(&obj.BumpSeed); err != nil {
		return fmt.Errorf("error while unmarshaling BumpSeed:%w", err)
	}
	// Deserialize `BinStepSeed`:
	if err = decoder.Decode(&obj.BinStepSeed); err != nil {
		return fmt.Errorf("error while unmarshaling BinStepSeed:%w", err)
	}
	// Deserialize `PairType`:
	if err = decoder.Decode(&obj.PairType); err != nil {
		return fmt.Errorf("error while unmarshaling PairType:%w", err)
	}
	// Deserialize `ActiveId`:
	if err = decoder.Decode(&obj.ActiveId); err != nil {
		return fmt.Errorf("error while unmarshaling ActiveId:%w", err)
	}
	// Deserialize `BinStep`:
	if err = decoder.Decode(&obj.BinStep); err != nil {
		return fmt.Errorf("error while unmarshaling BinStep:%w", err)
	}
	// Deserialize `Status`:
	if err = decoder.Decode(&obj.Status); err != nil {
		return fmt.Errorf("error while unmarshaling Status:%w", err)
	}
	// Deserialize `RequireBaseFactorSeed`:
	if err = decoder.Decode(&obj.RequireBaseFactorSeed); err != nil {
		return fmt.Errorf("error while unmarshaling RequireBaseFactorSeed:%w", err)
	}
	// Deserialize `BaseFactorSeed`:
	if err = decoder.Decode(&obj.BaseFactorSeed); err != nil {
		return fmt.Errorf("error while unmarshaling BaseFactorSeed:%w", err)
	}
	// Deserialize `ActivationType`:
	if err = decoder.Decode(&obj.ActivationType); err != nil {
		return fmt.Errorf("error while unmarshaling ActivationType:%w", err)
	}
	// Deserialize `CreatorPoolOnOffControl`:
	if err = decoder.Decode(&obj.CreatorPoolOnOffControl); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorPoolOnOffControl:%w", err)
	}
	// Deserialize `TokenXMint`:
	if err = decoder.Decode(&obj.TokenXMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenXMint:%w", err)
	}
	// Deserialize `TokenYMint`:
	if err = decoder.Decode(&obj.TokenYMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenYMint:%w", err)
	}
	// Deserialize `ReserveX`:
	if err = decoder.Decode(&obj.ReserveX); err != nil {
		return fmt.Errorf("error while unmarshaling ReserveX:%w", err)
	}
	// Deserialize `ReserveY`:
	if err = decoder.Decode(&obj.ReserveY); err != nil {
		return fmt.Errorf("error while unmarshaling ReserveY:%w", err)
	}
	// Deserialize `ProtocolFee`:
	if err = decoder.Decode(&obj.ProtocolFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFee:%w", err)
	}
	// Deserialize `Padding1`:
	if err = decoder.Decode(&obj.Padding1); err != nil {
		return fmt.Errorf("error while unmarshaling Padding1:%w", err)
	}
	// Deserialize `RewardInfos`:
	if err = decoder.Decode(&obj.RewardInfos); err != nil {
		return fmt.Errorf("error while unmarshaling RewardInfos:%w", err)
	}
	// Deserialize `Oracle`:
	if err = decoder.Decode(&obj.Oracle); err != nil {
		return fmt.Errorf("error while unmarshaling Oracle:%w", err)
	}
	// Deserialize `BinArrayBitmap`:
	if err = decoder.Decode(&obj.BinArrayBitmap); err != nil {
		return fmt.Errorf("error while unmarshaling BinArrayBitmap:%w", err)
	}
	// Deserialize `LastUpdatedAt`:
	if err = decoder.Decode(&obj.LastUpdatedAt); err != nil {
		return fmt.Errorf("error while unmarshaling LastUpdatedAt:%w", err)
	}
	// Deserialize `Padding2`:
	if err = decoder.Decode(&obj.Padding2); err != nil {
		return fmt.Errorf("error while unmarshaling Padding2:%w", err)
	}
	// Deserialize `PreActivationSwapAddress`:
	if err = decoder.Decode(&obj.PreActivationSwapAddress); err != nil {
		return fmt.Errorf("error while unmarshaling PreActivationSwapAddress:%w", err)
	}
	// Deserialize `BaseKey`:
	if err = decoder.Decode(&obj.BaseKey); err != nil {
		return fmt.Errorf("error while unmarshaling BaseKey:%w", err)
	}
	// Deserialize `ActivationPoint`:
	if err = decoder.Decode(&obj.ActivationPoint); err != nil {
		return fmt.Errorf("error while unmarshaling ActivationPoint:%w", err)
	}
	// Deserialize `PreActivationDuration`:
	if err = decoder.Decode(&obj.PreActivationDuration); err != nil {
		return fmt.Errorf("error while unmarshaling PreActivationDuration:%w", err)
	}
	// Deserialize `Padding3`:
	if err = decoder.Decode(&obj.Padding3); err != nil {
		return fmt.Errorf("error while unmarshaling Padding3:%w", err)
	}
	// Deserialize `Padding4`:
	if err = decoder.Decode(&obj.Padding4); err != nil {
		return fmt.Errorf("error while unmarshaling Padding4:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
	}
	// Deserialize `TokenMintXProgramFlag`:
	if err = decoder.Decode(&obj.TokenMintXProgramFlag); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMintXProgramFlag:%w", err)
	}
	// Deserialize `TokenMintYProgramFlag`:
	if err = decoder.Decode(&obj.TokenMintYProgramFlag); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMintYProgramFlag:%w", err)
	}
	// Deserialize `Reserved`:
	if err = decoder.Decode(&obj.Reserved); err != nil {
		return fmt.Errorf("error while unmarshaling Reserved:%w", err)
	}
	return nil
}

func (obj *LbPair) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling LbPair: %w", err)
	}
	return nil
}

func UnmarshalLbPair(buf []byte) (*LbPair, error) {
	obj := new(LbPair)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type Oracle struct {
	// Index of latest observation
	Idx uint64

	// Size of active sample. Active sample is initialized observation.
	ActiveSize uint64

	// Number of observations
	Length uint64
}

func (obj Oracle) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Oracle[:], false)
	if err != nil {
		return err
	}
	// Serialize `Idx`:
	if err = encoder.Encode(obj.Idx); err != nil {
		return fmt.Errorf("error while marshaling Idx:%w", err)
	}
	// Serialize `ActiveSize`:
	if err = encoder.Encode(obj.ActiveSize); err != nil {
		return fmt.Errorf("error while marshaling ActiveSize:%w", err)
	}
	// Serialize `Length`:
	if err = encoder.Encode(obj.Length); err != nil {
		return fmt.Errorf("error while marshaling Length:%w", err)
	}
	return nil
}

func (obj Oracle) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Oracle: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Oracle) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Oracle[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Oracle[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Idx`:
	if err = decoder.Decode(&obj.Idx); err != nil {
		return fmt.Errorf("error while unmarshaling Idx:%w", err)
	}
	// Deserialize `ActiveSize`:
	if err = decoder.Decode(&obj.ActiveSize); err != nil {
		return fmt.Errorf("error while unmarshaling ActiveSize:%w", err)
	}
	// Deserialize `Length`:
	if err = decoder.Decode(&obj.Length); err != nil {
		return fmt.Errorf("error while unmarshaling Length:%w", err)
	}
	return nil
}

func (obj *Oracle) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Oracle: %w", err)
	}
	return nil
}

func UnmarshalOracle(buf []byte) (*Oracle, error) {
	obj := new(Oracle)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type Position struct {
	// The LB pair of this position
	LbPair solanago.PublicKey

	// Owner of the position. Client rely on this to to fetch their positions.
	Owner solanago.PublicKey

	// Liquidity shares of this position in bins (lower_bin_id <-> upper_bin_id). This is the same as LP concept.
	LiquidityShares [70]uint64

	// Farming reward information
	RewardInfos [70]UserRewardInfo

	// Swap fee to claim information
	FeeInfos [70]FeeInfo

	// Lower bin ID
	LowerBinId int32

	// Upper bin ID
	UpperBinId int32

	// Last updated timestamp
	LastUpdatedAt int64

	// Total claimed token fee X
	TotalClaimedFeeXAmount uint64

	// Total claimed token fee Y
	TotalClaimedFeeYAmount uint64

	// Total claimed rewards
	TotalClaimedRewards [2]uint64

	// Reserved space for future use
	Reserved [160]uint8
}

func (obj Position) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Position[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `LiquidityShares`:
	if err = encoder.Encode(obj.LiquidityShares); err != nil {
		return fmt.Errorf("error while marshaling LiquidityShares:%w", err)
	}
	// Serialize `RewardInfos`:
	if err = encoder.Encode(obj.RewardInfos); err != nil {
		return fmt.Errorf("error while marshaling RewardInfos:%w", err)
	}
	// Serialize `FeeInfos`:
	if err = encoder.Encode(obj.FeeInfos); err != nil {
		return fmt.Errorf("error while marshaling FeeInfos:%w", err)
	}
	// Serialize `LowerBinId`:
	if err = encoder.Encode(obj.LowerBinId); err != nil {
		return fmt.Errorf("error while marshaling LowerBinId:%w", err)
	}
	// Serialize `UpperBinId`:
	if err = encoder.Encode(obj.UpperBinId); err != nil {
		return fmt.Errorf("error while marshaling UpperBinId:%w", err)
	}
	// Serialize `LastUpdatedAt`:
	if err = encoder.Encode(obj.LastUpdatedAt); err != nil {
		return fmt.Errorf("error while marshaling LastUpdatedAt:%w", err)
	}
	// Serialize `TotalClaimedFeeXAmount`:
	if err = encoder.Encode(obj.TotalClaimedFeeXAmount); err != nil {
		return fmt.Errorf("error while marshaling TotalClaimedFeeXAmount:%w", err)
	}
	// Serialize `TotalClaimedFeeYAmount`:
	if err = encoder.Encode(obj.TotalClaimedFeeYAmount); err != nil {
		return fmt.Errorf("error while marshaling TotalClaimedFeeYAmount:%w", err)
	}
	// Serialize `TotalClaimedRewards`:
	if err = encoder.Encode(obj.TotalClaimedRewards); err != nil {
		return fmt.Errorf("error while marshaling TotalClaimedRewards:%w", err)
	}
	// Serialize `Reserved`:
	if err = encoder.Encode(obj.Reserved); err != nil {
		return fmt.Errorf("error while marshaling Reserved:%w", err)
	}
	return nil
}

func (obj Position) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Position: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Position) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Position[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Position[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `LiquidityShares`:
	if err = decoder.Decode(&obj.LiquidityShares); err != nil {
		return fmt.Errorf("error while unmarshaling LiquidityShares:%w", err)
	}
	// Deserialize `RewardInfos`:
	if err = decoder.Decode(&obj.RewardInfos); err != nil {
		return fmt.Errorf("error while unmarshaling RewardInfos:%w", err)
	}
	// Deserialize `FeeInfos`:
	if err = decoder.Decode(&obj.FeeInfos); err != nil {
		return fmt.Errorf("error while unmarshaling FeeInfos:%w", err)
	}
	// Deserialize `LowerBinId`:
	if err = decoder.Decode(&obj.LowerBinId); err != nil {
		return fmt.Errorf("error while unmarshaling LowerBinId:%w", err)
	}
	// Deserialize `UpperBinId`:
	if err = decoder.Decode(&obj.UpperBinId); err != nil {
		return fmt.Errorf("error while unmarshaling UpperBinId:%w", err)
	}
	// Deserialize `LastUpdatedAt`:
	if err = decoder.Decode(&obj.LastUpdatedAt); err != nil {
		return fmt.Errorf("error while unmarshaling LastUpdatedAt:%w", err)
	}
	// Deserialize `TotalClaimedFeeXAmount`:
	if err = decoder.Decode(&obj.TotalClaimedFeeXAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TotalClaimedFeeXAmount:%w", err)
	}
	// Deserialize `TotalClaimedFeeYAmount`:
	if err = decoder.Decode(&obj.TotalClaimedFeeYAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TotalClaimedFeeYAmount:%w", err)
	}
	// Deserialize `TotalClaimedRewards`:
	if err = decoder.Decode(&obj.TotalClaimedRewards); err != nil {
		return fmt.Errorf("error while unmarshaling TotalClaimedRewards:%w", err)
	}
	// Deserialize `Reserved`:
	if err = decoder.Decode(&obj.Reserved); err != nil {
		return fmt.Errorf("error while unmarshaling Reserved:%w", err)
	}
	return nil
}

func (obj *Position) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Position: %w", err)
	}
	return nil
}

func UnmarshalPosition(buf []byte) (*Position, error) {
	obj := new(Position)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type PositionV2 struct {
	// The LB pair of this position
	LbPair solanago.PublicKey

	// Owner of the position. Client rely on this to to fetch their positions.
	Owner solanago.PublicKey

	// Liquidity shares of this position in bins (lower_bin_id <-> upper_bin_id). This is the same as LP concept.
	LiquidityShares [70]binary.Uint128

	// Farming reward information
	RewardInfos [70]UserRewardInfo

	// Swap fee to claim information
	FeeInfos [70]FeeInfo

	// Lower bin ID
	LowerBinId int32

	// Upper bin ID
	UpperBinId int32

	// Last updated timestamp
	LastUpdatedAt int64

	// Total claimed token fee X
	TotalClaimedFeeXAmount uint64

	// Total claimed token fee Y
	TotalClaimedFeeYAmount uint64

	// Total claimed rewards
	TotalClaimedRewards [2]uint64

	// Operator of position
	Operator solanago.PublicKey

	// Time point which the locked liquidity can be withdraw
	LockReleasePoint uint64

	// _padding_0, previous subjected_to_bootstrap_liquidity_locking, BE CAREFUL FOR TOMBSTONE WHEN REUSE !!
	Padding0 uint8

	// Address is able to claim fee in this position, only valid for bootstrap_liquidity_position
	FeeOwner solanago.PublicKey

	// Reserved space for future use
	Reserved [87]uint8
}

func (obj PositionV2) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_PositionV2[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `LiquidityShares`:
	if err = encoder.Encode(obj.LiquidityShares); err != nil {
		return fmt.Errorf("error while marshaling LiquidityShares:%w", err)
	}
	// Serialize `RewardInfos`:
	if err = encoder.Encode(obj.RewardInfos); err != nil {
		return fmt.Errorf("error while marshaling RewardInfos:%w", err)
	}
	// Serialize `FeeInfos`:
	if err = encoder.Encode(obj.FeeInfos); err != nil {
		return fmt.Errorf("error while marshaling FeeInfos:%w", err)
	}
	// Serialize `LowerBinId`:
	if err = encoder.Encode(obj.LowerBinId); err != nil {
		return fmt.Errorf("error while marshaling LowerBinId:%w", err)
	}
	// Serialize `UpperBinId`:
	if err = encoder.Encode(obj.UpperBinId); err != nil {
		return fmt.Errorf("error while marshaling UpperBinId:%w", err)
	}
	// Serialize `LastUpdatedAt`:
	if err = encoder.Encode(obj.LastUpdatedAt); err != nil {
		return fmt.Errorf("error while marshaling LastUpdatedAt:%w", err)
	}
	// Serialize `TotalClaimedFeeXAmount`:
	if err = encoder.Encode(obj.TotalClaimedFeeXAmount); err != nil {
		return fmt.Errorf("error while marshaling TotalClaimedFeeXAmount:%w", err)
	}
	// Serialize `TotalClaimedFeeYAmount`:
	if err = encoder.Encode(obj.TotalClaimedFeeYAmount); err != nil {
		return fmt.Errorf("error while marshaling TotalClaimedFeeYAmount:%w", err)
	}
	// Serialize `TotalClaimedRewards`:
	if err = encoder.Encode(obj.TotalClaimedRewards); err != nil {
		return fmt.Errorf("error while marshaling TotalClaimedRewards:%w", err)
	}
	// Serialize `Operator`:
	if err = encoder.Encode(obj.Operator); err != nil {
		return fmt.Errorf("error while marshaling Operator:%w", err)
	}
	// Serialize `LockReleasePoint`:
	if err = encoder.Encode(obj.LockReleasePoint); err != nil {
		return fmt.Errorf("error while marshaling LockReleasePoint:%w", err)
	}
	// Serialize `Padding0`:
	if err = encoder.Encode(obj.Padding0); err != nil {
		return fmt.Errorf("error while marshaling Padding0:%w", err)
	}
	// Serialize `FeeOwner`:
	if err = encoder.Encode(obj.FeeOwner); err != nil {
		return fmt.Errorf("error while marshaling FeeOwner:%w", err)
	}
	// Serialize `Reserved`:
	if err = encoder.Encode(obj.Reserved); err != nil {
		return fmt.Errorf("error while marshaling Reserved:%w", err)
	}
	return nil
}

func (obj PositionV2) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding PositionV2: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *PositionV2) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_PositionV2[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_PositionV2[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `LiquidityShares`:
	if err = decoder.Decode(&obj.LiquidityShares); err != nil {
		return fmt.Errorf("error while unmarshaling LiquidityShares:%w", err)
	}
	// Deserialize `RewardInfos`:
	if err = decoder.Decode(&obj.RewardInfos); err != nil {
		return fmt.Errorf("error while unmarshaling RewardInfos:%w", err)
	}
	// Deserialize `FeeInfos`:
	if err = decoder.Decode(&obj.FeeInfos); err != nil {
		return fmt.Errorf("error while unmarshaling FeeInfos:%w", err)
	}
	// Deserialize `LowerBinId`:
	if err = decoder.Decode(&obj.LowerBinId); err != nil {
		return fmt.Errorf("error while unmarshaling LowerBinId:%w", err)
	}
	// Deserialize `UpperBinId`:
	if err = decoder.Decode(&obj.UpperBinId); err != nil {
		return fmt.Errorf("error while unmarshaling UpperBinId:%w", err)
	}
	// Deserialize `LastUpdatedAt`:
	if err = decoder.Decode(&obj.LastUpdatedAt); err != nil {
		return fmt.Errorf("error while unmarshaling LastUpdatedAt:%w", err)
	}
	// Deserialize `TotalClaimedFeeXAmount`:
	if err = decoder.Decode(&obj.TotalClaimedFeeXAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TotalClaimedFeeXAmount:%w", err)
	}
	// Deserialize `TotalClaimedFeeYAmount`:
	if err = decoder.Decode(&obj.TotalClaimedFeeYAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TotalClaimedFeeYAmount:%w", err)
	}
	// Deserialize `TotalClaimedRewards`:
	if err = decoder.Decode(&obj.TotalClaimedRewards); err != nil {
		return fmt.Errorf("error while unmarshaling TotalClaimedRewards:%w", err)
	}
	// Deserialize `Operator`:
	if err = decoder.Decode(&obj.Operator); err != nil {
		return fmt.Errorf("error while unmarshaling Operator:%w", err)
	}
	// Deserialize `LockReleasePoint`:
	if err = decoder.Decode(&obj.LockReleasePoint); err != nil {
		return fmt.Errorf("error while unmarshaling LockReleasePoint:%w", err)
	}
	// Deserialize `Padding0`:
	if err = decoder.Decode(&obj.Padding0); err != nil {
		return fmt.Errorf("error while unmarshaling Padding0:%w", err)
	}
	// Deserialize `FeeOwner`:
	if err = decoder.Decode(&obj.FeeOwner); err != nil {
		return fmt.Errorf("error while unmarshaling FeeOwner:%w", err)
	}
	// Deserialize `Reserved`:
	if err = decoder.Decode(&obj.Reserved); err != nil {
		return fmt.Errorf("error while unmarshaling Reserved:%w", err)
	}
	return nil
}

func (obj *PositionV2) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling PositionV2: %w", err)
	}
	return nil
}

func UnmarshalPositionV2(buf []byte) (*PositionV2, error) {
	obj := new(PositionV2)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type PresetParameter struct {
	// Bin step. Represent the price increment / decrement.
	BinStep uint16

	// Used for base fee calculation. base_fee_rate = base_factor * bin_step * 10 * 10^base_fee_power_factor
	BaseFactor uint16

	// Filter period determine high frequency trading time window.
	FilterPeriod uint16

	// Decay period determine when the volatile fee start decay / decrease.
	DecayPeriod uint16

	// Reduction factor controls the volatile fee rate decrement rate.
	ReductionFactor uint16

	// Used to scale the variable fee component depending on the dynamic of the market
	VariableFeeControl uint32

	// Maximum number of bin crossed can be accumulated. Used to cap volatile fee rate.
	MaxVolatilityAccumulator uint32

	// Min bin id supported by the pool based on the configured bin step.
	MinBinId int32

	// Max bin id supported by the pool based on the configured bin step.
	MaxBinId int32

	// Portion of swap fees retained by the protocol by controlling protocol_share parameter. protocol_swap_fee = protocol_share * total_swap_fee
	ProtocolShare uint16
}

func (obj PresetParameter) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_PresetParameter[:], false)
	if err != nil {
		return err
	}
	// Serialize `BinStep`:
	if err = encoder.Encode(obj.BinStep); err != nil {
		return fmt.Errorf("error while marshaling BinStep:%w", err)
	}
	// Serialize `BaseFactor`:
	if err = encoder.Encode(obj.BaseFactor); err != nil {
		return fmt.Errorf("error while marshaling BaseFactor:%w", err)
	}
	// Serialize `FilterPeriod`:
	if err = encoder.Encode(obj.FilterPeriod); err != nil {
		return fmt.Errorf("error while marshaling FilterPeriod:%w", err)
	}
	// Serialize `DecayPeriod`:
	if err = encoder.Encode(obj.DecayPeriod); err != nil {
		return fmt.Errorf("error while marshaling DecayPeriod:%w", err)
	}
	// Serialize `ReductionFactor`:
	if err = encoder.Encode(obj.ReductionFactor); err != nil {
		return fmt.Errorf("error while marshaling ReductionFactor:%w", err)
	}
	// Serialize `VariableFeeControl`:
	if err = encoder.Encode(obj.VariableFeeControl); err != nil {
		return fmt.Errorf("error while marshaling VariableFeeControl:%w", err)
	}
	// Serialize `MaxVolatilityAccumulator`:
	if err = encoder.Encode(obj.MaxVolatilityAccumulator); err != nil {
		return fmt.Errorf("error while marshaling MaxVolatilityAccumulator:%w", err)
	}
	// Serialize `MinBinId`:
	if err = encoder.Encode(obj.MinBinId); err != nil {
		return fmt.Errorf("error while marshaling MinBinId:%w", err)
	}
	// Serialize `MaxBinId`:
	if err = encoder.Encode(obj.MaxBinId); err != nil {
		return fmt.Errorf("error while marshaling MaxBinId:%w", err)
	}
	// Serialize `ProtocolShare`:
	if err = encoder.Encode(obj.ProtocolShare); err != nil {
		return fmt.Errorf("error while marshaling ProtocolShare:%w", err)
	}
	return nil
}

func (obj PresetParameter) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding PresetParameter: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *PresetParameter) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_PresetParameter[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_PresetParameter[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `BinStep`:
	if err = decoder.Decode(&obj.BinStep); err != nil {
		return fmt.Errorf("error while unmarshaling BinStep:%w", err)
	}
	// Deserialize `BaseFactor`:
	if err = decoder.Decode(&obj.BaseFactor); err != nil {
		return fmt.Errorf("error while unmarshaling BaseFactor:%w", err)
	}
	// Deserialize `FilterPeriod`:
	if err = decoder.Decode(&obj.FilterPeriod); err != nil {
		return fmt.Errorf("error while unmarshaling FilterPeriod:%w", err)
	}
	// Deserialize `DecayPeriod`:
	if err = decoder.Decode(&obj.DecayPeriod); err != nil {
		return fmt.Errorf("error while unmarshaling DecayPeriod:%w", err)
	}
	// Deserialize `ReductionFactor`:
	if err = decoder.Decode(&obj.ReductionFactor); err != nil {
		return fmt.Errorf("error while unmarshaling ReductionFactor:%w", err)
	}
	// Deserialize `VariableFeeControl`:
	if err = decoder.Decode(&obj.VariableFeeControl); err != nil {
		return fmt.Errorf("error while unmarshaling VariableFeeControl:%w", err)
	}
	// Deserialize `MaxVolatilityAccumulator`:
	if err = decoder.Decode(&obj.MaxVolatilityAccumulator); err != nil {
		return fmt.Errorf("error while unmarshaling MaxVolatilityAccumulator:%w", err)
	}
	// Deserialize `MinBinId`:
	if err = decoder.Decode(&obj.MinBinId); err != nil {
		return fmt.Errorf("error while unmarshaling MinBinId:%w", err)
	}
	// Deserialize `MaxBinId`:
	if err = decoder.Decode(&obj.MaxBinId); err != nil {
		return fmt.Errorf("error while unmarshaling MaxBinId:%w", err)
	}
	// Deserialize `ProtocolShare`:
	if err = decoder.Decode(&obj.ProtocolShare); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolShare:%w", err)
	}
	return nil
}

func (obj *PresetParameter) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling PresetParameter: %w", err)
	}
	return nil
}

func UnmarshalPresetParameter(buf []byte) (*PresetParameter, error) {
	obj := new(PresetParameter)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type PresetParameter2 struct {
	// Bin step. Represent the price increment / decrement.
	BinStep uint16

	// Used for base fee calculation. base_fee_rate = base_factor * bin_step * 10 * 10^base_fee_power_factor
	BaseFactor uint16

	// Filter period determine high frequency trading time window.
	FilterPeriod uint16

	// Decay period determine when the volatile fee start decay / decrease.
	DecayPeriod uint16

	// Used to scale the variable fee component depending on the dynamic of the market
	VariableFeeControl uint32

	// Maximum number of bin crossed can be accumulated. Used to cap volatile fee rate.
	MaxVolatilityAccumulator uint32

	// Reduction factor controls the volatile fee rate decrement rate.
	ReductionFactor uint16

	// Portion of swap fees retained by the protocol by controlling protocol_share parameter. protocol_swap_fee = protocol_share * total_swap_fee
	ProtocolShare uint16

	// index
	Index uint16

	// Base fee power factor
	BaseFeePowerFactor uint8

	// Padding 0 for future use
	Padding0 uint8

	// Padding 1 for future use
	Padding1 [20]uint64
}

func (obj PresetParameter2) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_PresetParameter2[:], false)
	if err != nil {
		return err
	}
	// Serialize `BinStep`:
	if err = encoder.Encode(obj.BinStep); err != nil {
		return fmt.Errorf("error while marshaling BinStep:%w", err)
	}
	// Serialize `BaseFactor`:
	if err = encoder.Encode(obj.BaseFactor); err != nil {
		return fmt.Errorf("error while marshaling BaseFactor:%w", err)
	}
	// Serialize `FilterPeriod`:
	if err = encoder.Encode(obj.FilterPeriod); err != nil {
		return fmt.Errorf("error while marshaling FilterPeriod:%w", err)
	}
	// Serialize `DecayPeriod`:
	if err = encoder.Encode(obj.DecayPeriod); err != nil {
		return fmt.Errorf("error while marshaling DecayPeriod:%w", err)
	}
	// Serialize `VariableFeeControl`:
	if err = encoder.Encode(obj.VariableFeeControl); err != nil {
		return fmt.Errorf("error while marshaling VariableFeeControl:%w", err)
	}
	// Serialize `MaxVolatilityAccumulator`:
	if err = encoder.Encode(obj.MaxVolatilityAccumulator); err != nil {
		return fmt.Errorf("error while marshaling MaxVolatilityAccumulator:%w", err)
	}
	// Serialize `ReductionFactor`:
	if err = encoder.Encode(obj.ReductionFactor); err != nil {
		return fmt.Errorf("error while marshaling ReductionFactor:%w", err)
	}
	// Serialize `ProtocolShare`:
	if err = encoder.Encode(obj.ProtocolShare); err != nil {
		return fmt.Errorf("error while marshaling ProtocolShare:%w", err)
	}
	// Serialize `Index`:
	if err = encoder.Encode(obj.Index); err != nil {
		return fmt.Errorf("error while marshaling Index:%w", err)
	}
	// Serialize `BaseFeePowerFactor`:
	if err = encoder.Encode(obj.BaseFeePowerFactor); err != nil {
		return fmt.Errorf("error while marshaling BaseFeePowerFactor:%w", err)
	}
	// Serialize `Padding0`:
	if err = encoder.Encode(obj.Padding0); err != nil {
		return fmt.Errorf("error while marshaling Padding0:%w", err)
	}
	// Serialize `Padding1`:
	if err = encoder.Encode(obj.Padding1); err != nil {
		return fmt.Errorf("error while marshaling Padding1:%w", err)
	}
	return nil
}

func (obj PresetParameter2) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding PresetParameter2: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *PresetParameter2) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_PresetParameter2[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_PresetParameter2[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `BinStep`:
	if err = decoder.Decode(&obj.BinStep); err != nil {
		return fmt.Errorf("error while unmarshaling BinStep:%w", err)
	}
	// Deserialize `BaseFactor`:
	if err = decoder.Decode(&obj.BaseFactor); err != nil {
		return fmt.Errorf("error while unmarshaling BaseFactor:%w", err)
	}
	// Deserialize `FilterPeriod`:
	if err = decoder.Decode(&obj.FilterPeriod); err != nil {
		return fmt.Errorf("error while unmarshaling FilterPeriod:%w", err)
	}
	// Deserialize `DecayPeriod`:
	if err = decoder.Decode(&obj.DecayPeriod); err != nil {
		return fmt.Errorf("error while unmarshaling DecayPeriod:%w", err)
	}
	// Deserialize `VariableFeeControl`:
	if err = decoder.Decode(&obj.VariableFeeControl); err != nil {
		return fmt.Errorf("error while unmarshaling VariableFeeControl:%w", err)
	}
	// Deserialize `MaxVolatilityAccumulator`:
	if err = decoder.Decode(&obj.MaxVolatilityAccumulator); err != nil {
		return fmt.Errorf("error while unmarshaling MaxVolatilityAccumulator:%w", err)
	}
	// Deserialize `ReductionFactor`:
	if err = decoder.Decode(&obj.ReductionFactor); err != nil {
		return fmt.Errorf("error while unmarshaling ReductionFactor:%w", err)
	}
	// Deserialize `ProtocolShare`:
	if err = decoder.Decode(&obj.ProtocolShare); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolShare:%w", err)
	}
	// Deserialize `Index`:
	if err = decoder.Decode(&obj.Index); err != nil {
		return fmt.Errorf("error while unmarshaling Index:%w", err)
	}
	// Deserialize `BaseFeePowerFactor`:
	if err = decoder.Decode(&obj.BaseFeePowerFactor); err != nil {
		return fmt.Errorf("error while unmarshaling BaseFeePowerFactor:%w", err)
	}
	// Deserialize `Padding0`:
	if err = decoder.Decode(&obj.Padding0); err != nil {
		return fmt.Errorf("error while unmarshaling Padding0:%w", err)
	}
	// Deserialize `Padding1`:
	if err = decoder.Decode(&obj.Padding1); err != nil {
		return fmt.Errorf("error while unmarshaling Padding1:%w", err)
	}
	return nil
}

func (obj *PresetParameter2) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling PresetParameter2: %w", err)
	}
	return nil
}

func UnmarshalPresetParameter2(buf []byte) (*PresetParameter2, error) {
	obj := new(PresetParameter2)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Parameter that set by the protocol
type TokenBadge struct {
	// token mint
	TokenMint solanago.PublicKey

	// Reserve
	Padding [128]uint8
}

func (obj TokenBadge) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_TokenBadge[:], false)
	if err != nil {
		return err
	}
	// Serialize `TokenMint`:
	if err = encoder.Encode(obj.TokenMint); err != nil {
		return fmt.Errorf("error while marshaling TokenMint:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj TokenBadge) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding TokenBadge: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *TokenBadge) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_TokenBadge[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_TokenBadge[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TokenMint`:
	if err = decoder.Decode(&obj.TokenMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMint:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *TokenBadge) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling TokenBadge: %w", err)
	}
	return nil
}

func UnmarshalTokenBadge(buf []byte) (*TokenBadge, error) {
	obj := new(TokenBadge)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
