package orca_whirlpool

import (
	"bytes"
	"fmt"

	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type AdaptiveFeeTier struct {
	WhirlpoolsConfig         solanago.PublicKey
	FeeTierIndex             uint16
	TickSpacing              uint16
	InitializePoolAuthority  solanago.PublicKey
	DelegatedFeeAuthority    solanago.PublicKey
	DefaultBaseFeeRate       uint16
	FilterPeriod             uint16
	DecayPeriod              uint16
	ReductionFactor          uint16
	AdaptiveFeeControlFactor uint32
	MaxVolatilityAccumulator uint32
	TickGroupSize            uint16
	MajorSwapThresholdTicks  uint16
}

func (obj AdaptiveFeeTier) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_AdaptiveFeeTier[:], false)
	if err != nil {
		return err
	}
	// Serialize `WhirlpoolsConfig`:
	if err = encoder.Encode(obj.WhirlpoolsConfig); err != nil {
		return fmt.Errorf("error while marshaling WhirlpoolsConfig:%w", err)
	}
	// Serialize `FeeTierIndex`:
	if err = encoder.Encode(obj.FeeTierIndex); err != nil {
		return fmt.Errorf("error while marshaling FeeTierIndex:%w", err)
	}
	// Serialize `TickSpacing`:
	if err = encoder.Encode(obj.TickSpacing); err != nil {
		return fmt.Errorf("error while marshaling TickSpacing:%w", err)
	}
	// Serialize `InitializePoolAuthority`:
	if err = encoder.Encode(obj.InitializePoolAuthority); err != nil {
		return fmt.Errorf("error while marshaling InitializePoolAuthority:%w", err)
	}
	// Serialize `DelegatedFeeAuthority`:
	if err = encoder.Encode(obj.DelegatedFeeAuthority); err != nil {
		return fmt.Errorf("error while marshaling DelegatedFeeAuthority:%w", err)
	}
	// Serialize `DefaultBaseFeeRate`:
	if err = encoder.Encode(obj.DefaultBaseFeeRate); err != nil {
		return fmt.Errorf("error while marshaling DefaultBaseFeeRate:%w", err)
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
	// Serialize `AdaptiveFeeControlFactor`:
	if err = encoder.Encode(obj.AdaptiveFeeControlFactor); err != nil {
		return fmt.Errorf("error while marshaling AdaptiveFeeControlFactor:%w", err)
	}
	// Serialize `MaxVolatilityAccumulator`:
	if err = encoder.Encode(obj.MaxVolatilityAccumulator); err != nil {
		return fmt.Errorf("error while marshaling MaxVolatilityAccumulator:%w", err)
	}
	// Serialize `TickGroupSize`:
	if err = encoder.Encode(obj.TickGroupSize); err != nil {
		return fmt.Errorf("error while marshaling TickGroupSize:%w", err)
	}
	// Serialize `MajorSwapThresholdTicks`:
	if err = encoder.Encode(obj.MajorSwapThresholdTicks); err != nil {
		return fmt.Errorf("error while marshaling MajorSwapThresholdTicks:%w", err)
	}
	return nil
}

func (obj AdaptiveFeeTier) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding AdaptiveFeeTier: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *AdaptiveFeeTier) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_AdaptiveFeeTier[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_AdaptiveFeeTier[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `WhirlpoolsConfig`:
	if err = decoder.Decode(&obj.WhirlpoolsConfig); err != nil {
		return fmt.Errorf("error while unmarshaling WhirlpoolsConfig:%w", err)
	}
	// Deserialize `FeeTierIndex`:
	if err = decoder.Decode(&obj.FeeTierIndex); err != nil {
		return fmt.Errorf("error while unmarshaling FeeTierIndex:%w", err)
	}
	// Deserialize `TickSpacing`:
	if err = decoder.Decode(&obj.TickSpacing); err != nil {
		return fmt.Errorf("error while unmarshaling TickSpacing:%w", err)
	}
	// Deserialize `InitializePoolAuthority`:
	if err = decoder.Decode(&obj.InitializePoolAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling InitializePoolAuthority:%w", err)
	}
	// Deserialize `DelegatedFeeAuthority`:
	if err = decoder.Decode(&obj.DelegatedFeeAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling DelegatedFeeAuthority:%w", err)
	}
	// Deserialize `DefaultBaseFeeRate`:
	if err = decoder.Decode(&obj.DefaultBaseFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling DefaultBaseFeeRate:%w", err)
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
	// Deserialize `AdaptiveFeeControlFactor`:
	if err = decoder.Decode(&obj.AdaptiveFeeControlFactor); err != nil {
		return fmt.Errorf("error while unmarshaling AdaptiveFeeControlFactor:%w", err)
	}
	// Deserialize `MaxVolatilityAccumulator`:
	if err = decoder.Decode(&obj.MaxVolatilityAccumulator); err != nil {
		return fmt.Errorf("error while unmarshaling MaxVolatilityAccumulator:%w", err)
	}
	// Deserialize `TickGroupSize`:
	if err = decoder.Decode(&obj.TickGroupSize); err != nil {
		return fmt.Errorf("error while unmarshaling TickGroupSize:%w", err)
	}
	// Deserialize `MajorSwapThresholdTicks`:
	if err = decoder.Decode(&obj.MajorSwapThresholdTicks); err != nil {
		return fmt.Errorf("error while unmarshaling MajorSwapThresholdTicks:%w", err)
	}
	return nil
}

func (obj *AdaptiveFeeTier) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling AdaptiveFeeTier: %w", err)
	}
	return nil
}

func UnmarshalAdaptiveFeeTier(buf []byte) (*AdaptiveFeeTier, error) {
	obj := new(AdaptiveFeeTier)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type WhirlpoolsConfig struct {
	FeeAuthority                  solanago.PublicKey
	CollectProtocolFeesAuthority  solanago.PublicKey
	RewardEmissionsSuperAuthority solanago.PublicKey
	DefaultProtocolFeeRate        uint16
}

func (obj WhirlpoolsConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_WhirlpoolsConfig[:], false)
	if err != nil {
		return err
	}
	// Serialize `FeeAuthority`:
	if err = encoder.Encode(obj.FeeAuthority); err != nil {
		return fmt.Errorf("error while marshaling FeeAuthority:%w", err)
	}
	// Serialize `CollectProtocolFeesAuthority`:
	if err = encoder.Encode(obj.CollectProtocolFeesAuthority); err != nil {
		return fmt.Errorf("error while marshaling CollectProtocolFeesAuthority:%w", err)
	}
	// Serialize `RewardEmissionsSuperAuthority`:
	if err = encoder.Encode(obj.RewardEmissionsSuperAuthority); err != nil {
		return fmt.Errorf("error while marshaling RewardEmissionsSuperAuthority:%w", err)
	}
	// Serialize `DefaultProtocolFeeRate`:
	if err = encoder.Encode(obj.DefaultProtocolFeeRate); err != nil {
		return fmt.Errorf("error while marshaling DefaultProtocolFeeRate:%w", err)
	}
	return nil
}

func (obj WhirlpoolsConfig) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding WhirlpoolsConfig: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *WhirlpoolsConfig) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_WhirlpoolsConfig[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_WhirlpoolsConfig[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `FeeAuthority`:
	if err = decoder.Decode(&obj.FeeAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling FeeAuthority:%w", err)
	}
	// Deserialize `CollectProtocolFeesAuthority`:
	if err = decoder.Decode(&obj.CollectProtocolFeesAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling CollectProtocolFeesAuthority:%w", err)
	}
	// Deserialize `RewardEmissionsSuperAuthority`:
	if err = decoder.Decode(&obj.RewardEmissionsSuperAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling RewardEmissionsSuperAuthority:%w", err)
	}
	// Deserialize `DefaultProtocolFeeRate`:
	if err = decoder.Decode(&obj.DefaultProtocolFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling DefaultProtocolFeeRate:%w", err)
	}
	return nil
}

func (obj *WhirlpoolsConfig) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling WhirlpoolsConfig: %w", err)
	}
	return nil
}

func UnmarshalWhirlpoolsConfig(buf []byte) (*WhirlpoolsConfig, error) {
	obj := new(WhirlpoolsConfig)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type WhirlpoolsConfigExtension struct {
	WhirlpoolsConfig         solanago.PublicKey
	ConfigExtensionAuthority solanago.PublicKey
	TokenBadgeAuthority      solanago.PublicKey
}

func (obj WhirlpoolsConfigExtension) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_WhirlpoolsConfigExtension[:], false)
	if err != nil {
		return err
	}
	// Serialize `WhirlpoolsConfig`:
	if err = encoder.Encode(obj.WhirlpoolsConfig); err != nil {
		return fmt.Errorf("error while marshaling WhirlpoolsConfig:%w", err)
	}
	// Serialize `ConfigExtensionAuthority`:
	if err = encoder.Encode(obj.ConfigExtensionAuthority); err != nil {
		return fmt.Errorf("error while marshaling ConfigExtensionAuthority:%w", err)
	}
	// Serialize `TokenBadgeAuthority`:
	if err = encoder.Encode(obj.TokenBadgeAuthority); err != nil {
		return fmt.Errorf("error while marshaling TokenBadgeAuthority:%w", err)
	}
	return nil
}

func (obj WhirlpoolsConfigExtension) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding WhirlpoolsConfigExtension: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *WhirlpoolsConfigExtension) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_WhirlpoolsConfigExtension[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_WhirlpoolsConfigExtension[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `WhirlpoolsConfig`:
	if err = decoder.Decode(&obj.WhirlpoolsConfig); err != nil {
		return fmt.Errorf("error while unmarshaling WhirlpoolsConfig:%w", err)
	}
	// Deserialize `ConfigExtensionAuthority`:
	if err = decoder.Decode(&obj.ConfigExtensionAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling ConfigExtensionAuthority:%w", err)
	}
	// Deserialize `TokenBadgeAuthority`:
	if err = decoder.Decode(&obj.TokenBadgeAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBadgeAuthority:%w", err)
	}
	return nil
}

func (obj *WhirlpoolsConfigExtension) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling WhirlpoolsConfigExtension: %w", err)
	}
	return nil
}

func UnmarshalWhirlpoolsConfigExtension(buf []byte) (*WhirlpoolsConfigExtension, error) {
	obj := new(WhirlpoolsConfigExtension)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type DynamicTickArray struct {
	StartTickIndex int32
	Whirlpool      solanago.PublicKey
	TickBitmap     binary.Uint128
	Ticks          [88]DynamicTick
}

func (obj DynamicTickArray) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_DynamicTickArray[:], false)
	if err != nil {
		return err
	}
	// Serialize `StartTickIndex`:
	if err = encoder.Encode(obj.StartTickIndex); err != nil {
		return fmt.Errorf("error while marshaling StartTickIndex:%w", err)
	}
	// Serialize `Whirlpool`:
	if err = encoder.Encode(obj.Whirlpool); err != nil {
		return fmt.Errorf("error while marshaling Whirlpool:%w", err)
	}
	// Serialize `TickBitmap`:
	if err = encoder.Encode(obj.TickBitmap); err != nil {
		return fmt.Errorf("error while marshaling TickBitmap:%w", err)
	}
	// Serialize `Ticks`:
	{
		for i := range len(obj.Ticks) {
			// err = EncodeDynamicTick(encoder, obj.Ticks[i])
			if err = EncodeDynamicTick(encoder, obj.Ticks[i]); err != nil {
				return fmt.Errorf("error while marshaling Ticks-%d: %w", i, err)
			}
		}
	}
	return nil
}

func (obj DynamicTickArray) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding DynamicTickArray: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *DynamicTickArray) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_DynamicTickArray[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_DynamicTickArray[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `StartTickIndex`:
	if err = decoder.Decode(&obj.StartTickIndex); err != nil {
		return fmt.Errorf("error while unmarshaling StartTickIndex:%w", err)
	}
	// Deserialize `Whirlpool`:
	if err = decoder.Decode(&obj.Whirlpool); err != nil {
		return fmt.Errorf("error while unmarshaling Whirlpool:%w", err)
	}
	// Deserialize `TickBitmap`:
	if err = decoder.Decode(&obj.TickBitmap); err != nil {
		return fmt.Errorf("error while unmarshaling TickBitmap:%w", err)
	}
	// Deserialize `Ticks`:
	{
		for i := 0; i < len(obj.Ticks); i++ {
			obj.Ticks[i], err = DecodeDynamicTick(decoder)
			if err != nil {
				return fmt.Errorf("failed to reading field [%s-%d]: %w", "Ticks", i, err)
			}
		}
	}
	return nil
}

func (obj *DynamicTickArray) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling DynamicTickArray: %w", err)
	}
	return nil
}

func UnmarshalDynamicTickArray(buf []byte) (*DynamicTickArray, error) {
	obj := new(DynamicTickArray)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type FeeTier struct {
	WhirlpoolsConfig solanago.PublicKey
	TickSpacing      uint16
	DefaultFeeRate   uint16
}

func (obj FeeTier) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_FeeTier[:], false)
	if err != nil {
		return err
	}
	// Serialize `WhirlpoolsConfig`:
	if err = encoder.Encode(obj.WhirlpoolsConfig); err != nil {
		return fmt.Errorf("error while marshaling WhirlpoolsConfig:%w", err)
	}
	// Serialize `TickSpacing`:
	if err = encoder.Encode(obj.TickSpacing); err != nil {
		return fmt.Errorf("error while marshaling TickSpacing:%w", err)
	}
	// Serialize `DefaultFeeRate`:
	if err = encoder.Encode(obj.DefaultFeeRate); err != nil {
		return fmt.Errorf("error while marshaling DefaultFeeRate:%w", err)
	}
	return nil
}

func (obj FeeTier) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding FeeTier: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *FeeTier) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_FeeTier[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_FeeTier[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `WhirlpoolsConfig`:
	if err = decoder.Decode(&obj.WhirlpoolsConfig); err != nil {
		return fmt.Errorf("error while unmarshaling WhirlpoolsConfig:%w", err)
	}
	// Deserialize `TickSpacing`:
	if err = decoder.Decode(&obj.TickSpacing); err != nil {
		return fmt.Errorf("error while unmarshaling TickSpacing:%w", err)
	}
	// Deserialize `DefaultFeeRate`:
	if err = decoder.Decode(&obj.DefaultFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling DefaultFeeRate:%w", err)
	}
	return nil
}

func (obj *FeeTier) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling FeeTier: %w", err)
	}
	return nil
}

func UnmarshalFeeTier(buf []byte) (*FeeTier, error) {
	obj := new(FeeTier)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type TickArray struct {
	StartTickIndex int32
	Ticks          [88]Tick
	Whirlpool      solanago.PublicKey
}

func (obj TickArray) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_TickArray[:], false)
	if err != nil {
		return err
	}
	// Serialize `StartTickIndex`:
	if err = encoder.Encode(obj.StartTickIndex); err != nil {
		return fmt.Errorf("error while marshaling StartTickIndex:%w", err)
	}
	// Serialize `Ticks`:
	if err = encoder.Encode(obj.Ticks); err != nil {
		return fmt.Errorf("error while marshaling Ticks:%w", err)
	}
	// Serialize `Whirlpool`:
	if err = encoder.Encode(obj.Whirlpool); err != nil {
		return fmt.Errorf("error while marshaling Whirlpool:%w", err)
	}
	return nil
}

func (obj TickArray) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding TickArray: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *TickArray) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_TickArray[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_TickArray[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `StartTickIndex`:
	if err = decoder.Decode(&obj.StartTickIndex); err != nil {
		return fmt.Errorf("error while unmarshaling StartTickIndex:%w", err)
	}
	// Deserialize `Ticks`:
	if err = decoder.Decode(&obj.Ticks); err != nil {
		return fmt.Errorf("error while unmarshaling Ticks:%w", err)
	}
	// Deserialize `Whirlpool`:
	if err = decoder.Decode(&obj.Whirlpool); err != nil {
		return fmt.Errorf("error while unmarshaling Whirlpool:%w", err)
	}
	return nil
}

func (obj *TickArray) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling TickArray: %w", err)
	}
	return nil
}

func UnmarshalTickArray(buf []byte) (*TickArray, error) {
	obj := new(TickArray)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type LockConfig struct {
	Position        solanago.PublicKey
	PositionOwner   solanago.PublicKey
	Whirlpool       solanago.PublicKey
	LockedTimestamp uint64
	LockType        LockTypeLabel
}

func (obj LockConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_LockConfig[:], false)
	if err != nil {
		return err
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `PositionOwner`:
	if err = encoder.Encode(obj.PositionOwner); err != nil {
		return fmt.Errorf("error while marshaling PositionOwner:%w", err)
	}
	// Serialize `Whirlpool`:
	if err = encoder.Encode(obj.Whirlpool); err != nil {
		return fmt.Errorf("error while marshaling Whirlpool:%w", err)
	}
	// Serialize `LockedTimestamp`:
	if err = encoder.Encode(obj.LockedTimestamp); err != nil {
		return fmt.Errorf("error while marshaling LockedTimestamp:%w", err)
	}
	// Serialize `LockType`:
	if err = encoder.Encode(obj.LockType); err != nil {
		return fmt.Errorf("error while marshaling LockType:%w", err)
	}
	return nil
}

func (obj LockConfig) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding LockConfig: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *LockConfig) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_LockConfig[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_LockConfig[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `PositionOwner`:
	if err = decoder.Decode(&obj.PositionOwner); err != nil {
		return fmt.Errorf("error while unmarshaling PositionOwner:%w", err)
	}
	// Deserialize `Whirlpool`:
	if err = decoder.Decode(&obj.Whirlpool); err != nil {
		return fmt.Errorf("error while unmarshaling Whirlpool:%w", err)
	}
	// Deserialize `LockedTimestamp`:
	if err = decoder.Decode(&obj.LockedTimestamp); err != nil {
		return fmt.Errorf("error while unmarshaling LockedTimestamp:%w", err)
	}
	// Deserialize `LockType`:
	if err = decoder.Decode(&obj.LockType); err != nil {
		return fmt.Errorf("error while unmarshaling LockType:%w", err)
	}
	return nil
}

func (obj *LockConfig) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling LockConfig: %w", err)
	}
	return nil
}

func UnmarshalLockConfig(buf []byte) (*LockConfig, error) {
	obj := new(LockConfig)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type Oracle struct {
	Whirlpool            solanago.PublicKey
	TradeEnableTimestamp uint64
	AdaptiveFeeConstants AdaptiveFeeConstants
	AdaptiveFeeVariables AdaptiveFeeVariables
	Reserved             [128]uint8
}

func (obj Oracle) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Oracle[:], false)
	if err != nil {
		return err
	}
	// Serialize `Whirlpool`:
	if err = encoder.Encode(obj.Whirlpool); err != nil {
		return fmt.Errorf("error while marshaling Whirlpool:%w", err)
	}
	// Serialize `TradeEnableTimestamp`:
	if err = encoder.Encode(obj.TradeEnableTimestamp); err != nil {
		return fmt.Errorf("error while marshaling TradeEnableTimestamp:%w", err)
	}
	// Serialize `AdaptiveFeeConstants`:
	if err = encoder.Encode(obj.AdaptiveFeeConstants); err != nil {
		return fmt.Errorf("error while marshaling AdaptiveFeeConstants:%w", err)
	}
	// Serialize `AdaptiveFeeVariables`:
	if err = encoder.Encode(obj.AdaptiveFeeVariables); err != nil {
		return fmt.Errorf("error while marshaling AdaptiveFeeVariables:%w", err)
	}
	// Serialize `Reserved`:
	if err = encoder.Encode(obj.Reserved); err != nil {
		return fmt.Errorf("error while marshaling Reserved:%w", err)
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
	// Deserialize `Whirlpool`:
	if err = decoder.Decode(&obj.Whirlpool); err != nil {
		return fmt.Errorf("error while unmarshaling Whirlpool:%w", err)
	}
	// Deserialize `TradeEnableTimestamp`:
	if err = decoder.Decode(&obj.TradeEnableTimestamp); err != nil {
		return fmt.Errorf("error while unmarshaling TradeEnableTimestamp:%w", err)
	}
	// Deserialize `AdaptiveFeeConstants`:
	if err = decoder.Decode(&obj.AdaptiveFeeConstants); err != nil {
		return fmt.Errorf("error while unmarshaling AdaptiveFeeConstants:%w", err)
	}
	// Deserialize `AdaptiveFeeVariables`:
	if err = decoder.Decode(&obj.AdaptiveFeeVariables); err != nil {
		return fmt.Errorf("error while unmarshaling AdaptiveFeeVariables:%w", err)
	}
	// Deserialize `Reserved`:
	if err = decoder.Decode(&obj.Reserved); err != nil {
		return fmt.Errorf("error while unmarshaling Reserved:%w", err)
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
	Whirlpool            solanago.PublicKey
	PositionMint         solanago.PublicKey
	Liquidity            binary.Uint128
	TickLowerIndex       int32
	TickUpperIndex       int32
	FeeGrowthCheckpointA binary.Uint128
	FeeOwedA             uint64
	FeeGrowthCheckpointB binary.Uint128
	FeeOwedB             uint64
	RewardInfos          [3]PositionRewardInfo
}

func (obj Position) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Position[:], false)
	if err != nil {
		return err
	}
	// Serialize `Whirlpool`:
	if err = encoder.Encode(obj.Whirlpool); err != nil {
		return fmt.Errorf("error while marshaling Whirlpool:%w", err)
	}
	// Serialize `PositionMint`:
	if err = encoder.Encode(obj.PositionMint); err != nil {
		return fmt.Errorf("error while marshaling PositionMint:%w", err)
	}
	// Serialize `Liquidity`:
	if err = encoder.Encode(obj.Liquidity); err != nil {
		return fmt.Errorf("error while marshaling Liquidity:%w", err)
	}
	// Serialize `TickLowerIndex`:
	if err = encoder.Encode(obj.TickLowerIndex); err != nil {
		return fmt.Errorf("error while marshaling TickLowerIndex:%w", err)
	}
	// Serialize `TickUpperIndex`:
	if err = encoder.Encode(obj.TickUpperIndex); err != nil {
		return fmt.Errorf("error while marshaling TickUpperIndex:%w", err)
	}
	// Serialize `FeeGrowthCheckpointA`:
	if err = encoder.Encode(obj.FeeGrowthCheckpointA); err != nil {
		return fmt.Errorf("error while marshaling FeeGrowthCheckpointA:%w", err)
	}
	// Serialize `FeeOwedA`:
	if err = encoder.Encode(obj.FeeOwedA); err != nil {
		return fmt.Errorf("error while marshaling FeeOwedA:%w", err)
	}
	// Serialize `FeeGrowthCheckpointB`:
	if err = encoder.Encode(obj.FeeGrowthCheckpointB); err != nil {
		return fmt.Errorf("error while marshaling FeeGrowthCheckpointB:%w", err)
	}
	// Serialize `FeeOwedB`:
	if err = encoder.Encode(obj.FeeOwedB); err != nil {
		return fmt.Errorf("error while marshaling FeeOwedB:%w", err)
	}
	// Serialize `RewardInfos`:
	if err = encoder.Encode(obj.RewardInfos); err != nil {
		return fmt.Errorf("error while marshaling RewardInfos:%w", err)
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
	// Deserialize `Whirlpool`:
	if err = decoder.Decode(&obj.Whirlpool); err != nil {
		return fmt.Errorf("error while unmarshaling Whirlpool:%w", err)
	}
	// Deserialize `PositionMint`:
	if err = decoder.Decode(&obj.PositionMint); err != nil {
		return fmt.Errorf("error while unmarshaling PositionMint:%w", err)
	}
	// Deserialize `Liquidity`:
	if err = decoder.Decode(&obj.Liquidity); err != nil {
		return fmt.Errorf("error while unmarshaling Liquidity:%w", err)
	}
	// Deserialize `TickLowerIndex`:
	if err = decoder.Decode(&obj.TickLowerIndex); err != nil {
		return fmt.Errorf("error while unmarshaling TickLowerIndex:%w", err)
	}
	// Deserialize `TickUpperIndex`:
	if err = decoder.Decode(&obj.TickUpperIndex); err != nil {
		return fmt.Errorf("error while unmarshaling TickUpperIndex:%w", err)
	}
	// Deserialize `FeeGrowthCheckpointA`:
	if err = decoder.Decode(&obj.FeeGrowthCheckpointA); err != nil {
		return fmt.Errorf("error while unmarshaling FeeGrowthCheckpointA:%w", err)
	}
	// Deserialize `FeeOwedA`:
	if err = decoder.Decode(&obj.FeeOwedA); err != nil {
		return fmt.Errorf("error while unmarshaling FeeOwedA:%w", err)
	}
	// Deserialize `FeeGrowthCheckpointB`:
	if err = decoder.Decode(&obj.FeeGrowthCheckpointB); err != nil {
		return fmt.Errorf("error while unmarshaling FeeGrowthCheckpointB:%w", err)
	}
	// Deserialize `FeeOwedB`:
	if err = decoder.Decode(&obj.FeeOwedB); err != nil {
		return fmt.Errorf("error while unmarshaling FeeOwedB:%w", err)
	}
	// Deserialize `RewardInfos`:
	if err = decoder.Decode(&obj.RewardInfos); err != nil {
		return fmt.Errorf("error while unmarshaling RewardInfos:%w", err)
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

type PositionBundle struct {
	PositionBundleMint solanago.PublicKey
	PositionBitmap     [32]uint8
}

func (obj PositionBundle) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_PositionBundle[:], false)
	if err != nil {
		return err
	}
	// Serialize `PositionBundleMint`:
	if err = encoder.Encode(obj.PositionBundleMint); err != nil {
		return fmt.Errorf("error while marshaling PositionBundleMint:%w", err)
	}
	// Serialize `PositionBitmap`:
	if err = encoder.Encode(obj.PositionBitmap); err != nil {
		return fmt.Errorf("error while marshaling PositionBitmap:%w", err)
	}
	return nil
}

func (obj PositionBundle) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding PositionBundle: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *PositionBundle) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_PositionBundle[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_PositionBundle[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PositionBundleMint`:
	if err = decoder.Decode(&obj.PositionBundleMint); err != nil {
		return fmt.Errorf("error while unmarshaling PositionBundleMint:%w", err)
	}
	// Deserialize `PositionBitmap`:
	if err = decoder.Decode(&obj.PositionBitmap); err != nil {
		return fmt.Errorf("error while unmarshaling PositionBitmap:%w", err)
	}
	return nil
}

func (obj *PositionBundle) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling PositionBundle: %w", err)
	}
	return nil
}

func UnmarshalPositionBundle(buf []byte) (*PositionBundle, error) {
	obj := new(PositionBundle)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type TokenBadge struct {
	WhirlpoolsConfig solanago.PublicKey
	TokenMint        solanago.PublicKey
}

func (obj TokenBadge) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_TokenBadge[:], false)
	if err != nil {
		return err
	}
	// Serialize `WhirlpoolsConfig`:
	if err = encoder.Encode(obj.WhirlpoolsConfig); err != nil {
		return fmt.Errorf("error while marshaling WhirlpoolsConfig:%w", err)
	}
	// Serialize `TokenMint`:
	if err = encoder.Encode(obj.TokenMint); err != nil {
		return fmt.Errorf("error while marshaling TokenMint:%w", err)
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
	// Deserialize `WhirlpoolsConfig`:
	if err = decoder.Decode(&obj.WhirlpoolsConfig); err != nil {
		return fmt.Errorf("error while unmarshaling WhirlpoolsConfig:%w", err)
	}
	// Deserialize `TokenMint`:
	if err = decoder.Decode(&obj.TokenMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMint:%w", err)
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

type Whirlpool struct {
	WhirlpoolsConfig           solanago.PublicKey
	WhirlpoolBump              [1]uint8
	TickSpacing                uint16
	FeeTierIndexSeed           [2]uint8
	FeeRate                    uint16
	ProtocolFeeRate            uint16
	Liquidity                  binary.Uint128
	SqrtPrice                  binary.Uint128
	TickCurrentIndex           int32
	ProtocolFeeOwedA           uint64
	ProtocolFeeOwedB           uint64
	TokenMintA                 solanago.PublicKey
	TokenVaultA                solanago.PublicKey
	FeeGrowthGlobalA           binary.Uint128
	TokenMintB                 solanago.PublicKey
	TokenVaultB                solanago.PublicKey
	FeeGrowthGlobalB           binary.Uint128
	RewardLastUpdatedTimestamp uint64
	RewardInfos                [3]WhirlpoolRewardInfo
}

func (obj Whirlpool) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Whirlpool[:], false)
	if err != nil {
		return err
	}
	// Serialize `WhirlpoolsConfig`:
	if err = encoder.Encode(obj.WhirlpoolsConfig); err != nil {
		return fmt.Errorf("error while marshaling WhirlpoolsConfig:%w", err)
	}
	// Serialize `WhirlpoolBump`:
	if err = encoder.Encode(obj.WhirlpoolBump); err != nil {
		return fmt.Errorf("error while marshaling WhirlpoolBump:%w", err)
	}
	// Serialize `TickSpacing`:
	if err = encoder.Encode(obj.TickSpacing); err != nil {
		return fmt.Errorf("error while marshaling TickSpacing:%w", err)
	}
	// Serialize `FeeTierIndexSeed`:
	if err = encoder.Encode(obj.FeeTierIndexSeed); err != nil {
		return fmt.Errorf("error while marshaling FeeTierIndexSeed:%w", err)
	}
	// Serialize `FeeRate`:
	if err = encoder.Encode(obj.FeeRate); err != nil {
		return fmt.Errorf("error while marshaling FeeRate:%w", err)
	}
	// Serialize `ProtocolFeeRate`:
	if err = encoder.Encode(obj.ProtocolFeeRate); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeRate:%w", err)
	}
	// Serialize `Liquidity`:
	if err = encoder.Encode(obj.Liquidity); err != nil {
		return fmt.Errorf("error while marshaling Liquidity:%w", err)
	}
	// Serialize `SqrtPrice`:
	if err = encoder.Encode(obj.SqrtPrice); err != nil {
		return fmt.Errorf("error while marshaling SqrtPrice:%w", err)
	}
	// Serialize `TickCurrentIndex`:
	if err = encoder.Encode(obj.TickCurrentIndex); err != nil {
		return fmt.Errorf("error while marshaling TickCurrentIndex:%w", err)
	}
	// Serialize `ProtocolFeeOwedA`:
	if err = encoder.Encode(obj.ProtocolFeeOwedA); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeOwedA:%w", err)
	}
	// Serialize `ProtocolFeeOwedB`:
	if err = encoder.Encode(obj.ProtocolFeeOwedB); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeOwedB:%w", err)
	}
	// Serialize `TokenMintA`:
	if err = encoder.Encode(obj.TokenMintA); err != nil {
		return fmt.Errorf("error while marshaling TokenMintA:%w", err)
	}
	// Serialize `TokenVaultA`:
	if err = encoder.Encode(obj.TokenVaultA); err != nil {
		return fmt.Errorf("error while marshaling TokenVaultA:%w", err)
	}
	// Serialize `FeeGrowthGlobalA`:
	if err = encoder.Encode(obj.FeeGrowthGlobalA); err != nil {
		return fmt.Errorf("error while marshaling FeeGrowthGlobalA:%w", err)
	}
	// Serialize `TokenMintB`:
	if err = encoder.Encode(obj.TokenMintB); err != nil {
		return fmt.Errorf("error while marshaling TokenMintB:%w", err)
	}
	// Serialize `TokenVaultB`:
	if err = encoder.Encode(obj.TokenVaultB); err != nil {
		return fmt.Errorf("error while marshaling TokenVaultB:%w", err)
	}
	// Serialize `FeeGrowthGlobalB`:
	if err = encoder.Encode(obj.FeeGrowthGlobalB); err != nil {
		return fmt.Errorf("error while marshaling FeeGrowthGlobalB:%w", err)
	}
	// Serialize `RewardLastUpdatedTimestamp`:
	if err = encoder.Encode(obj.RewardLastUpdatedTimestamp); err != nil {
		return fmt.Errorf("error while marshaling RewardLastUpdatedTimestamp:%w", err)
	}
	// Serialize `RewardInfos`:
	if err = encoder.Encode(obj.RewardInfos); err != nil {
		return fmt.Errorf("error while marshaling RewardInfos:%w", err)
	}
	return nil
}

func (obj Whirlpool) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Whirlpool: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Whirlpool) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Whirlpool[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Whirlpool[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `WhirlpoolsConfig`:
	if err = decoder.Decode(&obj.WhirlpoolsConfig); err != nil {
		return fmt.Errorf("error while unmarshaling WhirlpoolsConfig:%w", err)
	}
	// Deserialize `WhirlpoolBump`:
	if err = decoder.Decode(&obj.WhirlpoolBump); err != nil {
		return fmt.Errorf("error while unmarshaling WhirlpoolBump:%w", err)
	}
	// Deserialize `TickSpacing`:
	if err = decoder.Decode(&obj.TickSpacing); err != nil {
		return fmt.Errorf("error while unmarshaling TickSpacing:%w", err)
	}
	// Deserialize `FeeTierIndexSeed`:
	if err = decoder.Decode(&obj.FeeTierIndexSeed); err != nil {
		return fmt.Errorf("error while unmarshaling FeeTierIndexSeed:%w", err)
	}
	// Deserialize `FeeRate`:
	if err = decoder.Decode(&obj.FeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling FeeRate:%w", err)
	}
	// Deserialize `ProtocolFeeRate`:
	if err = decoder.Decode(&obj.ProtocolFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeRate:%w", err)
	}
	// Deserialize `Liquidity`:
	if err = decoder.Decode(&obj.Liquidity); err != nil {
		return fmt.Errorf("error while unmarshaling Liquidity:%w", err)
	}
	// Deserialize `SqrtPrice`:
	if err = decoder.Decode(&obj.SqrtPrice); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtPrice:%w", err)
	}
	// Deserialize `TickCurrentIndex`:
	if err = decoder.Decode(&obj.TickCurrentIndex); err != nil {
		return fmt.Errorf("error while unmarshaling TickCurrentIndex:%w", err)
	}
	// Deserialize `ProtocolFeeOwedA`:
	if err = decoder.Decode(&obj.ProtocolFeeOwedA); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeOwedA:%w", err)
	}
	// Deserialize `ProtocolFeeOwedB`:
	if err = decoder.Decode(&obj.ProtocolFeeOwedB); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeOwedB:%w", err)
	}
	// Deserialize `TokenMintA`:
	if err = decoder.Decode(&obj.TokenMintA); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMintA:%w", err)
	}
	// Deserialize `TokenVaultA`:
	if err = decoder.Decode(&obj.TokenVaultA); err != nil {
		return fmt.Errorf("error while unmarshaling TokenVaultA:%w", err)
	}
	// Deserialize `FeeGrowthGlobalA`:
	if err = decoder.Decode(&obj.FeeGrowthGlobalA); err != nil {
		return fmt.Errorf("error while unmarshaling FeeGrowthGlobalA:%w", err)
	}
	// Deserialize `TokenMintB`:
	if err = decoder.Decode(&obj.TokenMintB); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMintB:%w", err)
	}
	// Deserialize `TokenVaultB`:
	if err = decoder.Decode(&obj.TokenVaultB); err != nil {
		return fmt.Errorf("error while unmarshaling TokenVaultB:%w", err)
	}
	// Deserialize `FeeGrowthGlobalB`:
	if err = decoder.Decode(&obj.FeeGrowthGlobalB); err != nil {
		return fmt.Errorf("error while unmarshaling FeeGrowthGlobalB:%w", err)
	}
	// Deserialize `RewardLastUpdatedTimestamp`:
	if err = decoder.Decode(&obj.RewardLastUpdatedTimestamp); err != nil {
		return fmt.Errorf("error while unmarshaling RewardLastUpdatedTimestamp:%w", err)
	}
	// Deserialize `RewardInfos`:
	if err = decoder.Decode(&obj.RewardInfos); err != nil {
		return fmt.Errorf("error while unmarshaling RewardInfos:%w", err)
	}
	return nil
}

func (obj *Whirlpool) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Whirlpool: %w", err)
	}
	return nil
}

func UnmarshalWhirlpool(buf []byte) (*Whirlpool, error) {
	obj := new(Whirlpool)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
