package token2022_program

import (
	"bytes"
	"fmt"

	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type Multisig struct {
	// Number of signers required
	M uint8

	// Number of valid signers
	N uint8

	// Is `true` if this structure has been initialized
	IsInitialized bool

	// Signer public keys
	Signers [11]solanago.PublicKey
}

func (obj Multisig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Serialize `M`:
	if err = encoder.Encode(obj.M); err != nil {
		return fmt.Errorf("error while marshaling M:%w", err)
	}
	// Serialize `N`:
	if err = encoder.Encode(obj.N); err != nil {
		return fmt.Errorf("error while marshaling N:%w", err)
	}
	// Serialize `IsInitialized`:
	if err = encoder.Encode(obj.IsInitialized); err != nil {
		return fmt.Errorf("error while marshaling IsInitialized:%w", err)
	}
	// Serialize `Signers`:
	if err = encoder.Encode(obj.Signers); err != nil {
		return fmt.Errorf("error while marshaling Signers:%w", err)
	}
	return nil
}

func (obj Multisig) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Multisig: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Multisig) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Deserialize `M`:
	if err = decoder.Decode(&obj.M); err != nil {
		return fmt.Errorf("error while unmarshaling M:%w", err)
	}
	// Deserialize `N`:
	if err = decoder.Decode(&obj.N); err != nil {
		return fmt.Errorf("error while unmarshaling N:%w", err)
	}
	// Deserialize `IsInitialized`:
	if err = decoder.Decode(&obj.IsInitialized); err != nil {
		return fmt.Errorf("error while unmarshaling IsInitialized:%w", err)
	}
	// Deserialize `Signers`:
	if err = decoder.Decode(&obj.Signers); err != nil {
		return fmt.Errorf("error while unmarshaling Signers:%w", err)
	}
	return nil
}

func (obj *Multisig) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Multisig: %w", err)
	}
	return nil
}

func UnmarshalMultisig(buf []byte) (*Multisig, error) {
	obj := new(Multisig)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type Mint struct {
	// Optional authority used to mint new tokens. The mint authority may only be provided during
	// mint creation. If no mint authority is present then the mint has a fixed supply and no
	// further tokens may be minted.
	MintAuthority *solanago.PublicKey `bin:"optional"`

	// Total supply of tokens.
	Supply uint64

	// Number of base 10 digits to the right of the decimal place.
	Decimals uint8

	// Is `true` if this structure has been initialized
	IsInitialized bool

	// Optional authority to freeze token accounts.
	FreezeAuthority *solanago.PublicKey `bin:"optional"`

	// extendes
	TransferFeeConfig        *Extension_TransferFeeConfig        `bin:"-"`
	MintCloseAuthority       *Extension_MintCloseAuthority       `bin:"-"`
	ConfidentialTransferMint *Extension_ConfidentialTransferMint `bin:"-"`
	DefaultAccountState      *Extension_DefaultAccountState      `bin:"-"`
	NonTransferable          *Extension_NonTransferable          `bin:"-"`
	InterestBearingConfig    *Extension_InterestBearingConfig    `bin:"-"`
	PermanentDelegate        *Extension_PermanentDelegate        `bin:"-"`
	TransferHook             *Extension_TransferHook             `bin:"-"`
	MetadataPointer          *Extension_MetadataPointer          `bin:"-"`
	TokenMetadata            *Extension_TokenMetadata            `bin:"-"`
	GroupPointer             *Extension_GroupPointer             `bin:"-"`
	GroupMemberPointer       *Extension_GroupMemberPointer       `bin:"-"`
	TokenGroup               *Extension_TokenGroup               `bin:"-"`
	TokenGroupMember         *Extension_TokenGroupMember         `bin:"-"`
	ScaledUiAmountConfig     *Extension_ScaledUiAmountConfig     `bin:"-"`
	PausableConfig           *Extension_PausableConfig           `bin:"-"`
}

func (obj Mint) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Serialize `MintAuthority` (optional):
	{
		if obj.MintAuthority == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling MintAuthority optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling MintAuthority optionality: %w", err)
			}
			if err = encoder.Encode(obj.MintAuthority); err != nil {
				return fmt.Errorf("error while marshaling MintAuthority: %w", err)
			}
		}
	}
	// Serialize `Supply`:
	if err = encoder.Encode(obj.Supply); err != nil {
		return fmt.Errorf("error while marshaling Supply:%w", err)
	}
	// Serialize `Decimals`:
	if err = encoder.Encode(obj.Decimals); err != nil {
		return fmt.Errorf("error while marshaling Decimals:%w", err)
	}
	// Serialize `IsInitialized`:
	if err = encoder.Encode(obj.IsInitialized); err != nil {
		return fmt.Errorf("error while marshaling IsInitialized:%w", err)
	}
	// Serialize `FreezeAuthority` (optional):
	{
		if obj.FreezeAuthority == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling FreezeAuthority optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling FreezeAuthority optionality: %w", err)
			}
			if err = encoder.Encode(obj.FreezeAuthority); err != nil {
				return fmt.Errorf("error while marshaling FreezeAuthority: %w", err)
			}
		}
	}
	// Serialize `TransferFeeConfig` (optional):
	if obj.TransferFeeConfig == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling TransferFeeConfig optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling TransferFeeConfig optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.TransferFeeConfig); err != nil {
			return fmt.Errorf("error while marshaling TransferFeeConfig: %w", err)
		}
	}

	// Serialize `MintCloseAuthority` (optional):
	if obj.MintCloseAuthority == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling MintCloseAuthority optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling MintCloseAuthority optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.MintCloseAuthority); err != nil {
			return fmt.Errorf("error while marshaling MintCloseAuthority: %w", err)
		}
	}

	// Serialize `ConfidentialTransferMint`
	if obj.ConfidentialTransferMint == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling ConfidentialTransferMint optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling ConfidentialTransferMint optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.ConfidentialTransferMint); err != nil {
			return fmt.Errorf("error while marshaling ConfidentialTransferMint: %w", err)
		}
	}

	// Serialize `DefaultAccountState`
	if obj.DefaultAccountState == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling DefaultAccountState optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling DefaultAccountState optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.DefaultAccountState); err != nil {
			return fmt.Errorf("error while marshaling DefaultAccountState: %w", err)
		}
	}

	// Serialize `NonTransferable`
	if obj.NonTransferable == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling NonTransferable optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling NonTransferable optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.NonTransferable); err != nil {
			return fmt.Errorf("error while marshaling NonTransferable: %w", err)
		}
	}

	// Serialize `InterestBearingConfig`
	if obj.InterestBearingConfig == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling InterestBearingConfig optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling InterestBearingConfig optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.InterestBearingConfig); err != nil {
			return fmt.Errorf("error while marshaling InterestBearingConfig: %w", err)
		}
	}

	// Serialize `PermanentDelegate`
	if obj.PermanentDelegate == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling PermanentDelegate optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling PermanentDelegate optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.PermanentDelegate); err != nil {
			return fmt.Errorf("error while marshaling PermanentDelegate: %w", err)
		}
	}

	// Serialize `TransferHook`
	if obj.TransferHook == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling TransferHook optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling TransferHook optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.TransferHook); err != nil {
			return fmt.Errorf("error while marshaling TransferHook: %w", err)
		}
	}

	// Serialize `MetadataPointer`
	if obj.MetadataPointer == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling MetadataPointer optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling MetadataPointer optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.MetadataPointer); err != nil {
			return fmt.Errorf("error while marshaling MetadataPointer: %w", err)
		}
	}

	// Serialize `TokenMetadata`
	if obj.TokenMetadata == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling TokenMetadata optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling TokenMetadata optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.TokenMetadata); err != nil {
			return fmt.Errorf("error while marshaling TokenMetadata: %w", err)
		}
	}

	// Serialize `GroupPointer`
	if obj.GroupPointer == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling GroupPointer optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling GroupPointer optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.GroupPointer); err != nil {
			return fmt.Errorf("error while marshaling GroupPointer: %w", err)
		}
	}

	// Serialize `GroupMemberPointer`
	if obj.GroupMemberPointer == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling GroupMemberPointer optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling GroupMemberPointer optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.GroupMemberPointer); err != nil {
			return fmt.Errorf("error while marshaling GroupMemberPointer: %w", err)
		}
	}

	// Serialize `TokenGroup`
	if obj.TokenGroup == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling TokenGroup optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling TokenGroup optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.TokenGroup); err != nil {
			return fmt.Errorf("error while marshaling TokenGroup: %w", err)
		}
	}

	// Serialize `TokenGroupMember`
	if obj.TokenGroupMember == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling TokenGroupMember optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling TokenGroupMember optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.TokenGroupMember); err != nil {
			return fmt.Errorf("error while marshaling TokenGroupMember: %w", err)
		}
	}

	// Serialize `ScaledUiAmountConfig`
	if obj.ScaledUiAmountConfig == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling ScaledUiAmountConfig optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling ScaledUiAmountConfig optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.ScaledUiAmountConfig); err != nil {
			return fmt.Errorf("error while marshaling ScaledUiAmountConfig: %w", err)
		}
	}

	// Serialize `PausableConfig`
	if obj.PausableConfig == nil {
		if err = encoder.WriteOption(false); err != nil {
			return fmt.Errorf("error while marshaling PausableConfig optionality: %w", err)
		}
	} else {
		if err = encoder.WriteOption(true); err != nil {
			return fmt.Errorf("error while marshaling PausableConfig optionality: %w", err)
		}
		if err = EncodeExtension(encoder, obj.PausableConfig); err != nil {
			return fmt.Errorf("error while marshaling PausableConfig: %w", err)
		}
	}
	return nil
}

func (obj Mint) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Mint: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Mint) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Deserialize `MintAuthority` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling MintAuthority:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.MintAuthority); err != nil {
				return fmt.Errorf("error while unmarshaling MintAuthority:%w", err)
			}
		}
	}
	// Deserialize `Supply`:
	if err = decoder.Decode(&obj.Supply); err != nil {
		return fmt.Errorf("error while unmarshaling Supply:%w", err)
	}
	// Deserialize `Decimals`:
	if err = decoder.Decode(&obj.Decimals); err != nil {
		return fmt.Errorf("error while unmarshaling Decimals:%w", err)
	}
	// Deserialize `IsInitialized`:
	if err = decoder.Decode(&obj.IsInitialized); err != nil {
		return fmt.Errorf("error while unmarshaling IsInitialized:%w", err)
	}
	// Deserialize `FreezeAuthority` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling FreezeAuthority:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.FreezeAuthority); err != nil {
				return fmt.Errorf("error while unmarshaling FreezeAuthority:%w", err)
			}
		}
	}
	// Deserialize `Extensions` (zero or more Extension objects)
	for decoder.Remaining() > 0 {
		ext, err := DecodeExtension(decoder)
		if err != nil {
			return fmt.Errorf("error while unmarshaling Extension: %w", err)
		}
		switch v := ext.(type) {
		case *Extension_TransferFeeConfig:
			obj.TransferFeeConfig = v
		case *Extension_MintCloseAuthority:
			obj.MintCloseAuthority = v
		case *Extension_ConfidentialTransferMint:
			obj.ConfidentialTransferMint = v
		case *Extension_DefaultAccountState:
			obj.DefaultAccountState = v
		case *Extension_NonTransferable:
			obj.NonTransferable = v
		case *Extension_InterestBearingConfig:
			obj.InterestBearingConfig = v
		case *Extension_PermanentDelegate:
			obj.PermanentDelegate = v
		case *Extension_TransferHook:
			obj.TransferHook = v
		case *Extension_MetadataPointer:
			obj.MetadataPointer = v
		case *Extension_TokenMetadata:
			obj.TokenMetadata = v
		case *Extension_GroupPointer:
			obj.GroupPointer = v
		case *Extension_GroupMemberPointer:
			obj.GroupMemberPointer = v
		case *Extension_TokenGroup:
			obj.TokenGroup = v
		case *Extension_TokenGroupMember:
			obj.TokenGroupMember = v
		case *Extension_ScaledUiAmountConfig:
			obj.ScaledUiAmountConfig = v
		case *Extension_PausableConfig:
			obj.PausableConfig = v
		default:
			return fmt.Errorf("unknown extension type: %T", ext)
		}
	}
	return nil
}

func (obj *Mint) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Mint: %w", err)
	}
	return nil
}

func UnmarshalMint(buf []byte) (*Mint, error) {
	obj := new(Mint)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type Account struct {
	// The mint associated with this account
	Mint solanago.PublicKey

	// The owner of this account.
	Owner solanago.PublicKey

	// The amount of tokens this account holds.
	Amount uint64

	// If `delegate` is `Some` then `delegated_amount` represents
	// the amount authorized by the delegate
	Delegate *solanago.PublicKey `bin:"optional"`

	// The account's state
	State AccountState

	// If is_some, this is a native token, and the value logs the rent-exempt reserve. An Account
	// is required to be rent-exempt, so the value is used by the Processor to ensure that wrapped
	// SOL accounts do not drop below this threshold.
	IsNative *uint64 `bin:"optional"`

	// The amount delegated
	DelegatedAmount uint64

	// Optional authority to close the account.
	CloseAuthority *solanago.PublicKey `bin:"optional"`
}

func (obj Account) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Account[:], false)
	if err != nil {
		return err
	}
	// Serialize `Mint`:
	if err = encoder.Encode(obj.Mint); err != nil {
		return fmt.Errorf("error while marshaling Mint:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `Amount`:
	if err = encoder.Encode(obj.Amount); err != nil {
		return fmt.Errorf("error while marshaling Amount:%w", err)
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
	// Serialize `State`:
	if err = encoder.Encode(obj.State); err != nil {
		return fmt.Errorf("error while marshaling State:%w", err)
	}
	// Serialize `IsNative` (optional):
	{
		if obj.IsNative == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling IsNative optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling IsNative optionality: %w", err)
			}
			if err = encoder.Encode(obj.IsNative); err != nil {
				return fmt.Errorf("error while marshaling IsNative: %w", err)
			}
		}
	}
	// Serialize `DelegatedAmount`:
	if err = encoder.Encode(obj.DelegatedAmount); err != nil {
		return fmt.Errorf("error while marshaling DelegatedAmount:%w", err)
	}
	// Serialize `CloseAuthority` (optional):
	{
		if obj.CloseAuthority == nil {
			if err = encoder.WriteOption(false); err != nil {
				return fmt.Errorf("error while marshaling CloseAuthority optionality: %w", err)
			}
		} else {
			if err = encoder.WriteOption(true); err != nil {
				return fmt.Errorf("error while marshaling CloseAuthority optionality: %w", err)
			}
			if err = encoder.Encode(obj.CloseAuthority); err != nil {
				return fmt.Errorf("error while marshaling CloseAuthority: %w", err)
			}
		}
	}
	return nil
}

func (obj Account) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Account: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Account) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Account[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Account[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Mint`:
	if err = decoder.Decode(&obj.Mint); err != nil {
		return fmt.Errorf("error while unmarshaling Mint:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `Amount`:
	if err = decoder.Decode(&obj.Amount); err != nil {
		return fmt.Errorf("error while unmarshaling Amount:%w", err)
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
	// Deserialize `State`:
	if err = decoder.Decode(&obj.State); err != nil {
		return fmt.Errorf("error while unmarshaling State:%w", err)
	}
	// Deserialize `IsNative` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling IsNative:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.IsNative); err != nil {
				return fmt.Errorf("error while unmarshaling IsNative:%w", err)
			}
		}
	}
	// Deserialize `DelegatedAmount`:
	if err = decoder.Decode(&obj.DelegatedAmount); err != nil {
		return fmt.Errorf("error while unmarshaling DelegatedAmount:%w", err)
	}
	// Deserialize `CloseAuthority` (optional):
	{
		ok, err := decoder.ReadOption()
		if err != nil {
			return fmt.Errorf("error while unmarshaling CloseAuthority:%w", err)
		}
		if ok {
			if err = decoder.Decode(&obj.CloseAuthority); err != nil {
				return fmt.Errorf("error while unmarshaling CloseAuthority:%w", err)
			}
		}
	}
	return nil
}

func (obj *Account) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Account: %w", err)
	}
	return nil
}

func UnmarshalAccount(buf []byte) (*Account, error) {
	obj := new(Account)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
