package system_program

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type Nonce struct {
	Version       uint32
	State         uint32
	Authorized    solanago.PublicKey
	Nonce         solanago.PublicKey
	FeeCalculator FeeCalculator
}

func (obj Nonce) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Serialize `Version`:
	if err = encoder.Encode(obj.Version); err != nil {
		return fmt.Errorf("error while marshaling Version:%w", err)
	}
	// Serialize `State`:
	if err = encoder.Encode(obj.State); err != nil {
		return fmt.Errorf("error while marshaling State:%w", err)
	}
	// Serialize `Authorized`:
	if err = encoder.Encode(obj.Authorized); err != nil {
		return fmt.Errorf("error while marshaling Authorized:%w", err)
	}
	// Serialize `Nonce`:
	if err = encoder.Encode(obj.Nonce); err != nil {
		return fmt.Errorf("error while marshaling Nonce:%w", err)
	}
	// Serialize `FeeCalculator`:
	if err = encoder.Encode(obj.FeeCalculator); err != nil {
		return fmt.Errorf("error while marshaling FeeCalculator:%w", err)
	}
	return nil
}

func (obj Nonce) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Nonce: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Nonce) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Deserialize `Version`:
	if err = decoder.Decode(&obj.Version); err != nil {
		return fmt.Errorf("error while unmarshaling Version:%w", err)
	}
	// Deserialize `State`:
	if err = decoder.Decode(&obj.State); err != nil {
		return fmt.Errorf("error while unmarshaling State:%w", err)
	}
	// Deserialize `Authorized`:
	if err = decoder.Decode(&obj.Authorized); err != nil {
		return fmt.Errorf("error while unmarshaling Authorized:%w", err)
	}
	// Deserialize `Nonce`:
	if err = decoder.Decode(&obj.Nonce); err != nil {
		return fmt.Errorf("error while unmarshaling Nonce:%w", err)
	}
	// Deserialize `FeeCalculator`:
	if err = decoder.Decode(&obj.FeeCalculator); err != nil {
		return fmt.Errorf("error while unmarshaling FeeCalculator:%w", err)
	}
	return nil
}

func (obj *Nonce) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Nonce: %w", err)
	}
	return nil
}

func UnmarshalNonce(buf []byte) (*Nonce, error) {
	obj := new(Nonce)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
