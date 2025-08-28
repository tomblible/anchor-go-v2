package token_metadata

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type CollectionAuthorityRecord struct {
	Key             Key
	Bump            uint8
	UpdateAuthority *solanago.PublicKey `bin:"optional"`
}

func (obj CollectionAuthorityRecord) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_CollectionAuthorityRecord[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key`:
	if err = encoder.Encode(obj.Key); err != nil {
		return fmt.Errorf("error while marshaling Key:%w", err)
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	// Serialize `UpdateAuthority` (optional):
	{
		if obj.UpdateAuthority == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling UpdateAuthority optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling UpdateAuthority optionality: %w", err)
			}
			if err = encoder.Encode(obj.UpdateAuthority); err != nil {
				return fmt.Errorf("error while marshaling UpdateAuthority: %w", err)
			}
		}
	}
	return nil
}

func (obj CollectionAuthorityRecord) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding CollectionAuthorityRecord: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *CollectionAuthorityRecord) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_CollectionAuthorityRecord[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_CollectionAuthorityRecord[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	if err = decoder.Decode(&obj.Key); err != nil {
		return fmt.Errorf("error while unmarshaling Key:%w", err)
	}
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	// Deserialize `UpdateAuthority` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling UpdateAuthority:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.UpdateAuthority); err != nil {
				return fmt.Errorf("error while unmarshaling UpdateAuthority:%w", err)
			}
		}
	}
	return nil
}

func (obj *CollectionAuthorityRecord) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling CollectionAuthorityRecord: %w", err)
	}
	return nil
}

func UnmarshalCollectionAuthorityRecord(buf []byte) (*CollectionAuthorityRecord, error) {
	obj := new(CollectionAuthorityRecord)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type MetadataDelegateRecord struct {
	Key             Key
	Bump            uint8
	Mint            solanago.PublicKey
	Delegate        solanago.PublicKey
	UpdateAuthority solanago.PublicKey
}

func (obj MetadataDelegateRecord) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_MetadataDelegateRecord[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key`:
	if err = encoder.Encode(obj.Key); err != nil {
		return fmt.Errorf("error while marshaling Key:%w", err)
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `Delegate`:
	if err = encoder.Encode(obj.Delegate); err != nil {
		return fmt.Errorf("error while marshaling Delegate:%w", err)
	}
	// Serialize `UpdateAuthority`:
	if err = encoder.Encode(obj.UpdateAuthority); err != nil {
		return fmt.Errorf("error while marshaling UpdateAuthority:%w", err)
	}
	return nil
}

func (obj MetadataDelegateRecord) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding MetadataDelegateRecord: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *MetadataDelegateRecord) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_MetadataDelegateRecord[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_MetadataDelegateRecord[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	if err = decoder.Decode(&obj.Key); err != nil {
		return fmt.Errorf("error while unmarshaling Key:%w", err)
	}
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `Delegate`:
	if err = decoder.Decode(&obj.Delegate); err != nil {
		return fmt.Errorf("error while unmarshaling Delegate:%w", err)
	}
	// Deserialize `UpdateAuthority`:
	if err = decoder.Decode(&obj.UpdateAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling UpdateAuthority:%w", err)
	}
	return nil
}

func (obj *MetadataDelegateRecord) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling MetadataDelegateRecord: %w", err)
	}
	return nil
}

func UnmarshalMetadataDelegateRecord(buf []byte) (*MetadataDelegateRecord, error) {
	obj := new(MetadataDelegateRecord)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type HolderDelegateRecord struct {
	Key             Key
	Bump            uint8
	Mint            solanago.PublicKey
	Delegate        solanago.PublicKey
	UpdateAuthority solanago.PublicKey
}

func (obj HolderDelegateRecord) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_HolderDelegateRecord[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key`:
	if err = encoder.Encode(obj.Key); err != nil {
		return fmt.Errorf("error while marshaling Key:%w", err)
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `Delegate`:
	if err = encoder.Encode(obj.Delegate); err != nil {
		return fmt.Errorf("error while marshaling Delegate:%w", err)
	}
	// Serialize `UpdateAuthority`:
	if err = encoder.Encode(obj.UpdateAuthority); err != nil {
		return fmt.Errorf("error while marshaling UpdateAuthority:%w", err)
	}
	return nil
}

func (obj HolderDelegateRecord) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding HolderDelegateRecord: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *HolderDelegateRecord) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_HolderDelegateRecord[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_HolderDelegateRecord[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	if err = decoder.Decode(&obj.Key); err != nil {
		return fmt.Errorf("error while unmarshaling Key:%w", err)
	}
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `Delegate`:
	if err = decoder.Decode(&obj.Delegate); err != nil {
		return fmt.Errorf("error while unmarshaling Delegate:%w", err)
	}
	// Deserialize `UpdateAuthority`:
	if err = decoder.Decode(&obj.UpdateAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling UpdateAuthority:%w", err)
	}
	return nil
}

func (obj *HolderDelegateRecord) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling HolderDelegateRecord: %w", err)
	}
	return nil
}

func UnmarshalHolderDelegateRecord(buf []byte) (*HolderDelegateRecord, error) {
	obj := new(HolderDelegateRecord)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type Edition struct {
	Key     Key
	Parent  solanago.PublicKey
	Edition uint64
}

func (obj Edition) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Edition[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key`:
	if err = encoder.Encode(obj.Key); err != nil {
		return fmt.Errorf("error while marshaling Key:%w", err)
	}
	// Serialize `Parent`:
	if err = encoder.Encode(obj.Parent); err != nil {
		return fmt.Errorf("error while marshaling Parent:%w", err)
	}
	// Serialize `Edition`:
	if err = encoder.Encode(obj.Edition); err != nil {
		return fmt.Errorf("error while marshaling Edition:%w", err)
	}
	return nil
}

func (obj Edition) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Edition: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Edition) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Edition[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Edition[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	if err = decoder.Decode(&obj.Key); err != nil {
		return fmt.Errorf("error while unmarshaling Key:%w", err)
	}
	// Deserialize `Parent`:
	if err = decoder.Decode(&obj.Parent); err != nil {
		return fmt.Errorf("error while unmarshaling Parent:%w", err)
	}
	// Deserialize `Edition`:
	if err = decoder.Decode(&obj.Edition); err != nil {
		return fmt.Errorf("error while unmarshaling Edition:%w", err)
	}
	return nil
}

func (obj *Edition) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Edition: %w", err)
	}
	return nil
}

func UnmarshalEdition(buf []byte) (*Edition, error) {
	obj := new(Edition)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EditionMarker struct {
	Key    Key
	Ledger [31]uint8
}

func (obj EditionMarker) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_EditionMarker[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key`:
	if err = encoder.Encode(obj.Key); err != nil {
		return fmt.Errorf("error while marshaling Key:%w", err)
	}
	// Serialize `Ledger`:
	if err = encoder.Encode(obj.Ledger); err != nil {
		return fmt.Errorf("error while marshaling Ledger:%w", err)
	}
	return nil
}

func (obj EditionMarker) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EditionMarker: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EditionMarker) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_EditionMarker[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_EditionMarker[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	if err = decoder.Decode(&obj.Key); err != nil {
		return fmt.Errorf("error while unmarshaling Key:%w", err)
	}
	// Deserialize `Ledger`:
	if err = decoder.Decode(&obj.Ledger); err != nil {
		return fmt.Errorf("error while unmarshaling Ledger:%w", err)
	}
	return nil
}

func (obj *EditionMarker) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EditionMarker: %w", err)
	}
	return nil
}

func UnmarshalEditionMarker(buf []byte) (*EditionMarker, error) {
	obj := new(EditionMarker)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type EditionMarkerV2 struct {
	Key    Key
	Ledger []byte
}

func (obj EditionMarkerV2) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_EditionMarkerV2[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key`:
	if err = encoder.Encode(obj.Key); err != nil {
		return fmt.Errorf("error while marshaling Key:%w", err)
	}
	// Serialize `Ledger`:
	if err = encoder.Encode(obj.Ledger); err != nil {
		return fmt.Errorf("error while marshaling Ledger:%w", err)
	}
	return nil
}

func (obj EditionMarkerV2) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding EditionMarkerV2: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *EditionMarkerV2) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_EditionMarkerV2[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_EditionMarkerV2[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	if err = decoder.Decode(&obj.Key); err != nil {
		return fmt.Errorf("error while unmarshaling Key:%w", err)
	}
	// Deserialize `Ledger`:
	if err = decoder.Decode(&obj.Ledger); err != nil {
		return fmt.Errorf("error while unmarshaling Ledger:%w", err)
	}
	return nil
}

func (obj *EditionMarkerV2) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling EditionMarkerV2: %w", err)
	}
	return nil
}

func UnmarshalEditionMarkerV2(buf []byte) (*EditionMarkerV2, error) {
	obj := new(EditionMarkerV2)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type TokenOwnedEscrow struct {
	Key       Key
	BaseToken solanago.PublicKey
	Authority EscrowAuthority
	Bump      uint8
}

func (obj TokenOwnedEscrow) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_TokenOwnedEscrow[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key`:
	if err = encoder.Encode(obj.Key); err != nil {
		return fmt.Errorf("error while marshaling Key:%w", err)
	}
	// Serialize `BaseToken`:
	if err = encoder.Encode(obj.BaseToken); err != nil {
		return fmt.Errorf("error while marshaling BaseToken:%w", err)
	}
	// Serialize `Authority`:
	{
		if err = EncodeEscrowAuthority(encoder, obj.Authority); err != nil {
			return fmt.Errorf("error while marshalingAuthority:%w", err)
		}
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	return nil
}

func (obj TokenOwnedEscrow) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding TokenOwnedEscrow: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *TokenOwnedEscrow) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_TokenOwnedEscrow[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_TokenOwnedEscrow[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	if err = decoder.Decode(&obj.Key); err != nil {
		return fmt.Errorf("error while unmarshaling Key:%w", err)
	}
	// Deserialize `BaseToken`:
	if err = decoder.Decode(&obj.BaseToken); err != nil {
		return fmt.Errorf("error while unmarshaling BaseToken:%w", err)
	}
	// Deserialize `Authority`:
	{
		var err error
		obj.Authority, err = DecodeEscrowAuthority(decoder)
		if err != nil {
			return err
		}
	}
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	return nil
}

func (obj *TokenOwnedEscrow) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling TokenOwnedEscrow: %w", err)
	}
	return nil
}

func UnmarshalTokenOwnedEscrow(buf []byte) (*TokenOwnedEscrow, error) {
	obj := new(TokenOwnedEscrow)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type MasterEditionV2 struct {
	Key       Key
	Supply    uint64
	MaxSupply *uint64 `bin:"optional"`
}

func (obj MasterEditionV2) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_MasterEditionV2[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key`:
	if err = encoder.Encode(obj.Key); err != nil {
		return fmt.Errorf("error while marshaling Key:%w", err)
	}
	// Serialize `Supply`:
	if err = encoder.Encode(obj.Supply); err != nil {
		return fmt.Errorf("error while marshaling Supply:%w", err)
	}
	// Serialize `MaxSupply` (optional):
	{
		if obj.MaxSupply == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling MaxSupply optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling MaxSupply optionality: %w", err)
			}
			if err = encoder.Encode(obj.MaxSupply); err != nil {
				return fmt.Errorf("error while marshaling MaxSupply: %w", err)
			}
		}
	}
	return nil
}

func (obj MasterEditionV2) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding MasterEditionV2: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *MasterEditionV2) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_MasterEditionV2[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_MasterEditionV2[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	if err = decoder.Decode(&obj.Key); err != nil {
		return fmt.Errorf("error while unmarshaling Key:%w", err)
	}
	// Deserialize `Supply`:
	if err = decoder.Decode(&obj.Supply); err != nil {
		return fmt.Errorf("error while unmarshaling Supply:%w", err)
	}
	// Deserialize `MaxSupply` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling MaxSupply:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.MaxSupply); err != nil {
				return fmt.Errorf("error while unmarshaling MaxSupply:%w", err)
			}
		}
	}
	return nil
}

func (obj *MasterEditionV2) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling MasterEditionV2: %w", err)
	}
	return nil
}

func UnmarshalMasterEditionV2(buf []byte) (*MasterEditionV2, error) {
	obj := new(MasterEditionV2)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type MasterEditionV1 struct {
	Key                              Key
	Supply                           uint64
	MaxSupply                        *uint64 `bin:"optional"`
	PrintingMint                     solanago.PublicKey
	OneTimePrintingAuthorizationMint solanago.PublicKey
}

func (obj MasterEditionV1) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_MasterEditionV1[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key`:
	if err = encoder.Encode(obj.Key); err != nil {
		return fmt.Errorf("error while marshaling Key:%w", err)
	}
	// Serialize `Supply`:
	if err = encoder.Encode(obj.Supply); err != nil {
		return fmt.Errorf("error while marshaling Supply:%w", err)
	}
	// Serialize `MaxSupply` (optional):
	{
		if obj.MaxSupply == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling MaxSupply optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling MaxSupply optionality: %w", err)
			}
			if err = encoder.Encode(obj.MaxSupply); err != nil {
				return fmt.Errorf("error while marshaling MaxSupply: %w", err)
			}
		}
	}
	// Serialize `PrintingMint`:
	if err = encoder.Encode(obj.PrintingMint); err != nil {
		return fmt.Errorf("error while marshaling PrintingMint:%w", err)
	}
	// Serialize `OneTimePrintingAuthorizationMint`:
	if err = encoder.Encode(obj.OneTimePrintingAuthorizationMint); err != nil {
		return fmt.Errorf("error while marshaling OneTimePrintingAuthorizationMint:%w", err)
	}
	return nil
}

func (obj MasterEditionV1) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding MasterEditionV1: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *MasterEditionV1) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_MasterEditionV1[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_MasterEditionV1[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	if err = decoder.Decode(&obj.Key); err != nil {
		return fmt.Errorf("error while unmarshaling Key:%w", err)
	}
	// Deserialize `Supply`:
	if err = decoder.Decode(&obj.Supply); err != nil {
		return fmt.Errorf("error while unmarshaling Supply:%w", err)
	}
	// Deserialize `MaxSupply` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling MaxSupply:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.MaxSupply); err != nil {
				return fmt.Errorf("error while unmarshaling MaxSupply:%w", err)
			}
		}
	}
	// Deserialize `PrintingMint`:
	if err = decoder.Decode(&obj.PrintingMint); err != nil {
		return fmt.Errorf("error while unmarshaling PrintingMint:%w", err)
	}
	// Deserialize `OneTimePrintingAuthorizationMint`:
	if err = decoder.Decode(&obj.OneTimePrintingAuthorizationMint); err != nil {
		return fmt.Errorf("error while unmarshaling OneTimePrintingAuthorizationMint:%w", err)
	}
	return nil
}

func (obj *MasterEditionV1) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling MasterEditionV1: %w", err)
	}
	return nil
}

func UnmarshalMasterEditionV1(buf []byte) (*MasterEditionV1, error) {
	obj := new(MasterEditionV1)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type Metadata struct {
	Key                 Key
	UpdateAuthority     solanago.PublicKey
	Mint                solanago.PublicKey
	Data                Data
	PrimarySaleHappened bool
	IsMutable           bool
	EditionNonce        *uint8              `bin:"optional"`
	TokenStandard       *TokenStandard      `bin:"optional"`
	Collection          *Collection         `bin:"optional"`
	Uses                *Uses               `bin:"optional"`
	CollectionDetails   *CollectionDetails  `bin:"optional"`
	ProgrammableConfig  *ProgrammableConfig `bin:"optional"`
}

func (obj Metadata) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Metadata[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key`:
	if err = encoder.Encode(obj.Key); err != nil {
		return fmt.Errorf("error while marshaling Key:%w", err)
	}
	// Serialize `UpdateAuthority`:
	if err = encoder.Encode(obj.UpdateAuthority); err != nil {
		return fmt.Errorf("error while marshaling UpdateAuthority:%w", err)
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `Data`:
	if err = encoder.Encode(obj.Data); err != nil {
		return fmt.Errorf("error while marshaling Data:%w", err)
	}
	// Serialize `PrimarySaleHappened`:
	if err = encoder.Encode(obj.PrimarySaleHappened); err != nil {
		return fmt.Errorf("error while marshaling PrimarySaleHappened:%w", err)
	}
	// Serialize `IsMutable`:
	if err = encoder.Encode(obj.IsMutable); err != nil {
		return fmt.Errorf("error while marshaling IsMutable:%w", err)
	}
	// Serialize `EditionNonce` (optional):
	{
		if obj.EditionNonce == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling EditionNonce optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling EditionNonce optionality: %w", err)
			}
			if err = encoder.Encode(obj.EditionNonce); err != nil {
				return fmt.Errorf("error while marshaling EditionNonce: %w", err)
			}
		}
	}
	// Serialize `TokenStandard` (optional):
	{
		if obj.TokenStandard == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling TokenStandard optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling TokenStandard optionality: %w", err)
			}
			if err = encoder.Encode(obj.TokenStandard); err != nil {
				return fmt.Errorf("error while marshaling TokenStandard: %w", err)
			}
		}
	}
	// Serialize `Collection` (optional):
	{
		if obj.Collection == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling Collection optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling Collection optionality: %w", err)
			}
			if err = encoder.Encode(obj.Collection); err != nil {
				return fmt.Errorf("error while marshaling Collection: %w", err)
			}
		}
	}
	// Serialize `Uses` (optional):
	{
		if obj.Uses == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling Uses optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling Uses optionality: %w", err)
			}
			if err = encoder.Encode(obj.Uses); err != nil {
				return fmt.Errorf("error while marshaling Uses: %w", err)
			}
		}
	}
	// Serialize `CollectionDetails` (optional):
	{
		if obj.CollectionDetails == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling CollectionDetails optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling CollectionDetails optionality: %w", err)
			}
			if err = encoder.Encode(obj.CollectionDetails); err != nil {
				return fmt.Errorf("error while marshaling CollectionDetails: %w", err)
			}
		}
	}
	// Serialize `ProgrammableConfig` (optional):
	{
		if obj.ProgrammableConfig == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling ProgrammableConfig optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling ProgrammableConfig optionality: %w", err)
			}
			if err = encoder.Encode(obj.ProgrammableConfig); err != nil {
				return fmt.Errorf("error while marshaling ProgrammableConfig: %w", err)
			}
		}
	}
	return nil
}

func (obj Metadata) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Metadata: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Metadata) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Metadata[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Metadata[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	if err = decoder.Decode(&obj.Key); err != nil {
		return fmt.Errorf("error while unmarshaling Key:%w", err)
	}
	// Deserialize `UpdateAuthority`:
	if err = decoder.Decode(&obj.UpdateAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling UpdateAuthority:%w", err)
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `Data`:
	if err = decoder.Decode(&obj.Data); err != nil {
		return fmt.Errorf("error while unmarshaling Data:%w", err)
	}
	// Deserialize `PrimarySaleHappened`:
	if err = decoder.Decode(&obj.PrimarySaleHappened); err != nil {
		return fmt.Errorf("error while unmarshaling PrimarySaleHappened:%w", err)
	}
	// Deserialize `IsMutable`:
	if err = decoder.Decode(&obj.IsMutable); err != nil {
		return fmt.Errorf("error while unmarshaling IsMutable:%w", err)
	}
	// Deserialize `EditionNonce` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling EditionNonce:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.EditionNonce); err != nil {
				return fmt.Errorf("error while unmarshaling EditionNonce:%w", err)
			}
		}
	}
	// Deserialize `TokenStandard` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling TokenStandard:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.TokenStandard); err != nil {
				return fmt.Errorf("error while unmarshaling TokenStandard:%w", err)
			}
		}
	}
	// Deserialize `Collection` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling Collection:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.Collection); err != nil {
				return fmt.Errorf("error while unmarshaling Collection:%w", err)
			}
		}
	}
	// Deserialize `Uses` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling Uses:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.Uses); err != nil {
				return fmt.Errorf("error while unmarshaling Uses:%w", err)
			}
		}
	}
	// Deserialize `CollectionDetails` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling CollectionDetails:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.CollectionDetails); err != nil {
				return fmt.Errorf("error while unmarshaling CollectionDetails:%w", err)
			}
		}
	}
	// Deserialize `ProgrammableConfig` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling ProgrammableConfig:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.ProgrammableConfig); err != nil {
				return fmt.Errorf("error while unmarshaling ProgrammableConfig:%w", err)
			}
		}
	}
	return nil
}

func (obj *Metadata) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Metadata: %w", err)
	}
	return nil
}

func UnmarshalMetadata(buf []byte) (*Metadata, error) {
	obj := new(Metadata)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type TokenRecord struct {
	Key             Key
	Bump            uint8
	State           TokenState
	RuleSetRevision *uint64             `bin:"optional"`
	Delegate        *solanago.PublicKey `bin:"optional"`
	DelegateRole    *TokenDelegateRole  `bin:"optional"`
	LockedTransfer  *solanago.PublicKey `bin:"optional"`
}

func (obj TokenRecord) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_TokenRecord[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key`:
	if err = encoder.Encode(obj.Key); err != nil {
		return fmt.Errorf("error while marshaling Key:%w", err)
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	// Serialize `State`:
	if err = encoder.Encode(obj.State); err != nil {
		return fmt.Errorf("error while marshaling State:%w", err)
	}
	// Serialize `RuleSetRevision` (optional):
	{
		if obj.RuleSetRevision == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling RuleSetRevision optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling RuleSetRevision optionality: %w", err)
			}
			if err = encoder.Encode(obj.RuleSetRevision); err != nil {
				return fmt.Errorf("error while marshaling RuleSetRevision: %w", err)
			}
		}
	}
	// Serialize `Delegate` (optional):
	{
		if obj.Delegate == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling Delegate optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling Delegate optionality: %w", err)
			}
			if err = encoder.Encode(obj.Delegate); err != nil {
				return fmt.Errorf("error while marshaling Delegate: %w", err)
			}
		}
	}
	// Serialize `DelegateRole` (optional):
	{
		if obj.DelegateRole == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling DelegateRole optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling DelegateRole optionality: %w", err)
			}
			if err = encoder.Encode(obj.DelegateRole); err != nil {
				return fmt.Errorf("error while marshaling DelegateRole: %w", err)
			}
		}
	}
	// Serialize `LockedTransfer` (optional):
	{
		if obj.LockedTransfer == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling LockedTransfer optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling LockedTransfer optionality: %w", err)
			}
			if err = encoder.Encode(obj.LockedTransfer); err != nil {
				return fmt.Errorf("error while marshaling LockedTransfer: %w", err)
			}
		}
	}
	return nil
}

func (obj TokenRecord) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding TokenRecord: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *TokenRecord) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_TokenRecord[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_TokenRecord[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	if err = decoder.Decode(&obj.Key); err != nil {
		return fmt.Errorf("error while unmarshaling Key:%w", err)
	}
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	// Deserialize `State`:
	if err = decoder.Decode(&obj.State); err != nil {
		return fmt.Errorf("error while unmarshaling State:%w", err)
	}
	// Deserialize `RuleSetRevision` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling RuleSetRevision:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.RuleSetRevision); err != nil {
				return fmt.Errorf("error while unmarshaling RuleSetRevision:%w", err)
			}
		}
	}
	// Deserialize `Delegate` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling Delegate:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.Delegate); err != nil {
				return fmt.Errorf("error while unmarshaling Delegate:%w", err)
			}
		}
	}
	// Deserialize `DelegateRole` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling DelegateRole:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.DelegateRole); err != nil {
				return fmt.Errorf("error while unmarshaling DelegateRole:%w", err)
			}
		}
	}
	// Deserialize `LockedTransfer` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling LockedTransfer:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.LockedTransfer); err != nil {
				return fmt.Errorf("error while unmarshaling LockedTransfer:%w", err)
			}
		}
	}
	return nil
}

func (obj *TokenRecord) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling TokenRecord: %w", err)
	}
	return nil
}

func UnmarshalTokenRecord(buf []byte) (*TokenRecord, error) {
	obj := new(TokenRecord)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type ReservationListV2 struct {
	Key                     Key
	MasterEdition           solanago.PublicKey
	SupplySnapshot          *uint64 `bin:"optional"`
	Reservations            []Reservation
	TotalReservationSpots   uint64
	CurrentReservationSpots uint64
}

func (obj ReservationListV2) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_ReservationListV2[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key`:
	if err = encoder.Encode(obj.Key); err != nil {
		return fmt.Errorf("error while marshaling Key:%w", err)
	}
	// Serialize `MasterEdition`:
	if err = encoder.Encode(obj.MasterEdition); err != nil {
		return fmt.Errorf("error while marshaling MasterEdition:%w", err)
	}
	// Serialize `SupplySnapshot` (optional):
	{
		if obj.SupplySnapshot == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling SupplySnapshot optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling SupplySnapshot optionality: %w", err)
			}
			if err = encoder.Encode(obj.SupplySnapshot); err != nil {
				return fmt.Errorf("error while marshaling SupplySnapshot: %w", err)
			}
		}
	}
	// Serialize `Reservations`:
	if err = encoder.Encode(obj.Reservations); err != nil {
		return fmt.Errorf("error while marshaling Reservations:%w", err)
	}
	// Serialize `TotalReservationSpots`:
	if err = encoder.Encode(obj.TotalReservationSpots); err != nil {
		return fmt.Errorf("error while marshaling TotalReservationSpots:%w", err)
	}
	// Serialize `CurrentReservationSpots`:
	if err = encoder.Encode(obj.CurrentReservationSpots); err != nil {
		return fmt.Errorf("error while marshaling CurrentReservationSpots:%w", err)
	}
	return nil
}

func (obj ReservationListV2) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding ReservationListV2: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *ReservationListV2) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_ReservationListV2[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_ReservationListV2[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	if err = decoder.Decode(&obj.Key); err != nil {
		return fmt.Errorf("error while unmarshaling Key:%w", err)
	}
	// Deserialize `MasterEdition`:
	if err = decoder.Decode(&obj.MasterEdition); err != nil {
		return fmt.Errorf("error while unmarshaling MasterEdition:%w", err)
	}
	// Deserialize `SupplySnapshot` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling SupplySnapshot:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.SupplySnapshot); err != nil {
				return fmt.Errorf("error while unmarshaling SupplySnapshot:%w", err)
			}
		}
	}
	// Deserialize `Reservations`:
	if err = decoder.Decode(&obj.Reservations); err != nil {
		return fmt.Errorf("error while unmarshaling Reservations:%w", err)
	}
	// Deserialize `TotalReservationSpots`:
	if err = decoder.Decode(&obj.TotalReservationSpots); err != nil {
		return fmt.Errorf("error while unmarshaling TotalReservationSpots:%w", err)
	}
	// Deserialize `CurrentReservationSpots`:
	if err = decoder.Decode(&obj.CurrentReservationSpots); err != nil {
		return fmt.Errorf("error while unmarshaling CurrentReservationSpots:%w", err)
	}
	return nil
}

func (obj *ReservationListV2) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling ReservationListV2: %w", err)
	}
	return nil
}

func UnmarshalReservationListV2(buf []byte) (*ReservationListV2, error) {
	obj := new(ReservationListV2)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type ReservationListV1 struct {
	Key            Key
	MasterEdition  solanago.PublicKey
	SupplySnapshot *uint64 `bin:"optional"`
	Reservations   []ReservationV1
}

func (obj ReservationListV1) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_ReservationListV1[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key`:
	if err = encoder.Encode(obj.Key); err != nil {
		return fmt.Errorf("error while marshaling Key:%w", err)
	}
	// Serialize `MasterEdition`:
	if err = encoder.Encode(obj.MasterEdition); err != nil {
		return fmt.Errorf("error while marshaling MasterEdition:%w", err)
	}
	// Serialize `SupplySnapshot` (optional):
	{
		if obj.SupplySnapshot == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling SupplySnapshot optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling SupplySnapshot optionality: %w", err)
			}
			if err = encoder.Encode(obj.SupplySnapshot); err != nil {
				return fmt.Errorf("error while marshaling SupplySnapshot: %w", err)
			}
		}
	}
	// Serialize `Reservations`:
	if err = encoder.Encode(obj.Reservations); err != nil {
		return fmt.Errorf("error while marshaling Reservations:%w", err)
	}
	return nil
}

func (obj ReservationListV1) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding ReservationListV1: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *ReservationListV1) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_ReservationListV1[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_ReservationListV1[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	if err = decoder.Decode(&obj.Key); err != nil {
		return fmt.Errorf("error while unmarshaling Key:%w", err)
	}
	// Deserialize `MasterEdition`:
	if err = decoder.Decode(&obj.MasterEdition); err != nil {
		return fmt.Errorf("error while unmarshaling MasterEdition:%w", err)
	}
	// Deserialize `SupplySnapshot` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling SupplySnapshot:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.SupplySnapshot); err != nil {
				return fmt.Errorf("error while unmarshaling SupplySnapshot:%w", err)
			}
		}
	}
	// Deserialize `Reservations`:
	if err = decoder.Decode(&obj.Reservations); err != nil {
		return fmt.Errorf("error while unmarshaling Reservations:%w", err)
	}
	return nil
}

func (obj *ReservationListV1) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling ReservationListV1: %w", err)
	}
	return nil
}

func UnmarshalReservationListV1(buf []byte) (*ReservationListV1, error) {
	obj := new(ReservationListV1)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type UseAuthorityRecord struct {
	Key         Key
	AllowedUses uint64
	Bump        uint8
}

func (obj UseAuthorityRecord) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_UseAuthorityRecord[:], false)
	if err != nil {
		return err
	}
	// Serialize `Key`:
	if err = encoder.Encode(obj.Key); err != nil {
		return fmt.Errorf("error while marshaling Key:%w", err)
	}
	// Serialize `AllowedUses`:
	if err = encoder.Encode(obj.AllowedUses); err != nil {
		return fmt.Errorf("error while marshaling AllowedUses:%w", err)
	}
	// Serialize `Bump`:
	if err = encoder.Encode(obj.Bump); err != nil {
		return fmt.Errorf("error while marshaling Bump:%w", err)
	}
	return nil
}

func (obj UseAuthorityRecord) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding UseAuthorityRecord: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *UseAuthorityRecord) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_UseAuthorityRecord[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_UseAuthorityRecord[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Key`:
	if err = decoder.Decode(&obj.Key); err != nil {
		return fmt.Errorf("error while unmarshaling Key:%w", err)
	}
	// Deserialize `AllowedUses`:
	if err = decoder.Decode(&obj.AllowedUses); err != nil {
		return fmt.Errorf("error while unmarshaling AllowedUses:%w", err)
	}
	// Deserialize `Bump`:
	if err = decoder.Decode(&obj.Bump); err != nil {
		return fmt.Errorf("error while unmarshaling Bump:%w", err)
	}
	return nil
}

func (obj *UseAuthorityRecord) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling UseAuthorityRecord: %w", err)
	}
	return nil
}

func UnmarshalUseAuthorityRecord(buf []byte) (*UseAuthorityRecord, error) {
	obj := new(UseAuthorityRecord)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
