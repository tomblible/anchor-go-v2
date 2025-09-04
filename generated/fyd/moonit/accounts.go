package moonit

import (
	"bytes"
	"fmt"

	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type ConfigAccount struct {
	MigrationAuthority        solanago.PublicKey
	BackendAuthority          solanago.PublicKey
	ConfigAuthority           solanago.PublicKey
	HelioFee                  solanago.PublicKey
	DexFee                    solanago.PublicKey
	FeeBps                    uint16
	DexFeeShare               uint8
	MigrationFee              uint64
	MarketcapThreshold        uint64
	MarketcapCurrency         Currency
	MinSupportedDecimalPlaces uint8
	MaxSupportedDecimalPlaces uint8
	MinSupportedTokenSupply   uint64
	MaxSupportedTokenSupply   uint64
	Bump                      uint8
	CoefB                     uint32
}

func (obj ConfigAccount) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_ConfigAccount[:], false)
	if err != nil {
		return err
	}
	// Serialize `MigrationAuthority`:
	if err = encoder.Encode(obj.MigrationAuthority); err != nil {
		return fmt.Errorf("error while marshaling MigrationAuthority:%w", err)
	}
	// Serialize `BackendAuthority`:
	if err = encoder.Encode(obj.BackendAuthority); err != nil {
		return fmt.Errorf("error while marshaling BackendAuthority:%w", err)
	}
	// Serialize `ConfigAuthority`:
	if err = encoder.Encode(obj.ConfigAuthority); err != nil {
		return fmt.Errorf("error while marshaling ConfigAuthority:%w", err)
	}
	// Serialize `HelioFee`:
	if err = encoder.Encode(obj.HelioFee); err != nil {
		return fmt.Errorf("error while marshaling HelioFee:%w", err)
	}
	// Serialize `DexFee`:
	if err = encoder.Encode(obj.DexFee); err != nil {
		return fmt.Errorf("error while marshaling DexFee:%w", err)
	}
	// Serialize `FeeBps`:
	if err = encoder.Encode(obj.FeeBps); err != nil {
		return fmt.Errorf("error while marshaling FeeBps:%w", err)
	}
	// Serialize `DexFeeShare`:
	if err = encoder.Encode(obj.DexFeeShare); err != nil {
		return fmt.Errorf("error while marshaling DexFeeShare:%w", err)
	}
	// Serialize `MigrationFee`:
	if err = encoder.Encode(obj.MigrationFee); err != nil {
		return fmt.Errorf("error while marshaling MigrationFee:%w", err)
	}
	// Serialize `MarketcapThreshold`:
	if err = encoder.Encode(obj.MarketcapThreshold); err != nil {
		return fmt.Errorf("error while marshaling MarketcapThreshold:%w", err)
	}
	// Serialize `MarketcapCurrency`:
	if err = encoder.Encode(obj.MarketcapCurrency); err != nil {
		return fmt.Errorf("error while marshaling MarketcapCurrency:%w", err)
	}
	// Serialize `MinSupportedDecimalPlaces`:
	if err = encoder.Encode(obj.MinSupportedDecimalPlaces); err != nil {
		return fmt.Errorf("error while marshaling MinSupportedDecimalPlaces:%w", err)
	}
	// Serialize `MaxSupportedDecimalPlaces`:
	if err = encoder.Encode(obj.MaxSupportedDecimalPlaces); err != nil {
		return fmt.Errorf("error while marshaling MaxSupportedDecimalPlaces:%w", err)
	}
	// Serialize `MinSupportedTokenSupply`:
	if err = encoder.Encode(obj.MinSupportedTokenSupply); err != nil {
		return fmt.Errorf("error while marshaling MinSupportedTokenSupply:%w", err)
	}
	// Serialize `MaxSupportedTokenSupply`:
	if err = encoder.Encode(obj.MaxSupportedTokenSupply); err != nil {
		return fmt.Errorf("error while marshaling MaxSupportedTokenSupply:%w", err)
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	// Serialize `CoefB`:
	if err = encoder.Encode(obj.CoefB); err != nil {
		return fmt.Errorf("error while marshaling CoefB:%w", err)
	}
	return nil
}

func (obj ConfigAccount) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding ConfigAccount: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *ConfigAccount) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_ConfigAccount[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_ConfigAccount[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `MigrationAuthority`:
	if err = decoder.Decode(&obj.MigrationAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationAuthority:%w", err)
	}
	// Deserialize `BackendAuthority`:
	if err = decoder.Decode(&obj.BackendAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling BackendAuthority:%w", err)
	}
	// Deserialize `ConfigAuthority`:
	if err = decoder.Decode(&obj.ConfigAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling ConfigAuthority:%w", err)
	}
	// Deserialize `HelioFee`:
	if err = decoder.Decode(&obj.HelioFee); err != nil {
		return fmt.Errorf("error while unmarshaling HelioFee:%w", err)
	}
	// Deserialize `DexFee`:
	if err = decoder.Decode(&obj.DexFee); err != nil {
		return fmt.Errorf("error while unmarshaling DexFee:%w", err)
	}
	// Deserialize `FeeBps`:
	if err = decoder.Decode(&obj.FeeBps); err != nil {
		return fmt.Errorf("error while unmarshaling FeeBps:%w", err)
	}
	// Deserialize `DexFeeShare`:
	if err = decoder.Decode(&obj.DexFeeShare); err != nil {
		return fmt.Errorf("error while unmarshaling DexFeeShare:%w", err)
	}
	// Deserialize `MigrationFee`:
	if err = decoder.Decode(&obj.MigrationFee); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationFee:%w", err)
	}
	// Deserialize `MarketcapThreshold`:
	if err = decoder.Decode(&obj.MarketcapThreshold); err != nil {
		return fmt.Errorf("error while unmarshaling MarketcapThreshold:%w", err)
	}
	// Deserialize `MarketcapCurrency`:
	if err = decoder.Decode(&obj.MarketcapCurrency); err != nil {
		return fmt.Errorf("error while unmarshaling MarketcapCurrency:%w", err)
	}
	// Deserialize `MinSupportedDecimalPlaces`:
	if err = decoder.Decode(&obj.MinSupportedDecimalPlaces); err != nil {
		return fmt.Errorf("error while unmarshaling MinSupportedDecimalPlaces:%w", err)
	}
	// Deserialize `MaxSupportedDecimalPlaces`:
	if err = decoder.Decode(&obj.MaxSupportedDecimalPlaces); err != nil {
		return fmt.Errorf("error while unmarshaling MaxSupportedDecimalPlaces:%w", err)
	}
	// Deserialize `MinSupportedTokenSupply`:
	if err = decoder.Decode(&obj.MinSupportedTokenSupply); err != nil {
		return fmt.Errorf("error while unmarshaling MinSupportedTokenSupply:%w", err)
	}
	// Deserialize `MaxSupportedTokenSupply`:
	if err = decoder.Decode(&obj.MaxSupportedTokenSupply); err != nil {
		return fmt.Errorf("error while unmarshaling MaxSupportedTokenSupply:%w", err)
	}
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	// Deserialize `CoefB`:
	if err = decoder.Decode(&obj.CoefB); err != nil {
		return fmt.Errorf("error while unmarshaling CoefB:%w", err)
	}
	return nil
}

func (obj *ConfigAccount) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling ConfigAccount: %w", err)
	}
	return nil
}

func UnmarshalConfigAccount(buf []byte) (*ConfigAccount, error) {
	obj := new(ConfigAccount)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type CurveAccount struct {
	TotalSupply        uint64
	CurveAmount        uint64
	Mint               solanago.PublicKey
	Decimals           uint8
	CollateralCurrency Currency
	CurveType          CurveType
	MarketcapThreshold uint64
	MarketcapCurrency  Currency
	MigrationFee       uint64
	CoefB              uint32
	Bump               uint8
	MigrationTarget    MigrationTarget
}

func (obj CurveAccount) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_CurveAccount[:], false)
	if err != nil {
		return err
	}
	// Serialize `TotalSupply`:
	if err = encoder.Encode(obj.TotalSupply); err != nil {
		return fmt.Errorf("error while marshaling TotalSupply:%w", err)
	}
	// Serialize `CurveAmount`:
	if err = encoder.Encode(obj.CurveAmount); err != nil {
		return fmt.Errorf("error while marshaling CurveAmount:%w", err)
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `Decimals`:
	if err = encoder.Encode(obj.Decimals); err != nil {
		return fmt.Errorf("error while marshaling Decimals:%w", err)
	}
	// Serialize `CollateralCurrency`:
	if err = encoder.Encode(obj.CollateralCurrency); err != nil {
		return fmt.Errorf("error while marshaling CollateralCurrency:%w", err)
	}
	// Serialize `CurveType`:
	if err = encoder.Encode(obj.CurveType); err != nil {
		return fmt.Errorf("error while marshaling CurveType:%w", err)
	}
	// Serialize `MarketcapThreshold`:
	if err = encoder.Encode(obj.MarketcapThreshold); err != nil {
		return fmt.Errorf("error while marshaling MarketcapThreshold:%w", err)
	}
	// Serialize `MarketcapCurrency`:
	if err = encoder.Encode(obj.MarketcapCurrency); err != nil {
		return fmt.Errorf("error while marshaling MarketcapCurrency:%w", err)
	}
	// Serialize `MigrationFee`:
	if err = encoder.Encode(obj.MigrationFee); err != nil {
		return fmt.Errorf("error while marshaling MigrationFee:%w", err)
	}
	// Serialize `CoefB`:
	if err = encoder.Encode(obj.CoefB); err != nil {
		return fmt.Errorf("error while marshaling CoefB:%w", err)
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	// Serialize `MigrationTarget`:
	if err = encoder.Encode(obj.MigrationTarget); err != nil {
		return fmt.Errorf("error while marshaling MigrationTarget:%w", err)
	}
	return nil
}

func (obj CurveAccount) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding CurveAccount: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *CurveAccount) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_CurveAccount[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_CurveAccount[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TotalSupply`:
	if err = decoder.Decode(&obj.TotalSupply); err != nil {
		return fmt.Errorf("error while unmarshaling TotalSupply:%w", err)
	}
	// Deserialize `CurveAmount`:
	if err = decoder.Decode(&obj.CurveAmount); err != nil {
		return fmt.Errorf("error while unmarshaling CurveAmount:%w", err)
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `Decimals`:
	if err = decoder.Decode(&obj.Decimals); err != nil {
		return fmt.Errorf("error while unmarshaling Decimals:%w", err)
	}
	// Deserialize `CollateralCurrency`:
	if err = decoder.Decode(&obj.CollateralCurrency); err != nil {
		return fmt.Errorf("error while unmarshaling CollateralCurrency:%w", err)
	}
	// Deserialize `CurveType`:
	if err = decoder.Decode(&obj.CurveType); err != nil {
		return fmt.Errorf("error while unmarshaling CurveType:%w", err)
	}
	// Deserialize `MarketcapThreshold`:
	if err = decoder.Decode(&obj.MarketcapThreshold); err != nil {
		return fmt.Errorf("error while unmarshaling MarketcapThreshold:%w", err)
	}
	// Deserialize `MarketcapCurrency`:
	if err = decoder.Decode(&obj.MarketcapCurrency); err != nil {
		return fmt.Errorf("error while unmarshaling MarketcapCurrency:%w", err)
	}
	// Deserialize `MigrationFee`:
	if err = decoder.Decode(&obj.MigrationFee); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationFee:%w", err)
	}
	// Deserialize `CoefB`:
	if err = decoder.Decode(&obj.CoefB); err != nil {
		return fmt.Errorf("error while unmarshaling CoefB:%w", err)
	}
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	// Deserialize `MigrationTarget`:
	if err = decoder.Decode(&obj.MigrationTarget); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationTarget:%w", err)
	}
	return nil
}

func (obj *CurveAccount) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling CurveAccount: %w", err)
	}
	return nil
}

func UnmarshalCurveAccount(buf []byte) (*CurveAccount, error) {
	obj := new(CurveAccount)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
