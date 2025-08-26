package meteora_damm

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type EvtAddLiquidity struct {
	Pool         solanago.PublicKey
	Position     solanago.PublicKey
	Owner        solanago.PublicKey
	Params       AddLiquidityParameters
	TokenAAmount uint64
	TokenBAmount uint64
	TotalAmountA uint64
	TotalAmountB uint64
}

func (obj EvtAddLiquidity) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtAddLiquidity[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `Params`:
	if err = encoder.Encode(obj.Params); err != nil {
		return fmt.Errorf("error while marshaling Params:%w", err)
	}
	// Serialize `TokenAAmount`:
	if err = encoder.Encode(obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenAAmount:%w", err)
	}
	// Serialize `TokenBAmount`:
	if err = encoder.Encode(obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenBAmount:%w", err)
	}
	// Serialize `TotalAmountA`:
	if err = encoder.Encode(obj.TotalAmountA); err != nil {
		return fmt.Errorf("error while marshaling TotalAmountA:%w", err)
	}
	// Serialize `TotalAmountB`:
	if err = encoder.Encode(obj.TotalAmountB); err != nil {
		return fmt.Errorf("error while marshaling TotalAmountB:%w", err)
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
		if !discriminator.Equal(Event_EvtAddLiquidity[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtAddLiquidity[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `Params`:
	if err = decoder.Decode(&obj.Params); err != nil {
		return fmt.Errorf("error while unmarshaling Params:%w", err)
	}
	// Deserialize `TokenAAmount`:
	if err = decoder.Decode(&obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAAmount:%w", err)
	}
	// Deserialize `TokenBAmount`:
	if err = decoder.Decode(&obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBAmount:%w", err)
	}
	// Deserialize `TotalAmountA`:
	if err = decoder.Decode(&obj.TotalAmountA); err != nil {
		return fmt.Errorf("error while unmarshaling TotalAmountA:%w", err)
	}
	// Deserialize `TotalAmountB`:
	if err = decoder.Decode(&obj.TotalAmountB); err != nil {
		return fmt.Errorf("error while unmarshaling TotalAmountB:%w", err)
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

type EvtClaimPartnerFee struct {
	Pool         solanago.PublicKey
	TokenAAmount uint64
	TokenBAmount uint64
}

func (obj EvtClaimPartnerFee) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtClaimPartnerFee[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
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

func (obj EvtClaimPartnerFee) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtClaimPartnerFee: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtClaimPartnerFee) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtClaimPartnerFee[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtClaimPartnerFee[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
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

func (obj *EvtClaimPartnerFee) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtClaimPartnerFee: %w", err)
	}
	return nil
}

func UnmarshalEvtClaimPartnerFee(buf []byte) (*EvtClaimPartnerFee, error) {
	obj := new(EvtClaimPartnerFee)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtClaimPositionFee struct {
	Pool        solanago.PublicKey
	Position    solanago.PublicKey
	Owner       solanago.PublicKey
	FeeAClaimed uint64
	FeeBClaimed uint64
}

func (obj EvtClaimPositionFee) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtClaimPositionFee[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `FeeAClaimed`:
	if err = encoder.Encode(obj.FeeAClaimed); err != nil {
		return fmt.Errorf("error while marshaling FeeAClaimed:%w", err)
	}
	// Serialize `FeeBClaimed`:
	if err = encoder.Encode(obj.FeeBClaimed); err != nil {
		return fmt.Errorf("error while marshaling FeeBClaimed:%w", err)
	}
	return nil
}

func (obj EvtClaimPositionFee) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtClaimPositionFee: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtClaimPositionFee) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtClaimPositionFee[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtClaimPositionFee[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `FeeAClaimed`:
	if err = decoder.Decode(&obj.FeeAClaimed); err != nil {
		return fmt.Errorf("error while unmarshaling FeeAClaimed:%w", err)
	}
	// Deserialize `FeeBClaimed`:
	if err = decoder.Decode(&obj.FeeBClaimed); err != nil {
		return fmt.Errorf("error while unmarshaling FeeBClaimed:%w", err)
	}
	return nil
}

func (obj *EvtClaimPositionFee) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtClaimPositionFee: %w", err)
	}
	return nil
}

func UnmarshalEvtClaimPositionFee(buf []byte) (*EvtClaimPositionFee, error) {
	obj := new(EvtClaimPositionFee)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtClaimProtocolFee struct {
	Pool         solanago.PublicKey
	TokenAAmount uint64
	TokenBAmount uint64
}

func (obj EvtClaimProtocolFee) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtClaimProtocolFee[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
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

func (obj EvtClaimProtocolFee) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtClaimProtocolFee: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtClaimProtocolFee) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtClaimProtocolFee[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtClaimProtocolFee[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
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

func (obj *EvtClaimProtocolFee) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtClaimProtocolFee: %w", err)
	}
	return nil
}

func UnmarshalEvtClaimProtocolFee(buf []byte) (*EvtClaimProtocolFee, error) {
	obj := new(EvtClaimProtocolFee)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtClaimReward struct {
	Pool        solanago.PublicKey
	Position    solanago.PublicKey
	Owner       solanago.PublicKey
	MintReward  solanago.PublicKey
	RewardIndex uint8
	TotalReward uint64
}

func (obj EvtClaimReward) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtClaimReward[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `MintReward`:
	if err = encoder.Encode(obj.MintReward); err != nil {
		return fmt.Errorf("error while marshaling MintReward:%w", err)
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
		if !discriminator.Equal(Event_EvtClaimReward[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtClaimReward[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `MintReward`:
	if err = decoder.Decode(&obj.MintReward); err != nil {
		return fmt.Errorf("error while unmarshaling MintReward:%w", err)
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

// Close claim fee operator
type EvtCloseClaimFeeOperator struct {
	ClaimFeeOperator solanago.PublicKey
	Operator         solanago.PublicKey
}

func (obj EvtCloseClaimFeeOperator) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtCloseClaimFeeOperator[:], false)
	if err != nil {
		return err
	}
	// Serialize `ClaimFeeOperator`:
	if err = encoder.Encode(obj.ClaimFeeOperator); err != nil {
		return fmt.Errorf("error while marshaling ClaimFeeOperator:%w", err)
	}
	// Serialize `Operator`:
	if err = encoder.Encode(obj.Operator); err != nil {
		return fmt.Errorf("error while marshaling Operator:%w", err)
	}
	return nil
}

func (obj EvtCloseClaimFeeOperator) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtCloseClaimFeeOperator: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtCloseClaimFeeOperator) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtCloseClaimFeeOperator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtCloseClaimFeeOperator[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `ClaimFeeOperator`:
	if err = decoder.Decode(&obj.ClaimFeeOperator); err != nil {
		return fmt.Errorf("error while unmarshaling ClaimFeeOperator:%w", err)
	}
	// Deserialize `Operator`:
	if err = decoder.Decode(&obj.Operator); err != nil {
		return fmt.Errorf("error while unmarshaling Operator:%w", err)
	}
	return nil
}

func (obj *EvtCloseClaimFeeOperator) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtCloseClaimFeeOperator: %w", err)
	}
	return nil
}

func UnmarshalEvtCloseClaimFeeOperator(buf []byte) (*EvtCloseClaimFeeOperator, error) {
	obj := new(EvtCloseClaimFeeOperator)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Close config
type EvtCloseConfig struct {
	// Config pubkey
	Config solanago.PublicKey

	// admin pk
	Admin solanago.PublicKey
}

func (obj EvtCloseConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtCloseConfig[:], false)
	if err != nil {
		return err
	}
	// Serialize `Config`:
	if err = encoder.Encode(obj.Config); err != nil {
		return fmt.Errorf("error while marshaling Config:%w", err)
	}
	// Serialize `Admin`:
	if err = encoder.Encode(obj.Admin); err != nil {
		return fmt.Errorf("error while marshaling Admin:%w", err)
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
		if !discriminator.Equal(Event_EvtCloseConfig[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtCloseConfig[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Config`:
	if err = decoder.Decode(&obj.Config); err != nil {
		return fmt.Errorf("error while unmarshaling Config:%w", err)
	}
	// Deserialize `Admin`:
	if err = decoder.Decode(&obj.Admin); err != nil {
		return fmt.Errorf("error while unmarshaling Admin:%w", err)
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

type EvtClosePosition struct {
	Pool            solanago.PublicKey
	Owner           solanago.PublicKey
	Position        solanago.PublicKey
	PositionNftMint solanago.PublicKey
}

func (obj EvtClosePosition) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtClosePosition[:], false)
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
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `PositionNftMint`:
	if err = encoder.Encode(obj.PositionNftMint); err != nil {
		return fmt.Errorf("error while marshaling PositionNftMint:%w", err)
	}
	return nil
}

func (obj EvtClosePosition) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtClosePosition: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtClosePosition) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtClosePosition[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtClosePosition[:],
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
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `PositionNftMint`:
	if err = decoder.Decode(&obj.PositionNftMint); err != nil {
		return fmt.Errorf("error while unmarshaling PositionNftMint:%w", err)
	}
	return nil
}

func (obj *EvtClosePosition) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtClosePosition: %w", err)
	}
	return nil
}

func UnmarshalEvtClosePosition(buf []byte) (*EvtClosePosition, error) {
	obj := new(EvtClosePosition)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create claim fee operator
type EvtCreateClaimFeeOperator struct {
	Operator solanago.PublicKey
}

func (obj EvtCreateClaimFeeOperator) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtCreateClaimFeeOperator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Operator`:
	if err = encoder.Encode(obj.Operator); err != nil {
		return fmt.Errorf("error while marshaling Operator:%w", err)
	}
	return nil
}

func (obj EvtCreateClaimFeeOperator) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtCreateClaimFeeOperator: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtCreateClaimFeeOperator) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtCreateClaimFeeOperator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtCreateClaimFeeOperator[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Operator`:
	if err = decoder.Decode(&obj.Operator); err != nil {
		return fmt.Errorf("error while unmarshaling Operator:%w", err)
	}
	return nil
}

func (obj *EvtCreateClaimFeeOperator) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtCreateClaimFeeOperator: %w", err)
	}
	return nil
}

func UnmarshalEvtCreateClaimFeeOperator(buf []byte) (*EvtCreateClaimFeeOperator, error) {
	obj := new(EvtCreateClaimFeeOperator)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create static config
type EvtCreateConfig struct {
	PoolFees             PoolFeeParameters
	VaultConfigKey       solanago.PublicKey
	PoolCreatorAuthority solanago.PublicKey
	ActivationType       uint8
	SqrtMinPrice         binary.Uint128
	SqrtMaxPrice         binary.Uint128
	CollectFeeMode       uint8
	Index                uint64
	Config               solanago.PublicKey
}

func (obj EvtCreateConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtCreateConfig[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolFees`:
	if err = encoder.Encode(obj.PoolFees); err != nil {
		return fmt.Errorf("error while marshaling PoolFees:%w", err)
	}
	// Serialize `VaultConfigKey`:
	if err = encoder.Encode(obj.VaultConfigKey); err != nil {
		return fmt.Errorf("error while marshaling VaultConfigKey:%w", err)
	}
	// Serialize `PoolCreatorAuthority`:
	if err = encoder.Encode(obj.PoolCreatorAuthority); err != nil {
		return fmt.Errorf("error while marshaling PoolCreatorAuthority:%w", err)
	}
	// Serialize `ActivationType`:
	if err = encoder.Encode(obj.ActivationType); err != nil {
		return fmt.Errorf("error while marshaling ActivationType:%w", err)
	}
	// Serialize `SqrtMinPrice`:
	if err = encoder.Encode(obj.SqrtMinPrice); err != nil {
		return fmt.Errorf("error while marshaling SqrtMinPrice:%w", err)
	}
	// Serialize `SqrtMaxPrice`:
	if err = encoder.Encode(obj.SqrtMaxPrice); err != nil {
		return fmt.Errorf("error while marshaling SqrtMaxPrice:%w", err)
	}
	// Serialize `CollectFeeMode`:
	if err = encoder.Encode(obj.CollectFeeMode); err != nil {
		return fmt.Errorf("error while marshaling CollectFeeMode:%w", err)
	}
	// Serialize `Index`:
	if err = encoder.Encode(obj.Index); err != nil {
		return fmt.Errorf("error while marshaling Index:%w", err)
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
		if !discriminator.Equal(Event_EvtCreateConfig[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtCreateConfig[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolFees`:
	if err = decoder.Decode(&obj.PoolFees); err != nil {
		return fmt.Errorf("error while unmarshaling PoolFees:%w", err)
	}
	// Deserialize `VaultConfigKey`:
	if err = decoder.Decode(&obj.VaultConfigKey); err != nil {
		return fmt.Errorf("error while unmarshaling VaultConfigKey:%w", err)
	}
	// Deserialize `PoolCreatorAuthority`:
	if err = decoder.Decode(&obj.PoolCreatorAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling PoolCreatorAuthority:%w", err)
	}
	// Deserialize `ActivationType`:
	if err = decoder.Decode(&obj.ActivationType); err != nil {
		return fmt.Errorf("error while unmarshaling ActivationType:%w", err)
	}
	// Deserialize `SqrtMinPrice`:
	if err = decoder.Decode(&obj.SqrtMinPrice); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtMinPrice:%w", err)
	}
	// Deserialize `SqrtMaxPrice`:
	if err = decoder.Decode(&obj.SqrtMaxPrice); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtMaxPrice:%w", err)
	}
	// Deserialize `CollectFeeMode`:
	if err = decoder.Decode(&obj.CollectFeeMode); err != nil {
		return fmt.Errorf("error while unmarshaling CollectFeeMode:%w", err)
	}
	// Deserialize `Index`:
	if err = decoder.Decode(&obj.Index); err != nil {
		return fmt.Errorf("error while unmarshaling Index:%w", err)
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

// Create dynamic config
type EvtCreateDynamicConfig struct {
	Config               solanago.PublicKey
	PoolCreatorAuthority solanago.PublicKey
	Index                uint64
}

func (obj EvtCreateDynamicConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtCreateDynamicConfig[:], false)
	if err != nil {
		return err
	}
	// Serialize `Config`:
	if err = encoder.Encode(obj.Config); err != nil {
		return fmt.Errorf("error while marshaling Config:%w", err)
	}
	// Serialize `PoolCreatorAuthority`:
	if err = encoder.Encode(obj.PoolCreatorAuthority); err != nil {
		return fmt.Errorf("error while marshaling PoolCreatorAuthority:%w", err)
	}
	// Serialize `Index`:
	if err = encoder.Encode(obj.Index); err != nil {
		return fmt.Errorf("error while marshaling Index:%w", err)
	}
	return nil
}

func (obj EvtCreateDynamicConfig) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtCreateDynamicConfig: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtCreateDynamicConfig) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtCreateDynamicConfig[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtCreateDynamicConfig[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Config`:
	if err = decoder.Decode(&obj.Config); err != nil {
		return fmt.Errorf("error while unmarshaling Config:%w", err)
	}
	// Deserialize `PoolCreatorAuthority`:
	if err = decoder.Decode(&obj.PoolCreatorAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling PoolCreatorAuthority:%w", err)
	}
	// Deserialize `Index`:
	if err = decoder.Decode(&obj.Index); err != nil {
		return fmt.Errorf("error while unmarshaling Index:%w", err)
	}
	return nil
}

func (obj *EvtCreateDynamicConfig) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtCreateDynamicConfig: %w", err)
	}
	return nil
}

func UnmarshalEvtCreateDynamicConfig(buf []byte) (*EvtCreateDynamicConfig, error) {
	obj := new(EvtCreateDynamicConfig)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtCreatePosition struct {
	Pool            solanago.PublicKey
	Owner           solanago.PublicKey
	Position        solanago.PublicKey
	PositionNftMint solanago.PublicKey
}

func (obj EvtCreatePosition) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtCreatePosition[:], false)
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
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `PositionNftMint`:
	if err = encoder.Encode(obj.PositionNftMint); err != nil {
		return fmt.Errorf("error while marshaling PositionNftMint:%w", err)
	}
	return nil
}

func (obj EvtCreatePosition) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtCreatePosition: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtCreatePosition) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtCreatePosition[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtCreatePosition[:],
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
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `PositionNftMint`:
	if err = decoder.Decode(&obj.PositionNftMint); err != nil {
		return fmt.Errorf("error while unmarshaling PositionNftMint:%w", err)
	}
	return nil
}

func (obj *EvtCreatePosition) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtCreatePosition: %w", err)
	}
	return nil
}

func UnmarshalEvtCreatePosition(buf []byte) (*EvtCreatePosition, error) {
	obj := new(EvtCreatePosition)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create token badge
type EvtCreateTokenBadge struct {
	TokenMint solanago.PublicKey
}

func (obj EvtCreateTokenBadge) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtCreateTokenBadge[:], false)
	if err != nil {
		return err
	}
	// Serialize `TokenMint`:
	if err = encoder.Encode(obj.TokenMint); err != nil {
		return fmt.Errorf("error while marshaling TokenMint:%w", err)
	}
	return nil
}

func (obj EvtCreateTokenBadge) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtCreateTokenBadge: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtCreateTokenBadge) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtCreateTokenBadge[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtCreateTokenBadge[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TokenMint`:
	if err = decoder.Decode(&obj.TokenMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMint:%w", err)
	}
	return nil
}

func (obj *EvtCreateTokenBadge) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtCreateTokenBadge: %w", err)
	}
	return nil
}

func UnmarshalEvtCreateTokenBadge(buf []byte) (*EvtCreateTokenBadge, error) {
	obj := new(EvtCreateTokenBadge)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtFundReward struct {
	Pool                        solanago.PublicKey
	Funder                      solanago.PublicKey
	MintReward                  solanago.PublicKey
	RewardIndex                 uint8
	Amount                      uint64
	TransferFeeExcludedAmountIn uint64
	RewardDurationEnd           uint64
	PreRewardRate               binary.Uint128
	PostRewardRate              binary.Uint128
}

func (obj EvtFundReward) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtFundReward[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Funder`:
	if err = encoder.Encode(obj.Funder); err != nil {
		return fmt.Errorf("error while marshaling Funder:%w", err)
	}
	// Serialize `MintReward`:
	if err = encoder.Encode(obj.MintReward); err != nil {
		return fmt.Errorf("error while marshaling MintReward:%w", err)
	}
	// Serialize `RewardIndex`:
	if err = encoder.Encode(obj.RewardIndex); err != nil {
		return fmt.Errorf("error while marshaling RewardIndex:%w", err)
	}
	// Serialize `Amount`:
	if err = encoder.Encode(obj.Amount); err != nil {
		return fmt.Errorf("error while marshaling Amount:%w", err)
	}
	// Serialize `TransferFeeExcludedAmountIn`:
	if err = encoder.Encode(obj.TransferFeeExcludedAmountIn); err != nil {
		return fmt.Errorf("error while marshaling TransferFeeExcludedAmountIn:%w", err)
	}
	// Serialize `RewardDurationEnd`:
	if err = encoder.Encode(obj.RewardDurationEnd); err != nil {
		return fmt.Errorf("error while marshaling RewardDurationEnd:%w", err)
	}
	// Serialize `PreRewardRate`:
	if err = encoder.Encode(obj.PreRewardRate); err != nil {
		return fmt.Errorf("error while marshaling PreRewardRate:%w", err)
	}
	// Serialize `PostRewardRate`:
	if err = encoder.Encode(obj.PostRewardRate); err != nil {
		return fmt.Errorf("error while marshaling PostRewardRate:%w", err)
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
		if !discriminator.Equal(Event_EvtFundReward[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtFundReward[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Funder`:
	if err = decoder.Decode(&obj.Funder); err != nil {
		return fmt.Errorf("error while unmarshaling Funder:%w", err)
	}
	// Deserialize `MintReward`:
	if err = decoder.Decode(&obj.MintReward); err != nil {
		return fmt.Errorf("error while unmarshaling MintReward:%w", err)
	}
	// Deserialize `RewardIndex`:
	if err = decoder.Decode(&obj.RewardIndex); err != nil {
		return fmt.Errorf("error while unmarshaling RewardIndex:%w", err)
	}
	// Deserialize `Amount`:
	if err = decoder.Decode(&obj.Amount); err != nil {
		return fmt.Errorf("error while unmarshaling Amount:%w", err)
	}
	// Deserialize `TransferFeeExcludedAmountIn`:
	if err = decoder.Decode(&obj.TransferFeeExcludedAmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling TransferFeeExcludedAmountIn:%w", err)
	}
	// Deserialize `RewardDurationEnd`:
	if err = decoder.Decode(&obj.RewardDurationEnd); err != nil {
		return fmt.Errorf("error while unmarshaling RewardDurationEnd:%w", err)
	}
	// Deserialize `PreRewardRate`:
	if err = decoder.Decode(&obj.PreRewardRate); err != nil {
		return fmt.Errorf("error while unmarshaling PreRewardRate:%w", err)
	}
	// Deserialize `PostRewardRate`:
	if err = decoder.Decode(&obj.PostRewardRate); err != nil {
		return fmt.Errorf("error while unmarshaling PostRewardRate:%w", err)
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

type EvtInitializePool struct {
	Pool            solanago.PublicKey
	TokenAMint      solanago.PublicKey
	TokenBMint      solanago.PublicKey
	Creator         solanago.PublicKey
	Payer           solanago.PublicKey
	AlphaVault      solanago.PublicKey
	PoolFees        PoolFeeParameters
	SqrtMinPrice    binary.Uint128
	SqrtMaxPrice    binary.Uint128
	ActivationType  uint8
	CollectFeeMode  uint8
	Liquidity       binary.Uint128
	SqrtPrice       binary.Uint128
	ActivationPoint uint64
	TokenAFlag      uint8
	TokenBFlag      uint8
	TokenAAmount    uint64
	TokenBAmount    uint64
	TotalAmountA    uint64
	TotalAmountB    uint64
	PoolType        uint8
}

func (obj EvtInitializePool) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtInitializePool[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `TokenAMint`:
	if err = encoder.Encode(obj.TokenAMint); err != nil {
		return fmt.Errorf("error while marshaling TokenAMint:%w", err)
	}
	// Serialize `TokenBMint`:
	if err = encoder.Encode(obj.TokenBMint); err != nil {
		return fmt.Errorf("error while marshaling TokenBMint:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
	}
	// Serialize `Payer`:
	if err = encoder.Encode(obj.Payer); err != nil {
		return fmt.Errorf("error while marshaling Payer:%w", err)
	}
	// Serialize `AlphaVault`:
	if err = encoder.Encode(obj.AlphaVault); err != nil {
		return fmt.Errorf("error while marshaling AlphaVault:%w", err)
	}
	// Serialize `PoolFees`:
	if err = encoder.Encode(obj.PoolFees); err != nil {
		return fmt.Errorf("error while marshaling PoolFees:%w", err)
	}
	// Serialize `SqrtMinPrice`:
	if err = encoder.Encode(obj.SqrtMinPrice); err != nil {
		return fmt.Errorf("error while marshaling SqrtMinPrice:%w", err)
	}
	// Serialize `SqrtMaxPrice`:
	if err = encoder.Encode(obj.SqrtMaxPrice); err != nil {
		return fmt.Errorf("error while marshaling SqrtMaxPrice:%w", err)
	}
	// Serialize `ActivationType`:
	if err = encoder.Encode(obj.ActivationType); err != nil {
		return fmt.Errorf("error while marshaling ActivationType:%w", err)
	}
	// Serialize `CollectFeeMode`:
	if err = encoder.Encode(obj.CollectFeeMode); err != nil {
		return fmt.Errorf("error while marshaling CollectFeeMode:%w", err)
	}
	// Serialize `Liquidity`:
	if err = encoder.Encode(obj.Liquidity); err != nil {
		return fmt.Errorf("error while marshaling Liquidity:%w", err)
	}
	// Serialize `SqrtPrice`:
	if err = encoder.Encode(obj.SqrtPrice); err != nil {
		return fmt.Errorf("error while marshaling SqrtPrice:%w", err)
	}
	// Serialize `ActivationPoint`:
	if err = encoder.Encode(obj.ActivationPoint); err != nil {
		return fmt.Errorf("error while marshaling ActivationPoint:%w", err)
	}
	// Serialize `TokenAFlag`:
	if err = encoder.Encode(obj.TokenAFlag); err != nil {
		return fmt.Errorf("error while marshaling TokenAFlag:%w", err)
	}
	// Serialize `TokenBFlag`:
	if err = encoder.Encode(obj.TokenBFlag); err != nil {
		return fmt.Errorf("error while marshaling TokenBFlag:%w", err)
	}
	// Serialize `TokenAAmount`:
	if err = encoder.Encode(obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenAAmount:%w", err)
	}
	// Serialize `TokenBAmount`:
	if err = encoder.Encode(obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenBAmount:%w", err)
	}
	// Serialize `TotalAmountA`:
	if err = encoder.Encode(obj.TotalAmountA); err != nil {
		return fmt.Errorf("error while marshaling TotalAmountA:%w", err)
	}
	// Serialize `TotalAmountB`:
	if err = encoder.Encode(obj.TotalAmountB); err != nil {
		return fmt.Errorf("error while marshaling TotalAmountB:%w", err)
	}
	// Serialize `PoolType`:
	if err = encoder.Encode(obj.PoolType); err != nil {
		return fmt.Errorf("error while marshaling PoolType:%w", err)
	}
	return nil
}

func (obj EvtInitializePool) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtInitializePool: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtInitializePool) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtInitializePool[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtInitializePool[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `TokenAMint`:
	if err = decoder.Decode(&obj.TokenAMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAMint:%w", err)
	}
	// Deserialize `TokenBMint`:
	if err = decoder.Decode(&obj.TokenBMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBMint:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
	}
	// Deserialize `Payer`:
	if err = decoder.Decode(&obj.Payer); err != nil {
		return fmt.Errorf("error while unmarshaling Payer:%w", err)
	}
	// Deserialize `AlphaVault`:
	if err = decoder.Decode(&obj.AlphaVault); err != nil {
		return fmt.Errorf("error while unmarshaling AlphaVault:%w", err)
	}
	// Deserialize `PoolFees`:
	if err = decoder.Decode(&obj.PoolFees); err != nil {
		return fmt.Errorf("error while unmarshaling PoolFees:%w", err)
	}
	// Deserialize `SqrtMinPrice`:
	if err = decoder.Decode(&obj.SqrtMinPrice); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtMinPrice:%w", err)
	}
	// Deserialize `SqrtMaxPrice`:
	if err = decoder.Decode(&obj.SqrtMaxPrice); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtMaxPrice:%w", err)
	}
	// Deserialize `ActivationType`:
	if err = decoder.Decode(&obj.ActivationType); err != nil {
		return fmt.Errorf("error while unmarshaling ActivationType:%w", err)
	}
	// Deserialize `CollectFeeMode`:
	if err = decoder.Decode(&obj.CollectFeeMode); err != nil {
		return fmt.Errorf("error while unmarshaling CollectFeeMode:%w", err)
	}
	// Deserialize `Liquidity`:
	if err = decoder.Decode(&obj.Liquidity); err != nil {
		return fmt.Errorf("error while unmarshaling Liquidity:%w", err)
	}
	// Deserialize `SqrtPrice`:
	if err = decoder.Decode(&obj.SqrtPrice); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtPrice:%w", err)
	}
	// Deserialize `ActivationPoint`:
	if err = decoder.Decode(&obj.ActivationPoint); err != nil {
		return fmt.Errorf("error while unmarshaling ActivationPoint:%w", err)
	}
	// Deserialize `TokenAFlag`:
	if err = decoder.Decode(&obj.TokenAFlag); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAFlag:%w", err)
	}
	// Deserialize `TokenBFlag`:
	if err = decoder.Decode(&obj.TokenBFlag); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBFlag:%w", err)
	}
	// Deserialize `TokenAAmount`:
	if err = decoder.Decode(&obj.TokenAAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAAmount:%w", err)
	}
	// Deserialize `TokenBAmount`:
	if err = decoder.Decode(&obj.TokenBAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBAmount:%w", err)
	}
	// Deserialize `TotalAmountA`:
	if err = decoder.Decode(&obj.TotalAmountA); err != nil {
		return fmt.Errorf("error while unmarshaling TotalAmountA:%w", err)
	}
	// Deserialize `TotalAmountB`:
	if err = decoder.Decode(&obj.TotalAmountB); err != nil {
		return fmt.Errorf("error while unmarshaling TotalAmountB:%w", err)
	}
	// Deserialize `PoolType`:
	if err = decoder.Decode(&obj.PoolType); err != nil {
		return fmt.Errorf("error while unmarshaling PoolType:%w", err)
	}
	return nil
}

func (obj *EvtInitializePool) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtInitializePool: %w", err)
	}
	return nil
}

func UnmarshalEvtInitializePool(buf []byte) (*EvtInitializePool, error) {
	obj := new(EvtInitializePool)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtInitializeReward struct {
	Pool           solanago.PublicKey
	RewardMint     solanago.PublicKey
	Funder         solanago.PublicKey
	Creator        solanago.PublicKey
	RewardIndex    uint8
	RewardDuration uint64
}

func (obj EvtInitializeReward) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtInitializeReward[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `RewardMint`:
	if err = encoder.Encode(obj.RewardMint); err != nil {
		return fmt.Errorf("error while marshaling RewardMint:%w", err)
	}
	// Serialize `Funder`:
	if err = encoder.Encode(obj.Funder); err != nil {
		return fmt.Errorf("error while marshaling Funder:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
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
		if !discriminator.Equal(Event_EvtInitializeReward[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtInitializeReward[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `RewardMint`:
	if err = decoder.Decode(&obj.RewardMint); err != nil {
		return fmt.Errorf("error while unmarshaling RewardMint:%w", err)
	}
	// Deserialize `Funder`:
	if err = decoder.Decode(&obj.Funder); err != nil {
		return fmt.Errorf("error while unmarshaling Funder:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
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

type EvtLockPosition struct {
	Pool                 solanago.PublicKey
	Position             solanago.PublicKey
	Owner                solanago.PublicKey
	Vesting              solanago.PublicKey
	CliffPoint           uint64
	PeriodFrequency      uint64
	CliffUnlockLiquidity binary.Uint128
	LiquidityPerPeriod   binary.Uint128
	NumberOfPeriod       uint16
}

func (obj EvtLockPosition) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtLockPosition[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `Vesting`:
	if err = encoder.Encode(obj.Vesting); err != nil {
		return fmt.Errorf("error while marshaling Vesting:%w", err)
	}
	// Serialize `CliffPoint`:
	if err = encoder.Encode(obj.CliffPoint); err != nil {
		return fmt.Errorf("error while marshaling CliffPoint:%w", err)
	}
	// Serialize `PeriodFrequency`:
	if err = encoder.Encode(obj.PeriodFrequency); err != nil {
		return fmt.Errorf("error while marshaling PeriodFrequency:%w", err)
	}
	// Serialize `CliffUnlockLiquidity`:
	if err = encoder.Encode(obj.CliffUnlockLiquidity); err != nil {
		return fmt.Errorf("error while marshaling CliffUnlockLiquidity:%w", err)
	}
	// Serialize `LiquidityPerPeriod`:
	if err = encoder.Encode(obj.LiquidityPerPeriod); err != nil {
		return fmt.Errorf("error while marshaling LiquidityPerPeriod:%w", err)
	}
	// Serialize `NumberOfPeriod`:
	if err = encoder.Encode(obj.NumberOfPeriod); err != nil {
		return fmt.Errorf("error while marshaling NumberOfPeriod:%w", err)
	}
	return nil
}

func (obj EvtLockPosition) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtLockPosition: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtLockPosition) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtLockPosition[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtLockPosition[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `Vesting`:
	if err = decoder.Decode(&obj.Vesting); err != nil {
		return fmt.Errorf("error while unmarshaling Vesting:%w", err)
	}
	// Deserialize `CliffPoint`:
	if err = decoder.Decode(&obj.CliffPoint); err != nil {
		return fmt.Errorf("error while unmarshaling CliffPoint:%w", err)
	}
	// Deserialize `PeriodFrequency`:
	if err = decoder.Decode(&obj.PeriodFrequency); err != nil {
		return fmt.Errorf("error while unmarshaling PeriodFrequency:%w", err)
	}
	// Deserialize `CliffUnlockLiquidity`:
	if err = decoder.Decode(&obj.CliffUnlockLiquidity); err != nil {
		return fmt.Errorf("error while unmarshaling CliffUnlockLiquidity:%w", err)
	}
	// Deserialize `LiquidityPerPeriod`:
	if err = decoder.Decode(&obj.LiquidityPerPeriod); err != nil {
		return fmt.Errorf("error while unmarshaling LiquidityPerPeriod:%w", err)
	}
	// Deserialize `NumberOfPeriod`:
	if err = decoder.Decode(&obj.NumberOfPeriod); err != nil {
		return fmt.Errorf("error while unmarshaling NumberOfPeriod:%w", err)
	}
	return nil
}

func (obj *EvtLockPosition) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtLockPosition: %w", err)
	}
	return nil
}

func UnmarshalEvtLockPosition(buf []byte) (*EvtLockPosition, error) {
	obj := new(EvtLockPosition)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtPermanentLockPosition struct {
	Pool                          solanago.PublicKey
	Position                      solanago.PublicKey
	LockLiquidityAmount           binary.Uint128
	TotalPermanentLockedLiquidity binary.Uint128
}

func (obj EvtPermanentLockPosition) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtPermanentLockPosition[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `LockLiquidityAmount`:
	if err = encoder.Encode(obj.LockLiquidityAmount); err != nil {
		return fmt.Errorf("error while marshaling LockLiquidityAmount:%w", err)
	}
	// Serialize `TotalPermanentLockedLiquidity`:
	if err = encoder.Encode(obj.TotalPermanentLockedLiquidity); err != nil {
		return fmt.Errorf("error while marshaling TotalPermanentLockedLiquidity:%w", err)
	}
	return nil
}

func (obj EvtPermanentLockPosition) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtPermanentLockPosition: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtPermanentLockPosition) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtPermanentLockPosition[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtPermanentLockPosition[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `LockLiquidityAmount`:
	if err = decoder.Decode(&obj.LockLiquidityAmount); err != nil {
		return fmt.Errorf("error while unmarshaling LockLiquidityAmount:%w", err)
	}
	// Deserialize `TotalPermanentLockedLiquidity`:
	if err = decoder.Decode(&obj.TotalPermanentLockedLiquidity); err != nil {
		return fmt.Errorf("error while unmarshaling TotalPermanentLockedLiquidity:%w", err)
	}
	return nil
}

func (obj *EvtPermanentLockPosition) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtPermanentLockPosition: %w", err)
	}
	return nil
}

func UnmarshalEvtPermanentLockPosition(buf []byte) (*EvtPermanentLockPosition, error) {
	obj := new(EvtPermanentLockPosition)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtRemoveLiquidity struct {
	Pool         solanago.PublicKey
	Position     solanago.PublicKey
	Owner        solanago.PublicKey
	Params       RemoveLiquidityParameters
	TokenAAmount uint64
	TokenBAmount uint64
}

func (obj EvtRemoveLiquidity) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtRemoveLiquidity[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `Params`:
	if err = encoder.Encode(obj.Params); err != nil {
		return fmt.Errorf("error while marshaling Params:%w", err)
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
		if !discriminator.Equal(Event_EvtRemoveLiquidity[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtRemoveLiquidity[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `Params`:
	if err = decoder.Decode(&obj.Params); err != nil {
		return fmt.Errorf("error while unmarshaling Params:%w", err)
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

type EvtSetPoolStatus struct {
	Pool   solanago.PublicKey
	Status uint8
}

func (obj EvtSetPoolStatus) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtSetPoolStatus[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Status`:
	if err = encoder.Encode(obj.Status); err != nil {
		return fmt.Errorf("error while marshaling Status:%w", err)
	}
	return nil
}

func (obj EvtSetPoolStatus) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtSetPoolStatus: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtSetPoolStatus) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtSetPoolStatus[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtSetPoolStatus[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Status`:
	if err = decoder.Decode(&obj.Status); err != nil {
		return fmt.Errorf("error while unmarshaling Status:%w", err)
	}
	return nil
}

func (obj *EvtSetPoolStatus) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtSetPoolStatus: %w", err)
	}
	return nil
}

func UnmarshalEvtSetPoolStatus(buf []byte) (*EvtSetPoolStatus, error) {
	obj := new(EvtSetPoolStatus)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtSplitPosition struct {
	Pool                    solanago.PublicKey
	FirstOwner              solanago.PublicKey
	SecondOwner             solanago.PublicKey
	FirstPosition           solanago.PublicKey
	SecondPosition          solanago.PublicKey
	CurrentSqrtPrice        binary.Uint128
	AmountSplits            SplitAmountInfo
	FirstPositionInfo       SplitPositionInfo
	SecondPositionInfo      SplitPositionInfo
	SplitPositionParameters SplitPositionParameters
}

func (obj EvtSplitPosition) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtSplitPosition[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `FirstOwner`:
	if err = encoder.Encode(obj.FirstOwner); err != nil {
		return fmt.Errorf("error while marshaling FirstOwner:%w", err)
	}
	// Serialize `SecondOwner`:
	if err = encoder.Encode(obj.SecondOwner); err != nil {
		return fmt.Errorf("error while marshaling SecondOwner:%w", err)
	}
	// Serialize `FirstPosition`:
	if err = encoder.Encode(obj.FirstPosition); err != nil {
		return fmt.Errorf("error while marshaling FirstPosition:%w", err)
	}
	// Serialize `SecondPosition`:
	if err = encoder.Encode(obj.SecondPosition); err != nil {
		return fmt.Errorf("error while marshaling SecondPosition:%w", err)
	}
	// Serialize `CurrentSqrtPrice`:
	if err = encoder.Encode(obj.CurrentSqrtPrice); err != nil {
		return fmt.Errorf("error while marshaling CurrentSqrtPrice:%w", err)
	}
	// Serialize `AmountSplits`:
	if err = encoder.Encode(obj.AmountSplits); err != nil {
		return fmt.Errorf("error while marshaling AmountSplits:%w", err)
	}
	// Serialize `FirstPositionInfo`:
	if err = encoder.Encode(obj.FirstPositionInfo); err != nil {
		return fmt.Errorf("error while marshaling FirstPositionInfo:%w", err)
	}
	// Serialize `SecondPositionInfo`:
	if err = encoder.Encode(obj.SecondPositionInfo); err != nil {
		return fmt.Errorf("error while marshaling SecondPositionInfo:%w", err)
	}
	// Serialize `SplitPositionParameters`:
	if err = encoder.Encode(obj.SplitPositionParameters); err != nil {
		return fmt.Errorf("error while marshaling SplitPositionParameters:%w", err)
	}
	return nil
}

func (obj EvtSplitPosition) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtSplitPosition: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtSplitPosition) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtSplitPosition[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtSplitPosition[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `FirstOwner`:
	if err = decoder.Decode(&obj.FirstOwner); err != nil {
		return fmt.Errorf("error while unmarshaling FirstOwner:%w", err)
	}
	// Deserialize `SecondOwner`:
	if err = decoder.Decode(&obj.SecondOwner); err != nil {
		return fmt.Errorf("error while unmarshaling SecondOwner:%w", err)
	}
	// Deserialize `FirstPosition`:
	if err = decoder.Decode(&obj.FirstPosition); err != nil {
		return fmt.Errorf("error while unmarshaling FirstPosition:%w", err)
	}
	// Deserialize `SecondPosition`:
	if err = decoder.Decode(&obj.SecondPosition); err != nil {
		return fmt.Errorf("error while unmarshaling SecondPosition:%w", err)
	}
	// Deserialize `CurrentSqrtPrice`:
	if err = decoder.Decode(&obj.CurrentSqrtPrice); err != nil {
		return fmt.Errorf("error while unmarshaling CurrentSqrtPrice:%w", err)
	}
	// Deserialize `AmountSplits`:
	if err = decoder.Decode(&obj.AmountSplits); err != nil {
		return fmt.Errorf("error while unmarshaling AmountSplits:%w", err)
	}
	// Deserialize `FirstPositionInfo`:
	if err = decoder.Decode(&obj.FirstPositionInfo); err != nil {
		return fmt.Errorf("error while unmarshaling FirstPositionInfo:%w", err)
	}
	// Deserialize `SecondPositionInfo`:
	if err = decoder.Decode(&obj.SecondPositionInfo); err != nil {
		return fmt.Errorf("error while unmarshaling SecondPositionInfo:%w", err)
	}
	// Deserialize `SplitPositionParameters`:
	if err = decoder.Decode(&obj.SplitPositionParameters); err != nil {
		return fmt.Errorf("error while unmarshaling SplitPositionParameters:%w", err)
	}
	return nil
}

func (obj *EvtSplitPosition) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtSplitPosition: %w", err)
	}
	return nil
}

func UnmarshalEvtSplitPosition(buf []byte) (*EvtSplitPosition, error) {
	obj := new(EvtSplitPosition)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtSwap struct {
	Pool             solanago.PublicKey
	TradeDirection   uint8
	HasReferral      bool
	Params           SwapParameters
	SwapResult       SwapResult
	ActualAmountIn   uint64
	CurrentTimestamp uint64
}

func (obj EvtSwap) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtSwap[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `TradeDirection`:
	if err = encoder.Encode(obj.TradeDirection); err != nil {
		return fmt.Errorf("error while marshaling TradeDirection:%w", err)
	}
	// Serialize `HasReferral`:
	if err = encoder.Encode(obj.HasReferral); err != nil {
		return fmt.Errorf("error while marshaling HasReferral:%w", err)
	}
	// Serialize `Params`:
	if err = encoder.Encode(obj.Params); err != nil {
		return fmt.Errorf("error while marshaling Params:%w", err)
	}
	// Serialize `SwapResult`:
	if err = encoder.Encode(obj.SwapResult); err != nil {
		return fmt.Errorf("error while marshaling SwapResult:%w", err)
	}
	// Serialize `ActualAmountIn`:
	if err = encoder.Encode(obj.ActualAmountIn); err != nil {
		return fmt.Errorf("error while marshaling ActualAmountIn:%w", err)
	}
	// Serialize `CurrentTimestamp`:
	if err = encoder.Encode(obj.CurrentTimestamp); err != nil {
		return fmt.Errorf("error while marshaling CurrentTimestamp:%w", err)
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
		if !discriminator.Equal(Event_EvtSwap[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtSwap[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `TradeDirection`:
	if err = decoder.Decode(&obj.TradeDirection); err != nil {
		return fmt.Errorf("error while unmarshaling TradeDirection:%w", err)
	}
	// Deserialize `HasReferral`:
	if err = decoder.Decode(&obj.HasReferral); err != nil {
		return fmt.Errorf("error while unmarshaling HasReferral:%w", err)
	}
	// Deserialize `Params`:
	if err = decoder.Decode(&obj.Params); err != nil {
		return fmt.Errorf("error while unmarshaling Params:%w", err)
	}
	// Deserialize `SwapResult`:
	if err = decoder.Decode(&obj.SwapResult); err != nil {
		return fmt.Errorf("error while unmarshaling SwapResult:%w", err)
	}
	// Deserialize `ActualAmountIn`:
	if err = decoder.Decode(&obj.ActualAmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling ActualAmountIn:%w", err)
	}
	// Deserialize `CurrentTimestamp`:
	if err = decoder.Decode(&obj.CurrentTimestamp); err != nil {
		return fmt.Errorf("error while unmarshaling CurrentTimestamp:%w", err)
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

type EvtUpdateRewardDuration struct {
	Pool              solanago.PublicKey
	RewardIndex       uint8
	OldRewardDuration uint64
	NewRewardDuration uint64
}

func (obj EvtUpdateRewardDuration) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtUpdateRewardDuration[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
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
		if !discriminator.Equal(Event_EvtUpdateRewardDuration[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtUpdateRewardDuration[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
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
	Pool        solanago.PublicKey
	RewardIndex uint8
	OldFunder   solanago.PublicKey
	NewFunder   solanago.PublicKey
}

func (obj EvtUpdateRewardFunder) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtUpdateRewardFunder[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
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
		if !discriminator.Equal(Event_EvtUpdateRewardFunder[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtUpdateRewardFunder[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
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
	Pool       solanago.PublicKey
	RewardMint solanago.PublicKey
	Amount     uint64
}

func (obj EvtWithdrawIneligibleReward) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtWithdrawIneligibleReward[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
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
		if !discriminator.Equal(Event_EvtWithdrawIneligibleReward[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtWithdrawIneligibleReward[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
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
