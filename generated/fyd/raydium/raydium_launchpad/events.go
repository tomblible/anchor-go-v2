package raydium_launchpad

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

// Emitted when vesting token claimed by beneficiary
type ClaimVestedEvent struct {
	PoolState   solanago.PublicKey
	Beneficiary solanago.PublicKey
	ClaimAmount uint64
}

func (obj ClaimVestedEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_ClaimVestedEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolState`:
	if err = encoder.Encode(obj.PoolState); err != nil {
		return fmt.Errorf("error while marshaling PoolState:%w", err)
	}
	// Serialize `Beneficiary`:
	if err = encoder.Encode(obj.Beneficiary); err != nil {
		return fmt.Errorf("error while marshaling Beneficiary:%w", err)
	}
	// Serialize `ClaimAmount`:
	if err = encoder.Encode(obj.ClaimAmount); err != nil {
		return fmt.Errorf("error while marshaling ClaimAmount:%w", err)
	}
	return nil
}

func (obj ClaimVestedEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding ClaimVestedEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *ClaimVestedEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_ClaimVestedEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_ClaimVestedEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolState`:
	if err = decoder.Decode(&obj.PoolState); err != nil {
		return fmt.Errorf("error while unmarshaling PoolState:%w", err)
	}
	// Deserialize `Beneficiary`:
	if err = decoder.Decode(&obj.Beneficiary); err != nil {
		return fmt.Errorf("error while unmarshaling Beneficiary:%w", err)
	}
	// Deserialize `ClaimAmount`:
	if err = decoder.Decode(&obj.ClaimAmount); err != nil {
		return fmt.Errorf("error while unmarshaling ClaimAmount:%w", err)
	}
	return nil
}

func (obj *ClaimVestedEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling ClaimVestedEvent: %w", err)
	}
	return nil
}

func UnmarshalClaimVestedEvent(buf []byte) (*ClaimVestedEvent, error) {
	obj := new(ClaimVestedEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Emitted when vest_account created
type CreateVestingEvent struct {
	PoolState   solanago.PublicKey
	Beneficiary solanago.PublicKey
	ShareAmount uint64
}

func (obj CreateVestingEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CreateVestingEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolState`:
	if err = encoder.Encode(obj.PoolState); err != nil {
		return fmt.Errorf("error while marshaling PoolState:%w", err)
	}
	// Serialize `Beneficiary`:
	if err = encoder.Encode(obj.Beneficiary); err != nil {
		return fmt.Errorf("error while marshaling Beneficiary:%w", err)
	}
	// Serialize `ShareAmount`:
	if err = encoder.Encode(obj.ShareAmount); err != nil {
		return fmt.Errorf("error while marshaling ShareAmount:%w", err)
	}
	return nil
}

func (obj CreateVestingEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding CreateVestingEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *CreateVestingEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CreateVestingEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CreateVestingEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolState`:
	if err = decoder.Decode(&obj.PoolState); err != nil {
		return fmt.Errorf("error while unmarshaling PoolState:%w", err)
	}
	// Deserialize `Beneficiary`:
	if err = decoder.Decode(&obj.Beneficiary); err != nil {
		return fmt.Errorf("error while unmarshaling Beneficiary:%w", err)
	}
	// Deserialize `ShareAmount`:
	if err = decoder.Decode(&obj.ShareAmount); err != nil {
		return fmt.Errorf("error while unmarshaling ShareAmount:%w", err)
	}
	return nil
}

func (obj *CreateVestingEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling CreateVestingEvent: %w", err)
	}
	return nil
}

func UnmarshalCreateVestingEvent(buf []byte) (*CreateVestingEvent, error) {
	obj := new(CreateVestingEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Emitted when pool created
type PoolCreateEvent struct {
	PoolState     solanago.PublicKey
	Creator       solanago.PublicKey
	Config        solanago.PublicKey
	BaseMintParam MintParams
	CurveParam    CurveParams
	VestingParam  VestingParams
}

func (obj PoolCreateEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_PoolCreateEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolState`:
	if err = encoder.Encode(obj.PoolState); err != nil {
		return fmt.Errorf("error while marshaling PoolState:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
	}
	// Serialize `Config`:
	if err = encoder.Encode(obj.Config); err != nil {
		return fmt.Errorf("error while marshaling Config:%w", err)
	}
	// Serialize `BaseMintParam`:
	if err = encoder.Encode(obj.BaseMintParam); err != nil {
		return fmt.Errorf("error while marshaling BaseMintParam:%w", err)
	}
	// Serialize `CurveParam`:
	{
		if err = EncodeCurveParams(encoder, obj.CurveParam); err != nil {
			return fmt.Errorf("error while marshalingCurveParam:%w", err)
		}
	}
	// Serialize `VestingParam`:
	if err = encoder.Encode(obj.VestingParam); err != nil {
		return fmt.Errorf("error while marshaling VestingParam:%w", err)
	}
	return nil
}

func (obj PoolCreateEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding PoolCreateEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *PoolCreateEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_PoolCreateEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_PoolCreateEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolState`:
	if err = decoder.Decode(&obj.PoolState); err != nil {
		return fmt.Errorf("error while unmarshaling PoolState:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
	}
	// Deserialize `Config`:
	if err = decoder.Decode(&obj.Config); err != nil {
		return fmt.Errorf("error while unmarshaling Config:%w", err)
	}
	// Deserialize `BaseMintParam`:
	if err = decoder.Decode(&obj.BaseMintParam); err != nil {
		return fmt.Errorf("error while unmarshaling BaseMintParam:%w", err)
	}
	// Deserialize `CurveParam`:
	{
		var err error
		obj.CurveParam, err = DecodeCurveParams(decoder)
		if err != nil {
			return err
		}
	}
	// Deserialize `VestingParam`:
	if err = decoder.Decode(&obj.VestingParam); err != nil {
		return fmt.Errorf("error while unmarshaling VestingParam:%w", err)
	}
	return nil
}

func (obj *PoolCreateEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling PoolCreateEvent: %w", err)
	}
	return nil
}

func UnmarshalPoolCreateEvent(buf []byte) (*PoolCreateEvent, error) {
	obj := new(PoolCreateEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Emitted when trade process
type TradeEvent struct {
	PoolState       solanago.PublicKey
	TotalBaseSell   uint64
	VirtualBase     uint64
	VirtualQuote    uint64
	RealBaseBefore  uint64
	RealQuoteBefore uint64
	RealBaseAfter   uint64
	RealQuoteAfter  uint64
	AmountIn        uint64
	AmountOut       uint64
	ProtocolFee     uint64
	PlatformFee     uint64
	CreatorFee      uint64
	ShareFee        uint64
	TradeDirection  TradeDirection
	PoolStatus      PoolStatus
	ExactIn         bool
}

func (obj TradeEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_TradeEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolState`:
	if err = encoder.Encode(obj.PoolState); err != nil {
		return fmt.Errorf("error while marshaling PoolState:%w", err)
	}
	// Serialize `TotalBaseSell`:
	if err = encoder.Encode(obj.TotalBaseSell); err != nil {
		return fmt.Errorf("error while marshaling TotalBaseSell:%w", err)
	}
	// Serialize `VirtualBase`:
	if err = encoder.Encode(obj.VirtualBase); err != nil {
		return fmt.Errorf("error while marshaling VirtualBase:%w", err)
	}
	// Serialize `VirtualQuote`:
	if err = encoder.Encode(obj.VirtualQuote); err != nil {
		return fmt.Errorf("error while marshaling VirtualQuote:%w", err)
	}
	// Serialize `RealBaseBefore`:
	if err = encoder.Encode(obj.RealBaseBefore); err != nil {
		return fmt.Errorf("error while marshaling RealBaseBefore:%w", err)
	}
	// Serialize `RealQuoteBefore`:
	if err = encoder.Encode(obj.RealQuoteBefore); err != nil {
		return fmt.Errorf("error while marshaling RealQuoteBefore:%w", err)
	}
	// Serialize `RealBaseAfter`:
	if err = encoder.Encode(obj.RealBaseAfter); err != nil {
		return fmt.Errorf("error while marshaling RealBaseAfter:%w", err)
	}
	// Serialize `RealQuoteAfter`:
	if err = encoder.Encode(obj.RealQuoteAfter); err != nil {
		return fmt.Errorf("error while marshaling RealQuoteAfter:%w", err)
	}
	// Serialize `AmountIn`:
	if err = encoder.Encode(obj.AmountIn); err != nil {
		return fmt.Errorf("error while marshaling AmountIn:%w", err)
	}
	// Serialize `AmountOut`:
	if err = encoder.Encode(obj.AmountOut); err != nil {
		return fmt.Errorf("error while marshaling AmountOut:%w", err)
	}
	// Serialize `ProtocolFee`:
	if err = encoder.Encode(obj.ProtocolFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFee:%w", err)
	}
	// Serialize `PlatformFee`:
	if err = encoder.Encode(obj.PlatformFee); err != nil {
		return fmt.Errorf("error while marshaling PlatformFee:%w", err)
	}
	// Serialize `CreatorFee`:
	if err = encoder.Encode(obj.CreatorFee); err != nil {
		return fmt.Errorf("error while marshaling CreatorFee:%w", err)
	}
	// Serialize `ShareFee`:
	if err = encoder.Encode(obj.ShareFee); err != nil {
		return fmt.Errorf("error while marshaling ShareFee:%w", err)
	}
	// Serialize `TradeDirection`:
	if err = encoder.Encode(obj.TradeDirection); err != nil {
		return fmt.Errorf("error while marshaling TradeDirection:%w", err)
	}
	// Serialize `PoolStatus`:
	if err = encoder.Encode(obj.PoolStatus); err != nil {
		return fmt.Errorf("error while marshaling PoolStatus:%w", err)
	}
	// Serialize `ExactIn`:
	if err = encoder.Encode(obj.ExactIn); err != nil {
		return fmt.Errorf("error while marshaling ExactIn:%w", err)
	}
	return nil
}

func (obj TradeEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding TradeEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *TradeEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_TradeEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_TradeEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolState`:
	if err = decoder.Decode(&obj.PoolState); err != nil {
		return fmt.Errorf("error while unmarshaling PoolState:%w", err)
	}
	// Deserialize `TotalBaseSell`:
	if err = decoder.Decode(&obj.TotalBaseSell); err != nil {
		return fmt.Errorf("error while unmarshaling TotalBaseSell:%w", err)
	}
	// Deserialize `VirtualBase`:
	if err = decoder.Decode(&obj.VirtualBase); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualBase:%w", err)
	}
	// Deserialize `VirtualQuote`:
	if err = decoder.Decode(&obj.VirtualQuote); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualQuote:%w", err)
	}
	// Deserialize `RealBaseBefore`:
	if err = decoder.Decode(&obj.RealBaseBefore); err != nil {
		return fmt.Errorf("error while unmarshaling RealBaseBefore:%w", err)
	}
	// Deserialize `RealQuoteBefore`:
	if err = decoder.Decode(&obj.RealQuoteBefore); err != nil {
		return fmt.Errorf("error while unmarshaling RealQuoteBefore:%w", err)
	}
	// Deserialize `RealBaseAfter`:
	if err = decoder.Decode(&obj.RealBaseAfter); err != nil {
		return fmt.Errorf("error while unmarshaling RealBaseAfter:%w", err)
	}
	// Deserialize `RealQuoteAfter`:
	if err = decoder.Decode(&obj.RealQuoteAfter); err != nil {
		return fmt.Errorf("error while unmarshaling RealQuoteAfter:%w", err)
	}
	// Deserialize `AmountIn`:
	if err = decoder.Decode(&obj.AmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling AmountIn:%w", err)
	}
	// Deserialize `AmountOut`:
	if err = decoder.Decode(&obj.AmountOut); err != nil {
		return fmt.Errorf("error while unmarshaling AmountOut:%w", err)
	}
	// Deserialize `ProtocolFee`:
	if err = decoder.Decode(&obj.ProtocolFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFee:%w", err)
	}
	// Deserialize `PlatformFee`:
	if err = decoder.Decode(&obj.PlatformFee); err != nil {
		return fmt.Errorf("error while unmarshaling PlatformFee:%w", err)
	}
	// Deserialize `CreatorFee`:
	if err = decoder.Decode(&obj.CreatorFee); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorFee:%w", err)
	}
	// Deserialize `ShareFee`:
	if err = decoder.Decode(&obj.ShareFee); err != nil {
		return fmt.Errorf("error while unmarshaling ShareFee:%w", err)
	}
	// Deserialize `TradeDirection`:
	if err = decoder.Decode(&obj.TradeDirection); err != nil {
		return fmt.Errorf("error while unmarshaling TradeDirection:%w", err)
	}
	// Deserialize `PoolStatus`:
	if err = decoder.Decode(&obj.PoolStatus); err != nil {
		return fmt.Errorf("error while unmarshaling PoolStatus:%w", err)
	}
	// Deserialize `ExactIn`:
	if err = decoder.Decode(&obj.ExactIn); err != nil {
		return fmt.Errorf("error while unmarshaling ExactIn:%w", err)
	}
	return nil
}

func (obj *TradeEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling TradeEvent: %w", err)
	}
	return nil
}

func UnmarshalTradeEvent(buf []byte) (*TradeEvent, error) {
	obj := new(TradeEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
