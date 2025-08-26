package raydium_cpmm

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

// Emitted when deposit and withdraw
type LpChangeEvent struct {
	PoolId         solanago.PublicKey
	LpAmountBefore uint64

	// pool vault sub trade fees
	Token0VaultBefore uint64

	// pool vault sub trade fees
	Token1VaultBefore uint64

	// calculate result without transfer fee
	Token0Amount uint64

	// calculate result without transfer fee
	Token1Amount      uint64
	Token0TransferFee uint64
	Token1TransferFee uint64
	ChangeType        uint8
}

func (obj LpChangeEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_LpChangeEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolId`:
	if err = encoder.Encode(obj.PoolId); err != nil {
		return fmt.Errorf("error while marshaling PoolId:%w", err)
	}
	// Serialize `LpAmountBefore`:
	if err = encoder.Encode(obj.LpAmountBefore); err != nil {
		return fmt.Errorf("error while marshaling LpAmountBefore:%w", err)
	}
	// Serialize `Token0VaultBefore`:
	if err = encoder.Encode(obj.Token0VaultBefore); err != nil {
		return fmt.Errorf("error while marshaling Token0VaultBefore:%w", err)
	}
	// Serialize `Token1VaultBefore`:
	if err = encoder.Encode(obj.Token1VaultBefore); err != nil {
		return fmt.Errorf("error while marshaling Token1VaultBefore:%w", err)
	}
	// Serialize `Token0Amount`:
	if err = encoder.Encode(obj.Token0Amount); err != nil {
		return fmt.Errorf("error while marshaling Token0Amount:%w", err)
	}
	// Serialize `Token1Amount`:
	if err = encoder.Encode(obj.Token1Amount); err != nil {
		return fmt.Errorf("error while marshaling Token1Amount:%w", err)
	}
	// Serialize `Token0TransferFee`:
	if err = encoder.Encode(obj.Token0TransferFee); err != nil {
		return fmt.Errorf("error while marshaling Token0TransferFee:%w", err)
	}
	// Serialize `Token1TransferFee`:
	if err = encoder.Encode(obj.Token1TransferFee); err != nil {
		return fmt.Errorf("error while marshaling Token1TransferFee:%w", err)
	}
	// Serialize `ChangeType`:
	if err = encoder.Encode(obj.ChangeType); err != nil {
		return fmt.Errorf("error while marshaling ChangeType:%w", err)
	}
	return nil
}

func (obj LpChangeEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding LpChangeEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *LpChangeEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_LpChangeEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_LpChangeEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolId`:
	if err = decoder.Decode(&obj.PoolId); err != nil {
		return fmt.Errorf("error while unmarshaling PoolId:%w", err)
	}
	// Deserialize `LpAmountBefore`:
	if err = decoder.Decode(&obj.LpAmountBefore); err != nil {
		return fmt.Errorf("error while unmarshaling LpAmountBefore:%w", err)
	}
	// Deserialize `Token0VaultBefore`:
	if err = decoder.Decode(&obj.Token0VaultBefore); err != nil {
		return fmt.Errorf("error while unmarshaling Token0VaultBefore:%w", err)
	}
	// Deserialize `Token1VaultBefore`:
	if err = decoder.Decode(&obj.Token1VaultBefore); err != nil {
		return fmt.Errorf("error while unmarshaling Token1VaultBefore:%w", err)
	}
	// Deserialize `Token0Amount`:
	if err = decoder.Decode(&obj.Token0Amount); err != nil {
		return fmt.Errorf("error while unmarshaling Token0Amount:%w", err)
	}
	// Deserialize `Token1Amount`:
	if err = decoder.Decode(&obj.Token1Amount); err != nil {
		return fmt.Errorf("error while unmarshaling Token1Amount:%w", err)
	}
	// Deserialize `Token0TransferFee`:
	if err = decoder.Decode(&obj.Token0TransferFee); err != nil {
		return fmt.Errorf("error while unmarshaling Token0TransferFee:%w", err)
	}
	// Deserialize `Token1TransferFee`:
	if err = decoder.Decode(&obj.Token1TransferFee); err != nil {
		return fmt.Errorf("error while unmarshaling Token1TransferFee:%w", err)
	}
	// Deserialize `ChangeType`:
	if err = decoder.Decode(&obj.ChangeType); err != nil {
		return fmt.Errorf("error while unmarshaling ChangeType:%w", err)
	}
	return nil
}

func (obj *LpChangeEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling LpChangeEvent: %w", err)
	}
	return nil
}

func UnmarshalLpChangeEvent(buf []byte) (*LpChangeEvent, error) {
	obj := new(LpChangeEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Emitted when swap
type SwapEvent struct {
	PoolId solanago.PublicKey

	// pool vault sub trade fees
	InputVaultBefore uint64

	// pool vault sub trade fees
	OutputVaultBefore uint64

	// calculate result without transfer fee
	InputAmount uint64

	// calculate result without transfer fee
	OutputAmount      uint64
	InputTransferFee  uint64
	OutputTransferFee uint64
	BaseInput         bool
}

func (obj SwapEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_SwapEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolId`:
	if err = encoder.Encode(obj.PoolId); err != nil {
		return fmt.Errorf("error while marshaling PoolId:%w", err)
	}
	// Serialize `InputVaultBefore`:
	if err = encoder.Encode(obj.InputVaultBefore); err != nil {
		return fmt.Errorf("error while marshaling InputVaultBefore:%w", err)
	}
	// Serialize `OutputVaultBefore`:
	if err = encoder.Encode(obj.OutputVaultBefore); err != nil {
		return fmt.Errorf("error while marshaling OutputVaultBefore:%w", err)
	}
	// Serialize `InputAmount`:
	if err = encoder.Encode(obj.InputAmount); err != nil {
		return fmt.Errorf("error while marshaling InputAmount:%w", err)
	}
	// Serialize `OutputAmount`:
	if err = encoder.Encode(obj.OutputAmount); err != nil {
		return fmt.Errorf("error while marshaling OutputAmount:%w", err)
	}
	// Serialize `InputTransferFee`:
	if err = encoder.Encode(obj.InputTransferFee); err != nil {
		return fmt.Errorf("error while marshaling InputTransferFee:%w", err)
	}
	// Serialize `OutputTransferFee`:
	if err = encoder.Encode(obj.OutputTransferFee); err != nil {
		return fmt.Errorf("error while marshaling OutputTransferFee:%w", err)
	}
	// Serialize `BaseInput`:
	if err = encoder.Encode(obj.BaseInput); err != nil {
		return fmt.Errorf("error while marshaling BaseInput:%w", err)
	}
	return nil
}

func (obj SwapEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding SwapEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *SwapEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_SwapEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_SwapEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolId`:
	if err = decoder.Decode(&obj.PoolId); err != nil {
		return fmt.Errorf("error while unmarshaling PoolId:%w", err)
	}
	// Deserialize `InputVaultBefore`:
	if err = decoder.Decode(&obj.InputVaultBefore); err != nil {
		return fmt.Errorf("error while unmarshaling InputVaultBefore:%w", err)
	}
	// Deserialize `OutputVaultBefore`:
	if err = decoder.Decode(&obj.OutputVaultBefore); err != nil {
		return fmt.Errorf("error while unmarshaling OutputVaultBefore:%w", err)
	}
	// Deserialize `InputAmount`:
	if err = decoder.Decode(&obj.InputAmount); err != nil {
		return fmt.Errorf("error while unmarshaling InputAmount:%w", err)
	}
	// Deserialize `OutputAmount`:
	if err = decoder.Decode(&obj.OutputAmount); err != nil {
		return fmt.Errorf("error while unmarshaling OutputAmount:%w", err)
	}
	// Deserialize `InputTransferFee`:
	if err = decoder.Decode(&obj.InputTransferFee); err != nil {
		return fmt.Errorf("error while unmarshaling InputTransferFee:%w", err)
	}
	// Deserialize `OutputTransferFee`:
	if err = decoder.Decode(&obj.OutputTransferFee); err != nil {
		return fmt.Errorf("error while unmarshaling OutputTransferFee:%w", err)
	}
	// Deserialize `BaseInput`:
	if err = decoder.Decode(&obj.BaseInput); err != nil {
		return fmt.Errorf("error while unmarshaling BaseInput:%w", err)
	}
	return nil
}

func (obj *SwapEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling SwapEvent: %w", err)
	}
	return nil
}

func UnmarshalSwapEvent(buf []byte) (*SwapEvent, error) {
	obj := new(SwapEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
