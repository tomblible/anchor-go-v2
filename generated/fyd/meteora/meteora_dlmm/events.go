package meteora_dlmm

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type EvtAddLiquidity struct {
	LbPair      solanago.PublicKey
	From        solanago.PublicKey
	Position    solanago.PublicKey
	Amounts     [2]uint64
	ActiveBinId int32
}

func (obj EvtAddLiquidity) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_AddLiquidity[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `From`:
	if err = encoder.Encode(obj.From); err != nil {
		return fmt.Errorf("error while marshaling From:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `Amounts`:
	if err = encoder.Encode(obj.Amounts); err != nil {
		return fmt.Errorf("error while marshaling Amounts:%w", err)
	}
	// Serialize `ActiveBinId`:
	if err = encoder.Encode(obj.ActiveBinId); err != nil {
		return fmt.Errorf("error while marshaling ActiveBinId:%w", err)
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
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `From`:
	if err = decoder.Decode(&obj.From); err != nil {
		return fmt.Errorf("error while unmarshaling From:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `Amounts`:
	if err = decoder.Decode(&obj.Amounts); err != nil {
		return fmt.Errorf("error while unmarshaling Amounts:%w", err)
	}
	// Deserialize `ActiveBinId`:
	if err = decoder.Decode(&obj.ActiveBinId); err != nil {
		return fmt.Errorf("error while unmarshaling ActiveBinId:%w", err)
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

type EvtClaimFee struct {
	LbPair   solanago.PublicKey
	Position solanago.PublicKey
	Owner    solanago.PublicKey
	FeeX     uint64
	FeeY     uint64
}

func (obj EvtClaimFee) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_ClaimFee[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `FeeX`:
	if err = encoder.Encode(obj.FeeX); err != nil {
		return fmt.Errorf("error while marshaling FeeX:%w", err)
	}
	// Serialize `FeeY`:
	if err = encoder.Encode(obj.FeeY); err != nil {
		return fmt.Errorf("error while marshaling FeeY:%w", err)
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
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `FeeX`:
	if err = decoder.Decode(&obj.FeeX); err != nil {
		return fmt.Errorf("error while unmarshaling FeeX:%w", err)
	}
	// Deserialize `FeeY`:
	if err = decoder.Decode(&obj.FeeY); err != nil {
		return fmt.Errorf("error while unmarshaling FeeY:%w", err)
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

type EvtClaimReward struct {
	LbPair      solanago.PublicKey
	Position    solanago.PublicKey
	Owner       solanago.PublicKey
	RewardIndex uint64
	TotalReward uint64
}

func (obj EvtClaimReward) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_ClaimReward[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `RewardIndex`:
	if err = encoder.Encode(obj.RewardIndex); err != nil {
		return fmt.Errorf("error while marshaling RewardIndex:%w", err)
	}
	// Serialize `TotalReward`:
	if err = encoder.Encode(obj.TotalReward); err != nil {
		return fmt.Errorf("error while marshaling TotalReward:%w", err)
	}
	return nil
}

func (obj EvtClaimReward) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtClaimReward: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtClaimReward) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_ClaimReward[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_ClaimReward[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `RewardIndex`:
	if err = decoder.Decode(&obj.RewardIndex); err != nil {
		return fmt.Errorf("error while unmarshaling RewardIndex:%w", err)
	}
	// Deserialize `TotalReward`:
	if err = decoder.Decode(&obj.TotalReward); err != nil {
		return fmt.Errorf("error while unmarshaling TotalReward:%w", err)
	}
	return nil
}

func (obj *EvtClaimReward) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtClaimReward: %w", err)
	}
	return nil
}

func UnmarshalEvtClaimReward(buf []byte) (*EvtClaimReward, error) {
	obj := new(EvtClaimReward)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtCompositionFee struct {
	From                    solanago.PublicKey
	BinId                   int16
	TokenXFeeAmount         uint64
	TokenYFeeAmount         uint64
	ProtocolTokenXFeeAmount uint64
	ProtocolTokenYFeeAmount uint64
}

func (obj EvtCompositionFee) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CompositionFee[:], false)
	if err != nil {
		return err
	}
	// Serialize `From`:
	if err = encoder.Encode(obj.From); err != nil {
		return fmt.Errorf("error while marshaling From:%w", err)
	}
	// Serialize `BinId`:
	if err = encoder.Encode(obj.BinId); err != nil {
		return fmt.Errorf("error while marshaling BinId:%w", err)
	}
	// Serialize `TokenXFeeAmount`:
	if err = encoder.Encode(obj.TokenXFeeAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenXFeeAmount:%w", err)
	}
	// Serialize `TokenYFeeAmount`:
	if err = encoder.Encode(obj.TokenYFeeAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenYFeeAmount:%w", err)
	}
	// Serialize `ProtocolTokenXFeeAmount`:
	if err = encoder.Encode(obj.ProtocolTokenXFeeAmount); err != nil {
		return fmt.Errorf("error while marshaling ProtocolTokenXFeeAmount:%w", err)
	}
	// Serialize `ProtocolTokenYFeeAmount`:
	if err = encoder.Encode(obj.ProtocolTokenYFeeAmount); err != nil {
		return fmt.Errorf("error while marshaling ProtocolTokenYFeeAmount:%w", err)
	}
	return nil
}

func (obj EvtCompositionFee) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtCompositionFee: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtCompositionFee) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CompositionFee[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CompositionFee[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `From`:
	if err = decoder.Decode(&obj.From); err != nil {
		return fmt.Errorf("error while unmarshaling From:%w", err)
	}
	// Deserialize `BinId`:
	if err = decoder.Decode(&obj.BinId); err != nil {
		return fmt.Errorf("error while unmarshaling BinId:%w", err)
	}
	// Deserialize `TokenXFeeAmount`:
	if err = decoder.Decode(&obj.TokenXFeeAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenXFeeAmount:%w", err)
	}
	// Deserialize `TokenYFeeAmount`:
	if err = decoder.Decode(&obj.TokenYFeeAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenYFeeAmount:%w", err)
	}
	// Deserialize `ProtocolTokenXFeeAmount`:
	if err = decoder.Decode(&obj.ProtocolTokenXFeeAmount); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolTokenXFeeAmount:%w", err)
	}
	// Deserialize `ProtocolTokenYFeeAmount`:
	if err = decoder.Decode(&obj.ProtocolTokenYFeeAmount); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolTokenYFeeAmount:%w", err)
	}
	return nil
}

func (obj *EvtCompositionFee) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtCompositionFee: %w", err)
	}
	return nil
}

func UnmarshalEvtCompositionFee(buf []byte) (*EvtCompositionFee, error) {
	obj := new(EvtCompositionFee)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtDecreasePositionLength struct {
	LbPair         solanago.PublicKey
	Position       solanago.PublicKey
	Owner          solanago.PublicKey
	LengthToRemove uint16
	Side           uint8
}

func (obj EvtDecreasePositionLength) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_DecreasePositionLength[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `LengthToRemove`:
	if err = encoder.Encode(obj.LengthToRemove); err != nil {
		return fmt.Errorf("error while marshaling LengthToRemove:%w", err)
	}
	// Serialize `Side`:
	if err = encoder.Encode(obj.Side); err != nil {
		return fmt.Errorf("error while marshaling Side:%w", err)
	}
	return nil
}

func (obj EvtDecreasePositionLength) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtDecreasePositionLength: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtDecreasePositionLength) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_DecreasePositionLength[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_DecreasePositionLength[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `LengthToRemove`:
	if err = decoder.Decode(&obj.LengthToRemove); err != nil {
		return fmt.Errorf("error while unmarshaling LengthToRemove:%w", err)
	}
	// Deserialize `Side`:
	if err = decoder.Decode(&obj.Side); err != nil {
		return fmt.Errorf("error while unmarshaling Side:%w", err)
	}
	return nil
}

func (obj *EvtDecreasePositionLength) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtDecreasePositionLength: %w", err)
	}
	return nil
}

func UnmarshalEvtDecreasePositionLength(buf []byte) (*EvtDecreasePositionLength, error) {
	obj := new(EvtDecreasePositionLength)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtDynamicFeeParameterUpdate struct {
	LbPair solanago.PublicKey

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
}

func (obj EvtDynamicFeeParameterUpdate) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_DynamicFeeParameterUpdate[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
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
	return nil
}

func (obj EvtDynamicFeeParameterUpdate) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtDynamicFeeParameterUpdate: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtDynamicFeeParameterUpdate) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_DynamicFeeParameterUpdate[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_DynamicFeeParameterUpdate[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
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
	return nil
}

func (obj *EvtDynamicFeeParameterUpdate) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtDynamicFeeParameterUpdate: %w", err)
	}
	return nil
}

func UnmarshalEvtDynamicFeeParameterUpdate(buf []byte) (*EvtDynamicFeeParameterUpdate, error) {
	obj := new(EvtDynamicFeeParameterUpdate)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtFeeParameterUpdate struct {
	LbPair        solanago.PublicKey
	ProtocolShare uint16
	BaseFactor    uint16
}

func (obj EvtFeeParameterUpdate) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_FeeParameterUpdate[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `ProtocolShare`:
	if err = encoder.Encode(obj.ProtocolShare); err != nil {
		return fmt.Errorf("error while marshaling ProtocolShare:%w", err)
	}
	// Serialize `BaseFactor`:
	if err = encoder.Encode(obj.BaseFactor); err != nil {
		return fmt.Errorf("error while marshaling BaseFactor:%w", err)
	}
	return nil
}

func (obj EvtFeeParameterUpdate) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtFeeParameterUpdate: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtFeeParameterUpdate) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_FeeParameterUpdate[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_FeeParameterUpdate[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `ProtocolShare`:
	if err = decoder.Decode(&obj.ProtocolShare); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolShare:%w", err)
	}
	// Deserialize `BaseFactor`:
	if err = decoder.Decode(&obj.BaseFactor); err != nil {
		return fmt.Errorf("error while unmarshaling BaseFactor:%w", err)
	}
	return nil
}

func (obj *EvtFeeParameterUpdate) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtFeeParameterUpdate: %w", err)
	}
	return nil
}

func UnmarshalEvtFeeParameterUpdate(buf []byte) (*EvtFeeParameterUpdate, error) {
	obj := new(EvtFeeParameterUpdate)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtFundReward struct {
	LbPair      solanago.PublicKey
	Funder      solanago.PublicKey
	RewardIndex uint64
	Amount      uint64
}

func (obj EvtFundReward) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_FundReward[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `Funder`:
	if err = encoder.Encode(obj.Funder); err != nil {
		return fmt.Errorf("error while marshaling Funder:%w", err)
	}
	// Serialize `RewardIndex`:
	if err = encoder.Encode(obj.RewardIndex); err != nil {
		return fmt.Errorf("error while marshaling RewardIndex:%w", err)
	}
	// Serialize `Amount`:
	if err = encoder.Encode(obj.Amount); err != nil {
		return fmt.Errorf("error while marshaling Amount:%w", err)
	}
	return nil
}

func (obj EvtFundReward) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtFundReward: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtFundReward) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_FundReward[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_FundReward[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `Funder`:
	if err = decoder.Decode(&obj.Funder); err != nil {
		return fmt.Errorf("error while unmarshaling Funder:%w", err)
	}
	// Deserialize `RewardIndex`:
	if err = decoder.Decode(&obj.RewardIndex); err != nil {
		return fmt.Errorf("error while unmarshaling RewardIndex:%w", err)
	}
	// Deserialize `Amount`:
	if err = decoder.Decode(&obj.Amount); err != nil {
		return fmt.Errorf("error while unmarshaling Amount:%w", err)
	}
	return nil
}

func (obj *EvtFundReward) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtFundReward: %w", err)
	}
	return nil
}

func UnmarshalEvtFundReward(buf []byte) (*EvtFundReward, error) {
	obj := new(EvtFundReward)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtGoToABin struct {
	LbPair    solanago.PublicKey
	FromBinId int32
	ToBinId   int32
}

func (obj EvtGoToABin) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_GoToABin[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `FromBinId`:
	if err = encoder.Encode(obj.FromBinId); err != nil {
		return fmt.Errorf("error while marshaling FromBinId:%w", err)
	}
	// Serialize `ToBinId`:
	if err = encoder.Encode(obj.ToBinId); err != nil {
		return fmt.Errorf("error while marshaling ToBinId:%w", err)
	}
	return nil
}

func (obj EvtGoToABin) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtGoToABin: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtGoToABin) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_GoToABin[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_GoToABin[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `FromBinId`:
	if err = decoder.Decode(&obj.FromBinId); err != nil {
		return fmt.Errorf("error while unmarshaling FromBinId:%w", err)
	}
	// Deserialize `ToBinId`:
	if err = decoder.Decode(&obj.ToBinId); err != nil {
		return fmt.Errorf("error while unmarshaling ToBinId:%w", err)
	}
	return nil
}

func (obj *EvtGoToABin) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtGoToABin: %w", err)
	}
	return nil
}

func UnmarshalEvtGoToABin(buf []byte) (*EvtGoToABin, error) {
	obj := new(EvtGoToABin)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtIncreaseObservation struct {
	Oracle               solanago.PublicKey
	NewObservationLength uint64
}

func (obj EvtIncreaseObservation) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_IncreaseObservation[:], false)
	if err != nil {
		return err
	}
	// Serialize `Oracle`:
	if err = encoder.Encode(obj.Oracle); err != nil {
		return fmt.Errorf("error while marshaling Oracle:%w", err)
	}
	// Serialize `NewObservationLength`:
	if err = encoder.Encode(obj.NewObservationLength); err != nil {
		return fmt.Errorf("error while marshaling NewObservationLength:%w", err)
	}
	return nil
}

func (obj EvtIncreaseObservation) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtIncreaseObservation: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtIncreaseObservation) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_IncreaseObservation[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_IncreaseObservation[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Oracle`:
	if err = decoder.Decode(&obj.Oracle); err != nil {
		return fmt.Errorf("error while unmarshaling Oracle:%w", err)
	}
	// Deserialize `NewObservationLength`:
	if err = decoder.Decode(&obj.NewObservationLength); err != nil {
		return fmt.Errorf("error while unmarshaling NewObservationLength:%w", err)
	}
	return nil
}

func (obj *EvtIncreaseObservation) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtIncreaseObservation: %w", err)
	}
	return nil
}

func UnmarshalEvtIncreaseObservation(buf []byte) (*EvtIncreaseObservation, error) {
	obj := new(EvtIncreaseObservation)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtIncreasePositionLength struct {
	LbPair      solanago.PublicKey
	Position    solanago.PublicKey
	Owner       solanago.PublicKey
	LengthToAdd uint16
	Side        uint8
}

func (obj EvtIncreasePositionLength) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_IncreasePositionLength[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `LengthToAdd`:
	if err = encoder.Encode(obj.LengthToAdd); err != nil {
		return fmt.Errorf("error while marshaling LengthToAdd:%w", err)
	}
	// Serialize `Side`:
	if err = encoder.Encode(obj.Side); err != nil {
		return fmt.Errorf("error while marshaling Side:%w", err)
	}
	return nil
}

func (obj EvtIncreasePositionLength) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtIncreasePositionLength: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtIncreasePositionLength) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_IncreasePositionLength[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_IncreasePositionLength[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `LengthToAdd`:
	if err = decoder.Decode(&obj.LengthToAdd); err != nil {
		return fmt.Errorf("error while unmarshaling LengthToAdd:%w", err)
	}
	// Deserialize `Side`:
	if err = decoder.Decode(&obj.Side); err != nil {
		return fmt.Errorf("error while unmarshaling Side:%w", err)
	}
	return nil
}

func (obj *EvtIncreasePositionLength) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtIncreasePositionLength: %w", err)
	}
	return nil
}

func UnmarshalEvtIncreasePositionLength(buf []byte) (*EvtIncreasePositionLength, error) {
	obj := new(EvtIncreasePositionLength)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtInitializeReward struct {
	LbPair         solanago.PublicKey
	RewardMint     solanago.PublicKey
	Funder         solanago.PublicKey
	RewardIndex    uint64
	RewardDuration uint64
}

func (obj EvtInitializeReward) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_InitializeReward[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `RewardMint`:
	if err = encoder.Encode(obj.RewardMint); err != nil {
		return fmt.Errorf("error while marshaling RewardMint:%w", err)
	}
	// Serialize `Funder`:
	if err = encoder.Encode(obj.Funder); err != nil {
		return fmt.Errorf("error while marshaling Funder:%w", err)
	}
	// Serialize `RewardIndex`:
	if err = encoder.Encode(obj.RewardIndex); err != nil {
		return fmt.Errorf("error while marshaling RewardIndex:%w", err)
	}
	// Serialize `RewardDuration`:
	if err = encoder.Encode(obj.RewardDuration); err != nil {
		return fmt.Errorf("error while marshaling RewardDuration:%w", err)
	}
	return nil
}

func (obj EvtInitializeReward) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtInitializeReward: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtInitializeReward) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_InitializeReward[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_InitializeReward[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `RewardMint`:
	if err = decoder.Decode(&obj.RewardMint); err != nil {
		return fmt.Errorf("error while unmarshaling RewardMint:%w", err)
	}
	// Deserialize `Funder`:
	if err = decoder.Decode(&obj.Funder); err != nil {
		return fmt.Errorf("error while unmarshaling Funder:%w", err)
	}
	// Deserialize `RewardIndex`:
	if err = decoder.Decode(&obj.RewardIndex); err != nil {
		return fmt.Errorf("error while unmarshaling RewardIndex:%w", err)
	}
	// Deserialize `RewardDuration`:
	if err = decoder.Decode(&obj.RewardDuration); err != nil {
		return fmt.Errorf("error while unmarshaling RewardDuration:%w", err)
	}
	return nil
}

func (obj *EvtInitializeReward) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtInitializeReward: %w", err)
	}
	return nil
}

func UnmarshalEvtInitializeReward(buf []byte) (*EvtInitializeReward, error) {
	obj := new(EvtInitializeReward)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtLbPairCreate struct {
	LbPair  solanago.PublicKey
	BinStep uint16
	TokenX  solanago.PublicKey
	TokenY  solanago.PublicKey
}

func (obj EvtLbPairCreate) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_LbPairCreate[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `BinStep`:
	if err = encoder.Encode(obj.BinStep); err != nil {
		return fmt.Errorf("error while marshaling BinStep:%w", err)
	}
	// Serialize `TokenX`:
	if err = encoder.Encode(obj.TokenX); err != nil {
		return fmt.Errorf("error while marshaling TokenX:%w", err)
	}
	// Serialize `TokenY`:
	if err = encoder.Encode(obj.TokenY); err != nil {
		return fmt.Errorf("error while marshaling TokenY:%w", err)
	}
	return nil
}

func (obj EvtLbPairCreate) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtLbPairCreate: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtLbPairCreate) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_LbPairCreate[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_LbPairCreate[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `BinStep`:
	if err = decoder.Decode(&obj.BinStep); err != nil {
		return fmt.Errorf("error while unmarshaling BinStep:%w", err)
	}
	// Deserialize `TokenX`:
	if err = decoder.Decode(&obj.TokenX); err != nil {
		return fmt.Errorf("error while unmarshaling TokenX:%w", err)
	}
	// Deserialize `TokenY`:
	if err = decoder.Decode(&obj.TokenY); err != nil {
		return fmt.Errorf("error while unmarshaling TokenY:%w", err)
	}
	return nil
}

func (obj *EvtLbPairCreate) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtLbPairCreate: %w", err)
	}
	return nil
}

func UnmarshalEvtLbPairCreate(buf []byte) (*EvtLbPairCreate, error) {
	obj := new(EvtLbPairCreate)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtPositionClose struct {
	Position solanago.PublicKey
	Owner    solanago.PublicKey
}

func (obj EvtPositionClose) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_PositionClose[:], false)
	if err != nil {
		return err
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	return nil
}

func (obj EvtPositionClose) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtPositionClose: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtPositionClose) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_PositionClose[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_PositionClose[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	return nil
}

func (obj *EvtPositionClose) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtPositionClose: %w", err)
	}
	return nil
}

func UnmarshalEvtPositionClose(buf []byte) (*EvtPositionClose, error) {
	obj := new(EvtPositionClose)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtPositionCreate struct {
	LbPair   solanago.PublicKey
	Position solanago.PublicKey
	Owner    solanago.PublicKey
}

func (obj EvtPositionCreate) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_PositionCreate[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	return nil
}

func (obj EvtPositionCreate) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtPositionCreate: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtPositionCreate) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_PositionCreate[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_PositionCreate[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	return nil
}

func (obj *EvtPositionCreate) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtPositionCreate: %w", err)
	}
	return nil
}

func UnmarshalEvtPositionCreate(buf []byte) (*EvtPositionCreate, error) {
	obj := new(EvtPositionCreate)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtRebalancing struct {
	LbPair           solanago.PublicKey
	Position         solanago.PublicKey
	XWithdrawnAmount uint64
	XAddedAmount     uint64
	YWithdrawnAmount uint64
	YAddedAmount     uint64
	XFeeAmount       uint64
	YFeeAmount       uint64
	OldMinId         int32
	OldMaxId         int32
	NewMinId         int32
	NewMaxId         int32
	Rewards          [2]uint64
}

func (obj EvtRebalancing) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_Rebalancing[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `XWithdrawnAmount`:
	if err = encoder.Encode(obj.XWithdrawnAmount); err != nil {
		return fmt.Errorf("error while marshaling XWithdrawnAmount:%w", err)
	}
	// Serialize `XAddedAmount`:
	if err = encoder.Encode(obj.XAddedAmount); err != nil {
		return fmt.Errorf("error while marshaling XAddedAmount:%w", err)
	}
	// Serialize `YWithdrawnAmount`:
	if err = encoder.Encode(obj.YWithdrawnAmount); err != nil {
		return fmt.Errorf("error while marshaling YWithdrawnAmount:%w", err)
	}
	// Serialize `YAddedAmount`:
	if err = encoder.Encode(obj.YAddedAmount); err != nil {
		return fmt.Errorf("error while marshaling YAddedAmount:%w", err)
	}
	// Serialize `XFeeAmount`:
	if err = encoder.Encode(obj.XFeeAmount); err != nil {
		return fmt.Errorf("error while marshaling XFeeAmount:%w", err)
	}
	// Serialize `YFeeAmount`:
	if err = encoder.Encode(obj.YFeeAmount); err != nil {
		return fmt.Errorf("error while marshaling YFeeAmount:%w", err)
	}
	// Serialize `OldMinId`:
	if err = encoder.Encode(obj.OldMinId); err != nil {
		return fmt.Errorf("error while marshaling OldMinId:%w", err)
	}
	// Serialize `OldMaxId`:
	if err = encoder.Encode(obj.OldMaxId); err != nil {
		return fmt.Errorf("error while marshaling OldMaxId:%w", err)
	}
	// Serialize `NewMinId`:
	if err = encoder.Encode(obj.NewMinId); err != nil {
		return fmt.Errorf("error while marshaling NewMinId:%w", err)
	}
	// Serialize `NewMaxId`:
	if err = encoder.Encode(obj.NewMaxId); err != nil {
		return fmt.Errorf("error while marshaling NewMaxId:%w", err)
	}
	// Serialize `Rewards`:
	if err = encoder.Encode(obj.Rewards); err != nil {
		return fmt.Errorf("error while marshaling Rewards:%w", err)
	}
	return nil
}

func (obj EvtRebalancing) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtRebalancing: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtRebalancing) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_Rebalancing[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_Rebalancing[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `XWithdrawnAmount`:
	if err = decoder.Decode(&obj.XWithdrawnAmount); err != nil {
		return fmt.Errorf("error while unmarshaling XWithdrawnAmount:%w", err)
	}
	// Deserialize `XAddedAmount`:
	if err = decoder.Decode(&obj.XAddedAmount); err != nil {
		return fmt.Errorf("error while unmarshaling XAddedAmount:%w", err)
	}
	// Deserialize `YWithdrawnAmount`:
	if err = decoder.Decode(&obj.YWithdrawnAmount); err != nil {
		return fmt.Errorf("error while unmarshaling YWithdrawnAmount:%w", err)
	}
	// Deserialize `YAddedAmount`:
	if err = decoder.Decode(&obj.YAddedAmount); err != nil {
		return fmt.Errorf("error while unmarshaling YAddedAmount:%w", err)
	}
	// Deserialize `XFeeAmount`:
	if err = decoder.Decode(&obj.XFeeAmount); err != nil {
		return fmt.Errorf("error while unmarshaling XFeeAmount:%w", err)
	}
	// Deserialize `YFeeAmount`:
	if err = decoder.Decode(&obj.YFeeAmount); err != nil {
		return fmt.Errorf("error while unmarshaling YFeeAmount:%w", err)
	}
	// Deserialize `OldMinId`:
	if err = decoder.Decode(&obj.OldMinId); err != nil {
		return fmt.Errorf("error while unmarshaling OldMinId:%w", err)
	}
	// Deserialize `OldMaxId`:
	if err = decoder.Decode(&obj.OldMaxId); err != nil {
		return fmt.Errorf("error while unmarshaling OldMaxId:%w", err)
	}
	// Deserialize `NewMinId`:
	if err = decoder.Decode(&obj.NewMinId); err != nil {
		return fmt.Errorf("error while unmarshaling NewMinId:%w", err)
	}
	// Deserialize `NewMaxId`:
	if err = decoder.Decode(&obj.NewMaxId); err != nil {
		return fmt.Errorf("error while unmarshaling NewMaxId:%w", err)
	}
	// Deserialize `Rewards`:
	if err = decoder.Decode(&obj.Rewards); err != nil {
		return fmt.Errorf("error while unmarshaling Rewards:%w", err)
	}
	return nil
}

func (obj *EvtRebalancing) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtRebalancing: %w", err)
	}
	return nil
}

func UnmarshalEvtRebalancing(buf []byte) (*EvtRebalancing, error) {
	obj := new(EvtRebalancing)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtRemoveLiquidity struct {
	LbPair      solanago.PublicKey
	From        solanago.PublicKey
	Position    solanago.PublicKey
	Amounts     [2]uint64
	ActiveBinId int32
}

func (obj EvtRemoveLiquidity) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_RemoveLiquidity[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `From`:
	if err = encoder.Encode(obj.From); err != nil {
		return fmt.Errorf("error while marshaling From:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `Amounts`:
	if err = encoder.Encode(obj.Amounts); err != nil {
		return fmt.Errorf("error while marshaling Amounts:%w", err)
	}
	// Serialize `ActiveBinId`:
	if err = encoder.Encode(obj.ActiveBinId); err != nil {
		return fmt.Errorf("error while marshaling ActiveBinId:%w", err)
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
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `From`:
	if err = decoder.Decode(&obj.From); err != nil {
		return fmt.Errorf("error while unmarshaling From:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `Amounts`:
	if err = decoder.Decode(&obj.Amounts); err != nil {
		return fmt.Errorf("error while unmarshaling Amounts:%w", err)
	}
	// Deserialize `ActiveBinId`:
	if err = decoder.Decode(&obj.ActiveBinId); err != nil {
		return fmt.Errorf("error while unmarshaling ActiveBinId:%w", err)
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

type EvtSwap struct {
	LbPair      solanago.PublicKey
	From        solanago.PublicKey
	StartBinId  int32
	EndBinId    int32
	AmountIn    uint64
	AmountOut   uint64
	SwapForY    bool
	Fee         uint64
	ProtocolFee uint64
	FeeBps      binary.Uint128
	HostFee     uint64
}

func (obj EvtSwap) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_Swap[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `From`:
	if err = encoder.Encode(obj.From); err != nil {
		return fmt.Errorf("error while marshaling From:%w", err)
	}
	// Serialize `StartBinId`:
	if err = encoder.Encode(obj.StartBinId); err != nil {
		return fmt.Errorf("error while marshaling StartBinId:%w", err)
	}
	// Serialize `EndBinId`:
	if err = encoder.Encode(obj.EndBinId); err != nil {
		return fmt.Errorf("error while marshaling EndBinId:%w", err)
	}
	// Serialize `AmountIn`:
	if err = encoder.Encode(obj.AmountIn); err != nil {
		return fmt.Errorf("error while marshaling AmountIn:%w", err)
	}
	// Serialize `AmountOut`:
	if err = encoder.Encode(obj.AmountOut); err != nil {
		return fmt.Errorf("error while marshaling AmountOut:%w", err)
	}
	// Serialize `SwapForY`:
	if err = encoder.Encode(obj.SwapForY); err != nil {
		return fmt.Errorf("error while marshaling SwapForY:%w", err)
	}
	// Serialize `Fee`:
	if err = encoder.Encode(obj.Fee); err != nil {
		return fmt.Errorf("error while marshaling Fee:%w", err)
	}
	// Serialize `ProtocolFee`:
	if err = encoder.Encode(obj.ProtocolFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFee:%w", err)
	}
	// Serialize `FeeBps`:
	if err = encoder.Encode(obj.FeeBps); err != nil {
		return fmt.Errorf("error while marshaling FeeBps:%w", err)
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
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `From`:
	if err = decoder.Decode(&obj.From); err != nil {
		return fmt.Errorf("error while unmarshaling From:%w", err)
	}
	// Deserialize `StartBinId`:
	if err = decoder.Decode(&obj.StartBinId); err != nil {
		return fmt.Errorf("error while unmarshaling StartBinId:%w", err)
	}
	// Deserialize `EndBinId`:
	if err = decoder.Decode(&obj.EndBinId); err != nil {
		return fmt.Errorf("error while unmarshaling EndBinId:%w", err)
	}
	// Deserialize `AmountIn`:
	if err = decoder.Decode(&obj.AmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling AmountIn:%w", err)
	}
	// Deserialize `AmountOut`:
	if err = decoder.Decode(&obj.AmountOut); err != nil {
		return fmt.Errorf("error while unmarshaling AmountOut:%w", err)
	}
	// Deserialize `SwapForY`:
	if err = decoder.Decode(&obj.SwapForY); err != nil {
		return fmt.Errorf("error while unmarshaling SwapForY:%w", err)
	}
	// Deserialize `Fee`:
	if err = decoder.Decode(&obj.Fee); err != nil {
		return fmt.Errorf("error while unmarshaling Fee:%w", err)
	}
	// Deserialize `ProtocolFee`:
	if err = decoder.Decode(&obj.ProtocolFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFee:%w", err)
	}
	// Deserialize `FeeBps`:
	if err = decoder.Decode(&obj.FeeBps); err != nil {
		return fmt.Errorf("error while unmarshaling FeeBps:%w", err)
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

type EvtUpdatePositionLockReleasePoint struct {
	Position            solanago.PublicKey
	CurrentPoint        uint64
	NewLockReleasePoint uint64
	OldLockReleasePoint uint64
	Sender              solanago.PublicKey
}

func (obj EvtUpdatePositionLockReleasePoint) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_UpdatePositionLockReleasePoint[:], false)
	if err != nil {
		return err
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `CurrentPoint`:
	if err = encoder.Encode(obj.CurrentPoint); err != nil {
		return fmt.Errorf("error while marshaling CurrentPoint:%w", err)
	}
	// Serialize `NewLockReleasePoint`:
	if err = encoder.Encode(obj.NewLockReleasePoint); err != nil {
		return fmt.Errorf("error while marshaling NewLockReleasePoint:%w", err)
	}
	// Serialize `OldLockReleasePoint`:
	if err = encoder.Encode(obj.OldLockReleasePoint); err != nil {
		return fmt.Errorf("error while marshaling OldLockReleasePoint:%w", err)
	}
	// Serialize `Sender`:
	if err = encoder.Encode(obj.Sender); err != nil {
		return fmt.Errorf("error while marshaling Sender:%w", err)
	}
	return nil
}

func (obj EvtUpdatePositionLockReleasePoint) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtUpdatePositionLockReleasePoint: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtUpdatePositionLockReleasePoint) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_UpdatePositionLockReleasePoint[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_UpdatePositionLockReleasePoint[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `CurrentPoint`:
	if err = decoder.Decode(&obj.CurrentPoint); err != nil {
		return fmt.Errorf("error while unmarshaling CurrentPoint:%w", err)
	}
	// Deserialize `NewLockReleasePoint`:
	if err = decoder.Decode(&obj.NewLockReleasePoint); err != nil {
		return fmt.Errorf("error while unmarshaling NewLockReleasePoint:%w", err)
	}
	// Deserialize `OldLockReleasePoint`:
	if err = decoder.Decode(&obj.OldLockReleasePoint); err != nil {
		return fmt.Errorf("error while unmarshaling OldLockReleasePoint:%w", err)
	}
	// Deserialize `Sender`:
	if err = decoder.Decode(&obj.Sender); err != nil {
		return fmt.Errorf("error while unmarshaling Sender:%w", err)
	}
	return nil
}

func (obj *EvtUpdatePositionLockReleasePoint) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtUpdatePositionLockReleasePoint: %w", err)
	}
	return nil
}

func UnmarshalEvtUpdatePositionLockReleasePoint(buf []byte) (*EvtUpdatePositionLockReleasePoint, error) {
	obj := new(EvtUpdatePositionLockReleasePoint)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtUpdatePositionOperator struct {
	Position    solanago.PublicKey
	OldOperator solanago.PublicKey
	NewOperator solanago.PublicKey
}

func (obj EvtUpdatePositionOperator) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_UpdatePositionOperator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `OldOperator`:
	if err = encoder.Encode(obj.OldOperator); err != nil {
		return fmt.Errorf("error while marshaling OldOperator:%w", err)
	}
	// Serialize `NewOperator`:
	if err = encoder.Encode(obj.NewOperator); err != nil {
		return fmt.Errorf("error while marshaling NewOperator:%w", err)
	}
	return nil
}

func (obj EvtUpdatePositionOperator) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtUpdatePositionOperator: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtUpdatePositionOperator) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_UpdatePositionOperator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_UpdatePositionOperator[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `OldOperator`:
	if err = decoder.Decode(&obj.OldOperator); err != nil {
		return fmt.Errorf("error while unmarshaling OldOperator:%w", err)
	}
	// Deserialize `NewOperator`:
	if err = decoder.Decode(&obj.NewOperator); err != nil {
		return fmt.Errorf("error while unmarshaling NewOperator:%w", err)
	}
	return nil
}

func (obj *EvtUpdatePositionOperator) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtUpdatePositionOperator: %w", err)
	}
	return nil
}

func UnmarshalEvtUpdatePositionOperator(buf []byte) (*EvtUpdatePositionOperator, error) {
	obj := new(EvtUpdatePositionOperator)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtUpdateRewardDuration struct {
	LbPair            solanago.PublicKey
	RewardIndex       uint64
	OldRewardDuration uint64
	NewRewardDuration uint64
}

func (obj EvtUpdateRewardDuration) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_UpdateRewardDuration[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `RewardIndex`:
	if err = encoder.Encode(obj.RewardIndex); err != nil {
		return fmt.Errorf("error while marshaling RewardIndex:%w", err)
	}
	// Serialize `OldRewardDuration`:
	if err = encoder.Encode(obj.OldRewardDuration); err != nil {
		return fmt.Errorf("error while marshaling OldRewardDuration:%w", err)
	}
	// Serialize `NewRewardDuration`:
	if err = encoder.Encode(obj.NewRewardDuration); err != nil {
		return fmt.Errorf("error while marshaling NewRewardDuration:%w", err)
	}
	return nil
}

func (obj EvtUpdateRewardDuration) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtUpdateRewardDuration: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtUpdateRewardDuration) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_UpdateRewardDuration[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_UpdateRewardDuration[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `RewardIndex`:
	if err = decoder.Decode(&obj.RewardIndex); err != nil {
		return fmt.Errorf("error while unmarshaling RewardIndex:%w", err)
	}
	// Deserialize `OldRewardDuration`:
	if err = decoder.Decode(&obj.OldRewardDuration); err != nil {
		return fmt.Errorf("error while unmarshaling OldRewardDuration:%w", err)
	}
	// Deserialize `NewRewardDuration`:
	if err = decoder.Decode(&obj.NewRewardDuration); err != nil {
		return fmt.Errorf("error while unmarshaling NewRewardDuration:%w", err)
	}
	return nil
}

func (obj *EvtUpdateRewardDuration) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtUpdateRewardDuration: %w", err)
	}
	return nil
}

func UnmarshalEvtUpdateRewardDuration(buf []byte) (*EvtUpdateRewardDuration, error) {
	obj := new(EvtUpdateRewardDuration)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtUpdateRewardFunder struct {
	LbPair      solanago.PublicKey
	RewardIndex uint64
	OldFunder   solanago.PublicKey
	NewFunder   solanago.PublicKey
}

func (obj EvtUpdateRewardFunder) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_UpdateRewardFunder[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `RewardIndex`:
	if err = encoder.Encode(obj.RewardIndex); err != nil {
		return fmt.Errorf("error while marshaling RewardIndex:%w", err)
	}
	// Serialize `OldFunder`:
	if err = encoder.Encode(obj.OldFunder); err != nil {
		return fmt.Errorf("error while marshaling OldFunder:%w", err)
	}
	// Serialize `NewFunder`:
	if err = encoder.Encode(obj.NewFunder); err != nil {
		return fmt.Errorf("error while marshaling NewFunder:%w", err)
	}
	return nil
}

func (obj EvtUpdateRewardFunder) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtUpdateRewardFunder: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtUpdateRewardFunder) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_UpdateRewardFunder[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_UpdateRewardFunder[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `RewardIndex`:
	if err = decoder.Decode(&obj.RewardIndex); err != nil {
		return fmt.Errorf("error while unmarshaling RewardIndex:%w", err)
	}
	// Deserialize `OldFunder`:
	if err = decoder.Decode(&obj.OldFunder); err != nil {
		return fmt.Errorf("error while unmarshaling OldFunder:%w", err)
	}
	// Deserialize `NewFunder`:
	if err = decoder.Decode(&obj.NewFunder); err != nil {
		return fmt.Errorf("error while unmarshaling NewFunder:%w", err)
	}
	return nil
}

func (obj *EvtUpdateRewardFunder) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtUpdateRewardFunder: %w", err)
	}
	return nil
}

func UnmarshalEvtUpdateRewardFunder(buf []byte) (*EvtUpdateRewardFunder, error) {
	obj := new(EvtUpdateRewardFunder)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtWithdrawIneligibleReward struct {
	LbPair     solanago.PublicKey
	RewardMint solanago.PublicKey
	Amount     uint64
}

func (obj EvtWithdrawIneligibleReward) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_WithdrawIneligibleReward[:], false)
	if err != nil {
		return err
	}
	// Serialize `LbPair`:
	if err = encoder.Encode(obj.LbPair); err != nil {
		return fmt.Errorf("error while marshaling LbPair:%w", err)
	}
	// Serialize `RewardMint`:
	if err = encoder.Encode(obj.RewardMint); err != nil {
		return fmt.Errorf("error while marshaling RewardMint:%w", err)
	}
	// Serialize `Amount`:
	if err = encoder.Encode(obj.Amount); err != nil {
		return fmt.Errorf("error while marshaling Amount:%w", err)
	}
	return nil
}

func (obj EvtWithdrawIneligibleReward) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtWithdrawIneligibleReward: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtWithdrawIneligibleReward) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_WithdrawIneligibleReward[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_WithdrawIneligibleReward[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LbPair`:
	if err = decoder.Decode(&obj.LbPair); err != nil {
		return fmt.Errorf("error while unmarshaling LbPair:%w", err)
	}
	// Deserialize `RewardMint`:
	if err = decoder.Decode(&obj.RewardMint); err != nil {
		return fmt.Errorf("error while unmarshaling RewardMint:%w", err)
	}
	// Deserialize `Amount`:
	if err = decoder.Decode(&obj.Amount); err != nil {
		return fmt.Errorf("error while unmarshaling Amount:%w", err)
	}
	return nil
}

func (obj *EvtWithdrawIneligibleReward) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtWithdrawIneligibleReward: %w", err)
	}
	return nil
}

func UnmarshalEvtWithdrawIneligibleReward(buf []byte) (*EvtWithdrawIneligibleReward, error) {
	obj := new(EvtWithdrawIneligibleReward)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
