package pump_amm

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type AdminSetCoinCreatorEvent struct {
	Timestamp                    int64
	AdminSetCoinCreatorAuthority solanago.PublicKey
	BaseMint                     solanago.PublicKey
	Pool                         solanago.PublicKey
	OldCoinCreator               solanago.PublicKey
	NewCoinCreator               solanago.PublicKey
}

func (obj AdminSetCoinCreatorEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_AdminSetCoinCreatorEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `AdminSetCoinCreatorAuthority`:
	if err = encoder.Encode(obj.AdminSetCoinCreatorAuthority); err != nil {
		return fmt.Errorf("error while marshaling AdminSetCoinCreatorAuthority:%w", err)
	}
	// Serialize `BaseMint`:
	if err = encoder.Encode(obj.BaseMint); err != nil {
		return fmt.Errorf("error while marshaling BaseMint:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `OldCoinCreator`:
	if err = encoder.Encode(obj.OldCoinCreator); err != nil {
		return fmt.Errorf("error while marshaling OldCoinCreator:%w", err)
	}
	// Serialize `NewCoinCreator`:
	if err = encoder.Encode(obj.NewCoinCreator); err != nil {
		return fmt.Errorf("error while marshaling NewCoinCreator:%w", err)
	}
	return nil
}

func (obj AdminSetCoinCreatorEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding AdminSetCoinCreatorEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *AdminSetCoinCreatorEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_AdminSetCoinCreatorEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_AdminSetCoinCreatorEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `AdminSetCoinCreatorAuthority`:
	if err = decoder.Decode(&obj.AdminSetCoinCreatorAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling AdminSetCoinCreatorAuthority:%w", err)
	}
	// Deserialize `BaseMint`:
	if err = decoder.Decode(&obj.BaseMint); err != nil {
		return fmt.Errorf("error while unmarshaling BaseMint:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `OldCoinCreator`:
	if err = decoder.Decode(&obj.OldCoinCreator); err != nil {
		return fmt.Errorf("error while unmarshaling OldCoinCreator:%w", err)
	}
	// Deserialize `NewCoinCreator`:
	if err = decoder.Decode(&obj.NewCoinCreator); err != nil {
		return fmt.Errorf("error while unmarshaling NewCoinCreator:%w", err)
	}
	return nil
}

func (obj *AdminSetCoinCreatorEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling AdminSetCoinCreatorEvent: %w", err)
	}
	return nil
}

func UnmarshalAdminSetCoinCreatorEvent(buf []byte) (*AdminSetCoinCreatorEvent, error) {
	obj := new(AdminSetCoinCreatorEvent)
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

type BuyEvent struct {
	Timestamp                        int64
	BaseAmountOut                    uint64
	MaxQuoteAmountIn                 uint64
	UserBaseTokenReserves            uint64
	UserQuoteTokenReserves           uint64
	PoolBaseTokenReserves            uint64
	PoolQuoteTokenReserves           uint64
	QuoteAmountIn                    uint64
	LpFeeBasisPoints                 uint64
	LpFee                            uint64
	ProtocolFeeBasisPoints           uint64
	ProtocolFee                      uint64
	QuoteAmountInWithLpFee           uint64
	UserQuoteAmountIn                uint64
	Pool                             solanago.PublicKey
	User                             solanago.PublicKey
	UserBaseTokenAccount             solanago.PublicKey
	UserQuoteTokenAccount            solanago.PublicKey
	ProtocolFeeRecipient             solanago.PublicKey
	ProtocolFeeRecipientTokenAccount solanago.PublicKey
	CoinCreator                      solanago.PublicKey
	CoinCreatorFeeBasisPoints        uint64
	CoinCreatorFee                   uint64
	TrackVolume                      bool
	TotalUnclaimedTokens             uint64
	TotalClaimedTokens               uint64
	CurrentSolVolume                 uint64
	LastUpdateTimestamp              int64
}

func (obj BuyEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_BuyEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `BaseAmountOut`:
	if err = encoder.Encode(obj.BaseAmountOut); err != nil {
		return fmt.Errorf("error while marshaling BaseAmountOut:%w", err)
	}
	// Serialize `MaxQuoteAmountIn`:
	if err = encoder.Encode(obj.MaxQuoteAmountIn); err != nil {
		return fmt.Errorf("error while marshaling MaxQuoteAmountIn:%w", err)
	}
	// Serialize `UserBaseTokenReserves`:
	if err = encoder.Encode(obj.UserBaseTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling UserBaseTokenReserves:%w", err)
	}
	// Serialize `UserQuoteTokenReserves`:
	if err = encoder.Encode(obj.UserQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling UserQuoteTokenReserves:%w", err)
	}
	// Serialize `PoolBaseTokenReserves`:
	if err = encoder.Encode(obj.PoolBaseTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling PoolBaseTokenReserves:%w", err)
	}
	// Serialize `PoolQuoteTokenReserves`:
	if err = encoder.Encode(obj.PoolQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling PoolQuoteTokenReserves:%w", err)
	}
	// Serialize `QuoteAmountIn`:
	if err = encoder.Encode(obj.QuoteAmountIn); err != nil {
		return fmt.Errorf("error while marshaling QuoteAmountIn:%w", err)
	}
	// Serialize `LpFeeBasisPoints`:
	if err = encoder.Encode(obj.LpFeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling LpFeeBasisPoints:%w", err)
	}
	// Serialize `LpFee`:
	if err = encoder.Encode(obj.LpFee); err != nil {
		return fmt.Errorf("error while marshaling LpFee:%w", err)
	}
	// Serialize `ProtocolFeeBasisPoints`:
	if err = encoder.Encode(obj.ProtocolFeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeBasisPoints:%w", err)
	}
	// Serialize `ProtocolFee`:
	if err = encoder.Encode(obj.ProtocolFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFee:%w", err)
	}
	// Serialize `QuoteAmountInWithLpFee`:
	if err = encoder.Encode(obj.QuoteAmountInWithLpFee); err != nil {
		return fmt.Errorf("error while marshaling QuoteAmountInWithLpFee:%w", err)
	}
	// Serialize `UserQuoteAmountIn`:
	if err = encoder.Encode(obj.UserQuoteAmountIn); err != nil {
		return fmt.Errorf("error while marshaling UserQuoteAmountIn:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `User`:
	if err = encoder.Encode(obj.User); err != nil {
		return fmt.Errorf("error while marshaling User:%w", err)
	}
	// Serialize `UserBaseTokenAccount`:
	if err = encoder.Encode(obj.UserBaseTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling UserBaseTokenAccount:%w", err)
	}
	// Serialize `UserQuoteTokenAccount`:
	if err = encoder.Encode(obj.UserQuoteTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling UserQuoteTokenAccount:%w", err)
	}
	// Serialize `ProtocolFeeRecipient`:
	if err = encoder.Encode(obj.ProtocolFeeRecipient); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeRecipient:%w", err)
	}
	// Serialize `ProtocolFeeRecipientTokenAccount`:
	if err = encoder.Encode(obj.ProtocolFeeRecipientTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeRecipientTokenAccount:%w", err)
	}
	// Serialize `CoinCreator`:
	if err = encoder.Encode(obj.CoinCreator); err != nil {
		return fmt.Errorf("error while marshaling CoinCreator:%w", err)
	}
	// Serialize `CoinCreatorFeeBasisPoints`:
	if err = encoder.Encode(obj.CoinCreatorFeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling CoinCreatorFeeBasisPoints:%w", err)
	}
	// Serialize `CoinCreatorFee`:
	if err = encoder.Encode(obj.CoinCreatorFee); err != nil {
		return fmt.Errorf("error while marshaling CoinCreatorFee:%w", err)
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

func (obj BuyEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding BuyEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *BuyEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_BuyEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_BuyEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `BaseAmountOut`:
	if err = decoder.Decode(&obj.BaseAmountOut); err != nil {
		return fmt.Errorf("error while unmarshaling BaseAmountOut:%w", err)
	}
	// Deserialize `MaxQuoteAmountIn`:
	if err = decoder.Decode(&obj.MaxQuoteAmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling MaxQuoteAmountIn:%w", err)
	}
	// Deserialize `UserBaseTokenReserves`:
	if err = decoder.Decode(&obj.UserBaseTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling UserBaseTokenReserves:%w", err)
	}
	// Deserialize `UserQuoteTokenReserves`:
	if err = decoder.Decode(&obj.UserQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling UserQuoteTokenReserves:%w", err)
	}
	// Deserialize `PoolBaseTokenReserves`:
	if err = decoder.Decode(&obj.PoolBaseTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling PoolBaseTokenReserves:%w", err)
	}
	// Deserialize `PoolQuoteTokenReserves`:
	if err = decoder.Decode(&obj.PoolQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling PoolQuoteTokenReserves:%w", err)
	}
	// Deserialize `QuoteAmountIn`:
	if err = decoder.Decode(&obj.QuoteAmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteAmountIn:%w", err)
	}
	// Deserialize `LpFeeBasisPoints`:
	if err = decoder.Decode(&obj.LpFeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling LpFeeBasisPoints:%w", err)
	}
	// Deserialize `LpFee`:
	if err = decoder.Decode(&obj.LpFee); err != nil {
		return fmt.Errorf("error while unmarshaling LpFee:%w", err)
	}
	// Deserialize `ProtocolFeeBasisPoints`:
	if err = decoder.Decode(&obj.ProtocolFeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeBasisPoints:%w", err)
	}
	// Deserialize `ProtocolFee`:
	if err = decoder.Decode(&obj.ProtocolFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFee:%w", err)
	}
	// Deserialize `QuoteAmountInWithLpFee`:
	if err = decoder.Decode(&obj.QuoteAmountInWithLpFee); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteAmountInWithLpFee:%w", err)
	}
	// Deserialize `UserQuoteAmountIn`:
	if err = decoder.Decode(&obj.UserQuoteAmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling UserQuoteAmountIn:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `User`:
	if err = decoder.Decode(&obj.User); err != nil {
		return fmt.Errorf("error while unmarshaling User:%w", err)
	}
	// Deserialize `UserBaseTokenAccount`:
	if err = decoder.Decode(&obj.UserBaseTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling UserBaseTokenAccount:%w", err)
	}
	// Deserialize `UserQuoteTokenAccount`:
	if err = decoder.Decode(&obj.UserQuoteTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling UserQuoteTokenAccount:%w", err)
	}
	// Deserialize `ProtocolFeeRecipient`:
	if err = decoder.Decode(&obj.ProtocolFeeRecipient); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeRecipient:%w", err)
	}
	// Deserialize `ProtocolFeeRecipientTokenAccount`:
	if err = decoder.Decode(&obj.ProtocolFeeRecipientTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeRecipientTokenAccount:%w", err)
	}
	// Deserialize `CoinCreator`:
	if err = decoder.Decode(&obj.CoinCreator); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreator:%w", err)
	}
	// Deserialize `CoinCreatorFeeBasisPoints`:
	if err = decoder.Decode(&obj.CoinCreatorFeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreatorFeeBasisPoints:%w", err)
	}
	// Deserialize `CoinCreatorFee`:
	if err = decoder.Decode(&obj.CoinCreatorFee); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreatorFee:%w", err)
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

func (obj *BuyEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling BuyEvent: %w", err)
	}
	return nil
}

func UnmarshalBuyEvent(buf []byte) (*BuyEvent, error) {
	obj := new(BuyEvent)
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

type CollectCoinCreatorFeeEvent struct {
	Timestamp               int64
	CoinCreator             solanago.PublicKey
	CoinCreatorFee          uint64
	CoinCreatorVaultAta     solanago.PublicKey
	CoinCreatorTokenAccount solanago.PublicKey
}

func (obj CollectCoinCreatorFeeEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CollectCoinCreatorFeeEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `CoinCreator`:
	if err = encoder.Encode(obj.CoinCreator); err != nil {
		return fmt.Errorf("error while marshaling CoinCreator:%w", err)
	}
	// Serialize `CoinCreatorFee`:
	if err = encoder.Encode(obj.CoinCreatorFee); err != nil {
		return fmt.Errorf("error while marshaling CoinCreatorFee:%w", err)
	}
	// Serialize `CoinCreatorVaultAta`:
	if err = encoder.Encode(obj.CoinCreatorVaultAta); err != nil {
		return fmt.Errorf("error while marshaling CoinCreatorVaultAta:%w", err)
	}
	// Serialize `CoinCreatorTokenAccount`:
	if err = encoder.Encode(obj.CoinCreatorTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling CoinCreatorTokenAccount:%w", err)
	}
	return nil
}

func (obj CollectCoinCreatorFeeEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding CollectCoinCreatorFeeEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *CollectCoinCreatorFeeEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CollectCoinCreatorFeeEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CollectCoinCreatorFeeEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `CoinCreator`:
	if err = decoder.Decode(&obj.CoinCreator); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreator:%w", err)
	}
	// Deserialize `CoinCreatorFee`:
	if err = decoder.Decode(&obj.CoinCreatorFee); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreatorFee:%w", err)
	}
	// Deserialize `CoinCreatorVaultAta`:
	if err = decoder.Decode(&obj.CoinCreatorVaultAta); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreatorVaultAta:%w", err)
	}
	// Deserialize `CoinCreatorTokenAccount`:
	if err = decoder.Decode(&obj.CoinCreatorTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreatorTokenAccount:%w", err)
	}
	return nil
}

func (obj *CollectCoinCreatorFeeEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling CollectCoinCreatorFeeEvent: %w", err)
	}
	return nil
}

func UnmarshalCollectCoinCreatorFeeEvent(buf []byte) (*CollectCoinCreatorFeeEvent, error) {
	obj := new(CollectCoinCreatorFeeEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type CreateConfigEvent struct {
	Timestamp                    int64
	Admin                        solanago.PublicKey
	LpFeeBasisPoints             uint64
	ProtocolFeeBasisPoints       uint64
	ProtocolFeeRecipients        [8]solanago.PublicKey
	CoinCreatorFeeBasisPoints    uint64
	AdminSetCoinCreatorAuthority solanago.PublicKey
}

func (obj CreateConfigEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CreateConfigEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `Admin`:
	if err = encoder.Encode(obj.Admin); err != nil {
		return fmt.Errorf("error while marshaling Admin:%w", err)
	}
	// Serialize `LpFeeBasisPoints`:
	if err = encoder.Encode(obj.LpFeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling LpFeeBasisPoints:%w", err)
	}
	// Serialize `ProtocolFeeBasisPoints`:
	if err = encoder.Encode(obj.ProtocolFeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeBasisPoints:%w", err)
	}
	// Serialize `ProtocolFeeRecipients`:
	if err = encoder.Encode(obj.ProtocolFeeRecipients); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeRecipients:%w", err)
	}
	// Serialize `CoinCreatorFeeBasisPoints`:
	if err = encoder.Encode(obj.CoinCreatorFeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling CoinCreatorFeeBasisPoints:%w", err)
	}
	// Serialize `AdminSetCoinCreatorAuthority`:
	if err = encoder.Encode(obj.AdminSetCoinCreatorAuthority); err != nil {
		return fmt.Errorf("error while marshaling AdminSetCoinCreatorAuthority:%w", err)
	}
	return nil
}

func (obj CreateConfigEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding CreateConfigEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *CreateConfigEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CreateConfigEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CreateConfigEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `Admin`:
	if err = decoder.Decode(&obj.Admin); err != nil {
		return fmt.Errorf("error while unmarshaling Admin:%w", err)
	}
	// Deserialize `LpFeeBasisPoints`:
	if err = decoder.Decode(&obj.LpFeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling LpFeeBasisPoints:%w", err)
	}
	// Deserialize `ProtocolFeeBasisPoints`:
	if err = decoder.Decode(&obj.ProtocolFeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeBasisPoints:%w", err)
	}
	// Deserialize `ProtocolFeeRecipients`:
	if err = decoder.Decode(&obj.ProtocolFeeRecipients); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeRecipients:%w", err)
	}
	// Deserialize `CoinCreatorFeeBasisPoints`:
	if err = decoder.Decode(&obj.CoinCreatorFeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreatorFeeBasisPoints:%w", err)
	}
	// Deserialize `AdminSetCoinCreatorAuthority`:
	if err = decoder.Decode(&obj.AdminSetCoinCreatorAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling AdminSetCoinCreatorAuthority:%w", err)
	}
	return nil
}

func (obj *CreateConfigEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling CreateConfigEvent: %w", err)
	}
	return nil
}

func UnmarshalCreateConfigEvent(buf []byte) (*CreateConfigEvent, error) {
	obj := new(CreateConfigEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type CreatePoolEvent struct {
	Timestamp             int64
	Index                 uint16
	Creator               solanago.PublicKey
	BaseMint              solanago.PublicKey
	QuoteMint             solanago.PublicKey
	BaseMintDecimals      uint8
	QuoteMintDecimals     uint8
	BaseAmountIn          uint64
	QuoteAmountIn         uint64
	PoolBaseAmount        uint64
	PoolQuoteAmount       uint64
	MinimumLiquidity      uint64
	InitialLiquidity      uint64
	LpTokenAmountOut      uint64
	PoolBump              uint8
	Pool                  solanago.PublicKey
	LpMint                solanago.PublicKey
	UserBaseTokenAccount  solanago.PublicKey
	UserQuoteTokenAccount solanago.PublicKey
	CoinCreator           solanago.PublicKey
}

func (obj CreatePoolEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CreatePoolEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `Index`:
	if err = encoder.Encode(obj.Index); err != nil {
		return fmt.Errorf("error while marshaling Index:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
	}
	// Serialize `BaseMint`:
	if err = encoder.Encode(obj.BaseMint); err != nil {
		return fmt.Errorf("error while marshaling BaseMint:%w", err)
	}
	// Serialize `QuoteMint`:
	if err = encoder.Encode(obj.QuoteMint); err != nil {
		return fmt.Errorf("error while marshaling QuoteMint:%w", err)
	}
	// Serialize `BaseMintDecimals`:
	if err = encoder.Encode(obj.BaseMintDecimals); err != nil {
		return fmt.Errorf("error while marshaling BaseMintDecimals:%w", err)
	}
	// Serialize `QuoteMintDecimals`:
	if err = encoder.Encode(obj.QuoteMintDecimals); err != nil {
		return fmt.Errorf("error while marshaling QuoteMintDecimals:%w", err)
	}
	// Serialize `BaseAmountIn`:
	if err = encoder.Encode(obj.BaseAmountIn); err != nil {
		return fmt.Errorf("error while marshaling BaseAmountIn:%w", err)
	}
	// Serialize `QuoteAmountIn`:
	if err = encoder.Encode(obj.QuoteAmountIn); err != nil {
		return fmt.Errorf("error while marshaling QuoteAmountIn:%w", err)
	}
	// Serialize `PoolBaseAmount`:
	if err = encoder.Encode(obj.PoolBaseAmount); err != nil {
		return fmt.Errorf("error while marshaling PoolBaseAmount:%w", err)
	}
	// Serialize `PoolQuoteAmount`:
	if err = encoder.Encode(obj.PoolQuoteAmount); err != nil {
		return fmt.Errorf("error while marshaling PoolQuoteAmount:%w", err)
	}
	// Serialize `MinimumLiquidity`:
	if err = encoder.Encode(obj.MinimumLiquidity); err != nil {
		return fmt.Errorf("error while marshaling MinimumLiquidity:%w", err)
	}
	// Serialize `InitialLiquidity`:
	if err = encoder.Encode(obj.InitialLiquidity); err != nil {
		return fmt.Errorf("error while marshaling InitialLiquidity:%w", err)
	}
	// Serialize `LpTokenAmountOut`:
	if err = encoder.Encode(obj.LpTokenAmountOut); err != nil {
		return fmt.Errorf("error while marshaling LpTokenAmountOut:%w", err)
	}
	// Serialize `PoolBump`:
	if err = encoder.Encode(obj.PoolBump); err != nil {
		return fmt.Errorf("error while marshaling PoolBump:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `LpMint`:
	if err = encoder.Encode(obj.LpMint); err != nil {
		return fmt.Errorf("error while marshaling LpMint:%w", err)
	}
	// Serialize `UserBaseTokenAccount`:
	if err = encoder.Encode(obj.UserBaseTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling UserBaseTokenAccount:%w", err)
	}
	// Serialize `UserQuoteTokenAccount`:
	if err = encoder.Encode(obj.UserQuoteTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling UserQuoteTokenAccount:%w", err)
	}
	// Serialize `CoinCreator`:
	if err = encoder.Encode(obj.CoinCreator); err != nil {
		return fmt.Errorf("error while marshaling CoinCreator:%w", err)
	}
	return nil
}

func (obj CreatePoolEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding CreatePoolEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *CreatePoolEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CreatePoolEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CreatePoolEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `Index`:
	if err = decoder.Decode(&obj.Index); err != nil {
		return fmt.Errorf("error while unmarshaling Index:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
	}
	// Deserialize `BaseMint`:
	if err = decoder.Decode(&obj.BaseMint); err != nil {
		return fmt.Errorf("error while unmarshaling BaseMint:%w", err)
	}
	// Deserialize `QuoteMint`:
	if err = decoder.Decode(&obj.QuoteMint); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteMint:%w", err)
	}
	// Deserialize `BaseMintDecimals`:
	if err = decoder.Decode(&obj.BaseMintDecimals); err != nil {
		return fmt.Errorf("error while unmarshaling BaseMintDecimals:%w", err)
	}
	// Deserialize `QuoteMintDecimals`:
	if err = decoder.Decode(&obj.QuoteMintDecimals); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteMintDecimals:%w", err)
	}
	// Deserialize `BaseAmountIn`:
	if err = decoder.Decode(&obj.BaseAmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling BaseAmountIn:%w", err)
	}
	// Deserialize `QuoteAmountIn`:
	if err = decoder.Decode(&obj.QuoteAmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteAmountIn:%w", err)
	}
	// Deserialize `PoolBaseAmount`:
	if err = decoder.Decode(&obj.PoolBaseAmount); err != nil {
		return fmt.Errorf("error while unmarshaling PoolBaseAmount:%w", err)
	}
	// Deserialize `PoolQuoteAmount`:
	if err = decoder.Decode(&obj.PoolQuoteAmount); err != nil {
		return fmt.Errorf("error while unmarshaling PoolQuoteAmount:%w", err)
	}
	// Deserialize `MinimumLiquidity`:
	if err = decoder.Decode(&obj.MinimumLiquidity); err != nil {
		return fmt.Errorf("error while unmarshaling MinimumLiquidity:%w", err)
	}
	// Deserialize `InitialLiquidity`:
	if err = decoder.Decode(&obj.InitialLiquidity); err != nil {
		return fmt.Errorf("error while unmarshaling InitialLiquidity:%w", err)
	}
	// Deserialize `LpTokenAmountOut`:
	if err = decoder.Decode(&obj.LpTokenAmountOut); err != nil {
		return fmt.Errorf("error while unmarshaling LpTokenAmountOut:%w", err)
	}
	// Deserialize `PoolBump`:
	if err = decoder.Decode(&obj.PoolBump); err != nil {
		return fmt.Errorf("error while unmarshaling PoolBump:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `LpMint`:
	if err = decoder.Decode(&obj.LpMint); err != nil {
		return fmt.Errorf("error while unmarshaling LpMint:%w", err)
	}
	// Deserialize `UserBaseTokenAccount`:
	if err = decoder.Decode(&obj.UserBaseTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling UserBaseTokenAccount:%w", err)
	}
	// Deserialize `UserQuoteTokenAccount`:
	if err = decoder.Decode(&obj.UserQuoteTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling UserQuoteTokenAccount:%w", err)
	}
	// Deserialize `CoinCreator`:
	if err = decoder.Decode(&obj.CoinCreator); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreator:%w", err)
	}
	return nil
}

func (obj *CreatePoolEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling CreatePoolEvent: %w", err)
	}
	return nil
}

func UnmarshalCreatePoolEvent(buf []byte) (*CreatePoolEvent, error) {
	obj := new(CreatePoolEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type DepositEvent struct {
	Timestamp              int64
	LpTokenAmountOut       uint64
	MaxBaseAmountIn        uint64
	MaxQuoteAmountIn       uint64
	UserBaseTokenReserves  uint64
	UserQuoteTokenReserves uint64
	PoolBaseTokenReserves  uint64
	PoolQuoteTokenReserves uint64
	BaseAmountIn           uint64
	QuoteAmountIn          uint64
	LpMintSupply           uint64
	Pool                   solanago.PublicKey
	User                   solanago.PublicKey
	UserBaseTokenAccount   solanago.PublicKey
	UserQuoteTokenAccount  solanago.PublicKey
	UserPoolTokenAccount   solanago.PublicKey
}

func (obj DepositEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_DepositEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `LpTokenAmountOut`:
	if err = encoder.Encode(obj.LpTokenAmountOut); err != nil {
		return fmt.Errorf("error while marshaling LpTokenAmountOut:%w", err)
	}
	// Serialize `MaxBaseAmountIn`:
	if err = encoder.Encode(obj.MaxBaseAmountIn); err != nil {
		return fmt.Errorf("error while marshaling MaxBaseAmountIn:%w", err)
	}
	// Serialize `MaxQuoteAmountIn`:
	if err = encoder.Encode(obj.MaxQuoteAmountIn); err != nil {
		return fmt.Errorf("error while marshaling MaxQuoteAmountIn:%w", err)
	}
	// Serialize `UserBaseTokenReserves`:
	if err = encoder.Encode(obj.UserBaseTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling UserBaseTokenReserves:%w", err)
	}
	// Serialize `UserQuoteTokenReserves`:
	if err = encoder.Encode(obj.UserQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling UserQuoteTokenReserves:%w", err)
	}
	// Serialize `PoolBaseTokenReserves`:
	if err = encoder.Encode(obj.PoolBaseTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling PoolBaseTokenReserves:%w", err)
	}
	// Serialize `PoolQuoteTokenReserves`:
	if err = encoder.Encode(obj.PoolQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling PoolQuoteTokenReserves:%w", err)
	}
	// Serialize `BaseAmountIn`:
	if err = encoder.Encode(obj.BaseAmountIn); err != nil {
		return fmt.Errorf("error while marshaling BaseAmountIn:%w", err)
	}
	// Serialize `QuoteAmountIn`:
	if err = encoder.Encode(obj.QuoteAmountIn); err != nil {
		return fmt.Errorf("error while marshaling QuoteAmountIn:%w", err)
	}
	// Serialize `LpMintSupply`:
	if err = encoder.Encode(obj.LpMintSupply); err != nil {
		return fmt.Errorf("error while marshaling LpMintSupply:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `User`:
	if err = encoder.Encode(obj.User); err != nil {
		return fmt.Errorf("error while marshaling User:%w", err)
	}
	// Serialize `UserBaseTokenAccount`:
	if err = encoder.Encode(obj.UserBaseTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling UserBaseTokenAccount:%w", err)
	}
	// Serialize `UserQuoteTokenAccount`:
	if err = encoder.Encode(obj.UserQuoteTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling UserQuoteTokenAccount:%w", err)
	}
	// Serialize `UserPoolTokenAccount`:
	if err = encoder.Encode(obj.UserPoolTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling UserPoolTokenAccount:%w", err)
	}
	return nil
}

func (obj DepositEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding DepositEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *DepositEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_DepositEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_DepositEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `LpTokenAmountOut`:
	if err = decoder.Decode(&obj.LpTokenAmountOut); err != nil {
		return fmt.Errorf("error while unmarshaling LpTokenAmountOut:%w", err)
	}
	// Deserialize `MaxBaseAmountIn`:
	if err = decoder.Decode(&obj.MaxBaseAmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling MaxBaseAmountIn:%w", err)
	}
	// Deserialize `MaxQuoteAmountIn`:
	if err = decoder.Decode(&obj.MaxQuoteAmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling MaxQuoteAmountIn:%w", err)
	}
	// Deserialize `UserBaseTokenReserves`:
	if err = decoder.Decode(&obj.UserBaseTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling UserBaseTokenReserves:%w", err)
	}
	// Deserialize `UserQuoteTokenReserves`:
	if err = decoder.Decode(&obj.UserQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling UserQuoteTokenReserves:%w", err)
	}
	// Deserialize `PoolBaseTokenReserves`:
	if err = decoder.Decode(&obj.PoolBaseTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling PoolBaseTokenReserves:%w", err)
	}
	// Deserialize `PoolQuoteTokenReserves`:
	if err = decoder.Decode(&obj.PoolQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling PoolQuoteTokenReserves:%w", err)
	}
	// Deserialize `BaseAmountIn`:
	if err = decoder.Decode(&obj.BaseAmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling BaseAmountIn:%w", err)
	}
	// Deserialize `QuoteAmountIn`:
	if err = decoder.Decode(&obj.QuoteAmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteAmountIn:%w", err)
	}
	// Deserialize `LpMintSupply`:
	if err = decoder.Decode(&obj.LpMintSupply); err != nil {
		return fmt.Errorf("error while unmarshaling LpMintSupply:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `User`:
	if err = decoder.Decode(&obj.User); err != nil {
		return fmt.Errorf("error while unmarshaling User:%w", err)
	}
	// Deserialize `UserBaseTokenAccount`:
	if err = decoder.Decode(&obj.UserBaseTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling UserBaseTokenAccount:%w", err)
	}
	// Deserialize `UserQuoteTokenAccount`:
	if err = decoder.Decode(&obj.UserQuoteTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling UserQuoteTokenAccount:%w", err)
	}
	// Deserialize `UserPoolTokenAccount`:
	if err = decoder.Decode(&obj.UserPoolTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling UserPoolTokenAccount:%w", err)
	}
	return nil
}

func (obj *DepositEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling DepositEvent: %w", err)
	}
	return nil
}

func UnmarshalDepositEvent(buf []byte) (*DepositEvent, error) {
	obj := new(DepositEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type DisableEvent struct {
	Timestamp         int64
	Admin             solanago.PublicKey
	DisableCreatePool bool
	DisableDeposit    bool
	DisableWithdraw   bool
	DisableBuy        bool
	DisableSell       bool
}

func (obj DisableEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_DisableEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `Admin`:
	if err = encoder.Encode(obj.Admin); err != nil {
		return fmt.Errorf("error while marshaling Admin:%w", err)
	}
	// Serialize `DisableCreatePool`:
	if err = encoder.Encode(obj.DisableCreatePool); err != nil {
		return fmt.Errorf("error while marshaling DisableCreatePool:%w", err)
	}
	// Serialize `DisableDeposit`:
	if err = encoder.Encode(obj.DisableDeposit); err != nil {
		return fmt.Errorf("error while marshaling DisableDeposit:%w", err)
	}
	// Serialize `DisableWithdraw`:
	if err = encoder.Encode(obj.DisableWithdraw); err != nil {
		return fmt.Errorf("error while marshaling DisableWithdraw:%w", err)
	}
	// Serialize `DisableBuy`:
	if err = encoder.Encode(obj.DisableBuy); err != nil {
		return fmt.Errorf("error while marshaling DisableBuy:%w", err)
	}
	// Serialize `DisableSell`:
	if err = encoder.Encode(obj.DisableSell); err != nil {
		return fmt.Errorf("error while marshaling DisableSell:%w", err)
	}
	return nil
}

func (obj DisableEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding DisableEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *DisableEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_DisableEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_DisableEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `Admin`:
	if err = decoder.Decode(&obj.Admin); err != nil {
		return fmt.Errorf("error while unmarshaling Admin:%w", err)
	}
	// Deserialize `DisableCreatePool`:
	if err = decoder.Decode(&obj.DisableCreatePool); err != nil {
		return fmt.Errorf("error while unmarshaling DisableCreatePool:%w", err)
	}
	// Deserialize `DisableDeposit`:
	if err = decoder.Decode(&obj.DisableDeposit); err != nil {
		return fmt.Errorf("error while unmarshaling DisableDeposit:%w", err)
	}
	// Deserialize `DisableWithdraw`:
	if err = decoder.Decode(&obj.DisableWithdraw); err != nil {
		return fmt.Errorf("error while unmarshaling DisableWithdraw:%w", err)
	}
	// Deserialize `DisableBuy`:
	if err = decoder.Decode(&obj.DisableBuy); err != nil {
		return fmt.Errorf("error while unmarshaling DisableBuy:%w", err)
	}
	// Deserialize `DisableSell`:
	if err = decoder.Decode(&obj.DisableSell); err != nil {
		return fmt.Errorf("error while unmarshaling DisableSell:%w", err)
	}
	return nil
}

func (obj *DisableEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling DisableEvent: %w", err)
	}
	return nil
}

func UnmarshalDisableEvent(buf []byte) (*DisableEvent, error) {
	obj := new(DisableEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type ExtendAccountEvent struct {
	Timestamp   int64
	Account     solanago.PublicKey
	User        solanago.PublicKey
	CurrentSize uint64
	NewSize     uint64
}

func (obj ExtendAccountEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_ExtendAccountEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
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
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
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

type SellEvent struct {
	Timestamp                        int64
	BaseAmountIn                     uint64
	MinQuoteAmountOut                uint64
	UserBaseTokenReserves            uint64
	UserQuoteTokenReserves           uint64
	PoolBaseTokenReserves            uint64
	PoolQuoteTokenReserves           uint64
	QuoteAmountOut                   uint64
	LpFeeBasisPoints                 uint64
	LpFee                            uint64
	ProtocolFeeBasisPoints           uint64
	ProtocolFee                      uint64
	QuoteAmountOutWithoutLpFee       uint64
	UserQuoteAmountOut               uint64
	Pool                             solanago.PublicKey
	User                             solanago.PublicKey
	UserBaseTokenAccount             solanago.PublicKey
	UserQuoteTokenAccount            solanago.PublicKey
	ProtocolFeeRecipient             solanago.PublicKey
	ProtocolFeeRecipientTokenAccount solanago.PublicKey
	CoinCreator                      solanago.PublicKey
	CoinCreatorFeeBasisPoints        uint64
	CoinCreatorFee                   uint64
}

func (obj SellEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_SellEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `BaseAmountIn`:
	if err = encoder.Encode(obj.BaseAmountIn); err != nil {
		return fmt.Errorf("error while marshaling BaseAmountIn:%w", err)
	}
	// Serialize `MinQuoteAmountOut`:
	if err = encoder.Encode(obj.MinQuoteAmountOut); err != nil {
		return fmt.Errorf("error while marshaling MinQuoteAmountOut:%w", err)
	}
	// Serialize `UserBaseTokenReserves`:
	if err = encoder.Encode(obj.UserBaseTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling UserBaseTokenReserves:%w", err)
	}
	// Serialize `UserQuoteTokenReserves`:
	if err = encoder.Encode(obj.UserQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling UserQuoteTokenReserves:%w", err)
	}
	// Serialize `PoolBaseTokenReserves`:
	if err = encoder.Encode(obj.PoolBaseTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling PoolBaseTokenReserves:%w", err)
	}
	// Serialize `PoolQuoteTokenReserves`:
	if err = encoder.Encode(obj.PoolQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling PoolQuoteTokenReserves:%w", err)
	}
	// Serialize `QuoteAmountOut`:
	if err = encoder.Encode(obj.QuoteAmountOut); err != nil {
		return fmt.Errorf("error while marshaling QuoteAmountOut:%w", err)
	}
	// Serialize `LpFeeBasisPoints`:
	if err = encoder.Encode(obj.LpFeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling LpFeeBasisPoints:%w", err)
	}
	// Serialize `LpFee`:
	if err = encoder.Encode(obj.LpFee); err != nil {
		return fmt.Errorf("error while marshaling LpFee:%w", err)
	}
	// Serialize `ProtocolFeeBasisPoints`:
	if err = encoder.Encode(obj.ProtocolFeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeBasisPoints:%w", err)
	}
	// Serialize `ProtocolFee`:
	if err = encoder.Encode(obj.ProtocolFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFee:%w", err)
	}
	// Serialize `QuoteAmountOutWithoutLpFee`:
	if err = encoder.Encode(obj.QuoteAmountOutWithoutLpFee); err != nil {
		return fmt.Errorf("error while marshaling QuoteAmountOutWithoutLpFee:%w", err)
	}
	// Serialize `UserQuoteAmountOut`:
	if err = encoder.Encode(obj.UserQuoteAmountOut); err != nil {
		return fmt.Errorf("error while marshaling UserQuoteAmountOut:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `User`:
	if err = encoder.Encode(obj.User); err != nil {
		return fmt.Errorf("error while marshaling User:%w", err)
	}
	// Serialize `UserBaseTokenAccount`:
	if err = encoder.Encode(obj.UserBaseTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling UserBaseTokenAccount:%w", err)
	}
	// Serialize `UserQuoteTokenAccount`:
	if err = encoder.Encode(obj.UserQuoteTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling UserQuoteTokenAccount:%w", err)
	}
	// Serialize `ProtocolFeeRecipient`:
	if err = encoder.Encode(obj.ProtocolFeeRecipient); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeRecipient:%w", err)
	}
	// Serialize `ProtocolFeeRecipientTokenAccount`:
	if err = encoder.Encode(obj.ProtocolFeeRecipientTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeRecipientTokenAccount:%w", err)
	}
	// Serialize `CoinCreator`:
	if err = encoder.Encode(obj.CoinCreator); err != nil {
		return fmt.Errorf("error while marshaling CoinCreator:%w", err)
	}
	// Serialize `CoinCreatorFeeBasisPoints`:
	if err = encoder.Encode(obj.CoinCreatorFeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling CoinCreatorFeeBasisPoints:%w", err)
	}
	// Serialize `CoinCreatorFee`:
	if err = encoder.Encode(obj.CoinCreatorFee); err != nil {
		return fmt.Errorf("error while marshaling CoinCreatorFee:%w", err)
	}
	return nil
}

func (obj SellEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding SellEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *SellEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_SellEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_SellEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `BaseAmountIn`:
	if err = decoder.Decode(&obj.BaseAmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling BaseAmountIn:%w", err)
	}
	// Deserialize `MinQuoteAmountOut`:
	if err = decoder.Decode(&obj.MinQuoteAmountOut); err != nil {
		return fmt.Errorf("error while unmarshaling MinQuoteAmountOut:%w", err)
	}
	// Deserialize `UserBaseTokenReserves`:
	if err = decoder.Decode(&obj.UserBaseTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling UserBaseTokenReserves:%w", err)
	}
	// Deserialize `UserQuoteTokenReserves`:
	if err = decoder.Decode(&obj.UserQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling UserQuoteTokenReserves:%w", err)
	}
	// Deserialize `PoolBaseTokenReserves`:
	if err = decoder.Decode(&obj.PoolBaseTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling PoolBaseTokenReserves:%w", err)
	}
	// Deserialize `PoolQuoteTokenReserves`:
	if err = decoder.Decode(&obj.PoolQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling PoolQuoteTokenReserves:%w", err)
	}
	// Deserialize `QuoteAmountOut`:
	if err = decoder.Decode(&obj.QuoteAmountOut); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteAmountOut:%w", err)
	}
	// Deserialize `LpFeeBasisPoints`:
	if err = decoder.Decode(&obj.LpFeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling LpFeeBasisPoints:%w", err)
	}
	// Deserialize `LpFee`:
	if err = decoder.Decode(&obj.LpFee); err != nil {
		return fmt.Errorf("error while unmarshaling LpFee:%w", err)
	}
	// Deserialize `ProtocolFeeBasisPoints`:
	if err = decoder.Decode(&obj.ProtocolFeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeBasisPoints:%w", err)
	}
	// Deserialize `ProtocolFee`:
	if err = decoder.Decode(&obj.ProtocolFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFee:%w", err)
	}
	// Deserialize `QuoteAmountOutWithoutLpFee`:
	if err = decoder.Decode(&obj.QuoteAmountOutWithoutLpFee); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteAmountOutWithoutLpFee:%w", err)
	}
	// Deserialize `UserQuoteAmountOut`:
	if err = decoder.Decode(&obj.UserQuoteAmountOut); err != nil {
		return fmt.Errorf("error while unmarshaling UserQuoteAmountOut:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `User`:
	if err = decoder.Decode(&obj.User); err != nil {
		return fmt.Errorf("error while unmarshaling User:%w", err)
	}
	// Deserialize `UserBaseTokenAccount`:
	if err = decoder.Decode(&obj.UserBaseTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling UserBaseTokenAccount:%w", err)
	}
	// Deserialize `UserQuoteTokenAccount`:
	if err = decoder.Decode(&obj.UserQuoteTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling UserQuoteTokenAccount:%w", err)
	}
	// Deserialize `ProtocolFeeRecipient`:
	if err = decoder.Decode(&obj.ProtocolFeeRecipient); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeRecipient:%w", err)
	}
	// Deserialize `ProtocolFeeRecipientTokenAccount`:
	if err = decoder.Decode(&obj.ProtocolFeeRecipientTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeRecipientTokenAccount:%w", err)
	}
	// Deserialize `CoinCreator`:
	if err = decoder.Decode(&obj.CoinCreator); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreator:%w", err)
	}
	// Deserialize `CoinCreatorFeeBasisPoints`:
	if err = decoder.Decode(&obj.CoinCreatorFeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreatorFeeBasisPoints:%w", err)
	}
	// Deserialize `CoinCreatorFee`:
	if err = decoder.Decode(&obj.CoinCreatorFee); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreatorFee:%w", err)
	}
	return nil
}

func (obj *SellEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling SellEvent: %w", err)
	}
	return nil
}

func UnmarshalSellEvent(buf []byte) (*SellEvent, error) {
	obj := new(SellEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type SetBondingCurveCoinCreatorEvent struct {
	Timestamp    int64
	BaseMint     solanago.PublicKey
	Pool         solanago.PublicKey
	BondingCurve solanago.PublicKey
	CoinCreator  solanago.PublicKey
}

func (obj SetBondingCurveCoinCreatorEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_SetBondingCurveCoinCreatorEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `BaseMint`:
	if err = encoder.Encode(obj.BaseMint); err != nil {
		return fmt.Errorf("error while marshaling BaseMint:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `BondingCurve`:
	if err = encoder.Encode(obj.BondingCurve); err != nil {
		return fmt.Errorf("error while marshaling BondingCurve:%w", err)
	}
	// Serialize `CoinCreator`:
	if err = encoder.Encode(obj.CoinCreator); err != nil {
		return fmt.Errorf("error while marshaling CoinCreator:%w", err)
	}
	return nil
}

func (obj SetBondingCurveCoinCreatorEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding SetBondingCurveCoinCreatorEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *SetBondingCurveCoinCreatorEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_SetBondingCurveCoinCreatorEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_SetBondingCurveCoinCreatorEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `BaseMint`:
	if err = decoder.Decode(&obj.BaseMint); err != nil {
		return fmt.Errorf("error while unmarshaling BaseMint:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `BondingCurve`:
	if err = decoder.Decode(&obj.BondingCurve); err != nil {
		return fmt.Errorf("error while unmarshaling BondingCurve:%w", err)
	}
	// Deserialize `CoinCreator`:
	if err = decoder.Decode(&obj.CoinCreator); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreator:%w", err)
	}
	return nil
}

func (obj *SetBondingCurveCoinCreatorEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling SetBondingCurveCoinCreatorEvent: %w", err)
	}
	return nil
}

func UnmarshalSetBondingCurveCoinCreatorEvent(buf []byte) (*SetBondingCurveCoinCreatorEvent, error) {
	obj := new(SetBondingCurveCoinCreatorEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type SetMetaplexCoinCreatorEvent struct {
	Timestamp   int64
	BaseMint    solanago.PublicKey
	Pool        solanago.PublicKey
	Metadata    solanago.PublicKey
	CoinCreator solanago.PublicKey
}

func (obj SetMetaplexCoinCreatorEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_SetMetaplexCoinCreatorEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `BaseMint`:
	if err = encoder.Encode(obj.BaseMint); err != nil {
		return fmt.Errorf("error while marshaling BaseMint:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Metadata`:
	if err = encoder.Encode(obj.Metadata); err != nil {
		return fmt.Errorf("error while marshaling Metadata:%w", err)
	}
	// Serialize `CoinCreator`:
	if err = encoder.Encode(obj.CoinCreator); err != nil {
		return fmt.Errorf("error while marshaling CoinCreator:%w", err)
	}
	return nil
}

func (obj SetMetaplexCoinCreatorEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding SetMetaplexCoinCreatorEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *SetMetaplexCoinCreatorEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_SetMetaplexCoinCreatorEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_SetMetaplexCoinCreatorEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `BaseMint`:
	if err = decoder.Decode(&obj.BaseMint); err != nil {
		return fmt.Errorf("error while unmarshaling BaseMint:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Metadata`:
	if err = decoder.Decode(&obj.Metadata); err != nil {
		return fmt.Errorf("error while unmarshaling Metadata:%w", err)
	}
	// Deserialize `CoinCreator`:
	if err = decoder.Decode(&obj.CoinCreator); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreator:%w", err)
	}
	return nil
}

func (obj *SetMetaplexCoinCreatorEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling SetMetaplexCoinCreatorEvent: %w", err)
	}
	return nil
}

func UnmarshalSetMetaplexCoinCreatorEvent(buf []byte) (*SetMetaplexCoinCreatorEvent, error) {
	obj := new(SetMetaplexCoinCreatorEvent)
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

type UpdateAdminEvent struct {
	Timestamp int64
	Admin     solanago.PublicKey
	NewAdmin  solanago.PublicKey
}

func (obj UpdateAdminEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_UpdateAdminEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `Admin`:
	if err = encoder.Encode(obj.Admin); err != nil {
		return fmt.Errorf("error while marshaling Admin:%w", err)
	}
	// Serialize `NewAdmin`:
	if err = encoder.Encode(obj.NewAdmin); err != nil {
		return fmt.Errorf("error while marshaling NewAdmin:%w", err)
	}
	return nil
}

func (obj UpdateAdminEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding UpdateAdminEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *UpdateAdminEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_UpdateAdminEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_UpdateAdminEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `Admin`:
	if err = decoder.Decode(&obj.Admin); err != nil {
		return fmt.Errorf("error while unmarshaling Admin:%w", err)
	}
	// Deserialize `NewAdmin`:
	if err = decoder.Decode(&obj.NewAdmin); err != nil {
		return fmt.Errorf("error while unmarshaling NewAdmin:%w", err)
	}
	return nil
}

func (obj *UpdateAdminEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling UpdateAdminEvent: %w", err)
	}
	return nil
}

func UnmarshalUpdateAdminEvent(buf []byte) (*UpdateAdminEvent, error) {
	obj := new(UpdateAdminEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type UpdateFeeConfigEvent struct {
	Timestamp                    int64
	Admin                        solanago.PublicKey
	LpFeeBasisPoints             uint64
	ProtocolFeeBasisPoints       uint64
	ProtocolFeeRecipients        [8]solanago.PublicKey
	CoinCreatorFeeBasisPoints    uint64
	AdminSetCoinCreatorAuthority solanago.PublicKey
}

func (obj UpdateFeeConfigEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_UpdateFeeConfigEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `Admin`:
	if err = encoder.Encode(obj.Admin); err != nil {
		return fmt.Errorf("error while marshaling Admin:%w", err)
	}
	// Serialize `LpFeeBasisPoints`:
	if err = encoder.Encode(obj.LpFeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling LpFeeBasisPoints:%w", err)
	}
	// Serialize `ProtocolFeeBasisPoints`:
	if err = encoder.Encode(obj.ProtocolFeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeBasisPoints:%w", err)
	}
	// Serialize `ProtocolFeeRecipients`:
	if err = encoder.Encode(obj.ProtocolFeeRecipients); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeRecipients:%w", err)
	}
	// Serialize `CoinCreatorFeeBasisPoints`:
	if err = encoder.Encode(obj.CoinCreatorFeeBasisPoints); err != nil {
		return fmt.Errorf("error while marshaling CoinCreatorFeeBasisPoints:%w", err)
	}
	// Serialize `AdminSetCoinCreatorAuthority`:
	if err = encoder.Encode(obj.AdminSetCoinCreatorAuthority); err != nil {
		return fmt.Errorf("error while marshaling AdminSetCoinCreatorAuthority:%w", err)
	}
	return nil
}

func (obj UpdateFeeConfigEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding UpdateFeeConfigEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *UpdateFeeConfigEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_UpdateFeeConfigEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_UpdateFeeConfigEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `Admin`:
	if err = decoder.Decode(&obj.Admin); err != nil {
		return fmt.Errorf("error while unmarshaling Admin:%w", err)
	}
	// Deserialize `LpFeeBasisPoints`:
	if err = decoder.Decode(&obj.LpFeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling LpFeeBasisPoints:%w", err)
	}
	// Deserialize `ProtocolFeeBasisPoints`:
	if err = decoder.Decode(&obj.ProtocolFeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeBasisPoints:%w", err)
	}
	// Deserialize `ProtocolFeeRecipients`:
	if err = decoder.Decode(&obj.ProtocolFeeRecipients); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeRecipients:%w", err)
	}
	// Deserialize `CoinCreatorFeeBasisPoints`:
	if err = decoder.Decode(&obj.CoinCreatorFeeBasisPoints); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreatorFeeBasisPoints:%w", err)
	}
	// Deserialize `AdminSetCoinCreatorAuthority`:
	if err = decoder.Decode(&obj.AdminSetCoinCreatorAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling AdminSetCoinCreatorAuthority:%w", err)
	}
	return nil
}

func (obj *UpdateFeeConfigEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling UpdateFeeConfigEvent: %w", err)
	}
	return nil
}

func UnmarshalUpdateFeeConfigEvent(buf []byte) (*UpdateFeeConfigEvent, error) {
	obj := new(UpdateFeeConfigEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type WithdrawEvent struct {
	Timestamp              int64
	LpTokenAmountIn        uint64
	MinBaseAmountOut       uint64
	MinQuoteAmountOut      uint64
	UserBaseTokenReserves  uint64
	UserQuoteTokenReserves uint64
	PoolBaseTokenReserves  uint64
	PoolQuoteTokenReserves uint64
	BaseAmountOut          uint64
	QuoteAmountOut         uint64
	LpMintSupply           uint64
	Pool                   solanago.PublicKey
	User                   solanago.PublicKey
	UserBaseTokenAccount   solanago.PublicKey
	UserQuoteTokenAccount  solanago.PublicKey
	UserPoolTokenAccount   solanago.PublicKey
}

func (obj WithdrawEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_WithdrawEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Timestamp`:
	if err = encoder.Encode(obj.Timestamp); err != nil {
		return fmt.Errorf("error while marshaling Timestamp:%w", err)
	}
	// Serialize `LpTokenAmountIn`:
	if err = encoder.Encode(obj.LpTokenAmountIn); err != nil {
		return fmt.Errorf("error while marshaling LpTokenAmountIn:%w", err)
	}
	// Serialize `MinBaseAmountOut`:
	if err = encoder.Encode(obj.MinBaseAmountOut); err != nil {
		return fmt.Errorf("error while marshaling MinBaseAmountOut:%w", err)
	}
	// Serialize `MinQuoteAmountOut`:
	if err = encoder.Encode(obj.MinQuoteAmountOut); err != nil {
		return fmt.Errorf("error while marshaling MinQuoteAmountOut:%w", err)
	}
	// Serialize `UserBaseTokenReserves`:
	if err = encoder.Encode(obj.UserBaseTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling UserBaseTokenReserves:%w", err)
	}
	// Serialize `UserQuoteTokenReserves`:
	if err = encoder.Encode(obj.UserQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling UserQuoteTokenReserves:%w", err)
	}
	// Serialize `PoolBaseTokenReserves`:
	if err = encoder.Encode(obj.PoolBaseTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling PoolBaseTokenReserves:%w", err)
	}
	// Serialize `PoolQuoteTokenReserves`:
	if err = encoder.Encode(obj.PoolQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while marshaling PoolQuoteTokenReserves:%w", err)
	}
	// Serialize `BaseAmountOut`:
	if err = encoder.Encode(obj.BaseAmountOut); err != nil {
		return fmt.Errorf("error while marshaling BaseAmountOut:%w", err)
	}
	// Serialize `QuoteAmountOut`:
	if err = encoder.Encode(obj.QuoteAmountOut); err != nil {
		return fmt.Errorf("error while marshaling QuoteAmountOut:%w", err)
	}
	// Serialize `LpMintSupply`:
	if err = encoder.Encode(obj.LpMintSupply); err != nil {
		return fmt.Errorf("error while marshaling LpMintSupply:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `User`:
	if err = encoder.Encode(obj.User); err != nil {
		return fmt.Errorf("error while marshaling User:%w", err)
	}
	// Serialize `UserBaseTokenAccount`:
	if err = encoder.Encode(obj.UserBaseTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling UserBaseTokenAccount:%w", err)
	}
	// Serialize `UserQuoteTokenAccount`:
	if err = encoder.Encode(obj.UserQuoteTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling UserQuoteTokenAccount:%w", err)
	}
	// Serialize `UserPoolTokenAccount`:
	if err = encoder.Encode(obj.UserPoolTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling UserPoolTokenAccount:%w", err)
	}
	return nil
}

func (obj WithdrawEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding WithdrawEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *WithdrawEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_WithdrawEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_WithdrawEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Timestamp`:
	if err = decoder.Decode(&obj.Timestamp); err != nil {
		return fmt.Errorf("error while unmarshaling Timestamp:%w", err)
	}
	// Deserialize `LpTokenAmountIn`:
	if err = decoder.Decode(&obj.LpTokenAmountIn); err != nil {
		return fmt.Errorf("error while unmarshaling LpTokenAmountIn:%w", err)
	}
	// Deserialize `MinBaseAmountOut`:
	if err = decoder.Decode(&obj.MinBaseAmountOut); err != nil {
		return fmt.Errorf("error while unmarshaling MinBaseAmountOut:%w", err)
	}
	// Deserialize `MinQuoteAmountOut`:
	if err = decoder.Decode(&obj.MinQuoteAmountOut); err != nil {
		return fmt.Errorf("error while unmarshaling MinQuoteAmountOut:%w", err)
	}
	// Deserialize `UserBaseTokenReserves`:
	if err = decoder.Decode(&obj.UserBaseTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling UserBaseTokenReserves:%w", err)
	}
	// Deserialize `UserQuoteTokenReserves`:
	if err = decoder.Decode(&obj.UserQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling UserQuoteTokenReserves:%w", err)
	}
	// Deserialize `PoolBaseTokenReserves`:
	if err = decoder.Decode(&obj.PoolBaseTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling PoolBaseTokenReserves:%w", err)
	}
	// Deserialize `PoolQuoteTokenReserves`:
	if err = decoder.Decode(&obj.PoolQuoteTokenReserves); err != nil {
		return fmt.Errorf("error while unmarshaling PoolQuoteTokenReserves:%w", err)
	}
	// Deserialize `BaseAmountOut`:
	if err = decoder.Decode(&obj.BaseAmountOut); err != nil {
		return fmt.Errorf("error while unmarshaling BaseAmountOut:%w", err)
	}
	// Deserialize `QuoteAmountOut`:
	if err = decoder.Decode(&obj.QuoteAmountOut); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteAmountOut:%w", err)
	}
	// Deserialize `LpMintSupply`:
	if err = decoder.Decode(&obj.LpMintSupply); err != nil {
		return fmt.Errorf("error while unmarshaling LpMintSupply:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `User`:
	if err = decoder.Decode(&obj.User); err != nil {
		return fmt.Errorf("error while unmarshaling User:%w", err)
	}
	// Deserialize `UserBaseTokenAccount`:
	if err = decoder.Decode(&obj.UserBaseTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling UserBaseTokenAccount:%w", err)
	}
	// Deserialize `UserQuoteTokenAccount`:
	if err = decoder.Decode(&obj.UserQuoteTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling UserQuoteTokenAccount:%w", err)
	}
	// Deserialize `UserPoolTokenAccount`:
	if err = decoder.Decode(&obj.UserPoolTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling UserPoolTokenAccount:%w", err)
	}
	return nil
}

func (obj *WithdrawEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling WithdrawEvent: %w", err)
	}
	return nil
}

func UnmarshalWithdrawEvent(buf []byte) (*WithdrawEvent, error) {
	obj := new(WithdrawEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
