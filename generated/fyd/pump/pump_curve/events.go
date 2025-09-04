package pump_curve

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type AdminSetCreatorEvent struct {
	Timestamp                int64
	AdminSetCreatorAuthority solanago.PublicKey
	Mint                     solanago.PublicKey
	BondingCurve             solanago.PublicKey
	OldCreator               solanago.PublicKey
	NewCreator               solanago.PublicKey
}

func (obj AdminSetCreatorEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_AdminSetCreatorEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `AdminSetCreatorAuthority`:
	if err = encoder.Encode(obj.AdminSetCreatorAuthority); err != nil {
		return fmt.Errorf("error while marshaling AdminSetCreatorAuthority:%w", err)
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `BondingCurve`:
	if err = encoder.Encode(obj.BondingCurve); err != nil {
		return fmt.Errorf("error while marshaling BondingCurve:%w", err)
	}
	// Serialize `OldCreator`:
	if err = encoder.Encode(obj.OldCreator); err != nil {
		return fmt.Errorf("error while marshaling OldCreator:%w", err)
	}
	// Serialize `NewCreator`:
	if err = encoder.Encode(obj.NewCreator); err != nil {
		return fmt.Errorf("error while marshaling NewCreator:%w", err)
	}
	return nil
}

func (obj AdminSetCreatorEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding AdminSetCreatorEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *AdminSetCreatorEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_AdminSetCreatorEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_AdminSetCreatorEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `AdminSetCreatorAuthority`:
	if err = decoder.Decode(&obj.AdminSetCreatorAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling AdminSetCreatorAuthority:%w", err)
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `BondingCurve`:
	if err = decoder.Decode(&obj.BondingCurve); err != nil {
		return fmt.Errorf("error while unmarshaling BondingCurve:%w", err)
	}
	// Deserialize `OldCreator`:
	if err = decoder.Decode(&obj.OldCreator); err != nil {
		return fmt.Errorf("error while unmarshaling OldCreator:%w", err)
	}
	// Deserialize `NewCreator`:
	if err = decoder.Decode(&obj.NewCreator); err != nil {
		return fmt.Errorf("error while unmarshaling NewCreator:%w", err)
	}
	return nil
}

func (obj *AdminSetCreatorEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling AdminSetCreatorEvent: %w", err)
	}
	return nil
}

func UnmarshalAdminSetCreatorEvent(buf []byte) (*AdminSetCreatorEvent, error) {
	obj := new(AdminSetCreatorEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type AdminSetIdlAuthorityEvent struct {
	IdlAuthority solanago.PublicKey
}

func (obj AdminSetIdlAuthorityEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_AdminSetIdlAuthorityEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `IdlAuthority`:
	if err = encoder.Encode(obj.IdlAuthority); err != nil {
		return fmt.Errorf("error while marshaling IdlAuthority:%w", err)
	}
	return nil
}

func (obj AdminSetIdlAuthorityEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding AdminSetIdlAuthorityEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *AdminSetIdlAuthorityEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_AdminSetIdlAuthorityEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_AdminSetIdlAuthorityEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `IdlAuthority`:
	if err = decoder.Decode(&obj.IdlAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling IdlAuthority:%w", err)
	}
	return nil
}

func (obj *AdminSetIdlAuthorityEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling AdminSetIdlAuthorityEvent: %w", err)
	}
	return nil
}

func UnmarshalAdminSetIdlAuthorityEvent(buf []byte) (*AdminSetIdlAuthorityEvent, error) {
	obj := new(AdminSetIdlAuthorityEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type AdminUpdateTokenIncentivesEvent struct {
	StartTime         int64
	EndTime           int64
	DayNumber         uint64
	TokenSupplyPerDay uint64
	Mint              solanago.PublicKey
	SecondsInADay     int64
	Timestamp         int64
}

func (obj AdminUpdateTokenIncentivesEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_AdminUpdateTokenIncentivesEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `StartTime`:
	if err = encoder.Encode(obj.StartTime); err != nil {
		return fmt.Errorf("error while marshaling StartTime:%w", err)
	}
	// Serialize `EndTime`:
	if err = encoder.Encode(obj.EndTime); err != nil {
		return fmt.Errorf("error while marshaling EndTime:%w", err)
	}
	// Serialize `DayNumber`:
	if err = encoder.Encode(obj.DayNumber); err != nil {
		return fmt.Errorf("error while marshaling DayNumber:%w", err)
	}
	// Serialize `TokenSupplyPerDay`:
	if err = encoder.Encode(obj.TokenSupplyPerDay); err != nil {
		return fmt.Errorf("error while marshaling TokenSupplyPerDay:%w", err)
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `SecondsInADay`:
	if err = encoder.Encode(obj.SecondsInADay); err != nil {
		return fmt.Errorf("error while marshaling SecondsInADay:%w", err)
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	return nil
}

func (obj AdminUpdateTokenIncentivesEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding AdminUpdateTokenIncentivesEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *AdminUpdateTokenIncentivesEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_AdminUpdateTokenIncentivesEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_AdminUpdateTokenIncentivesEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `StartTime`:
	if err = decoder.Decode(&obj.StartTime); err != nil {
		return fmt.Errorf("error while unmarshaling StartTime:%w", err)
	}
	// Deserialize `EndTime`:
	if err = decoder.Decode(&obj.EndTime); err != nil {
		return fmt.Errorf("error while unmarshaling EndTime:%w", err)
	}
	// Deserialize `DayNumber`:
	if err = decoder.Decode(&obj.DayNumber); err != nil {
		return fmt.Errorf("error while unmarshaling DayNumber:%w", err)
	}
	// Deserialize `TokenSupplyPerDay`:
	if err = decoder.Decode(&obj.TokenSupplyPerDay); err != nil {
		return fmt.Errorf("error while unmarshaling TokenSupplyPerDay:%w", err)
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `SecondsInADay`:
	if err = decoder.Decode(&obj.SecondsInADay); err != nil {
		return fmt.Errorf("error while unmarshaling SecondsInADay:%w", err)
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	return nil
}

func (obj *AdminUpdateTokenIncentivesEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling AdminUpdateTokenIncentivesEvent: %w", err)
	}
	return nil
}

func UnmarshalAdminUpdateTokenIncentivesEvent(buf []byte) (*AdminUpdateTokenIncentivesEvent, error) {
	obj := new(AdminUpdateTokenIncentivesEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type ClaimTokenIncentivesEvent struct {
	User               solanago.PublicKey
	Mint               solanago.PublicKey
	Amount             uint64
	Timestamp          int64
	TotalClaimedTokens uint64
	CurrentSolVolume   uint64
}

func (obj ClaimTokenIncentivesEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_ClaimTokenIncentivesEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `User`:
	if err = encoder.Encode(obj.User); err != nil {
		return fmt.Errorf("error while marshaling User:%w", err)
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `Amount`:
	if err = encoder.Encode(obj.Amount); err != nil {
		return fmt.Errorf("error while marshaling Amount:%w", err)
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `TotalClaimedTokens`:
	if err = encoder.Encode(obj.TotalClaimedTokens); err != nil {
		return fmt.Errorf("error while marshaling TotalClaimedTokens:%w", err)
	}
	// Serialize `CurrentSolVolume`:
	if err = encoder.Encode(obj.CurrentSolVolume); err != nil {
		return fmt.Errorf("error while marshaling CurrentSolVolume:%w", err)
	}
	return nil
}

func (obj ClaimTokenIncentivesEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding ClaimTokenIncentivesEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *ClaimTokenIncentivesEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_ClaimTokenIncentivesEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_ClaimTokenIncentivesEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `User`:
	if err = decoder.Decode(&obj.User); err != nil {
		return fmt.Errorf("error while unmarshaling User:%w", err)
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `Amount`:
	if err = decoder.Decode(&obj.Amount); err != nil {
		return fmt.Errorf("error while unmarshaling Amount:%w", err)
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `TotalClaimedTokens`:
	if err = decoder.Decode(&obj.TotalClaimedTokens); err != nil {
		return fmt.Errorf("error while unmarshaling TotalClaimedTokens:%w", err)
	}
	// Deserialize `CurrentSolVolume`:
	if err = decoder.Decode(&obj.CurrentSolVolume); err != nil {
		return fmt.Errorf("error while unmarshaling CurrentSolVolume:%w", err)
	}
	return nil
}

func (obj *ClaimTokenIncentivesEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling ClaimTokenIncentivesEvent: %w", err)
	}
	return nil
}

func UnmarshalClaimTokenIncentivesEvent(buf []byte) (*ClaimTokenIncentivesEvent, error) {
	obj := new(ClaimTokenIncentivesEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type CloseUserVolumeAccumulatorEvent struct {
	User                 solanago.PublicKey
	Timestamp            int64
	TotalUnclaimedTokens uint64
	TotalClaimedTokens   uint64
	CurrentSolVolume     uint64
	LastUpdateTimestamp  int64
}

func (obj CloseUserVolumeAccumulatorEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CloseUserVolumeAccumulatorEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `User`:
	if err = encoder.Encode(obj.User); err != nil {
		return fmt.Errorf("error while marshaling User:%w", err)
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `TotalUnclaimedTokens`:
	if err = encoder.Encode(obj.TotalUnclaimedTokens); err != nil {
		return fmt.Errorf("error while marshaling TotalUnclaimedTokens:%w", err)
	}
	// Serialize `TotalClaimedTokens`:
	if err = encoder.Encode(obj.TotalClaimedTokens); err != nil {
		return fmt.Errorf("error while marshaling TotalClaimedTokens:%w", err)
	}
	// Serialize `CurrentSolVolume`:
	if err = encoder.Encode(obj.CurrentSolVolume); err != nil {
		return fmt.Errorf("error while marshaling CurrentSolVolume:%w", err)
	}
	// Serialize `LastUpdateTimestamp`:
	if err = encoder.Encode(obj.LastUpdateTimestamp); err != nil {
		return fmt.Errorf("error while marshaling LastUpdateTimestamp:%w", err)
	}
	return nil
}

func (obj CloseUserVolumeAccumulatorEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding CloseUserVolumeAccumulatorEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *CloseUserVolumeAccumulatorEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CloseUserVolumeAccumulatorEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CloseUserVolumeAccumulatorEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `User`:
	if err = decoder.Decode(&obj.User); err != nil {
		return fmt.Errorf("error while unmarshaling User:%w", err)
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `TotalUnclaimedTokens`:
	if err = decoder.Decode(&obj.TotalUnclaimedTokens); err != nil {
		return fmt.Errorf("error while unmarshaling TotalUnclaimedTokens:%w", err)
	}
	// Deserialize `TotalClaimedTokens`:
	if err = decoder.Decode(&obj.TotalClaimedTokens); err != nil {
		return fmt.Errorf("error while unmarshaling TotalClaimedTokens:%w", err)
	}
	// Deserialize `CurrentSolVolume`:
	if err = decoder.Decode(&obj.CurrentSolVolume); err != nil {
		return fmt.Errorf("error while unmarshaling CurrentSolVolume:%w", err)
	}
	// Deserialize `LastUpdateTimestamp`:
	if err = decoder.Decode(&obj.LastUpdateTimestamp); err != nil {
		return fmt.Errorf("error while unmarshaling LastUpdateTimestamp:%w", err)
	}
	return nil
}

func (obj *CloseUserVolumeAccumulatorEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling CloseUserVolumeAccumulatorEvent: %w", err)
	}
	return nil
}

func UnmarshalCloseUserVolumeAccumulatorEvent(buf []byte) (*CloseUserVolumeAccumulatorEvent, error) {
	obj := new(CloseUserVolumeAccumulatorEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type CollectCreatorFeeEvent struct {
	Timestamp  int64
	Creator    solanago.PublicKey
	CreatorFee uint64
}

func (obj CollectCreatorFeeEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CollectCreatorFeeEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
	}
	// Serialize `CreatorFee`:
	if err = encoder.Encode(obj.CreatorFee); err != nil {
		return fmt.Errorf("error while marshaling CreatorFee:%w", err)
	}
	return nil
}

func (obj CollectCreatorFeeEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding CollectCreatorFeeEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *CollectCreatorFeeEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CollectCreatorFeeEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CollectCreatorFeeEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
	}
	// Deserialize `CreatorFee`:
	if err = decoder.Decode(&obj.CreatorFee); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorFee:%w", err)
	}
	return nil
}

func (obj *CollectCreatorFeeEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling CollectCreatorFeeEvent: %w", err)
	}
	return nil
}

func UnmarshalCollectCreatorFeeEvent(buf []byte) (*CollectCreatorFeeEvent, error) {
	obj := new(CollectCreatorFeeEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type CompleteEvent struct {
	User         solanago.PublicKey
	Mint         solanago.PublicKey
	BondingCurve solanago.PublicKey
	Timestamp    int64
}

func (obj CompleteEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CompleteEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `User`:
	if err = encoder.Encode(obj.User); err != nil {
		return fmt.Errorf("error while marshaling User:%w", err)
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `BondingCurve`:
	if err = encoder.Encode(obj.BondingCurve); err != nil {
		return fmt.Errorf("error while marshaling BondingCurve:%w", err)
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	return nil
}

func (obj CompleteEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding CompleteEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *CompleteEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CompleteEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CompleteEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `User`:
	if err = decoder.Decode(&obj.User); err != nil {
		return fmt.Errorf("error while unmarshaling User:%w", err)
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `BondingCurve`:
	if err = decoder.Decode(&obj.BondingCurve); err != nil {
		return fmt.Errorf("error while unmarshaling BondingCurve:%w", err)
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	return nil
}

func (obj *CompleteEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling CompleteEvent: %w", err)
	}
	return nil
}

func UnmarshalCompleteEvent(buf []byte) (*CompleteEvent, error) {
	obj := new(CompleteEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type CompletePumpAmmMigrationEvent struct {
	User             solanago.PublicKey
	Mint             solanago.PublicKey
	MintAmount       uint64
	SolAmount        uint64
	PoolMigrationFee uint64
	BondingCurve     solanago.PublicKey
	Timestamp        int64
	Pool             solanago.PublicKey
}

func (obj CompletePumpAmmMigrationEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CompletePumpAmmMigrationEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `User`:
	if err = encoder.Encode(obj.User); err != nil {
		return fmt.Errorf("error while marshaling User:%w", err)
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `MintAmount`:
	if err = encoder.Encode(obj.MintAmount); err != nil {
		return fmt.Errorf("error while marshaling MintAmount:%w", err)
	}
	// Serialize `SolAmount`:
	if err = encoder.Encode(obj.SolAmount); err != nil {
		return fmt.Errorf("error while marshaling SolAmount:%w", err)
	}
	// Serialize `PoolMigrationFee`:
	if err = encoder.Encode(obj.PoolMigrationFee); err != nil {
		return fmt.Errorf("error while marshaling PoolMigrationFee:%w", err)
	}
	// Serialize `BondingCurve`:
	if err = encoder.Encode(obj.BondingCurve); err != nil {
		return fmt.Errorf("error while marshaling BondingCurve:%w", err)
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	return nil
}

func (obj CompletePumpAmmMigrationEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding CompletePumpAmmMigrationEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *CompletePumpAmmMigrationEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CompletePumpAmmMigrationEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CompletePumpAmmMigrationEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `User`:
	if err = decoder.Decode(&obj.User); err != nil {
		return fmt.Errorf("error while unmarshaling User:%w", err)
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `MintAmount`:
	if err = decoder.Decode(&obj.MintAmount); err != nil {
		return fmt.Errorf("error while unmarshaling MintAmount:%w", err)
	}
	// Deserialize `SolAmount`:
	if err = decoder.Decode(&obj.SolAmount); err != nil {
		return fmt.Errorf("error while unmarshaling SolAmount:%w", err)
	}
	// Deserialize `PoolMigrationFee`:
	if err = decoder.Decode(&obj.PoolMigrationFee); err != nil {
		return fmt.Errorf("error while unmarshaling PoolMigrationFee:%w", err)
	}
	// Deserialize `BondingCurve`:
	if err = decoder.Decode(&obj.BondingCurve); err != nil {
		return fmt.Errorf("error while unmarshaling BondingCurve:%w", err)
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	return nil
}

func (obj *CompletePumpAmmMigrationEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling CompletePumpAmmMigrationEvent: %w", err)
	}
	return nil
}

func UnmarshalCompletePumpAmmMigrationEvent(buf []byte) (*CompletePumpAmmMigrationEvent, error) {
	obj := new(CompletePumpAmmMigrationEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type CreateEvent struct {
	Name                 string
	Symbol               string
	Uri                  string
	Mint                 solanago.PublicKey
	BondingCurve         solanago.PublicKey
	User                 solanago.PublicKey
	Creator              solanago.PublicKey
	Timestamp            int64
	VirtualTokenReserves uint64
	VirtualSolReserves   uint64
	RealTokenReserves    uint64
	TokenTotalSupply     uint64
}

func (obj CreateEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CreateEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Name`:
	if err = encoder.Encode(obj.Name); err != nil {
		return fmt.Errorf("error while marshaling Name:%w", err)
	}
	// Serialize `Symbol`:
	if err = encoder.Encode(obj.Symbol); err != nil {
		return fmt.Errorf("error while marshaling Symbol:%w", err)
	}
	// Serialize `Uri`:
	if err = encoder.Encode(obj.Uri); err != nil {
		return fmt.Errorf("error while marshaling Uri:%w", err)
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `BondingCurve`:
	if err = encoder.Encode(obj.BondingCurve); err != nil {
		return fmt.Errorf("error while marshaling BondingCurve:%w", err)
	}
	// Serialize `User`:
	if err = encoder.Encode(obj.User); err != nil {
		return fmt.Errorf("error while marshaling User:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `VirtualTokenReserves`:
	if err = encoder.Encode(obj.VirtualTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling VirtualTokenReserves:%w", err)
	}
	// Serialize `VirtualSolReserves`:
	if err = encoder.Encode(obj.VirtualSolReserves); err != nil {
		return fmt.Errorf("error while marshaling VirtualSolReserves:%w", err)
	}
	// Serialize `RealTokenReserves`:
	if err = encoder.Encode(obj.RealTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling RealTokenReserves:%w", err)
	}
	// Serialize `TokenTotalSupply`:
	if err = encoder.Encode(obj.TokenTotalSupply); err != nil {
		return fmt.Errorf("error while marshaling TokenTotalSupply:%w", err)
	}
	return nil
}

func (obj CreateEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding CreateEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *CreateEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CreateEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CreateEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Name`:
	if err = decoder.Decode(&obj.Name); err != nil {
		return fmt.Errorf("error while unmarshaling Name:%w", err)
	}
	// Deserialize `Symbol`:
	if err = decoder.Decode(&obj.Symbol); err != nil {
		return fmt.Errorf("error while unmarshaling Symbol:%w", err)
	}
	// Deserialize `Uri`:
	if err = decoder.Decode(&obj.Uri); err != nil {
		return fmt.Errorf("error while unmarshaling Uri:%w", err)
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `BondingCurve`:
	if err = decoder.Decode(&obj.BondingCurve); err != nil {
		return fmt.Errorf("error while unmarshaling BondingCurve:%w", err)
	}
	// Deserialize `User`:
	if err = decoder.Decode(&obj.User); err != nil {
		return fmt.Errorf("error while unmarshaling User:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `VirtualTokenReserves`:
	if err = decoder.Decode(&obj.VirtualTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualTokenReserves:%w", err)
	}
	// Deserialize `VirtualSolReserves`:
	if err = decoder.Decode(&obj.VirtualSolReserves); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualSolReserves:%w", err)
	}
	// Deserialize `RealTokenReserves`:
	if err = decoder.Decode(&obj.RealTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling RealTokenReserves:%w", err)
	}
	// Deserialize `TokenTotalSupply`:
	if err = decoder.Decode(&obj.TokenTotalSupply); err != nil {
		return fmt.Errorf("error while unmarshaling TokenTotalSupply:%w", err)
	}
	return nil
}

func (obj *CreateEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling CreateEvent: %w", err)
	}
	return nil
}

func UnmarshalCreateEvent(buf []byte) (*CreateEvent, error) {
	obj := new(CreateEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type ExtendAccountEvent struct {
	Account     solanago.PublicKey
	User        solanago.PublicKey
	CurrentSize uint64
	NewSize     uint64
	Timestamp   int64
}

func (obj ExtendAccountEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_ExtendAccountEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Account`:
	if err = encoder.Encode(obj.Account); err != nil {
		return fmt.Errorf("error while marshaling Account:%w", err)
	}
	// Serialize `User`:
	if err = encoder.Encode(obj.User); err != nil {
		return fmt.Errorf("error while marshaling User:%w", err)
	}
	// Serialize `CurrentSize`:
	if err = encoder.Encode(obj.CurrentSize); err != nil {
		return fmt.Errorf("error while marshaling CurrentSize:%w", err)
	}
	// Serialize `NewSize`:
	if err = encoder.Encode(obj.NewSize); err != nil {
		return fmt.Errorf("error while marshaling NewSize:%w", err)
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	return nil
}

func (obj ExtendAccountEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding ExtendAccountEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *ExtendAccountEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_ExtendAccountEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_ExtendAccountEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Account`:
	if err = decoder.Decode(&obj.Account); err != nil {
		return fmt.Errorf("error while unmarshaling Account:%w", err)
	}
	// Deserialize `User`:
	if err = decoder.Decode(&obj.User); err != nil {
		return fmt.Errorf("error while unmarshaling User:%w", err)
	}
	// Deserialize `CurrentSize`:
	if err = decoder.Decode(&obj.CurrentSize); err != nil {
		return fmt.Errorf("error while unmarshaling CurrentSize:%w", err)
	}
	// Deserialize `NewSize`:
	if err = decoder.Decode(&obj.NewSize); err != nil {
		return fmt.Errorf("error while unmarshaling NewSize:%w", err)
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	return nil
}

func (obj *ExtendAccountEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling ExtendAccountEvent: %w", err)
	}
	return nil
}

func UnmarshalExtendAccountEvent(buf []byte) (*ExtendAccountEvent, error) {
	obj := new(ExtendAccountEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type InitUserVolumeAccumulatorEvent struct {
	Payer     solanago.PublicKey
	User      solanago.PublicKey
	Timestamp int64
}

func (obj InitUserVolumeAccumulatorEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_InitUserVolumeAccumulatorEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Payer`:
	if err = encoder.Encode(obj.Payer); err != nil {
		return fmt.Errorf("error while marshaling Payer:%w", err)
	}
	// Serialize `User`:
	if err = encoder.Encode(obj.User); err != nil {
		return fmt.Errorf("error while marshaling User:%w", err)
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	return nil
}

func (obj InitUserVolumeAccumulatorEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding InitUserVolumeAccumulatorEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *InitUserVolumeAccumulatorEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_InitUserVolumeAccumulatorEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_InitUserVolumeAccumulatorEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Payer`:
	if err = decoder.Decode(&obj.Payer); err != nil {
		return fmt.Errorf("error while unmarshaling Payer:%w", err)
	}
	// Deserialize `User`:
	if err = decoder.Decode(&obj.User); err != nil {
		return fmt.Errorf("error while unmarshaling User:%w", err)
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	return nil
}

func (obj *InitUserVolumeAccumulatorEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling InitUserVolumeAccumulatorEvent: %w", err)
	}
	return nil
}

func UnmarshalInitUserVolumeAccumulatorEvent(buf []byte) (*InitUserVolumeAccumulatorEvent, error) {
	obj := new(InitUserVolumeAccumulatorEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type SetCreatorEvent struct {
	Timestamp    int64
	Mint         solanago.PublicKey
	BondingCurve solanago.PublicKey
	Creator      solanago.PublicKey
}

func (obj SetCreatorEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_SetCreatorEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `BondingCurve`:
	if err = encoder.Encode(obj.BondingCurve); err != nil {
		return fmt.Errorf("error while marshaling BondingCurve:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
	}
	return nil
}

func (obj SetCreatorEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding SetCreatorEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *SetCreatorEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_SetCreatorEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_SetCreatorEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `BondingCurve`:
	if err = decoder.Decode(&obj.BondingCurve); err != nil {
		return fmt.Errorf("error while unmarshaling BondingCurve:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
	}
	return nil
}

func (obj *SetCreatorEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling SetCreatorEvent: %w", err)
	}
	return nil
}

func UnmarshalSetCreatorEvent(buf []byte) (*SetCreatorEvent, error) {
	obj := new(SetCreatorEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type SetMetaplexCreatorEvent struct {
	Timestamp    int64
	Mint         solanago.PublicKey
	BondingCurve solanago.PublicKey
	Metadata     solanago.PublicKey
	Creator      solanago.PublicKey
}

func (obj SetMetaplexCreatorEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_SetMetaplexCreatorEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `BondingCurve`:
	if err = encoder.Encode(obj.BondingCurve); err != nil {
		return fmt.Errorf("error while marshaling BondingCurve:%w", err)
	}
	// Serialize `Metadata`:
	if err = encoder.Encode(obj.Metadata); err != nil {
		return fmt.Errorf("error while marshaling Metadata:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
	}
	return nil
}

func (obj SetMetaplexCreatorEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding SetMetaplexCreatorEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *SetMetaplexCreatorEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_SetMetaplexCreatorEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_SetMetaplexCreatorEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `BondingCurve`:
	if err = decoder.Decode(&obj.BondingCurve); err != nil {
		return fmt.Errorf("error while unmarshaling BondingCurve:%w", err)
	}
	// Deserialize `Metadata`:
	if err = decoder.Decode(&obj.Metadata); err != nil {
		return fmt.Errorf("error while unmarshaling Metadata:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
	}
	return nil
}

func (obj *SetMetaplexCreatorEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling SetMetaplexCreatorEvent: %w", err)
	}
	return nil
}

func UnmarshalSetMetaplexCreatorEvent(buf []byte) (*SetMetaplexCreatorEvent, error) {
	obj := new(SetMetaplexCreatorEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type SetParamsEvent struct {
	InitialVirtualTokenReserves uint64
	InitialVirtualSolReserves   uint64
	InitialRealTokenReserves    uint64
	FinalRealSolReserves        uint64
	TokenTotalSupply            uint64
	FeeBasisPoints              uint64
	WithdrawAuthority           solanago.PublicKey
	EnableMigrate               bool
	PoolMigrationFee            uint64
	CreatorFeeBasisPoints       uint64
	FeeRecipients               [8]solanago.PublicKey
	Timestamp                   int64
	SetCreatorAuthority         solanago.PublicKey
	AdminSetCreatorAuthority    solanago.PublicKey
}

func (obj SetParamsEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_SetParamsEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `InitialVirtualTokenReserves`:
	if err = encoder.Encode(obj.InitialVirtualTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling InitialVirtualTokenReserves:%w", err)
	}
	// Serialize `InitialVirtualSolReserves`:
	if err = encoder.Encode(obj.InitialVirtualSolReserves); err != nil {
		return fmt.Errorf("error while marshaling InitialVirtualSolReserves:%w", err)
	}
	// Serialize `InitialRealTokenReserves`:
	if err = encoder.Encode(obj.InitialRealTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling InitialRealTokenReserves:%w", err)
	}
	// Serialize `FinalRealSolReserves`:
	if err = encoder.Encode(obj.FinalRealSolReserves); err != nil {
		return fmt.Errorf("error while marshaling FinalRealSolReserves:%w", err)
	}
	// Serialize `TokenTotalSupply`:
	if err = encoder.Encode(obj.TokenTotalSupply); err != nil {
		return fmt.Errorf("error while marshaling TokenTotalSupply:%w", err)
	}
	// Serialize `FeeBasisPoints`:
	if err = encoder.Encode(obj.FeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling FeeBasisPoints:%w", err)
	}
	// Serialize `WithdrawAuthority`:
	if err = encoder.Encode(obj.WithdrawAuthority); err != nil {
		return fmt.Errorf("error while marshaling WithdrawAuthority:%w", err)
	}
	// Serialize `EnableMigrate`:
	if err = encoder.Encode(obj.EnableMigrate); err != nil {
		return fmt.Errorf("error while marshaling EnableMigrate:%w", err)
	}
	// Serialize `PoolMigrationFee`:
	if err = encoder.Encode(obj.PoolMigrationFee); err != nil {
		return fmt.Errorf("error while marshaling PoolMigrationFee:%w", err)
	}
	// Serialize `CreatorFeeBasisPoints`:
	if err = encoder.Encode(obj.CreatorFeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling CreatorFeeBasisPoints:%w", err)
	}
	// Serialize `FeeRecipients`:
	if err = encoder.Encode(obj.FeeRecipients); err != nil {
		return fmt.Errorf("error while marshaling FeeRecipients:%w", err)
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `SetCreatorAuthority`:
	if err = encoder.Encode(obj.SetCreatorAuthority); err != nil {
		return fmt.Errorf("error while marshaling SetCreatorAuthority:%w", err)
	}
	// Serialize `AdminSetCreatorAuthority`:
	if err = encoder.Encode(obj.AdminSetCreatorAuthority); err != nil {
		return fmt.Errorf("error while marshaling AdminSetCreatorAuthority:%w", err)
	}
	return nil
}

func (obj SetParamsEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding SetParamsEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *SetParamsEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_SetParamsEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_SetParamsEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `InitialVirtualTokenReserves`:
	if err = decoder.Decode(&obj.InitialVirtualTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling InitialVirtualTokenReserves:%w", err)
	}
	// Deserialize `InitialVirtualSolReserves`:
	if err = decoder.Decode(&obj.InitialVirtualSolReserves); err != nil {
		return fmt.Errorf("error while unmarshaling InitialVirtualSolReserves:%w", err)
	}
	// Deserialize `InitialRealTokenReserves`:
	if err = decoder.Decode(&obj.InitialRealTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling InitialRealTokenReserves:%w", err)
	}
	// Deserialize `FinalRealSolReserves`:
	if err = decoder.Decode(&obj.FinalRealSolReserves); err != nil {
		return fmt.Errorf("error while unmarshaling FinalRealSolReserves:%w", err)
	}
	// Deserialize `TokenTotalSupply`:
	if err = decoder.Decode(&obj.TokenTotalSupply); err != nil {
		return fmt.Errorf("error while unmarshaling TokenTotalSupply:%w", err)
	}
	// Deserialize `FeeBasisPoints`:
	if err = decoder.Decode(&obj.FeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling FeeBasisPoints:%w", err)
	}
	// Deserialize `WithdrawAuthority`:
	if err = decoder.Decode(&obj.WithdrawAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling WithdrawAuthority:%w", err)
	}
	// Deserialize `EnableMigrate`:
	if err = decoder.Decode(&obj.EnableMigrate); err != nil {
		return fmt.Errorf("error while unmarshaling EnableMigrate:%w", err)
	}
	// Deserialize `PoolMigrationFee`:
	if err = decoder.Decode(&obj.PoolMigrationFee); err != nil {
		return fmt.Errorf("error while unmarshaling PoolMigrationFee:%w", err)
	}
	// Deserialize `CreatorFeeBasisPoints`:
	if err = decoder.Decode(&obj.CreatorFeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorFeeBasisPoints:%w", err)
	}
	// Deserialize `FeeRecipients`:
	if err = decoder.Decode(&obj.FeeRecipients); err != nil {
		return fmt.Errorf("error while unmarshaling FeeRecipients:%w", err)
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `SetCreatorAuthority`:
	if err = decoder.Decode(&obj.SetCreatorAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling SetCreatorAuthority:%w", err)
	}
	// Deserialize `AdminSetCreatorAuthority`:
	if err = decoder.Decode(&obj.AdminSetCreatorAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling AdminSetCreatorAuthority:%w", err)
	}
	return nil
}

func (obj *SetParamsEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling SetParamsEvent: %w", err)
	}
	return nil
}

func UnmarshalSetParamsEvent(buf []byte) (*SetParamsEvent, error) {
	obj := new(SetParamsEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type SyncUserVolumeAccumulatorEvent struct {
	User                     solanago.PublicKey
	TotalClaimedTokensBefore uint64
	TotalClaimedTokensAfter  uint64
	Timestamp                int64
}

func (obj SyncUserVolumeAccumulatorEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_SyncUserVolumeAccumulatorEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `User`:
	if err = encoder.Encode(obj.User); err != nil {
		return fmt.Errorf("error while marshaling User:%w", err)
	}
	// Serialize `TotalClaimedTokensBefore`:
	if err = encoder.Encode(obj.TotalClaimedTokensBefore); err != nil {
		return fmt.Errorf("error while marshaling TotalClaimedTokensBefore:%w", err)
	}
	// Serialize `TotalClaimedTokensAfter`:
	if err = encoder.Encode(obj.TotalClaimedTokensAfter); err != nil {
		return fmt.Errorf("error while marshaling TotalClaimedTokensAfter:%w", err)
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	return nil
}

func (obj SyncUserVolumeAccumulatorEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding SyncUserVolumeAccumulatorEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *SyncUserVolumeAccumulatorEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_SyncUserVolumeAccumulatorEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_SyncUserVolumeAccumulatorEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `User`:
	if err = decoder.Decode(&obj.User); err != nil {
		return fmt.Errorf("error while unmarshaling User:%w", err)
	}
	// Deserialize `TotalClaimedTokensBefore`:
	if err = decoder.Decode(&obj.TotalClaimedTokensBefore); err != nil {
		return fmt.Errorf("error while unmarshaling TotalClaimedTokensBefore:%w", err)
	}
	// Deserialize `TotalClaimedTokensAfter`:
	if err = decoder.Decode(&obj.TotalClaimedTokensAfter); err != nil {
		return fmt.Errorf("error while unmarshaling TotalClaimedTokensAfter:%w", err)
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	return nil
}

func (obj *SyncUserVolumeAccumulatorEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling SyncUserVolumeAccumulatorEvent: %w", err)
	}
	return nil
}

func UnmarshalSyncUserVolumeAccumulatorEvent(buf []byte) (*SyncUserVolumeAccumulatorEvent, error) {
	obj := new(SyncUserVolumeAccumulatorEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type TradeEvent struct {
	Mint                  solanago.PublicKey
	SolAmount             uint64
	TokenAmount           uint64
	IsBuy                 bool
	User                  solanago.PublicKey
	Timestamp             int64
	VirtualSolReserves    uint64
	VirtualTokenReserves  uint64
	RealSolReserves       uint64
	RealTokenReserves     uint64
	FeeRecipient          solanago.PublicKey
	FeeBasisPoints        uint64
	Fee                   uint64
	Creator               solanago.PublicKey
	CreatorFeeBasisPoints uint64
	CreatorFee            uint64
	TrackVolume           bool
	TotalUnclaimedTokens  uint64
	TotalClaimedTokens    uint64
	CurrentSolVolume      uint64
	LastUpdateTimestamp   int64
}

func (obj TradeEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_TradeEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `SolAmount`:
	if err = encoder.Encode(obj.SolAmount); err != nil {
		return fmt.Errorf("error while marshaling SolAmount:%w", err)
	}
	// Serialize `TokenAmount`:
	if err = encoder.Encode(obj.TokenAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenAmount:%w", err)
	}
	// Serialize `IsBuy`:
	if err = encoder.Encode(obj.IsBuy); err != nil {
		return fmt.Errorf("error while marshaling IsBuy:%w", err)
	}
	// Serialize `User`:
	if err = encoder.Encode(obj.User); err != nil {
		return fmt.Errorf("error while marshaling User:%w", err)
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `VirtualSolReserves`:
	if err = encoder.Encode(obj.VirtualSolReserves); err != nil {
		return fmt.Errorf("error while marshaling VirtualSolReserves:%w", err)
	}
	// Serialize `VirtualTokenReserves`:
	if err = encoder.Encode(obj.VirtualTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling VirtualTokenReserves:%w", err)
	}
	// Serialize `RealSolReserves`:
	if err = encoder.Encode(obj.RealSolReserves); err != nil {
		return fmt.Errorf("error while marshaling RealSolReserves:%w", err)
	}
	// Serialize `RealTokenReserves`:
	if err = encoder.Encode(obj.RealTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling RealTokenReserves:%w", err)
	}
	// Serialize `FeeRecipient`:
	if err = encoder.Encode(obj.FeeRecipient); err != nil {
		return fmt.Errorf("error while marshaling FeeRecipient:%w", err)
	}
	// Serialize `FeeBasisPoints`:
	if err = encoder.Encode(obj.FeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling FeeBasisPoints:%w", err)
	}
	// Serialize `Fee`:
	if err = encoder.Encode(obj.Fee); err != nil {
		return fmt.Errorf("error while marshaling Fee:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
	}
	// Serialize `CreatorFeeBasisPoints`:
	if err = encoder.Encode(obj.CreatorFeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling CreatorFeeBasisPoints:%w", err)
	}
	// Serialize `CreatorFee`:
	if err = encoder.Encode(obj.CreatorFee); err != nil {
		return fmt.Errorf("error while marshaling CreatorFee:%w", err)
	}
	// Serialize `TrackVolume`:
	if err = encoder.Encode(obj.TrackVolume); err != nil {
		return fmt.Errorf("error while marshaling TrackVolume:%w", err)
	}
	// Serialize `TotalUnclaimedTokens`:
	if err = encoder.Encode(obj.TotalUnclaimedTokens); err != nil {
		return fmt.Errorf("error while marshaling TotalUnclaimedTokens:%w", err)
	}
	// Serialize `TotalClaimedTokens`:
	if err = encoder.Encode(obj.TotalClaimedTokens); err != nil {
		return fmt.Errorf("error while marshaling TotalClaimedTokens:%w", err)
	}
	// Serialize `CurrentSolVolume`:
	if err = encoder.Encode(obj.CurrentSolVolume); err != nil {
		return fmt.Errorf("error while marshaling CurrentSolVolume:%w", err)
	}
	// Serialize `LastUpdateTimestamp`:
	if err = encoder.Encode(obj.LastUpdateTimestamp); err != nil {
		return fmt.Errorf("error while marshaling LastUpdateTimestamp:%w", err)
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
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `SolAmount`:
	if err = decoder.Decode(&obj.SolAmount); err != nil {
		return fmt.Errorf("error while unmarshaling SolAmount:%w", err)
	}
	// Deserialize `TokenAmount`:
	if err = decoder.Decode(&obj.TokenAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAmount:%w", err)
	}
	// Deserialize `IsBuy`:
	if err = decoder.Decode(&obj.IsBuy); err != nil {
		return fmt.Errorf("error while unmarshaling IsBuy:%w", err)
	}
	// Deserialize `User`:
	if err = decoder.Decode(&obj.User); err != nil {
		return fmt.Errorf("error while unmarshaling User:%w", err)
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `VirtualSolReserves`:
	if err = decoder.Decode(&obj.VirtualSolReserves); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualSolReserves:%w", err)
	}
	// Deserialize `VirtualTokenReserves`:
	if err = decoder.Decode(&obj.VirtualTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualTokenReserves:%w", err)
	}
	// Deserialize `RealSolReserves`:
	if err = decoder.Decode(&obj.RealSolReserves); err != nil {
		return fmt.Errorf("error while unmarshaling RealSolReserves:%w", err)
	}
	// Deserialize `RealTokenReserves`:
	if err = decoder.Decode(&obj.RealTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling RealTokenReserves:%w", err)
	}
	// Deserialize `FeeRecipient`:
	if err = decoder.Decode(&obj.FeeRecipient); err != nil {
		return fmt.Errorf("error while unmarshaling FeeRecipient:%w", err)
	}
	// Deserialize `FeeBasisPoints`:
	if err = decoder.Decode(&obj.FeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling FeeBasisPoints:%w", err)
	}
	// Deserialize `Fee`:
	if err = decoder.Decode(&obj.Fee); err != nil {
		return fmt.Errorf("error while unmarshaling Fee:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
	}
	// Deserialize `CreatorFeeBasisPoints`:
	if err = decoder.Decode(&obj.CreatorFeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorFeeBasisPoints:%w", err)
	}
	// Deserialize `CreatorFee`:
	if err = decoder.Decode(&obj.CreatorFee); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorFee:%w", err)
	}
	// Deserialize `TrackVolume`:
	if err = decoder.Decode(&obj.TrackVolume); err != nil {
		return fmt.Errorf("error while unmarshaling TrackVolume:%w", err)
	}
	// Deserialize `TotalUnclaimedTokens`:
	if err = decoder.Decode(&obj.TotalUnclaimedTokens); err != nil {
		return fmt.Errorf("error while unmarshaling TotalUnclaimedTokens:%w", err)
	}
	// Deserialize `TotalClaimedTokens`:
	if err = decoder.Decode(&obj.TotalClaimedTokens); err != nil {
		return fmt.Errorf("error while unmarshaling TotalClaimedTokens:%w", err)
	}
	// Deserialize `CurrentSolVolume`:
	if err = decoder.Decode(&obj.CurrentSolVolume); err != nil {
		return fmt.Errorf("error while unmarshaling CurrentSolVolume:%w", err)
	}
	// Deserialize `LastUpdateTimestamp`:
	if err = decoder.Decode(&obj.LastUpdateTimestamp); err != nil {
		return fmt.Errorf("error while unmarshaling LastUpdateTimestamp:%w", err)
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

type UpdateGlobalAuthorityEvent struct {
	Global       solanago.PublicKey
	Authority    solanago.PublicKey
	NewAuthority solanago.PublicKey
	Timestamp    int64
}

func (obj UpdateGlobalAuthorityEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_UpdateGlobalAuthorityEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Global`:
	if err = encoder.Encode(obj.Global); err != nil {
		return fmt.Errorf("error while marshaling Global:%w", err)
	}
	// Serialize `Authority`:
	if err = encoder.Encode(obj.Authority); err != nil {
		return fmt.Errorf("error while marshaling Authority:%w", err)
	}
	// Serialize `NewAuthority`:
	if err = encoder.Encode(obj.NewAuthority); err != nil {
		return fmt.Errorf("error while marshaling NewAuthority:%w", err)
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	return nil
}

func (obj UpdateGlobalAuthorityEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding UpdateGlobalAuthorityEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *UpdateGlobalAuthorityEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_UpdateGlobalAuthorityEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_UpdateGlobalAuthorityEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Global`:
	if err = decoder.Decode(&obj.Global); err != nil {
		return fmt.Errorf("error while unmarshaling Global:%w", err)
	}
	// Deserialize `Authority`:
	if err = decoder.Decode(&obj.Authority); err != nil {
		return fmt.Errorf("error while unmarshaling Authority:%w", err)
	}
	// Deserialize `NewAuthority`:
	if err = decoder.Decode(&obj.NewAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling NewAuthority:%w", err)
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	return nil
}

func (obj *UpdateGlobalAuthorityEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling UpdateGlobalAuthorityEvent: %w", err)
	}
	return nil
}

func UnmarshalUpdateGlobalAuthorityEvent(buf []byte) (*UpdateGlobalAuthorityEvent, error) {
	obj := new(UpdateGlobalAuthorityEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
