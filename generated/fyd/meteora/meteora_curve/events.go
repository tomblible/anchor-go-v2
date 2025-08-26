package meteora_curve

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type EvtClaimCreatorTradingFee struct {
	Pool             solanago.PublicKey
	TokenBaseAmount  uint64
	TokenQuoteAmount uint64
}

func (obj EvtClaimCreatorTradingFee) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtClaimCreatorTradingFee[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `TokenBaseAmount`:
	if err = encoder.Encode(obj.TokenBaseAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenBaseAmount:%w", err)
	}
	// Serialize `TokenQuoteAmount`:
	if err = encoder.Encode(obj.TokenQuoteAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenQuoteAmount:%w", err)
	}
	return nil
}

func (obj EvtClaimCreatorTradingFee) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtClaimCreatorTradingFee: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtClaimCreatorTradingFee) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtClaimCreatorTradingFee[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtClaimCreatorTradingFee[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `TokenBaseAmount`:
	if err = decoder.Decode(&obj.TokenBaseAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBaseAmount:%w", err)
	}
	// Deserialize `TokenQuoteAmount`:
	if err = decoder.Decode(&obj.TokenQuoteAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenQuoteAmount:%w", err)
	}
	return nil
}

func (obj *EvtClaimCreatorTradingFee) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtClaimCreatorTradingFee: %w", err)
	}
	return nil
}

func UnmarshalEvtClaimCreatorTradingFee(buf []byte) (*EvtClaimCreatorTradingFee, error) {
	obj := new(EvtClaimCreatorTradingFee)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtClaimProtocolFee struct {
	Pool             solanago.PublicKey
	TokenBaseAmount  uint64
	TokenQuoteAmount uint64
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
	// Serialize `TokenBaseAmount`:
	if err = encoder.Encode(obj.TokenBaseAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenBaseAmount:%w", err)
	}
	// Serialize `TokenQuoteAmount`:
	if err = encoder.Encode(obj.TokenQuoteAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenQuoteAmount:%w", err)
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
	// Deserialize `TokenBaseAmount`:
	if err = decoder.Decode(&obj.TokenBaseAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBaseAmount:%w", err)
	}
	// Deserialize `TokenQuoteAmount`:
	if err = decoder.Decode(&obj.TokenQuoteAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenQuoteAmount:%w", err)
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

type EvtClaimTradingFee struct {
	Pool             solanago.PublicKey
	TokenBaseAmount  uint64
	TokenQuoteAmount uint64
}

func (obj EvtClaimTradingFee) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtClaimTradingFee[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `TokenBaseAmount`:
	if err = encoder.Encode(obj.TokenBaseAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenBaseAmount:%w", err)
	}
	// Serialize `TokenQuoteAmount`:
	if err = encoder.Encode(obj.TokenQuoteAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenQuoteAmount:%w", err)
	}
	return nil
}

func (obj EvtClaimTradingFee) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtClaimTradingFee: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtClaimTradingFee) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtClaimTradingFee[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtClaimTradingFee[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `TokenBaseAmount`:
	if err = decoder.Decode(&obj.TokenBaseAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBaseAmount:%w", err)
	}
	// Deserialize `TokenQuoteAmount`:
	if err = decoder.Decode(&obj.TokenQuoteAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenQuoteAmount:%w", err)
	}
	return nil
}

func (obj *EvtClaimTradingFee) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtClaimTradingFee: %w", err)
	}
	return nil
}

func UnmarshalEvtClaimTradingFee(buf []byte) (*EvtClaimTradingFee, error) {
	obj := new(EvtClaimTradingFee)
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

// Create config
type EvtCreateConfig struct {
	Config                    solanago.PublicKey
	QuoteMint                 solanago.PublicKey
	FeeClaimer                solanago.PublicKey
	Owner                     solanago.PublicKey
	PoolFees                  PoolFeeParameters
	CollectFeeMode            uint8
	MigrationOption           uint8
	ActivationType            uint8
	TokenDecimal              uint8
	TokenType                 uint8
	PartnerLockedLpPercentage uint8
	PartnerLpPercentage       uint8
	CreatorLockedLpPercentage uint8
	CreatorLpPercentage       uint8
	SwapBaseAmount            uint64
	MigrationQuoteThreshold   uint64
	MigrationBaseAmount       uint64
	SqrtStartPrice            binary.Uint128
	LockedVesting             LockedVestingParams
	MigrationFeeOption        uint8
	FixedTokenSupplyFlag      uint8
	PreMigrationTokenSupply   uint64
	PostMigrationTokenSupply  uint64
	Curve                     []LiquidityDistributionParameters
}

func (obj EvtCreateConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtCreateConfig[:], false)
	if err != nil {
		return err
	}
	// Serialize `Config`:
	if err = encoder.Encode(obj.Config); err != nil {
		return fmt.Errorf("error while marshaling Config:%w", err)
	}
	// Serialize `QuoteMint`:
	if err = encoder.Encode(obj.QuoteMint); err != nil {
		return fmt.Errorf("error while marshaling QuoteMint:%w", err)
	}
	// Serialize `FeeClaimer`:
	if err = encoder.Encode(obj.FeeClaimer); err != nil {
		return fmt.Errorf("error while marshaling FeeClaimer:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `PoolFees`:
	if err = encoder.Encode(obj.PoolFees); err != nil {
		return fmt.Errorf("error while marshaling PoolFees:%w", err)
	}
	// Serialize `CollectFeeMode`:
	if err = encoder.Encode(obj.CollectFeeMode); err != nil {
		return fmt.Errorf("error while marshaling CollectFeeMode:%w", err)
	}
	// Serialize `MigrationOption`:
	if err = encoder.Encode(obj.MigrationOption); err != nil {
		return fmt.Errorf("error while marshaling MigrationOption:%w", err)
	}
	// Serialize `ActivationType`:
	if err = encoder.Encode(obj.ActivationType); err != nil {
		return fmt.Errorf("error while marshaling ActivationType:%w", err)
	}
	// Serialize `TokenDecimal`:
	if err = encoder.Encode(obj.TokenDecimal); err != nil {
		return fmt.Errorf("error while marshaling TokenDecimal:%w", err)
	}
	// Serialize `TokenType`:
	if err = encoder.Encode(obj.TokenType); err != nil {
		return fmt.Errorf("error while marshaling TokenType:%w", err)
	}
	// Serialize `PartnerLockedLpPercentage`:
	if err = encoder.Encode(obj.PartnerLockedLpPercentage); err != nil {
		return fmt.Errorf("error while marshaling PartnerLockedLpPercentage:%w", err)
	}
	// Serialize `PartnerLpPercentage`:
	if err = encoder.Encode(obj.PartnerLpPercentage); err != nil {
		return fmt.Errorf("error while marshaling PartnerLpPercentage:%w", err)
	}
	// Serialize `CreatorLockedLpPercentage`:
	if err = encoder.Encode(obj.CreatorLockedLpPercentage); err != nil {
		return fmt.Errorf("error while marshaling CreatorLockedLpPercentage:%w", err)
	}
	// Serialize `CreatorLpPercentage`:
	if err = encoder.Encode(obj.CreatorLpPercentage); err != nil {
		return fmt.Errorf("error while marshaling CreatorLpPercentage:%w", err)
	}
	// Serialize `SwapBaseAmount`:
	if err = encoder.Encode(obj.SwapBaseAmount); err != nil {
		return fmt.Errorf("error while marshaling SwapBaseAmount:%w", err)
	}
	// Serialize `MigrationQuoteThreshold`:
	if err = encoder.Encode(obj.MigrationQuoteThreshold); err != nil {
		return fmt.Errorf("error while marshaling MigrationQuoteThreshold:%w", err)
	}
	// Serialize `MigrationBaseAmount`:
	if err = encoder.Encode(obj.MigrationBaseAmount); err != nil {
		return fmt.Errorf("error while marshaling MigrationBaseAmount:%w", err)
	}
	// Serialize `SqrtStartPrice`:
	if err = encoder.Encode(obj.SqrtStartPrice); err != nil {
		return fmt.Errorf("error while marshaling SqrtStartPrice:%w", err)
	}
	// Serialize `LockedVesting`:
	if err = encoder.Encode(obj.LockedVesting); err != nil {
		return fmt.Errorf("error while marshaling LockedVesting:%w", err)
	}
	// Serialize `MigrationFeeOption`:
	if err = encoder.Encode(obj.MigrationFeeOption); err != nil {
		return fmt.Errorf("error while marshaling MigrationFeeOption:%w", err)
	}
	// Serialize `FixedTokenSupplyFlag`:
	if err = encoder.Encode(obj.FixedTokenSupplyFlag); err != nil {
		return fmt.Errorf("error while marshaling FixedTokenSupplyFlag:%w", err)
	}
	// Serialize `PreMigrationTokenSupply`:
	if err = encoder.Encode(obj.PreMigrationTokenSupply); err != nil {
		return fmt.Errorf("error while marshaling PreMigrationTokenSupply:%w", err)
	}
	// Serialize `PostMigrationTokenSupply`:
	if err = encoder.Encode(obj.PostMigrationTokenSupply); err != nil {
		return fmt.Errorf("error while marshaling PostMigrationTokenSupply:%w", err)
	}
	// Serialize `Curve`:
	if err = encoder.Encode(obj.Curve); err != nil {
		return fmt.Errorf("error while marshaling Curve:%w", err)
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
	// Deserialize `Config`:
	if err = decoder.Decode(&obj.Config); err != nil {
		return fmt.Errorf("error while unmarshaling Config:%w", err)
	}
	// Deserialize `QuoteMint`:
	if err = decoder.Decode(&obj.QuoteMint); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteMint:%w", err)
	}
	// Deserialize `FeeClaimer`:
	if err = decoder.Decode(&obj.FeeClaimer); err != nil {
		return fmt.Errorf("error while unmarshaling FeeClaimer:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `PoolFees`:
	if err = decoder.Decode(&obj.PoolFees); err != nil {
		return fmt.Errorf("error while unmarshaling PoolFees:%w", err)
	}
	// Deserialize `CollectFeeMode`:
	if err = decoder.Decode(&obj.CollectFeeMode); err != nil {
		return fmt.Errorf("error while unmarshaling CollectFeeMode:%w", err)
	}
	// Deserialize `MigrationOption`:
	if err = decoder.Decode(&obj.MigrationOption); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationOption:%w", err)
	}
	// Deserialize `ActivationType`:
	if err = decoder.Decode(&obj.ActivationType); err != nil {
		return fmt.Errorf("error while unmarshaling ActivationType:%w", err)
	}
	// Deserialize `TokenDecimal`:
	if err = decoder.Decode(&obj.TokenDecimal); err != nil {
		return fmt.Errorf("error while unmarshaling TokenDecimal:%w", err)
	}
	// Deserialize `TokenType`:
	if err = decoder.Decode(&obj.TokenType); err != nil {
		return fmt.Errorf("error while unmarshaling TokenType:%w", err)
	}
	// Deserialize `PartnerLockedLpPercentage`:
	if err = decoder.Decode(&obj.PartnerLockedLpPercentage); err != nil {
		return fmt.Errorf("error while unmarshaling PartnerLockedLpPercentage:%w", err)
	}
	// Deserialize `PartnerLpPercentage`:
	if err = decoder.Decode(&obj.PartnerLpPercentage); err != nil {
		return fmt.Errorf("error while unmarshaling PartnerLpPercentage:%w", err)
	}
	// Deserialize `CreatorLockedLpPercentage`:
	if err = decoder.Decode(&obj.CreatorLockedLpPercentage); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorLockedLpPercentage:%w", err)
	}
	// Deserialize `CreatorLpPercentage`:
	if err = decoder.Decode(&obj.CreatorLpPercentage); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorLpPercentage:%w", err)
	}
	// Deserialize `SwapBaseAmount`:
	if err = decoder.Decode(&obj.SwapBaseAmount); err != nil {
		return fmt.Errorf("error while unmarshaling SwapBaseAmount:%w", err)
	}
	// Deserialize `MigrationQuoteThreshold`:
	if err = decoder.Decode(&obj.MigrationQuoteThreshold); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationQuoteThreshold:%w", err)
	}
	// Deserialize `MigrationBaseAmount`:
	if err = decoder.Decode(&obj.MigrationBaseAmount); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationBaseAmount:%w", err)
	}
	// Deserialize `SqrtStartPrice`:
	if err = decoder.Decode(&obj.SqrtStartPrice); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtStartPrice:%w", err)
	}
	// Deserialize `LockedVesting`:
	if err = decoder.Decode(&obj.LockedVesting); err != nil {
		return fmt.Errorf("error while unmarshaling LockedVesting:%w", err)
	}
	// Deserialize `MigrationFeeOption`:
	if err = decoder.Decode(&obj.MigrationFeeOption); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationFeeOption:%w", err)
	}
	// Deserialize `FixedTokenSupplyFlag`:
	if err = decoder.Decode(&obj.FixedTokenSupplyFlag); err != nil {
		return fmt.Errorf("error while unmarshaling FixedTokenSupplyFlag:%w", err)
	}
	// Deserialize `PreMigrationTokenSupply`:
	if err = decoder.Decode(&obj.PreMigrationTokenSupply); err != nil {
		return fmt.Errorf("error while unmarshaling PreMigrationTokenSupply:%w", err)
	}
	// Deserialize `PostMigrationTokenSupply`:
	if err = decoder.Decode(&obj.PostMigrationTokenSupply); err != nil {
		return fmt.Errorf("error while unmarshaling PostMigrationTokenSupply:%w", err)
	}
	// Deserialize `Curve`:
	if err = decoder.Decode(&obj.Curve); err != nil {
		return fmt.Errorf("error while unmarshaling Curve:%w", err)
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

type EvtCreateConfigV2 struct {
	Config           solanago.PublicKey
	QuoteMint        solanago.PublicKey
	FeeClaimer       solanago.PublicKey
	LeftoverReceiver solanago.PublicKey
	ConfigParameters ConfigParameters
}

func (obj EvtCreateConfigV2) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtCreateConfigV2[:], false)
	if err != nil {
		return err
	}
	// Serialize `Config`:
	if err = encoder.Encode(obj.Config); err != nil {
		return fmt.Errorf("error while marshaling Config:%w", err)
	}
	// Serialize `QuoteMint`:
	if err = encoder.Encode(obj.QuoteMint); err != nil {
		return fmt.Errorf("error while marshaling QuoteMint:%w", err)
	}
	// Serialize `FeeClaimer`:
	if err = encoder.Encode(obj.FeeClaimer); err != nil {
		return fmt.Errorf("error while marshaling FeeClaimer:%w", err)
	}
	// Serialize `LeftoverReceiver`:
	if err = encoder.Encode(obj.LeftoverReceiver); err != nil {
		return fmt.Errorf("error while marshaling LeftoverReceiver:%w", err)
	}
	// Serialize `ConfigParameters`:
	if err = encoder.Encode(obj.ConfigParameters); err != nil {
		return fmt.Errorf("error while marshaling ConfigParameters:%w", err)
	}
	return nil
}

func (obj EvtCreateConfigV2) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtCreateConfigV2: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtCreateConfigV2) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtCreateConfigV2[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtCreateConfigV2[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Config`:
	if err = decoder.Decode(&obj.Config); err != nil {
		return fmt.Errorf("error while unmarshaling Config:%w", err)
	}
	// Deserialize `QuoteMint`:
	if err = decoder.Decode(&obj.QuoteMint); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteMint:%w", err)
	}
	// Deserialize `FeeClaimer`:
	if err = decoder.Decode(&obj.FeeClaimer); err != nil {
		return fmt.Errorf("error while unmarshaling FeeClaimer:%w", err)
	}
	// Deserialize `LeftoverReceiver`:
	if err = decoder.Decode(&obj.LeftoverReceiver); err != nil {
		return fmt.Errorf("error while unmarshaling LeftoverReceiver:%w", err)
	}
	// Deserialize `ConfigParameters`:
	if err = decoder.Decode(&obj.ConfigParameters); err != nil {
		return fmt.Errorf("error while unmarshaling ConfigParameters:%w", err)
	}
	return nil
}

func (obj *EvtCreateConfigV2) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtCreateConfigV2: %w", err)
	}
	return nil
}

func UnmarshalEvtCreateConfigV2(buf []byte) (*EvtCreateConfigV2, error) {
	obj := new(EvtCreateConfigV2)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtCreateDammV2MigrationMetadata struct {
	VirtualPool solanago.PublicKey
}

func (obj EvtCreateDammV2MigrationMetadata) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtCreateDammV2MigrationMetadata[:], false)
	if err != nil {
		return err
	}
	// Serialize `VirtualPool`:
	if err = encoder.Encode(obj.VirtualPool); err != nil {
		return fmt.Errorf("error while marshaling VirtualPool:%w", err)
	}
	return nil
}

func (obj EvtCreateDammV2MigrationMetadata) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtCreateDammV2MigrationMetadata: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtCreateDammV2MigrationMetadata) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtCreateDammV2MigrationMetadata[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtCreateDammV2MigrationMetadata[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `VirtualPool`:
	if err = decoder.Decode(&obj.VirtualPool); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualPool:%w", err)
	}
	return nil
}

func (obj *EvtCreateDammV2MigrationMetadata) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtCreateDammV2MigrationMetadata: %w", err)
	}
	return nil
}

func UnmarshalEvtCreateDammV2MigrationMetadata(buf []byte) (*EvtCreateDammV2MigrationMetadata, error) {
	obj := new(EvtCreateDammV2MigrationMetadata)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtCreateMeteoraMigrationMetadata struct {
	VirtualPool solanago.PublicKey
}

func (obj EvtCreateMeteoraMigrationMetadata) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtCreateMeteoraMigrationMetadata[:], false)
	if err != nil {
		return err
	}
	// Serialize `VirtualPool`:
	if err = encoder.Encode(obj.VirtualPool); err != nil {
		return fmt.Errorf("error while marshaling VirtualPool:%w", err)
	}
	return nil
}

func (obj EvtCreateMeteoraMigrationMetadata) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtCreateMeteoraMigrationMetadata: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtCreateMeteoraMigrationMetadata) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtCreateMeteoraMigrationMetadata[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtCreateMeteoraMigrationMetadata[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `VirtualPool`:
	if err = decoder.Decode(&obj.VirtualPool); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualPool:%w", err)
	}
	return nil
}

func (obj *EvtCreateMeteoraMigrationMetadata) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtCreateMeteoraMigrationMetadata: %w", err)
	}
	return nil
}

func UnmarshalEvtCreateMeteoraMigrationMetadata(buf []byte) (*EvtCreateMeteoraMigrationMetadata, error) {
	obj := new(EvtCreateMeteoraMigrationMetadata)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtCreatorWithdrawSurplus struct {
	Pool          solanago.PublicKey
	SurplusAmount uint64
}

func (obj EvtCreatorWithdrawSurplus) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtCreatorWithdrawSurplus[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `SurplusAmount`:
	if err = encoder.Encode(obj.SurplusAmount); err != nil {
		return fmt.Errorf("error while marshaling SurplusAmount:%w", err)
	}
	return nil
}

func (obj EvtCreatorWithdrawSurplus) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtCreatorWithdrawSurplus: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtCreatorWithdrawSurplus) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtCreatorWithdrawSurplus[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtCreatorWithdrawSurplus[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `SurplusAmount`:
	if err = decoder.Decode(&obj.SurplusAmount); err != nil {
		return fmt.Errorf("error while unmarshaling SurplusAmount:%w", err)
	}
	return nil
}

func (obj *EvtCreatorWithdrawSurplus) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtCreatorWithdrawSurplus: %w", err)
	}
	return nil
}

func UnmarshalEvtCreatorWithdrawSurplus(buf []byte) (*EvtCreatorWithdrawSurplus, error) {
	obj := new(EvtCreatorWithdrawSurplus)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtCurveComplete struct {
	Pool         solanago.PublicKey
	Config       solanago.PublicKey
	BaseReserve  uint64
	QuoteReserve uint64
}

func (obj EvtCurveComplete) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtCurveComplete[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Config`:
	if err = encoder.Encode(obj.Config); err != nil {
		return fmt.Errorf("error while marshaling Config:%w", err)
	}
	// Serialize `BaseReserve`:
	if err = encoder.Encode(obj.BaseReserve); err != nil {
		return fmt.Errorf("error while marshaling BaseReserve:%w", err)
	}
	// Serialize `QuoteReserve`:
	if err = encoder.Encode(obj.QuoteReserve); err != nil {
		return fmt.Errorf("error while marshaling QuoteReserve:%w", err)
	}
	return nil
}

func (obj EvtCurveComplete) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtCurveComplete: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtCurveComplete) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtCurveComplete[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtCurveComplete[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Config`:
	if err = decoder.Decode(&obj.Config); err != nil {
		return fmt.Errorf("error while unmarshaling Config:%w", err)
	}
	// Deserialize `BaseReserve`:
	if err = decoder.Decode(&obj.BaseReserve); err != nil {
		return fmt.Errorf("error while unmarshaling BaseReserve:%w", err)
	}
	// Deserialize `QuoteReserve`:
	if err = decoder.Decode(&obj.QuoteReserve); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteReserve:%w", err)
	}
	return nil
}

func (obj *EvtCurveComplete) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtCurveComplete: %w", err)
	}
	return nil
}

func UnmarshalEvtCurveComplete(buf []byte) (*EvtCurveComplete, error) {
	obj := new(EvtCurveComplete)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtInitializePool struct {
	Pool            solanago.PublicKey
	Config          solanago.PublicKey
	Creator         solanago.PublicKey
	BaseMint        solanago.PublicKey
	PoolType        uint8
	ActivationPoint uint64
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
	// Serialize `Config`:
	if err = encoder.Encode(obj.Config); err != nil {
		return fmt.Errorf("error while marshaling Config:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
	}
	// Serialize `BaseMint`:
	if err = encoder.Encode(obj.BaseMint); err != nil {
		return fmt.Errorf("error while marshaling BaseMint:%w", err)
	}
	// Serialize `PoolType`:
	if err = encoder.Encode(obj.PoolType); err != nil {
		return fmt.Errorf("error while marshaling PoolType:%w", err)
	}
	// Serialize `ActivationPoint`:
	if err = encoder.Encode(obj.ActivationPoint); err != nil {
		return fmt.Errorf("error while marshaling ActivationPoint:%w", err)
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
	// Deserialize `Config`:
	if err = decoder.Decode(&obj.Config); err != nil {
		return fmt.Errorf("error while unmarshaling Config:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
	}
	// Deserialize `BaseMint`:
	if err = decoder.Decode(&obj.BaseMint); err != nil {
		return fmt.Errorf("error while unmarshaling BaseMint:%w", err)
	}
	// Deserialize `PoolType`:
	if err = decoder.Decode(&obj.PoolType); err != nil {
		return fmt.Errorf("error while unmarshaling PoolType:%w", err)
	}
	// Deserialize `ActivationPoint`:
	if err = decoder.Decode(&obj.ActivationPoint); err != nil {
		return fmt.Errorf("error while unmarshaling ActivationPoint:%w", err)
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

// Create partner metadata
type EvtPartnerMetadata struct {
	PartnerMetadata solanago.PublicKey
	FeeClaimer      solanago.PublicKey
}

func (obj EvtPartnerMetadata) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtPartnerMetadata[:], false)
	if err != nil {
		return err
	}
	// Serialize `PartnerMetadata`:
	if err = encoder.Encode(obj.PartnerMetadata); err != nil {
		return fmt.Errorf("error while marshaling PartnerMetadata:%w", err)
	}
	// Serialize `FeeClaimer`:
	if err = encoder.Encode(obj.FeeClaimer); err != nil {
		return fmt.Errorf("error while marshaling FeeClaimer:%w", err)
	}
	return nil
}

func (obj EvtPartnerMetadata) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtPartnerMetadata: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtPartnerMetadata) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtPartnerMetadata[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtPartnerMetadata[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PartnerMetadata`:
	if err = decoder.Decode(&obj.PartnerMetadata); err != nil {
		return fmt.Errorf("error while unmarshaling PartnerMetadata:%w", err)
	}
	// Deserialize `FeeClaimer`:
	if err = decoder.Decode(&obj.FeeClaimer); err != nil {
		return fmt.Errorf("error while unmarshaling FeeClaimer:%w", err)
	}
	return nil
}

func (obj *EvtPartnerMetadata) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtPartnerMetadata: %w", err)
	}
	return nil
}

func UnmarshalEvtPartnerMetadata(buf []byte) (*EvtPartnerMetadata, error) {
	obj := new(EvtPartnerMetadata)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtPartnerWithdrawMigrationFee struct {
	Pool solanago.PublicKey
	Fee  uint64
}

func (obj EvtPartnerWithdrawMigrationFee) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtPartnerWithdrawMigrationFee[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Fee`:
	if err = encoder.Encode(obj.Fee); err != nil {
		return fmt.Errorf("error while marshaling Fee:%w", err)
	}
	return nil
}

func (obj EvtPartnerWithdrawMigrationFee) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtPartnerWithdrawMigrationFee: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtPartnerWithdrawMigrationFee) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtPartnerWithdrawMigrationFee[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtPartnerWithdrawMigrationFee[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Fee`:
	if err = decoder.Decode(&obj.Fee); err != nil {
		return fmt.Errorf("error while unmarshaling Fee:%w", err)
	}
	return nil
}

func (obj *EvtPartnerWithdrawMigrationFee) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtPartnerWithdrawMigrationFee: %w", err)
	}
	return nil
}

func UnmarshalEvtPartnerWithdrawMigrationFee(buf []byte) (*EvtPartnerWithdrawMigrationFee, error) {
	obj := new(EvtPartnerWithdrawMigrationFee)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtPartnerWithdrawSurplus struct {
	Pool          solanago.PublicKey
	SurplusAmount uint64
}

func (obj EvtPartnerWithdrawSurplus) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtPartnerWithdrawSurplus[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `SurplusAmount`:
	if err = encoder.Encode(obj.SurplusAmount); err != nil {
		return fmt.Errorf("error while marshaling SurplusAmount:%w", err)
	}
	return nil
}

func (obj EvtPartnerWithdrawSurplus) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtPartnerWithdrawSurplus: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtPartnerWithdrawSurplus) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtPartnerWithdrawSurplus[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtPartnerWithdrawSurplus[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `SurplusAmount`:
	if err = decoder.Decode(&obj.SurplusAmount); err != nil {
		return fmt.Errorf("error while unmarshaling SurplusAmount:%w", err)
	}
	return nil
}

func (obj *EvtPartnerWithdrawSurplus) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtPartnerWithdrawSurplus: %w", err)
	}
	return nil
}

func UnmarshalEvtPartnerWithdrawSurplus(buf []byte) (*EvtPartnerWithdrawSurplus, error) {
	obj := new(EvtPartnerWithdrawSurplus)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtProtocolWithdrawSurplus struct {
	Pool          solanago.PublicKey
	SurplusAmount uint64
}

func (obj EvtProtocolWithdrawSurplus) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtProtocolWithdrawSurplus[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `SurplusAmount`:
	if err = encoder.Encode(obj.SurplusAmount); err != nil {
		return fmt.Errorf("error while marshaling SurplusAmount:%w", err)
	}
	return nil
}

func (obj EvtProtocolWithdrawSurplus) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtProtocolWithdrawSurplus: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtProtocolWithdrawSurplus) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtProtocolWithdrawSurplus[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtProtocolWithdrawSurplus[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `SurplusAmount`:
	if err = decoder.Decode(&obj.SurplusAmount); err != nil {
		return fmt.Errorf("error while unmarshaling SurplusAmount:%w", err)
	}
	return nil
}

func (obj *EvtProtocolWithdrawSurplus) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtProtocolWithdrawSurplus: %w", err)
	}
	return nil
}

func UnmarshalEvtProtocolWithdrawSurplus(buf []byte) (*EvtProtocolWithdrawSurplus, error) {
	obj := new(EvtProtocolWithdrawSurplus)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtSwap struct {
	Pool             solanago.PublicKey
	Config           solanago.PublicKey
	TradeDirection   uint8
	HasReferral      bool
	Params           SwapParameters
	SwapResult       SwapResult
	AmountIn         uint64
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
	// Serialize `Config`:
	if err = encoder.Encode(obj.Config); err != nil {
		return fmt.Errorf("error while marshaling Config:%w", err)
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
	// Serialize `AmountIn`:
	if err = encoder.Encode(obj.AmountIn); err != nil {
		return fmt.Errorf("error while marshaling AmountIn:%w", err)
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
	// Deserialize `Config`:
	if err = decoder.Decode(&obj.Config); err != nil {
		return fmt.Errorf("error while unmarshaling Config:%w", err)
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
	// Deserialize `AmountIn`:
	if err = decoder.Decode(&obj.AmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling AmountIn:%w", err)
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

type EvtSwap2 struct {
	Pool               solanago.PublicKey
	Config             solanago.PublicKey
	TradeDirection     uint8
	HasReferral        bool
	SwapParameters     SwapParameters2
	SwapResult         SwapResult2
	QuoteReserveAmount uint64
	MigrationThreshold uint64
	CurrentTimestamp   uint64
}

func (obj EvtSwap2) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtSwap2[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Config`:
	if err = encoder.Encode(obj.Config); err != nil {
		return fmt.Errorf("error while marshaling Config:%w", err)
	}
	// Serialize `TradeDirection`:
	if err = encoder.Encode(obj.TradeDirection); err != nil {
		return fmt.Errorf("error while marshaling TradeDirection:%w", err)
	}
	// Serialize `HasReferral`:
	if err = encoder.Encode(obj.HasReferral); err != nil {
		return fmt.Errorf("error while marshaling HasReferral:%w", err)
	}
	// Serialize `SwapParameters`:
	if err = encoder.Encode(obj.SwapParameters); err != nil {
		return fmt.Errorf("error while marshaling SwapParameters:%w", err)
	}
	// Serialize `SwapResult`:
	if err = encoder.Encode(obj.SwapResult); err != nil {
		return fmt.Errorf("error while marshaling SwapResult:%w", err)
	}
	// Serialize `QuoteReserveAmount`:
	if err = encoder.Encode(obj.QuoteReserveAmount); err != nil {
		return fmt.Errorf("error while marshaling QuoteReserveAmount:%w", err)
	}
	// Serialize `MigrationThreshold`:
	if err = encoder.Encode(obj.MigrationThreshold); err != nil {
		return fmt.Errorf("error while marshaling MigrationThreshold:%w", err)
	}
	// Serialize `CurrentTimestamp`:
	if err = encoder.Encode(obj.CurrentTimestamp); err != nil {
		return fmt.Errorf("error while marshaling CurrentTimestamp:%w", err)
	}
	return nil
}

func (obj EvtSwap2) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtSwap2: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtSwap2) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtSwap2[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtSwap2[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Config`:
	if err = decoder.Decode(&obj.Config); err != nil {
		return fmt.Errorf("error while unmarshaling Config:%w", err)
	}
	// Deserialize `TradeDirection`:
	if err = decoder.Decode(&obj.TradeDirection); err != nil {
		return fmt.Errorf("error while unmarshaling TradeDirection:%w", err)
	}
	// Deserialize `HasReferral`:
	if err = decoder.Decode(&obj.HasReferral); err != nil {
		return fmt.Errorf("error while unmarshaling HasReferral:%w", err)
	}
	// Deserialize `SwapParameters`:
	if err = decoder.Decode(&obj.SwapParameters); err != nil {
		return fmt.Errorf("error while unmarshaling SwapParameters:%w", err)
	}
	// Deserialize `SwapResult`:
	if err = decoder.Decode(&obj.SwapResult); err != nil {
		return fmt.Errorf("error while unmarshaling SwapResult:%w", err)
	}
	// Deserialize `QuoteReserveAmount`:
	if err = decoder.Decode(&obj.QuoteReserveAmount); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteReserveAmount:%w", err)
	}
	// Deserialize `MigrationThreshold`:
	if err = decoder.Decode(&obj.MigrationThreshold); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationThreshold:%w", err)
	}
	// Deserialize `CurrentTimestamp`:
	if err = decoder.Decode(&obj.CurrentTimestamp); err != nil {
		return fmt.Errorf("error while unmarshaling CurrentTimestamp:%w", err)
	}
	return nil
}

func (obj *EvtSwap2) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtSwap2: %w", err)
	}
	return nil
}

func UnmarshalEvtSwap2(buf []byte) (*EvtSwap2, error) {
	obj := new(EvtSwap2)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtUpdatePoolCreator struct {
	Pool       solanago.PublicKey
	Creator    solanago.PublicKey
	NewCreator solanago.PublicKey
}

func (obj EvtUpdatePoolCreator) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtUpdatePoolCreator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
	}
	// Serialize `NewCreator`:
	if err = encoder.Encode(obj.NewCreator); err != nil {
		return fmt.Errorf("error while marshaling NewCreator:%w", err)
	}
	return nil
}

func (obj EvtUpdatePoolCreator) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtUpdatePoolCreator: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtUpdatePoolCreator) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtUpdatePoolCreator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtUpdatePoolCreator[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
	}
	// Deserialize `NewCreator`:
	if err = decoder.Decode(&obj.NewCreator); err != nil {
		return fmt.Errorf("error while unmarshaling NewCreator:%w", err)
	}
	return nil
}

func (obj *EvtUpdatePoolCreator) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtUpdatePoolCreator: %w", err)
	}
	return nil
}

func UnmarshalEvtUpdatePoolCreator(buf []byte) (*EvtUpdatePoolCreator, error) {
	obj := new(EvtUpdatePoolCreator)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create virtual pool metadata
type EvtVirtualPoolMetadata struct {
	VirtualPoolMetadata solanago.PublicKey
	VirtualPool         solanago.PublicKey
}

func (obj EvtVirtualPoolMetadata) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtVirtualPoolMetadata[:], false)
	if err != nil {
		return err
	}
	// Serialize `VirtualPoolMetadata`:
	if err = encoder.Encode(obj.VirtualPoolMetadata); err != nil {
		return fmt.Errorf("error while marshaling VirtualPoolMetadata:%w", err)
	}
	// Serialize `VirtualPool`:
	if err = encoder.Encode(obj.VirtualPool); err != nil {
		return fmt.Errorf("error while marshaling VirtualPool:%w", err)
	}
	return nil
}

func (obj EvtVirtualPoolMetadata) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtVirtualPoolMetadata: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtVirtualPoolMetadata) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtVirtualPoolMetadata[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtVirtualPoolMetadata[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `VirtualPoolMetadata`:
	if err = decoder.Decode(&obj.VirtualPoolMetadata); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualPoolMetadata:%w", err)
	}
	// Deserialize `VirtualPool`:
	if err = decoder.Decode(&obj.VirtualPool); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualPool:%w", err)
	}
	return nil
}

func (obj *EvtVirtualPoolMetadata) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtVirtualPoolMetadata: %w", err)
	}
	return nil
}

func UnmarshalEvtVirtualPoolMetadata(buf []byte) (*EvtVirtualPoolMetadata, error) {
	obj := new(EvtVirtualPoolMetadata)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtWithdrawLeftover struct {
	Pool             solanago.PublicKey
	LeftoverReceiver solanago.PublicKey
	LeftoverAmount   uint64
}

func (obj EvtWithdrawLeftover) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtWithdrawLeftover[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `LeftoverReceiver`:
	if err = encoder.Encode(obj.LeftoverReceiver); err != nil {
		return fmt.Errorf("error while marshaling LeftoverReceiver:%w", err)
	}
	// Serialize `LeftoverAmount`:
	if err = encoder.Encode(obj.LeftoverAmount); err != nil {
		return fmt.Errorf("error while marshaling LeftoverAmount:%w", err)
	}
	return nil
}

func (obj EvtWithdrawLeftover) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtWithdrawLeftover: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtWithdrawLeftover) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtWithdrawLeftover[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtWithdrawLeftover[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `LeftoverReceiver`:
	if err = decoder.Decode(&obj.LeftoverReceiver); err != nil {
		return fmt.Errorf("error while unmarshaling LeftoverReceiver:%w", err)
	}
	// Deserialize `LeftoverAmount`:
	if err = decoder.Decode(&obj.LeftoverAmount); err != nil {
		return fmt.Errorf("error while unmarshaling LeftoverAmount:%w", err)
	}
	return nil
}

func (obj *EvtWithdrawLeftover) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtWithdrawLeftover: %w", err)
	}
	return nil
}

func UnmarshalEvtWithdrawLeftover(buf []byte) (*EvtWithdrawLeftover, error) {
	obj := new(EvtWithdrawLeftover)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EvtWithdrawMigrationFee struct {
	Pool solanago.PublicKey
	Fee  uint64
	Flag uint8
}

func (obj EvtWithdrawMigrationFee) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_EvtWithdrawMigrationFee[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Fee`:
	if err = encoder.Encode(obj.Fee); err != nil {
		return fmt.Errorf("error while marshaling Fee:%w", err)
	}
	// Serialize `Flag`:
	if err = encoder.Encode(obj.Flag); err != nil {
		return fmt.Errorf("error while marshaling Flag:%w", err)
	}
	return nil
}

func (obj EvtWithdrawMigrationFee) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EvtWithdrawMigrationFee: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EvtWithdrawMigrationFee) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_EvtWithdrawMigrationFee[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_EvtWithdrawMigrationFee[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Fee`:
	if err = decoder.Decode(&obj.Fee); err != nil {
		return fmt.Errorf("error while unmarshaling Fee:%w", err)
	}
	// Deserialize `Flag`:
	if err = decoder.Decode(&obj.Flag); err != nil {
		return fmt.Errorf("error while unmarshaling Flag:%w", err)
	}
	return nil
}

func (obj *EvtWithdrawMigrationFee) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EvtWithdrawMigrationFee: %w", err)
	}
	return nil
}

func UnmarshalEvtWithdrawMigrationFee(buf []byte) (*EvtWithdrawMigrationFee, error) {
	obj := new(EvtWithdrawMigrationFee)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
