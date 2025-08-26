package meteora_vault

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type EvtAddLiquidity struct {
	LpMintAmount uint64
	TokenAmount  uint64
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
	// Serialize `TokenAmount`:
	if err = encoder.Encode(obj.TokenAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenAmount:%w", err)
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
	// Deserialize `TokenAmount`:
	if err = decoder.Decode(&obj.TokenAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAmount:%w", err)
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
	LpUnmintAmount uint64
	TokenAmount    uint64
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
	// Serialize `TokenAmount`:
	if err = encoder.Encode(obj.TokenAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenAmount:%w", err)
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
	// Deserialize `TokenAmount`:
	if err = decoder.Decode(&obj.TokenAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAmount:%w", err)
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

type EvtStrategyDeposit struct {
	StrategyType StrategyType
	TokenAmount  uint64
}

func (obj EvtStrategyDeposit) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_StrategyDeposit[:], false)
	if err != nil {
		return err
	}
	// Serialize `StrategyType`:
	if err = encoder.Encode(obj.StrategyType); err != nil {
		return fmt.Errorf("error while marshaling StrategyType:%w", err)
	}
	// Serialize `TokenAmount`:
	if err = encoder.Encode(obj.TokenAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenAmount:%w", err)
	}
	return nil
}

func (obj EvtStrategyDeposit) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtStrategyDeposit: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtStrategyDeposit) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_StrategyDeposit[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_StrategyDeposit[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `StrategyType`:
	if err = decoder.Decode(&obj.StrategyType); err != nil {
		return fmt.Errorf("error while unmarshaling StrategyType:%w", err)
	}
	// Deserialize `TokenAmount`:
	if err = decoder.Decode(&obj.TokenAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAmount:%w", err)
	}
	return nil
}

func (obj *EvtStrategyDeposit) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtStrategyDeposit: %w", err)
	}
	return nil
}

func UnmarshalEvtStrategyDeposit(buf []byte) (*EvtStrategyDeposit, error) {
	obj := new(EvtStrategyDeposit)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtStrategyWithdraw struct {
	StrategyType         StrategyType
	CollateralAmount     uint64
	EstimatedTokenAmount uint64
}

func (obj EvtStrategyWithdraw) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_StrategyWithdraw[:], false)
	if err != nil {
		return err
	}
	// Serialize `StrategyType`:
	if err = encoder.Encode(obj.StrategyType); err != nil {
		return fmt.Errorf("error while marshaling StrategyType:%w", err)
	}
	// Serialize `CollateralAmount`:
	if err = encoder.Encode(obj.CollateralAmount); err != nil {
		return fmt.Errorf("error while marshaling CollateralAmount:%w", err)
	}
	// Serialize `EstimatedTokenAmount`:
	if err = encoder.Encode(obj.EstimatedTokenAmount); err != nil {
		return fmt.Errorf("error while marshaling EstimatedTokenAmount:%w", err)
	}
	return nil
}

func (obj EvtStrategyWithdraw) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtStrategyWithdraw: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtStrategyWithdraw) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_StrategyWithdraw[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_StrategyWithdraw[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `StrategyType`:
	if err = decoder.Decode(&obj.StrategyType); err != nil {
		return fmt.Errorf("error while unmarshaling StrategyType:%w", err)
	}
	// Deserialize `CollateralAmount`:
	if err = decoder.Decode(&obj.CollateralAmount); err != nil {
		return fmt.Errorf("error while unmarshaling CollateralAmount:%w", err)
	}
	// Deserialize `EstimatedTokenAmount`:
	if err = decoder.Decode(&obj.EstimatedTokenAmount); err != nil {
		return fmt.Errorf("error while unmarshaling EstimatedTokenAmount:%w", err)
	}
	return nil
}

func (obj *EvtStrategyWithdraw) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtStrategyWithdraw: %w", err)
	}
	return nil
}

func UnmarshalEvtStrategyWithdraw(buf []byte) (*EvtStrategyWithdraw, error) {
	obj := new(EvtStrategyWithdraw)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtClaimReward struct {
	StrategyType StrategyType
	TokenAmount  uint64
	MintAccount  solanago.PublicKey
}

func (obj EvtClaimReward) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_ClaimReward[:], false)
	if err != nil {
		return err
	}
	// Serialize `StrategyType`:
	if err = encoder.Encode(obj.StrategyType); err != nil {
		return fmt.Errorf("error while marshaling StrategyType:%w", err)
	}
	// Serialize `TokenAmount`:
	if err = encoder.Encode(obj.TokenAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenAmount:%w", err)
	}
	// Serialize `MintAccount`:
	if err = encoder.Encode(obj.MintAccount); err != nil {
		return fmt.Errorf("error while marshaling MintAccount:%w", err)
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
	// Deserialize `StrategyType`:
	if err = decoder.Decode(&obj.StrategyType); err != nil {
		return fmt.Errorf("error while unmarshaling StrategyType:%w", err)
	}
	// Deserialize `TokenAmount`:
	if err = decoder.Decode(&obj.TokenAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAmount:%w", err)
	}
	// Deserialize `MintAccount`:
	if err = decoder.Decode(&obj.MintAccount); err != nil {
		return fmt.Errorf("error while unmarshaling MintAccount:%w", err)
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

type EvtPerformanceFee struct {
	LpMintMore uint64
}

func (obj EvtPerformanceFee) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_PerformanceFee[:], false)
	if err != nil {
		return err
	}
	// Serialize `LpMintMore`:
	if err = encoder.Encode(obj.LpMintMore); err != nil {
		return fmt.Errorf("error while marshaling LpMintMore:%w", err)
	}
	return nil
}

func (obj EvtPerformanceFee) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtPerformanceFee: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtPerformanceFee) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_PerformanceFee[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_PerformanceFee[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `LpMintMore`:
	if err = decoder.Decode(&obj.LpMintMore); err != nil {
		return fmt.Errorf("error while unmarshaling LpMintMore:%w", err)
	}
	return nil
}

func (obj *EvtPerformanceFee) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtPerformanceFee: %w", err)
	}
	return nil
}

func UnmarshalEvtPerformanceFee(buf []byte) (*EvtPerformanceFee, error) {
	obj := new(EvtPerformanceFee)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtReportLoss struct {
	Strategy solanago.PublicKey
	Loss     uint64
}

func (obj EvtReportLoss) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_ReportLoss[:], false)
	if err != nil {
		return err
	}
	// Serialize `Strategy`:
	if err = encoder.Encode(obj.Strategy); err != nil {
		return fmt.Errorf("error while marshaling Strategy:%w", err)
	}
	// Serialize `Loss`:
	if err = encoder.Encode(obj.Loss); err != nil {
		return fmt.Errorf("error while marshaling Loss:%w", err)
	}
	return nil
}

func (obj EvtReportLoss) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtReportLoss: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtReportLoss) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_ReportLoss[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_ReportLoss[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Strategy`:
	if err = decoder.Decode(&obj.Strategy); err != nil {
		return fmt.Errorf("error while unmarshaling Strategy:%w", err)
	}
	// Deserialize `Loss`:
	if err = decoder.Decode(&obj.Loss); err != nil {
		return fmt.Errorf("error while unmarshaling Loss:%w", err)
	}
	return nil
}

func (obj *EvtReportLoss) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtReportLoss: %w", err)
	}
	return nil
}

func UnmarshalEvtReportLoss(buf []byte) (*EvtReportLoss, error) {
	obj := new(EvtReportLoss)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtTotalAmount struct {
	TotalAmount uint64
}

func (obj EvtTotalAmount) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_TotalAmount[:], false)
	if err != nil {
		return err
	}
	// Serialize `TotalAmount`:
	if err = encoder.Encode(obj.TotalAmount); err != nil {
		return fmt.Errorf("error while marshaling TotalAmount:%w", err)
	}
	return nil
}

func (obj EvtTotalAmount) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtTotalAmount: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtTotalAmount) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_TotalAmount[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_TotalAmount[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TotalAmount`:
	if err = decoder.Decode(&obj.TotalAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TotalAmount:%w", err)
	}
	return nil
}

func (obj *EvtTotalAmount) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtTotalAmount: %w", err)
	}
	return nil
}

func UnmarshalEvtTotalAmount(buf []byte) (*EvtTotalAmount, error) {
	obj := new(EvtTotalAmount)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
