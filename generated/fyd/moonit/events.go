package moonit

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type TradeEvent struct {
	Amount           uint64
	CollateralAmount uint64
	DexFee           uint64
	HelioFee         uint64
	Allocation       uint64
	Curve            solanago.PublicKey
	CostToken        solanago.PublicKey
	Sender           solanago.PublicKey
	Type             TradeType
	Label            string
}

func (obj TradeEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_TradeEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Amount`:
	if err = encoder.Encode(obj.Amount); err != nil {
		return fmt.Errorf("error while marshaling Amount:%w", err)
	}
	// Serialize `CollateralAmount`:
	if err = encoder.Encode(obj.CollateralAmount); err != nil {
		return fmt.Errorf("error while marshaling CollateralAmount:%w", err)
	}
	// Serialize `DexFee`:
	if err = encoder.Encode(obj.DexFee); err != nil {
		return fmt.Errorf("error while marshaling DexFee:%w", err)
	}
	// Serialize `HelioFee`:
	if err = encoder.Encode(obj.HelioFee); err != nil {
		return fmt.Errorf("error while marshaling HelioFee:%w", err)
	}
	// Serialize `Allocation`:
	if err = encoder.Encode(obj.Allocation); err != nil {
		return fmt.Errorf("error while marshaling Allocation:%w", err)
	}
	// Serialize `Curve`:
	if err = encoder.Encode(obj.Curve); err != nil {
		return fmt.Errorf("error while marshaling Curve:%w", err)
	}
	// Serialize `CostToken`:
	if err = encoder.Encode(obj.CostToken); err != nil {
		return fmt.Errorf("error while marshaling CostToken:%w", err)
	}
	// Serialize `Sender`:
	if err = encoder.Encode(obj.Sender); err != nil {
		return fmt.Errorf("error while marshaling Sender:%w", err)
	}
	// Serialize `Type`:
	if err = encoder.Encode(obj.Type); err != nil {
		return fmt.Errorf("error while marshaling Type:%w", err)
	}
	// Serialize `Label`:
	if err = encoder.Encode(obj.Label); err != nil {
		return fmt.Errorf("error while marshaling Label:%w", err)
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
	// Deserialize `Amount`:
	if err = decoder.Decode(&obj.Amount); err != nil {
		return fmt.Errorf("error while unmarshaling Amount:%w", err)
	}
	// Deserialize `CollateralAmount`:
	if err = decoder.Decode(&obj.CollateralAmount); err != nil {
		return fmt.Errorf("error while unmarshaling CollateralAmount:%w", err)
	}
	// Deserialize `DexFee`:
	if err = decoder.Decode(&obj.DexFee); err != nil {
		return fmt.Errorf("error while unmarshaling DexFee:%w", err)
	}
	// Deserialize `HelioFee`:
	if err = decoder.Decode(&obj.HelioFee); err != nil {
		return fmt.Errorf("error while unmarshaling HelioFee:%w", err)
	}
	// Deserialize `Allocation`:
	if err = decoder.Decode(&obj.Allocation); err != nil {
		return fmt.Errorf("error while unmarshaling Allocation:%w", err)
	}
	// Deserialize `Curve`:
	if err = decoder.Decode(&obj.Curve); err != nil {
		return fmt.Errorf("error while unmarshaling Curve:%w", err)
	}
	// Deserialize `CostToken`:
	if err = decoder.Decode(&obj.CostToken); err != nil {
		return fmt.Errorf("error while unmarshaling CostToken:%w", err)
	}
	// Deserialize `Sender`:
	if err = decoder.Decode(&obj.Sender); err != nil {
		return fmt.Errorf("error while unmarshaling Sender:%w", err)
	}
	// Deserialize `Type`:
	if err = decoder.Decode(&obj.Type); err != nil {
		return fmt.Errorf("error while unmarshaling Type:%w", err)
	}
	// Deserialize `Label`:
	if err = decoder.Decode(&obj.Label); err != nil {
		return fmt.Errorf("error while unmarshaling Label:%w", err)
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

type MigrationEvent struct {
	TokensMigrated     uint64
	TokensBurned       uint64
	CollateralMigrated uint64
	Fee                uint64
	Label              string
}

func (obj MigrationEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_MigrationEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `TokensMigrated`:
	if err = encoder.Encode(obj.TokensMigrated); err != nil {
		return fmt.Errorf("error while marshaling TokensMigrated:%w", err)
	}
	// Serialize `TokensBurned`:
	if err = encoder.Encode(obj.TokensBurned); err != nil {
		return fmt.Errorf("error while marshaling TokensBurned:%w", err)
	}
	// Serialize `CollateralMigrated`:
	if err = encoder.Encode(obj.CollateralMigrated); err != nil {
		return fmt.Errorf("error while marshaling CollateralMigrated:%w", err)
	}
	// Serialize `Fee`:
	if err = encoder.Encode(obj.Fee); err != nil {
		return fmt.Errorf("error while marshaling Fee:%w", err)
	}
	// Serialize `Label`:
	if err = encoder.Encode(obj.Label); err != nil {
		return fmt.Errorf("error while marshaling Label:%w", err)
	}
	return nil
}

func (obj MigrationEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding MigrationEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *MigrationEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_MigrationEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_MigrationEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TokensMigrated`:
	if err = decoder.Decode(&obj.TokensMigrated); err != nil {
		return fmt.Errorf("error while unmarshaling TokensMigrated:%w", err)
	}
	// Deserialize `TokensBurned`:
	if err = decoder.Decode(&obj.TokensBurned); err != nil {
		return fmt.Errorf("error while unmarshaling TokensBurned:%w", err)
	}
	// Deserialize `CollateralMigrated`:
	if err = decoder.Decode(&obj.CollateralMigrated); err != nil {
		return fmt.Errorf("error while unmarshaling CollateralMigrated:%w", err)
	}
	// Deserialize `Fee`:
	if err = decoder.Decode(&obj.Fee); err != nil {
		return fmt.Errorf("error while unmarshaling Fee:%w", err)
	}
	// Deserialize `Label`:
	if err = decoder.Decode(&obj.Label); err != nil {
		return fmt.Errorf("error while unmarshaling Label:%w", err)
	}
	return nil
}

func (obj *MigrationEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling MigrationEvent: %w", err)
	}
	return nil
}

func UnmarshalMigrationEvent(buf []byte) (*MigrationEvent, error) {
	obj := new(MigrationEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
