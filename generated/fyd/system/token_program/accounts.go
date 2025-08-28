package token_program

import (
	"bytes"
	"fmt"

	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type Mint struct {
	MintAuthority   *solanago.PublicKey `bin:"optional"`
	Supply          uint64
	Decimals        uint8
	IsInitialized   bool
	FreezeAuthority *solanago.PublicKey `bin:"optional"`
}

func (obj Mint) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Mint[:], false)
	if err != nil {
		return err
	}
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
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Mint[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Mint[:],
				fmt.Sprint(discriminator[:]))
		}
	}
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

type Token struct {
	Mint            solanago.PublicKey
	Owner           solanago.PublicKey
	Amount          uint64
	Delegate        *solanago.PublicKey `bin:"optional"`
	State           TokenState
	IsNative        *uint64 `bin:"optional"`
	DelegatedAmount uint64
	CloseAuthority  *solanago.PublicKey `bin:"optional"`
}

func (obj Token) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Token[:], false)
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

func (obj Token) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Token: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Token) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Token[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Token[:],
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

func (obj *Token) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Token: %w", err)
	}
	return nil
}

func UnmarshalToken(buf []byte) (*Token, error) {
	obj := new(Token)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type Multisig struct {
	M             uint8
	N             uint8
	IsInitialized bool
	Signers       [11]solanago.PublicKey
}

func (obj Multisig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Multisig[:], false)
	if err != nil {
		return err
	}
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
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Multisig[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Multisig[:],
				fmt.Sprint(discriminator[:]))
		}
	}
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
