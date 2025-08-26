package pump_curve

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type BondingCurve struct {
	VirtualTokenReserves uint64
	VirtualSolReserves   uint64
	RealTokenReserves    uint64
	RealSolReserves      uint64
	TokenTotalSupply     uint64
	Complete             bool
	Creator              solanago.PublicKey
}

func (obj BondingCurve) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_BondingCurve[:], false)
	if err != nil {
		return err
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
	// Serialize `RealSolReserves`:
	if err = encoder.Encode(obj.RealSolReserves); err != nil {
		return fmt.Errorf("error while marshaling RealSolReserves:%w", err)
	}
	// Serialize `TokenTotalSupply`:
	if err = encoder.Encode(obj.TokenTotalSupply); err != nil {
		return fmt.Errorf("error while marshaling TokenTotalSupply:%w", err)
	}
	// Serialize `Complete`:
	if err = encoder.Encode(obj.Complete); err != nil {
		return fmt.Errorf("error while marshaling Complete:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
	}
	return nil
}

func (obj BondingCurve) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding BondingCurve: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *BondingCurve) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_BondingCurve[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_BondingCurve[:],
				fmt.Sprint(discriminator[:]))
		}
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
	// Deserialize `RealSolReserves`:
	if err = decoder.Decode(&obj.RealSolReserves); err != nil {
		return fmt.Errorf("error while unmarshaling RealSolReserves:%w", err)
	}
	// Deserialize `TokenTotalSupply`:
	if err = decoder.Decode(&obj.TokenTotalSupply); err != nil {
		return fmt.Errorf("error while unmarshaling TokenTotalSupply:%w", err)
	}
	// Deserialize `Complete`:
	if err = decoder.Decode(&obj.Complete); err != nil {
		return fmt.Errorf("error while unmarshaling Complete:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
	}
	return nil
}

func (obj *BondingCurve) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling BondingCurve: %w", err)
	}
	return nil
}

func UnmarshalBondingCurve(buf []byte) (*BondingCurve, error) {
	obj := new(BondingCurve)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type Global struct {
	// Unused
	Initialized                 bool
	Authority                   solanago.PublicKey
	FeeRecipient                solanago.PublicKey
	InitialVirtualTokenReserves uint64
	InitialVirtualSolReserves   uint64
	InitialRealTokenReserves    uint64
	TokenTotalSupply            uint64
	FeeBasisPoints              uint64
	WithdrawAuthority           solanago.PublicKey

	// Unused
	EnableMigrate            bool
	PoolMigrationFee         uint64
	CreatorFeeBasisPoints    uint64
	FeeRecipients            [7]solanago.PublicKey
	SetCreatorAuthority      solanago.PublicKey
	AdminSetCreatorAuthority solanago.PublicKey
}

func (obj Global) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Global[:], false)
	if err != nil {
		return err
	}
	// Serialize `Initialized`:
	if err = encoder.Encode(obj.Initialized); err != nil {
		return fmt.Errorf("error while marshaling Initialized:%w", err)
	}
	// Serialize `Authority`:
	if err = encoder.Encode(obj.Authority); err != nil {
		return fmt.Errorf("error while marshaling Authority:%w", err)
	}
	// Serialize `FeeRecipient`:
	if err = encoder.Encode(obj.FeeRecipient); err != nil {
		return fmt.Errorf("error while marshaling FeeRecipient:%w", err)
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

func (obj Global) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Global: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Global) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Global[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Global[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Initialized`:
	if err = decoder.Decode(&obj.Initialized); err != nil {
		return fmt.Errorf("error while unmarshaling Initialized:%w", err)
	}
	// Deserialize `Authority`:
	if err = decoder.Decode(&obj.Authority); err != nil {
		return fmt.Errorf("error while unmarshaling Authority:%w", err)
	}
	// Deserialize `FeeRecipient`:
	if err = decoder.Decode(&obj.FeeRecipient); err != nil {
		return fmt.Errorf("error while unmarshaling FeeRecipient:%w", err)
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

func (obj *Global) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Global: %w", err)
	}
	return nil
}

func UnmarshalGlobal(buf []byte) (*Global, error) {
	obj := new(Global)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type GlobalVolumeAccumulator struct {
	StartTime        int64
	EndTime          int64
	SecondsInADay    int64
	Mint             solanago.PublicKey
	TotalTokenSupply [30]uint64
	SolVolumes       [30]uint64
}

func (obj GlobalVolumeAccumulator) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_GlobalVolumeAccumulator[:], false)
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
	// Serialize `SecondsInADay`:
	if err = encoder.Encode(obj.SecondsInADay); err != nil {
		return fmt.Errorf("error while marshaling SecondsInADay:%w", err)
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `TotalTokenSupply`:
	if err = encoder.Encode(obj.TotalTokenSupply); err != nil {
		return fmt.Errorf("error while marshaling TotalTokenSupply:%w", err)
	}
	// Serialize `SolVolumes`:
	if err = encoder.Encode(obj.SolVolumes); err != nil {
		return fmt.Errorf("error while marshaling SolVolumes:%w", err)
	}
	return nil
}

func (obj GlobalVolumeAccumulator) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding GlobalVolumeAccumulator: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *GlobalVolumeAccumulator) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_GlobalVolumeAccumulator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_GlobalVolumeAccumulator[:],
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
	// Deserialize `SecondsInADay`:
	if err = decoder.Decode(&obj.SecondsInADay); err != nil {
		return fmt.Errorf("error while unmarshaling SecondsInADay:%w", err)
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `TotalTokenSupply`:
	if err = decoder.Decode(&obj.TotalTokenSupply); err != nil {
		return fmt.Errorf("error while unmarshaling TotalTokenSupply:%w", err)
	}
	// Deserialize `SolVolumes`:
	if err = decoder.Decode(&obj.SolVolumes); err != nil {
		return fmt.Errorf("error while unmarshaling SolVolumes:%w", err)
	}
	return nil
}

func (obj *GlobalVolumeAccumulator) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling GlobalVolumeAccumulator: %w", err)
	}
	return nil
}

func UnmarshalGlobalVolumeAccumulator(buf []byte) (*GlobalVolumeAccumulator, error) {
	obj := new(GlobalVolumeAccumulator)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type UserVolumeAccumulator struct {
	User                  solanago.PublicKey
	NeedsClaim            bool
	TotalUnclaimedTokens  uint64
	TotalClaimedTokens    uint64
	CurrentSolVolume      uint64
	LastUpdateTimestamp   int64
	HasTotalClaimedTokens bool
}

func (obj UserVolumeAccumulator) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_UserVolumeAccumulator[:], false)
	if err != nil {
		return err
	}
	// Serialize `User`:
	if err = encoder.Encode(obj.User); err != nil {
		return fmt.Errorf("error while marshaling User:%w", err)
	}
	// Serialize `NeedsClaim`:
	if err = encoder.Encode(obj.NeedsClaim); err != nil {
		return fmt.Errorf("error while marshaling NeedsClaim:%w", err)
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
	// Serialize `HasTotalClaimedTokens`:
	if err = encoder.Encode(obj.HasTotalClaimedTokens); err != nil {
		return fmt.Errorf("error while marshaling HasTotalClaimedTokens:%w", err)
	}
	return nil
}

func (obj UserVolumeAccumulator) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding UserVolumeAccumulator: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *UserVolumeAccumulator) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_UserVolumeAccumulator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_UserVolumeAccumulator[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `User`:
	if err = decoder.Decode(&obj.User); err != nil {
		return fmt.Errorf("error while unmarshaling User:%w", err)
	}
	// Deserialize `NeedsClaim`:
	if err = decoder.Decode(&obj.NeedsClaim); err != nil {
		return fmt.Errorf("error while unmarshaling NeedsClaim:%w", err)
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
	// Deserialize `HasTotalClaimedTokens`:
	if err = decoder.Decode(&obj.HasTotalClaimedTokens); err != nil {
		return fmt.Errorf("error while unmarshaling HasTotalClaimedTokens:%w", err)
	}
	return nil
}

func (obj *UserVolumeAccumulator) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling UserVolumeAccumulator: %w", err)
	}
	return nil
}

func UnmarshalUserVolumeAccumulator(buf []byte) (*UserVolumeAccumulator, error) {
	obj := new(UserVolumeAccumulator)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
