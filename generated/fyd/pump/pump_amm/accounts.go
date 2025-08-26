package pump_amm

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

type GlobalConfig struct {
	// The admin pubkey
	Admin solanago.PublicKey

	// The lp fee in basis points (0.01%)
	LpFeeBasisPoints uint64

	// The protocol fee in basis points (0.01%)
	ProtocolFeeBasisPoints uint64

	// Flags to disable certain functionality
	// bit 0 - Disable create pool
	// bit 1 - Disable deposit
	// bit 2 - Disable withdraw
	// bit 3 - Disable buy
	// bit 4 - Disable sell
	DisableFlags uint8

	// Addresses of the protocol fee recipients
	ProtocolFeeRecipients [8]solanago.PublicKey

	// The coin creator fee in basis points (0.01%)
	CoinCreatorFeeBasisPoints uint64

	// The admin authority for setting coin creators
	AdminSetCoinCreatorAuthority solanago.PublicKey
}

func (obj GlobalConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_GlobalConfig[:], false)
	if err != nil {
		return err
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
	// Serialize `DisableFlags`:
	if err = encoder.Encode(obj.DisableFlags); err != nil {
		return fmt.Errorf("error while marshaling DisableFlags:%w", err)
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

func (obj GlobalConfig) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding GlobalConfig: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *GlobalConfig) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_GlobalConfig[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_GlobalConfig[:],
				fmt.Sprint(discriminator[:]))
		}
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
	// Deserialize `DisableFlags`:
	if err = decoder.Decode(&obj.DisableFlags); err != nil {
		return fmt.Errorf("error while unmarshaling DisableFlags:%w", err)
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

func (obj *GlobalConfig) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling GlobalConfig: %w", err)
	}
	return nil
}

func UnmarshalGlobalConfig(buf []byte) (*GlobalConfig, error) {
	obj := new(GlobalConfig)
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

type Pool struct {
	PoolBump              uint8
	Index                 uint16
	Creator               solanago.PublicKey
	BaseMint              solanago.PublicKey
	QuoteMint             solanago.PublicKey
	LpMint                solanago.PublicKey
	PoolBaseTokenAccount  solanago.PublicKey
	PoolQuoteTokenAccount solanago.PublicKey

	// True circulating supply without burns and lock-ups
	LpSupply    uint64
	CoinCreator solanago.PublicKey
}

func (obj Pool) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Pool[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolBump`:
	if err = encoder.Encode(obj.PoolBump); err != nil {
		return fmt.Errorf("error while marshaling PoolBump:%w", err)
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
	// Serialize `LpMint`:
	if err = encoder.Encode(obj.LpMint); err != nil {
		return fmt.Errorf("error while marshaling LpMint:%w", err)
	}
	// Serialize `PoolBaseTokenAccount`:
	if err = encoder.Encode(obj.PoolBaseTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling PoolBaseTokenAccount:%w", err)
	}
	// Serialize `PoolQuoteTokenAccount`:
	if err = encoder.Encode(obj.PoolQuoteTokenAccount); err != nil {
		return fmt.Errorf("error while marshaling PoolQuoteTokenAccount:%w", err)
	}
	// Serialize `LpSupply`:
	if err = encoder.Encode(obj.LpSupply); err != nil {
		return fmt.Errorf("error while marshaling LpSupply:%w", err)
	}
	// Serialize `CoinCreator`:
	if err = encoder.Encode(obj.CoinCreator); err != nil {
		return fmt.Errorf("error while marshaling CoinCreator:%w", err)
	}
	return nil
}

func (obj Pool) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Pool: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Pool) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Pool[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Pool[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolBump`:
	if err = decoder.Decode(&obj.PoolBump); err != nil {
		return fmt.Errorf("error while unmarshaling PoolBump:%w", err)
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
	// Deserialize `LpMint`:
	if err = decoder.Decode(&obj.LpMint); err != nil {
		return fmt.Errorf("error while unmarshaling LpMint:%w", err)
	}
	// Deserialize `PoolBaseTokenAccount`:
	if err = decoder.Decode(&obj.PoolBaseTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling PoolBaseTokenAccount:%w", err)
	}
	// Deserialize `PoolQuoteTokenAccount`:
	if err = decoder.Decode(&obj.PoolQuoteTokenAccount); err != nil {
		return fmt.Errorf("error while unmarshaling PoolQuoteTokenAccount:%w", err)
	}
	// Deserialize `LpSupply`:
	if err = decoder.Decode(&obj.LpSupply); err != nil {
		return fmt.Errorf("error while unmarshaling LpSupply:%w", err)
	}
	// Deserialize `CoinCreator`:
	if err = decoder.Decode(&obj.CoinCreator); err != nil {
		return fmt.Errorf("error while unmarshaling CoinCreator:%w", err)
	}
	return nil
}

func (obj *Pool) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Pool: %w", err)
	}
	return nil
}

func UnmarshalPool(buf []byte) (*Pool, error) {
	obj := new(Pool)
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
