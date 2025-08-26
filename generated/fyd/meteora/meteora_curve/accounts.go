package meteora_curve

import (
	"bytes"
	"fmt"

	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

// Parameter that set by the protocol
type ClaimFeeOperator struct {
	// operator
	Operator solanago.PublicKey

	// Reserve
	Padding [128]uint8
}

func (obj ClaimFeeOperator) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_ClaimFeeOperator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Operator`:
	if err = encoder.Encode(obj.Operator); err != nil {
		return fmt.Errorf("error while marshaling Operator:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj ClaimFeeOperator) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding ClaimFeeOperator: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *ClaimFeeOperator) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_ClaimFeeOperator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_ClaimFeeOperator[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Operator`:
	if err = decoder.Decode(&obj.Operator); err != nil {
		return fmt.Errorf("error while unmarshaling Operator:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *ClaimFeeOperator) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling ClaimFeeOperator: %w", err)
	}
	return nil
}

func UnmarshalClaimFeeOperator(buf []byte) (*ClaimFeeOperator, error) {
	obj := new(ClaimFeeOperator)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type Config struct {
	PoolFees             PoolFees
	ActivationDuration   uint64
	VaultConfigKey       solanago.PublicKey
	PoolCreatorAuthority solanago.PublicKey
	ActivationType       uint8
	PartnerFeeNumerator  uint64
	Padding              [219]uint8
}

func (obj Config) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Config[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolFees`:
	if err = encoder.Encode(obj.PoolFees); err != nil {
		return fmt.Errorf("error while marshaling PoolFees:%w", err)
	}
	// Serialize `ActivationDuration`:
	if err = encoder.Encode(obj.ActivationDuration); err != nil {
		return fmt.Errorf("error while marshaling ActivationDuration:%w", err)
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
	// Serialize `PartnerFeeNumerator`:
	if err = encoder.Encode(obj.PartnerFeeNumerator); err != nil {
		return fmt.Errorf("error while marshaling PartnerFeeNumerator:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj Config) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Config: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Config) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Config[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Config[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolFees`:
	if err = decoder.Decode(&obj.PoolFees); err != nil {
		return fmt.Errorf("error while unmarshaling PoolFees:%w", err)
	}
	// Deserialize `ActivationDuration`:
	if err = decoder.Decode(&obj.ActivationDuration); err != nil {
		return fmt.Errorf("error while unmarshaling ActivationDuration:%w", err)
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
	// Deserialize `PartnerFeeNumerator`:
	if err = decoder.Decode(&obj.PartnerFeeNumerator); err != nil {
		return fmt.Errorf("error while unmarshaling PartnerFeeNumerator:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *Config) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Config: %w", err)
	}
	return nil
}

func UnmarshalConfig(buf []byte) (*Config, error) {
	obj := new(Config)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// State of lock escrow account
type LockEscrow struct {
	Pool                solanago.PublicKey
	Owner               solanago.PublicKey
	EscrowVault         solanago.PublicKey
	Bump                uint8
	TotalLockedAmount   uint64
	LpPerToken          binary.Uint128
	UnclaimedFeePending uint64
	AFee                uint64
	BFee                uint64
}

func (obj LockEscrow) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_LockEscrow[:], false)
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
	// Serialize `EscrowVault`:
	if err = encoder.Encode(obj.EscrowVault); err != nil {
		return fmt.Errorf("error while marshaling EscrowVault:%w", err)
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	// Serialize `TotalLockedAmount`:
	if err = encoder.Encode(obj.TotalLockedAmount); err != nil {
		return fmt.Errorf("error while marshaling TotalLockedAmount:%w", err)
	}
	// Serialize `LpPerToken`:
	if err = encoder.Encode(obj.LpPerToken); err != nil {
		return fmt.Errorf("error while marshaling LpPerToken:%w", err)
	}
	// Serialize `UnclaimedFeePending`:
	if err = encoder.Encode(obj.UnclaimedFeePending); err != nil {
		return fmt.Errorf("error while marshaling UnclaimedFeePending:%w", err)
	}
	// Serialize `AFee`:
	if err = encoder.Encode(obj.AFee); err != nil {
		return fmt.Errorf("error while marshaling AFee:%w", err)
	}
	// Serialize `BFee`:
	if err = encoder.Encode(obj.BFee); err != nil {
		return fmt.Errorf("error while marshaling BFee:%w", err)
	}
	return nil
}

func (obj LockEscrow) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding LockEscrow: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *LockEscrow) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_LockEscrow[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_LockEscrow[:],
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
	// Deserialize `EscrowVault`:
	if err = decoder.Decode(&obj.EscrowVault); err != nil {
		return fmt.Errorf("error while unmarshaling EscrowVault:%w", err)
	}
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	// Deserialize `TotalLockedAmount`:
	if err = decoder.Decode(&obj.TotalLockedAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TotalLockedAmount:%w", err)
	}
	// Deserialize `LpPerToken`:
	if err = decoder.Decode(&obj.LpPerToken); err != nil {
		return fmt.Errorf("error while unmarshaling LpPerToken:%w", err)
	}
	// Deserialize `UnclaimedFeePending`:
	if err = decoder.Decode(&obj.UnclaimedFeePending); err != nil {
		return fmt.Errorf("error while unmarshaling UnclaimedFeePending:%w", err)
	}
	// Deserialize `AFee`:
	if err = decoder.Decode(&obj.AFee); err != nil {
		return fmt.Errorf("error while unmarshaling AFee:%w", err)
	}
	// Deserialize `BFee`:
	if err = decoder.Decode(&obj.BFee); err != nil {
		return fmt.Errorf("error while unmarshaling BFee:%w", err)
	}
	return nil
}

func (obj *LockEscrow) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling LockEscrow: %w", err)
	}
	return nil
}

func UnmarshalLockEscrow(buf []byte) (*LockEscrow, error) {
	obj := new(LockEscrow)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type MeteoraDammMigrationMetadata struct {
	// pool
	VirtualPool solanago.PublicKey

	// !!! BE CAREFUL to use tomestone field, previous is pool creator
	Padding0 [32]uint8

	// partner
	Partner solanago.PublicKey

	// lp mint
	LpMint solanago.PublicKey

	// partner locked lp
	PartnerLockedLp uint64

	// partner lp
	PartnerLp uint64

	// creator locked lp
	CreatorLockedLp uint64

	// creator lp
	CreatorLp uint64

	// padding
	Padding1 uint8

	// flag to check whether lp is locked for creator
	CreatorLockedStatus uint8

	// flag to check whether lp is locked for partner
	PartnerLockedStatus uint8

	// flag to check whether creator has claimed lp token
	CreatorClaimStatus uint8

	// flag to check whether partner has claimed lp token
	PartnerClaimStatus uint8

	// Reserve
	Padding [107]uint8
}

func (obj MeteoraDammMigrationMetadata) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_MeteoraDammMigrationMetadata[:], false)
	if err != nil {
		return err
	}
	// Serialize `VirtualPool`:
	if err = encoder.Encode(obj.VirtualPool); err != nil {
		return fmt.Errorf("error while marshaling VirtualPool:%w", err)
	}
	// Serialize `Padding0`:
	if err = encoder.Encode(obj.Padding0); err != nil {
		return fmt.Errorf("error while marshaling Padding0:%w", err)
	}
	// Serialize `Partner`:
	if err = encoder.Encode(obj.Partner); err != nil {
		return fmt.Errorf("error while marshaling Partner:%w", err)
	}
	// Serialize `LpMint`:
	if err = encoder.Encode(obj.LpMint); err != nil {
		return fmt.Errorf("error while marshaling LpMint:%w", err)
	}
	// Serialize `PartnerLockedLp`:
	if err = encoder.Encode(obj.PartnerLockedLp); err != nil {
		return fmt.Errorf("error while marshaling PartnerLockedLp:%w", err)
	}
	// Serialize `PartnerLp`:
	if err = encoder.Encode(obj.PartnerLp); err != nil {
		return fmt.Errorf("error while marshaling PartnerLp:%w", err)
	}
	// Serialize `CreatorLockedLp`:
	if err = encoder.Encode(obj.CreatorLockedLp); err != nil {
		return fmt.Errorf("error while marshaling CreatorLockedLp:%w", err)
	}
	// Serialize `CreatorLp`:
	if err = encoder.Encode(obj.CreatorLp); err != nil {
		return fmt.Errorf("error while marshaling CreatorLp:%w", err)
	}
	// Serialize `Padding0`:
	if err = encoder.Encode(obj.Padding0); err != nil {
		return fmt.Errorf("error while marshaling Padding0:%w", err)
	}
	// Serialize `CreatorLockedStatus`:
	if err = encoder.Encode(obj.CreatorLockedStatus); err != nil {
		return fmt.Errorf("error while marshaling CreatorLockedStatus:%w", err)
	}
	// Serialize `PartnerLockedStatus`:
	if err = encoder.Encode(obj.PartnerLockedStatus); err != nil {
		return fmt.Errorf("error while marshaling PartnerLockedStatus:%w", err)
	}
	// Serialize `CreatorClaimStatus`:
	if err = encoder.Encode(obj.CreatorClaimStatus); err != nil {
		return fmt.Errorf("error while marshaling CreatorClaimStatus:%w", err)
	}
	// Serialize `PartnerClaimStatus`:
	if err = encoder.Encode(obj.PartnerClaimStatus); err != nil {
		return fmt.Errorf("error while marshaling PartnerClaimStatus:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj MeteoraDammMigrationMetadata) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding MeteoraDammMigrationMetadata: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *MeteoraDammMigrationMetadata) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_MeteoraDammMigrationMetadata[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_MeteoraDammMigrationMetadata[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `VirtualPool`:
	if err = decoder.Decode(&obj.VirtualPool); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualPool:%w", err)
	}
	// Deserialize `Padding0`:
	if err = decoder.Decode(&obj.Padding0); err != nil {
		return fmt.Errorf("error while unmarshaling Padding0:%w", err)
	}
	// Deserialize `Partner`:
	if err = decoder.Decode(&obj.Partner); err != nil {
		return fmt.Errorf("error while unmarshaling Partner:%w", err)
	}
	// Deserialize `LpMint`:
	if err = decoder.Decode(&obj.LpMint); err != nil {
		return fmt.Errorf("error while unmarshaling LpMint:%w", err)
	}
	// Deserialize `PartnerLockedLp`:
	if err = decoder.Decode(&obj.PartnerLockedLp); err != nil {
		return fmt.Errorf("error while unmarshaling PartnerLockedLp:%w", err)
	}
	// Deserialize `PartnerLp`:
	if err = decoder.Decode(&obj.PartnerLp); err != nil {
		return fmt.Errorf("error while unmarshaling PartnerLp:%w", err)
	}
	// Deserialize `CreatorLockedLp`:
	if err = decoder.Decode(&obj.CreatorLockedLp); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorLockedLp:%w", err)
	}
	// Deserialize `CreatorLp`:
	if err = decoder.Decode(&obj.CreatorLp); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorLp:%w", err)
	}
	// Deserialize `Padding0`:
	if err = decoder.Decode(&obj.Padding0); err != nil {
		return fmt.Errorf("error while unmarshaling Padding0:%w", err)
	}
	// Deserialize `CreatorLockedStatus`:
	if err = decoder.Decode(&obj.CreatorLockedStatus); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorLockedStatus:%w", err)
	}
	// Deserialize `PartnerLockedStatus`:
	if err = decoder.Decode(&obj.PartnerLockedStatus); err != nil {
		return fmt.Errorf("error while unmarshaling PartnerLockedStatus:%w", err)
	}
	// Deserialize `CreatorClaimStatus`:
	if err = decoder.Decode(&obj.CreatorClaimStatus); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorClaimStatus:%w", err)
	}
	// Deserialize `PartnerClaimStatus`:
	if err = decoder.Decode(&obj.PartnerClaimStatus); err != nil {
		return fmt.Errorf("error while unmarshaling PartnerClaimStatus:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *MeteoraDammMigrationMetadata) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling MeteoraDammMigrationMetadata: %w", err)
	}
	return nil
}

func UnmarshalMeteoraDammMigrationMetadata(buf []byte) (*MeteoraDammMigrationMetadata, error) {
	obj := new(MeteoraDammMigrationMetadata)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type MeteoraDammV2Metadata struct {
	// pool
	VirtualPool solanago.PublicKey

	// !!! BE CAREFUL to use tomestone field, previous is pool creator
	Padding0 [32]uint8

	// partner
	Partner solanago.PublicKey

	// Reserve
	Padding [126]uint8
}

func (obj MeteoraDammV2Metadata) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_MeteoraDammV2Metadata[:], false)
	if err != nil {
		return err
	}
	// Serialize `VirtualPool`:
	if err = encoder.Encode(obj.VirtualPool); err != nil {
		return fmt.Errorf("error while marshaling VirtualPool:%w", err)
	}
	// Serialize `Padding0`:
	if err = encoder.Encode(obj.Padding0); err != nil {
		return fmt.Errorf("error while marshaling Padding0:%w", err)
	}
	// Serialize `Partner`:
	if err = encoder.Encode(obj.Partner); err != nil {
		return fmt.Errorf("error while marshaling Partner:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj MeteoraDammV2Metadata) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding MeteoraDammV2Metadata: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *MeteoraDammV2Metadata) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_MeteoraDammV2Metadata[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_MeteoraDammV2Metadata[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `VirtualPool`:
	if err = decoder.Decode(&obj.VirtualPool); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualPool:%w", err)
	}
	// Deserialize `Padding0`:
	if err = decoder.Decode(&obj.Padding0); err != nil {
		return fmt.Errorf("error while unmarshaling Padding0:%w", err)
	}
	// Deserialize `Partner`:
	if err = decoder.Decode(&obj.Partner); err != nil {
		return fmt.Errorf("error while unmarshaling Partner:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *MeteoraDammV2Metadata) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling MeteoraDammV2Metadata: %w", err)
	}
	return nil
}

func UnmarshalMeteoraDammV2Metadata(buf []byte) (*MeteoraDammV2Metadata, error) {
	obj := new(MeteoraDammV2Metadata)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Metadata for a partner.
type PartnerMetadata struct {
	// fee claimer
	FeeClaimer solanago.PublicKey

	// padding for future use
	Padding [6]binary.Uint128

	// Name of partner.
	Name string

	// Website of partner.
	Website string

	// Logo of partner
	Logo string
}

func (obj PartnerMetadata) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_PartnerMetadata[:], false)
	if err != nil {
		return err
	}
	// Serialize `FeeClaimer`:
	if err = encoder.Encode(obj.FeeClaimer); err != nil {
		return fmt.Errorf("error while marshaling FeeClaimer:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	// Serialize `Name`:
	if err = encoder.Encode(obj.Name); err != nil {
		return fmt.Errorf("error while marshaling Name:%w", err)
	}
	// Serialize `Website`:
	if err = encoder.Encode(obj.Website); err != nil {
		return fmt.Errorf("error while marshaling Website:%w", err)
	}
	// Serialize `Logo`:
	if err = encoder.Encode(obj.Logo); err != nil {
		return fmt.Errorf("error while marshaling Logo:%w", err)
	}
	return nil
}

func (obj PartnerMetadata) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding PartnerMetadata: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *PartnerMetadata) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_PartnerMetadata[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_PartnerMetadata[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `FeeClaimer`:
	if err = decoder.Decode(&obj.FeeClaimer); err != nil {
		return fmt.Errorf("error while unmarshaling FeeClaimer:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	// Deserialize `Name`:
	if err = decoder.Decode(&obj.Name); err != nil {
		return fmt.Errorf("error while unmarshaling Name:%w", err)
	}
	// Deserialize `Website`:
	if err = decoder.Decode(&obj.Website); err != nil {
		return fmt.Errorf("error while unmarshaling Website:%w", err)
	}
	// Deserialize `Logo`:
	if err = decoder.Decode(&obj.Logo); err != nil {
		return fmt.Errorf("error while unmarshaling Logo:%w", err)
	}
	return nil
}

func (obj *PartnerMetadata) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling PartnerMetadata: %w", err)
	}
	return nil
}

func UnmarshalPartnerMetadata(buf []byte) (*PartnerMetadata, error) {
	obj := new(PartnerMetadata)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type PoolConfig struct {
	// quote mint
	QuoteMint solanago.PublicKey

	// Address to get the fee
	FeeClaimer solanago.PublicKey

	// Address to receive extra base token after migration, in case token is fixed supply
	LeftoverReceiver solanago.PublicKey

	// Pool fee
	PoolFees PoolFeesConfig

	// Collect fee mode
	CollectFeeMode uint8

	// migration option
	MigrationOption uint8

	// whether mode slot or timestamp
	ActivationType uint8

	// token decimals
	TokenDecimal uint8

	// version
	Version uint8

	// token type of base token
	TokenType uint8

	// quote token flag
	QuoteTokenFlag uint8

	// partner locked lp percentage
	PartnerLockedLpPercentage uint8

	// partner lp percentage
	PartnerLpPercentage uint8

	// creator post migration fee percentage
	CreatorLockedLpPercentage uint8

	// creator lp percentage
	CreatorLpPercentage uint8

	// migration fee option
	MigrationFeeOption uint8

	// flag to indicate whether token is dynamic supply (0) or fixed supply (1)
	FixedTokenSupplyFlag uint8

	// creator trading fee percentage
	CreatorTradingFeePercentage uint8

	// token update authority
	TokenUpdateAuthority uint8

	// migration fee percentage
	MigrationFeePercentage uint8

	// creator migration fee percentage
	CreatorMigrationFeePercentage uint8

	// padding 0
	Padding0 [7]uint8

	// swap base amount
	SwapBaseAmount uint64

	// migration quote threshold (in quote token)
	MigrationQuoteThreshold uint64

	// migration base threshold (in base token)
	MigrationBaseThreshold uint64

	// migration sqrt price
	MigrationSqrtPrice binary.Uint128

	// locked vesting config
	LockedVestingConfig LockedVestingConfig

	// pre migration token supply
	PreMigrationTokenSupply uint64

	// post migration token supply
	PostMigrationTokenSupply uint64

	// migrated pool collect fee mode
	MigratedCollectFeeMode uint8

	// migrated dynamic fee option.
	MigratedDynamicFee uint8

	// migrated pool fee in bps
	MigratedPoolFeeBps uint16

	// padding 1
	Padding1 [12]uint8

	// padding 2
	Padding2 binary.Uint128

	// minimum price
	SqrtStartPrice binary.Uint128

	// curve, only use 20 point firstly, we can extend that latter
	Curve [20]LiquidityDistributionConfig
}

func (obj PoolConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_PoolConfig[:], false)
	if err != nil {
		return err
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
	// Serialize `Version`:
	if err = encoder.Encode(obj.Version); err != nil {
		return fmt.Errorf("error while marshaling Version:%w", err)
	}
	// Serialize `TokenType`:
	if err = encoder.Encode(obj.TokenType); err != nil {
		return fmt.Errorf("error while marshaling TokenType:%w", err)
	}
	// Serialize `QuoteTokenFlag`:
	if err = encoder.Encode(obj.QuoteTokenFlag); err != nil {
		return fmt.Errorf("error while marshaling QuoteTokenFlag:%w", err)
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
	// Serialize `MigrationFeeOption`:
	if err = encoder.Encode(obj.MigrationFeeOption); err != nil {
		return fmt.Errorf("error while marshaling MigrationFeeOption:%w", err)
	}
	// Serialize `FixedTokenSupplyFlag`:
	if err = encoder.Encode(obj.FixedTokenSupplyFlag); err != nil {
		return fmt.Errorf("error while marshaling FixedTokenSupplyFlag:%w", err)
	}
	// Serialize `CreatorTradingFeePercentage`:
	if err = encoder.Encode(obj.CreatorTradingFeePercentage); err != nil {
		return fmt.Errorf("error while marshaling CreatorTradingFeePercentage:%w", err)
	}
	// Serialize `TokenUpdateAuthority`:
	if err = encoder.Encode(obj.TokenUpdateAuthority); err != nil {
		return fmt.Errorf("error while marshaling TokenUpdateAuthority:%w", err)
	}
	// Serialize `MigrationFeePercentage`:
	if err = encoder.Encode(obj.MigrationFeePercentage); err != nil {
		return fmt.Errorf("error while marshaling MigrationFeePercentage:%w", err)
	}
	// Serialize `CreatorMigrationFeePercentage`:
	if err = encoder.Encode(obj.CreatorMigrationFeePercentage); err != nil {
		return fmt.Errorf("error while marshaling CreatorMigrationFeePercentage:%w", err)
	}
	// Serialize `Padding0`:
	if err = encoder.Encode(obj.Padding0); err != nil {
		return fmt.Errorf("error while marshaling Padding0:%w", err)
	}
	// Serialize `SwapBaseAmount`:
	if err = encoder.Encode(obj.SwapBaseAmount); err != nil {
		return fmt.Errorf("error while marshaling SwapBaseAmount:%w", err)
	}
	// Serialize `MigrationQuoteThreshold`:
	if err = encoder.Encode(obj.MigrationQuoteThreshold); err != nil {
		return fmt.Errorf("error while marshaling MigrationQuoteThreshold:%w", err)
	}
	// Serialize `MigrationBaseThreshold`:
	if err = encoder.Encode(obj.MigrationBaseThreshold); err != nil {
		return fmt.Errorf("error while marshaling MigrationBaseThreshold:%w", err)
	}
	// Serialize `MigrationSqrtPrice`:
	if err = encoder.Encode(obj.MigrationSqrtPrice); err != nil {
		return fmt.Errorf("error while marshaling MigrationSqrtPrice:%w", err)
	}
	// Serialize `LockedVestingConfig`:
	if err = encoder.Encode(obj.LockedVestingConfig); err != nil {
		return fmt.Errorf("error while marshaling LockedVestingConfig:%w", err)
	}
	// Serialize `PreMigrationTokenSupply`:
	if err = encoder.Encode(obj.PreMigrationTokenSupply); err != nil {
		return fmt.Errorf("error while marshaling PreMigrationTokenSupply:%w", err)
	}
	// Serialize `PostMigrationTokenSupply`:
	if err = encoder.Encode(obj.PostMigrationTokenSupply); err != nil {
		return fmt.Errorf("error while marshaling PostMigrationTokenSupply:%w", err)
	}
	// Serialize `MigratedCollectFeeMode`:
	if err = encoder.Encode(obj.MigratedCollectFeeMode); err != nil {
		return fmt.Errorf("error while marshaling MigratedCollectFeeMode:%w", err)
	}
	// Serialize `MigratedDynamicFee`:
	if err = encoder.Encode(obj.MigratedDynamicFee); err != nil {
		return fmt.Errorf("error while marshaling MigratedDynamicFee:%w", err)
	}
	// Serialize `MigratedPoolFeeBps`:
	if err = encoder.Encode(obj.MigratedPoolFeeBps); err != nil {
		return fmt.Errorf("error while marshaling MigratedPoolFeeBps:%w", err)
	}
	// Serialize `Padding1`:
	if err = encoder.Encode(obj.Padding1); err != nil {
		return fmt.Errorf("error while marshaling Padding1:%w", err)
	}
	// Serialize `Padding2`:
	if err = encoder.Encode(obj.Padding2); err != nil {
		return fmt.Errorf("error while marshaling Padding2:%w", err)
	}
	// Serialize `SqrtStartPrice`:
	if err = encoder.Encode(obj.SqrtStartPrice); err != nil {
		return fmt.Errorf("error while marshaling SqrtStartPrice:%w", err)
	}
	// Serialize `Curve`:
	if err = encoder.Encode(obj.Curve); err != nil {
		return fmt.Errorf("error while marshaling Curve:%w", err)
	}
	return nil
}

func (obj PoolConfig) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding PoolConfig: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *PoolConfig) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_PoolConfig[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_PoolConfig[:],
				fmt.Sprint(discriminator[:]))
		}
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
	// Deserialize `Version`:
	if err = decoder.Decode(&obj.Version); err != nil {
		return fmt.Errorf("error while unmarshaling Version:%w", err)
	}
	// Deserialize `TokenType`:
	if err = decoder.Decode(&obj.TokenType); err != nil {
		return fmt.Errorf("error while unmarshaling TokenType:%w", err)
	}
	// Deserialize `QuoteTokenFlag`:
	if err = decoder.Decode(&obj.QuoteTokenFlag); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteTokenFlag:%w", err)
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
	// Deserialize `MigrationFeeOption`:
	if err = decoder.Decode(&obj.MigrationFeeOption); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationFeeOption:%w", err)
	}
	// Deserialize `FixedTokenSupplyFlag`:
	if err = decoder.Decode(&obj.FixedTokenSupplyFlag); err != nil {
		return fmt.Errorf("error while unmarshaling FixedTokenSupplyFlag:%w", err)
	}
	// Deserialize `CreatorTradingFeePercentage`:
	if err = decoder.Decode(&obj.CreatorTradingFeePercentage); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorTradingFeePercentage:%w", err)
	}
	// Deserialize `TokenUpdateAuthority`:
	if err = decoder.Decode(&obj.TokenUpdateAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling TokenUpdateAuthority:%w", err)
	}
	// Deserialize `MigrationFeePercentage`:
	if err = decoder.Decode(&obj.MigrationFeePercentage); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationFeePercentage:%w", err)
	}
	// Deserialize `CreatorMigrationFeePercentage`:
	if err = decoder.Decode(&obj.CreatorMigrationFeePercentage); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorMigrationFeePercentage:%w", err)
	}
	// Deserialize `Padding0`:
	if err = decoder.Decode(&obj.Padding0); err != nil {
		return fmt.Errorf("error while unmarshaling Padding0:%w", err)
	}
	// Deserialize `SwapBaseAmount`:
	if err = decoder.Decode(&obj.SwapBaseAmount); err != nil {
		return fmt.Errorf("error while unmarshaling SwapBaseAmount:%w", err)
	}
	// Deserialize `MigrationQuoteThreshold`:
	if err = decoder.Decode(&obj.MigrationQuoteThreshold); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationQuoteThreshold:%w", err)
	}
	// Deserialize `MigrationBaseThreshold`:
	if err = decoder.Decode(&obj.MigrationBaseThreshold); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationBaseThreshold:%w", err)
	}
	// Deserialize `MigrationSqrtPrice`:
	if err = decoder.Decode(&obj.MigrationSqrtPrice); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationSqrtPrice:%w", err)
	}
	// Deserialize `LockedVestingConfig`:
	if err = decoder.Decode(&obj.LockedVestingConfig); err != nil {
		return fmt.Errorf("error while unmarshaling LockedVestingConfig:%w", err)
	}
	// Deserialize `PreMigrationTokenSupply`:
	if err = decoder.Decode(&obj.PreMigrationTokenSupply); err != nil {
		return fmt.Errorf("error while unmarshaling PreMigrationTokenSupply:%w", err)
	}
	// Deserialize `PostMigrationTokenSupply`:
	if err = decoder.Decode(&obj.PostMigrationTokenSupply); err != nil {
		return fmt.Errorf("error while unmarshaling PostMigrationTokenSupply:%w", err)
	}
	// Deserialize `MigratedCollectFeeMode`:
	if err = decoder.Decode(&obj.MigratedCollectFeeMode); err != nil {
		return fmt.Errorf("error while unmarshaling MigratedCollectFeeMode:%w", err)
	}
	// Deserialize `MigratedDynamicFee`:
	if err = decoder.Decode(&obj.MigratedDynamicFee); err != nil {
		return fmt.Errorf("error while unmarshaling MigratedDynamicFee:%w", err)
	}
	// Deserialize `MigratedPoolFeeBps`:
	if err = decoder.Decode(&obj.MigratedPoolFeeBps); err != nil {
		return fmt.Errorf("error while unmarshaling MigratedPoolFeeBps:%w", err)
	}
	// Deserialize `Padding1`:
	if err = decoder.Decode(&obj.Padding1); err != nil {
		return fmt.Errorf("error while unmarshaling Padding1:%w", err)
	}
	// Deserialize `Padding2`:
	if err = decoder.Decode(&obj.Padding2); err != nil {
		return fmt.Errorf("error while unmarshaling Padding2:%w", err)
	}
	// Deserialize `SqrtStartPrice`:
	if err = decoder.Decode(&obj.SqrtStartPrice); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtStartPrice:%w", err)
	}
	// Deserialize `Curve`:
	if err = decoder.Decode(&obj.Curve); err != nil {
		return fmt.Errorf("error while unmarshaling Curve:%w", err)
	}
	return nil
}

func (obj *PoolConfig) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling PoolConfig: %w", err)
	}
	return nil
}

func UnmarshalPoolConfig(buf []byte) (*PoolConfig, error) {
	obj := new(PoolConfig)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type VirtualPool struct {
	// volatility tracker
	VolatilityTracker VolatilityTracker

	// config key
	Config solanago.PublicKey

	// creator
	Creator solanago.PublicKey

	// base mint
	BaseMint solanago.PublicKey

	// base vault
	BaseVault solanago.PublicKey

	// quote vault
	QuoteVault solanago.PublicKey

	// base reserve
	BaseReserve uint64

	// quote reserve
	QuoteReserve uint64

	// protocol base fee
	ProtocolBaseFee uint64

	// protocol quote fee
	ProtocolQuoteFee uint64

	// partner base fee
	PartnerBaseFee uint64

	// trading quote fee
	PartnerQuoteFee uint64

	// current price
	SqrtPrice binary.Uint128

	// Activation point
	ActivationPoint uint64

	// pool type, spl token or token2022
	PoolType uint8

	// is migrated
	IsMigrated uint8

	// is partner withdraw surplus
	IsPartnerWithdrawSurplus uint8

	// is protocol withdraw surplus
	IsProtocolWithdrawSurplus uint8

	// migration progress
	MigrationProgress uint8

	// is withdraw leftover
	IsWithdrawLeftover uint8

	// is creator withdraw surplus
	IsCreatorWithdrawSurplus uint8

	// migration fee withdraw status, first bit is for partner, second bit is for creator
	MigrationFeeWithdrawStatus uint8

	// pool metrics
	Metrics PoolMetrics

	// The time curve is finished
	FinishCurveTimestamp uint64

	// creator base fee
	CreatorBaseFee uint64

	// creator quote fee
	CreatorQuoteFee uint64

	// Padding for further use
	Padding1 [7]uint64
}

func (obj VirtualPool) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_VirtualPool[:], false)
	if err != nil {
		return err
	}
	// Serialize `VolatilityTracker`:
	if err = encoder.Encode(obj.VolatilityTracker); err != nil {
		return fmt.Errorf("error while marshaling VolatilityTracker:%w", err)
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
	// Serialize `BaseVault`:
	if err = encoder.Encode(obj.BaseVault); err != nil {
		return fmt.Errorf("error while marshaling BaseVault:%w", err)
	}
	// Serialize `QuoteVault`:
	if err = encoder.Encode(obj.QuoteVault); err != nil {
		return fmt.Errorf("error while marshaling QuoteVault:%w", err)
	}
	// Serialize `BaseReserve`:
	if err = encoder.Encode(obj.BaseReserve); err != nil {
		return fmt.Errorf("error while marshaling BaseReserve:%w", err)
	}
	// Serialize `QuoteReserve`:
	if err = encoder.Encode(obj.QuoteReserve); err != nil {
		return fmt.Errorf("error while marshaling QuoteReserve:%w", err)
	}
	// Serialize `ProtocolBaseFee`:
	if err = encoder.Encode(obj.ProtocolBaseFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolBaseFee:%w", err)
	}
	// Serialize `ProtocolQuoteFee`:
	if err = encoder.Encode(obj.ProtocolQuoteFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolQuoteFee:%w", err)
	}
	// Serialize `PartnerBaseFee`:
	if err = encoder.Encode(obj.PartnerBaseFee); err != nil {
		return fmt.Errorf("error while marshaling PartnerBaseFee:%w", err)
	}
	// Serialize `PartnerQuoteFee`:
	if err = encoder.Encode(obj.PartnerQuoteFee); err != nil {
		return fmt.Errorf("error while marshaling PartnerQuoteFee:%w", err)
	}
	// Serialize `SqrtPrice`:
	if err = encoder.Encode(obj.SqrtPrice); err != nil {
		return fmt.Errorf("error while marshaling SqrtPrice:%w", err)
	}
	// Serialize `ActivationPoint`:
	if err = encoder.Encode(obj.ActivationPoint); err != nil {
		return fmt.Errorf("error while marshaling ActivationPoint:%w", err)
	}
	// Serialize `PoolType`:
	if err = encoder.Encode(obj.PoolType); err != nil {
		return fmt.Errorf("error while marshaling PoolType:%w", err)
	}
	// Serialize `IsMigrated`:
	if err = encoder.Encode(obj.IsMigrated); err != nil {
		return fmt.Errorf("error while marshaling IsMigrated:%w", err)
	}
	// Serialize `IsPartnerWithdrawSurplus`:
	if err = encoder.Encode(obj.IsPartnerWithdrawSurplus); err != nil {
		return fmt.Errorf("error while marshaling IsPartnerWithdrawSurplus:%w", err)
	}
	// Serialize `IsProtocolWithdrawSurplus`:
	if err = encoder.Encode(obj.IsProtocolWithdrawSurplus); err != nil {
		return fmt.Errorf("error while marshaling IsProtocolWithdrawSurplus:%w", err)
	}
	// Serialize `MigrationProgress`:
	if err = encoder.Encode(obj.MigrationProgress); err != nil {
		return fmt.Errorf("error while marshaling MigrationProgress:%w", err)
	}
	// Serialize `IsWithdrawLeftover`:
	if err = encoder.Encode(obj.IsWithdrawLeftover); err != nil {
		return fmt.Errorf("error while marshaling IsWithdrawLeftover:%w", err)
	}
	// Serialize `IsCreatorWithdrawSurplus`:
	if err = encoder.Encode(obj.IsCreatorWithdrawSurplus); err != nil {
		return fmt.Errorf("error while marshaling IsCreatorWithdrawSurplus:%w", err)
	}
	// Serialize `MigrationFeeWithdrawStatus`:
	if err = encoder.Encode(obj.MigrationFeeWithdrawStatus); err != nil {
		return fmt.Errorf("error while marshaling MigrationFeeWithdrawStatus:%w", err)
	}
	// Serialize `Metrics`:
	if err = encoder.Encode(obj.Metrics); err != nil {
		return fmt.Errorf("error while marshaling Metrics:%w", err)
	}
	// Serialize `FinishCurveTimestamp`:
	if err = encoder.Encode(obj.FinishCurveTimestamp); err != nil {
		return fmt.Errorf("error while marshaling FinishCurveTimestamp:%w", err)
	}
	// Serialize `CreatorBaseFee`:
	if err = encoder.Encode(obj.CreatorBaseFee); err != nil {
		return fmt.Errorf("error while marshaling CreatorBaseFee:%w", err)
	}
	// Serialize `CreatorQuoteFee`:
	if err = encoder.Encode(obj.CreatorQuoteFee); err != nil {
		return fmt.Errorf("error while marshaling CreatorQuoteFee:%w", err)
	}
	// Serialize `Padding1`:
	if err = encoder.Encode(obj.Padding1); err != nil {
		return fmt.Errorf("error while marshaling Padding1:%w", err)
	}
	return nil
}

func (obj VirtualPool) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding VirtualPool: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *VirtualPool) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_VirtualPool[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_VirtualPool[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `VolatilityTracker`:
	if err = decoder.Decode(&obj.VolatilityTracker); err != nil {
		return fmt.Errorf("error while unmarshaling VolatilityTracker:%w", err)
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
	// Deserialize `BaseVault`:
	if err = decoder.Decode(&obj.BaseVault); err != nil {
		return fmt.Errorf("error while unmarshaling BaseVault:%w", err)
	}
	// Deserialize `QuoteVault`:
	if err = decoder.Decode(&obj.QuoteVault); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteVault:%w", err)
	}
	// Deserialize `BaseReserve`:
	if err = decoder.Decode(&obj.BaseReserve); err != nil {
		return fmt.Errorf("error while unmarshaling BaseReserve:%w", err)
	}
	// Deserialize `QuoteReserve`:
	if err = decoder.Decode(&obj.QuoteReserve); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteReserve:%w", err)
	}
	// Deserialize `ProtocolBaseFee`:
	if err = decoder.Decode(&obj.ProtocolBaseFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolBaseFee:%w", err)
	}
	// Deserialize `ProtocolQuoteFee`:
	if err = decoder.Decode(&obj.ProtocolQuoteFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolQuoteFee:%w", err)
	}
	// Deserialize `PartnerBaseFee`:
	if err = decoder.Decode(&obj.PartnerBaseFee); err != nil {
		return fmt.Errorf("error while unmarshaling PartnerBaseFee:%w", err)
	}
	// Deserialize `PartnerQuoteFee`:
	if err = decoder.Decode(&obj.PartnerQuoteFee); err != nil {
		return fmt.Errorf("error while unmarshaling PartnerQuoteFee:%w", err)
	}
	// Deserialize `SqrtPrice`:
	if err = decoder.Decode(&obj.SqrtPrice); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtPrice:%w", err)
	}
	// Deserialize `ActivationPoint`:
	if err = decoder.Decode(&obj.ActivationPoint); err != nil {
		return fmt.Errorf("error while unmarshaling ActivationPoint:%w", err)
	}
	// Deserialize `PoolType`:
	if err = decoder.Decode(&obj.PoolType); err != nil {
		return fmt.Errorf("error while unmarshaling PoolType:%w", err)
	}
	// Deserialize `IsMigrated`:
	if err = decoder.Decode(&obj.IsMigrated); err != nil {
		return fmt.Errorf("error while unmarshaling IsMigrated:%w", err)
	}
	// Deserialize `IsPartnerWithdrawSurplus`:
	if err = decoder.Decode(&obj.IsPartnerWithdrawSurplus); err != nil {
		return fmt.Errorf("error while unmarshaling IsPartnerWithdrawSurplus:%w", err)
	}
	// Deserialize `IsProtocolWithdrawSurplus`:
	if err = decoder.Decode(&obj.IsProtocolWithdrawSurplus); err != nil {
		return fmt.Errorf("error while unmarshaling IsProtocolWithdrawSurplus:%w", err)
	}
	// Deserialize `MigrationProgress`:
	if err = decoder.Decode(&obj.MigrationProgress); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationProgress:%w", err)
	}
	// Deserialize `IsWithdrawLeftover`:
	if err = decoder.Decode(&obj.IsWithdrawLeftover); err != nil {
		return fmt.Errorf("error while unmarshaling IsWithdrawLeftover:%w", err)
	}
	// Deserialize `IsCreatorWithdrawSurplus`:
	if err = decoder.Decode(&obj.IsCreatorWithdrawSurplus); err != nil {
		return fmt.Errorf("error while unmarshaling IsCreatorWithdrawSurplus:%w", err)
	}
	// Deserialize `MigrationFeeWithdrawStatus`:
	if err = decoder.Decode(&obj.MigrationFeeWithdrawStatus); err != nil {
		return fmt.Errorf("error while unmarshaling MigrationFeeWithdrawStatus:%w", err)
	}
	// Deserialize `Metrics`:
	if err = decoder.Decode(&obj.Metrics); err != nil {
		return fmt.Errorf("error while unmarshaling Metrics:%w", err)
	}
	// Deserialize `FinishCurveTimestamp`:
	if err = decoder.Decode(&obj.FinishCurveTimestamp); err != nil {
		return fmt.Errorf("error while unmarshaling FinishCurveTimestamp:%w", err)
	}
	// Deserialize `CreatorBaseFee`:
	if err = decoder.Decode(&obj.CreatorBaseFee); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorBaseFee:%w", err)
	}
	// Deserialize `CreatorQuoteFee`:
	if err = decoder.Decode(&obj.CreatorQuoteFee); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorQuoteFee:%w", err)
	}
	// Deserialize `Padding1`:
	if err = decoder.Decode(&obj.Padding1); err != nil {
		return fmt.Errorf("error while unmarshaling Padding1:%w", err)
	}
	return nil
}

func (obj *VirtualPool) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling VirtualPool: %w", err)
	}
	return nil
}

func UnmarshalVirtualPool(buf []byte) (*VirtualPool, error) {
	obj := new(VirtualPool)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Metadata for a virtual pool.
type VirtualPoolMetadata struct {
	// virtual pool
	VirtualPool solanago.PublicKey

	// padding for future use
	Padding [6]binary.Uint128

	// Name of project.
	Name string

	// Website of project.
	Website string

	// Logo of project
	Logo string
}

func (obj VirtualPoolMetadata) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_VirtualPoolMetadata[:], false)
	if err != nil {
		return err
	}
	// Serialize `VirtualPool`:
	if err = encoder.Encode(obj.VirtualPool); err != nil {
		return fmt.Errorf("error while marshaling VirtualPool:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	// Serialize `Name`:
	if err = encoder.Encode(obj.Name); err != nil {
		return fmt.Errorf("error while marshaling Name:%w", err)
	}
	// Serialize `Website`:
	if err = encoder.Encode(obj.Website); err != nil {
		return fmt.Errorf("error while marshaling Website:%w", err)
	}
	// Serialize `Logo`:
	if err = encoder.Encode(obj.Logo); err != nil {
		return fmt.Errorf("error while marshaling Logo:%w", err)
	}
	return nil
}

func (obj VirtualPoolMetadata) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding VirtualPoolMetadata: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *VirtualPoolMetadata) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_VirtualPoolMetadata[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_VirtualPoolMetadata[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `VirtualPool`:
	if err = decoder.Decode(&obj.VirtualPool); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualPool:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	// Deserialize `Name`:
	if err = decoder.Decode(&obj.Name); err != nil {
		return fmt.Errorf("error while unmarshaling Name:%w", err)
	}
	// Deserialize `Website`:
	if err = decoder.Decode(&obj.Website); err != nil {
		return fmt.Errorf("error while unmarshaling Website:%w", err)
	}
	// Deserialize `Logo`:
	if err = decoder.Decode(&obj.Logo); err != nil {
		return fmt.Errorf("error while unmarshaling Logo:%w", err)
	}
	return nil
}

func (obj *VirtualPoolMetadata) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling VirtualPoolMetadata: %w", err)
	}
	return nil
}

func UnmarshalVirtualPoolMetadata(buf []byte) (*VirtualPoolMetadata, error) {
	obj := new(VirtualPoolMetadata)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
