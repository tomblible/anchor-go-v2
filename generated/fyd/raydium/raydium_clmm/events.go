package raydium_clmm

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

// Emitted when tokens are collected for a position
type CollectPersonalFeeEvent struct {
	// The ID of the token for which underlying tokens were collected
	PositionNftMint solanago.PublicKey

	// The token account that received the collected token_0 tokens
	RecipientTokenAccount0 solanago.PublicKey

	// The token account that received the collected token_1 tokens
	RecipientTokenAccount1 solanago.PublicKey

	// The amount of token_0 owed to the position that was collected
	Amount0 uint64

	// The amount of token_1 owed to the position that was collected
	Amount1 uint64
}

func (obj CollectPersonalFeeEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CollectPersonalFeeEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `PositionNftMint`:
	if err = encoder.Encode(obj.PositionNftMint); err != nil {
		return fmt.Errorf("error while marshaling PositionNftMint:%w", err)
	}
	// Serialize `RecipientTokenAccount0`:
	if err = encoder.Encode(obj.RecipientTokenAccount0); err != nil {
		return fmt.Errorf("error while marshaling RecipientTokenAccount0:%w", err)
	}
	// Serialize `RecipientTokenAccount1`:
	if err = encoder.Encode(obj.RecipientTokenAccount1); err != nil {
		return fmt.Errorf("error while marshaling RecipientTokenAccount1:%w", err)
	}
	// Serialize `Amount0`:
	if err = encoder.Encode(obj.Amount0); err != nil {
		return fmt.Errorf("error while marshaling Amount0:%w", err)
	}
	// Serialize `Amount1`:
	if err = encoder.Encode(obj.Amount1); err != nil {
		return fmt.Errorf("error while marshaling Amount1:%w", err)
	}
	return nil
}

func (obj CollectPersonalFeeEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding CollectPersonalFeeEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *CollectPersonalFeeEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CollectPersonalFeeEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CollectPersonalFeeEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PositionNftMint`:
	if err = decoder.Decode(&obj.PositionNftMint); err != nil {
		return fmt.Errorf("error while unmarshaling PositionNftMint:%w", err)
	}
	// Deserialize `RecipientTokenAccount0`:
	if err = decoder.Decode(&obj.RecipientTokenAccount0); err != nil {
		return fmt.Errorf("error while unmarshaling RecipientTokenAccount0:%w", err)
	}
	// Deserialize `RecipientTokenAccount1`:
	if err = decoder.Decode(&obj.RecipientTokenAccount1); err != nil {
		return fmt.Errorf("error while unmarshaling RecipientTokenAccount1:%w", err)
	}
	// Deserialize `Amount0`:
	if err = decoder.Decode(&obj.Amount0); err != nil {
		return fmt.Errorf("error while unmarshaling Amount0:%w", err)
	}
	// Deserialize `Amount1`:
	if err = decoder.Decode(&obj.Amount1); err != nil {
		return fmt.Errorf("error while unmarshaling Amount1:%w", err)
	}
	return nil
}

func (obj *CollectPersonalFeeEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling CollectPersonalFeeEvent: %w", err)
	}
	return nil
}

func UnmarshalCollectPersonalFeeEvent(buf []byte) (*CollectPersonalFeeEvent, error) {
	obj := new(CollectPersonalFeeEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Emitted when the collected protocol fees are withdrawn by the factory owner
type CollectProtocolFeeEvent struct {
	// The pool whose protocol fee is collected
	PoolState solanago.PublicKey

	// The address that receives the collected token_0 protocol fees
	RecipientTokenAccount0 solanago.PublicKey

	// The address that receives the collected token_1 protocol fees
	RecipientTokenAccount1 solanago.PublicKey

	// The amount of token_0 protocol fees that is withdrawn
	Amount0 uint64

	// The amount of token_0 protocol fees that is withdrawn
	Amount1 uint64
}

func (obj CollectProtocolFeeEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CollectProtocolFeeEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolState`:
	if err = encoder.Encode(obj.PoolState); err != nil {
		return fmt.Errorf("error while marshaling PoolState:%w", err)
	}
	// Serialize `RecipientTokenAccount0`:
	if err = encoder.Encode(obj.RecipientTokenAccount0); err != nil {
		return fmt.Errorf("error while marshaling RecipientTokenAccount0:%w", err)
	}
	// Serialize `RecipientTokenAccount1`:
	if err = encoder.Encode(obj.RecipientTokenAccount1); err != nil {
		return fmt.Errorf("error while marshaling RecipientTokenAccount1:%w", err)
	}
	// Serialize `Amount0`:
	if err = encoder.Encode(obj.Amount0); err != nil {
		return fmt.Errorf("error while marshaling Amount0:%w", err)
	}
	// Serialize `Amount1`:
	if err = encoder.Encode(obj.Amount1); err != nil {
		return fmt.Errorf("error while marshaling Amount1:%w", err)
	}
	return nil
}

func (obj CollectProtocolFeeEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding CollectProtocolFeeEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *CollectProtocolFeeEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CollectProtocolFeeEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CollectProtocolFeeEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolState`:
	if err = decoder.Decode(&obj.PoolState); err != nil {
		return fmt.Errorf("error while unmarshaling PoolState:%w", err)
	}
	// Deserialize `RecipientTokenAccount0`:
	if err = decoder.Decode(&obj.RecipientTokenAccount0); err != nil {
		return fmt.Errorf("error while unmarshaling RecipientTokenAccount0:%w", err)
	}
	// Deserialize `RecipientTokenAccount1`:
	if err = decoder.Decode(&obj.RecipientTokenAccount1); err != nil {
		return fmt.Errorf("error while unmarshaling RecipientTokenAccount1:%w", err)
	}
	// Deserialize `Amount0`:
	if err = decoder.Decode(&obj.Amount0); err != nil {
		return fmt.Errorf("error while unmarshaling Amount0:%w", err)
	}
	// Deserialize `Amount1`:
	if err = decoder.Decode(&obj.Amount1); err != nil {
		return fmt.Errorf("error while unmarshaling Amount1:%w", err)
	}
	return nil
}

func (obj *CollectProtocolFeeEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling CollectProtocolFeeEvent: %w", err)
	}
	return nil
}

func UnmarshalCollectProtocolFeeEvent(buf []byte) (*CollectProtocolFeeEvent, error) {
	obj := new(CollectProtocolFeeEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Emitted when create or update a config
type ConfigChangeEvent struct {
	Index           uint16
	Owner           solanago.PublicKey
	ProtocolFeeRate uint32
	TradeFeeRate    uint32
	TickSpacing     uint16
	FundFeeRate     uint32
	FundOwner       solanago.PublicKey
}

func (obj ConfigChangeEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_ConfigChangeEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `Index`:
	if err = encoder.Encode(obj.Index); err != nil {
		return fmt.Errorf("error while marshaling Index:%w", err)
	}
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `ProtocolFeeRate`:
	if err = encoder.Encode(obj.ProtocolFeeRate); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeRate:%w", err)
	}
	// Serialize `TradeFeeRate`:
	if err = encoder.Encode(obj.TradeFeeRate); err != nil {
		return fmt.Errorf("error while marshaling TradeFeeRate:%w", err)
	}
	// Serialize `TickSpacing`:
	if err = encoder.Encode(obj.TickSpacing); err != nil {
		return fmt.Errorf("error while marshaling TickSpacing:%w", err)
	}
	// Serialize `FundFeeRate`:
	if err = encoder.Encode(obj.FundFeeRate); err != nil {
		return fmt.Errorf("error while marshaling FundFeeRate:%w", err)
	}
	// Serialize `FundOwner`:
	if err = encoder.Encode(obj.FundOwner); err != nil {
		return fmt.Errorf("error while marshaling FundOwner:%w", err)
	}
	return nil
}

func (obj ConfigChangeEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding ConfigChangeEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *ConfigChangeEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_ConfigChangeEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_ConfigChangeEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Index`:
	if err = decoder.Decode(&obj.Index); err != nil {
		return fmt.Errorf("error while unmarshaling Index:%w", err)
	}
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `ProtocolFeeRate`:
	if err = decoder.Decode(&obj.ProtocolFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeRate:%w", err)
	}
	// Deserialize `TradeFeeRate`:
	if err = decoder.Decode(&obj.TradeFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling TradeFeeRate:%w", err)
	}
	// Deserialize `TickSpacing`:
	if err = decoder.Decode(&obj.TickSpacing); err != nil {
		return fmt.Errorf("error while unmarshaling TickSpacing:%w", err)
	}
	// Deserialize `FundFeeRate`:
	if err = decoder.Decode(&obj.FundFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling FundFeeRate:%w", err)
	}
	// Deserialize `FundOwner`:
	if err = decoder.Decode(&obj.FundOwner); err != nil {
		return fmt.Errorf("error while unmarshaling FundOwner:%w", err)
	}
	return nil
}

func (obj *ConfigChangeEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling ConfigChangeEvent: %w", err)
	}
	return nil
}

func UnmarshalConfigChangeEvent(buf []byte) (*ConfigChangeEvent, error) {
	obj := new(ConfigChangeEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Emitted when create a new position
type CreatePersonalPositionEvent struct {
	// The pool for which liquidity was added
	PoolState solanago.PublicKey

	// The address that create the position
	Minter solanago.PublicKey

	// The owner of the position and recipient of any minted liquidity
	NftOwner solanago.PublicKey

	// The lower tick of the position
	TickLowerIndex int32

	// The upper tick of the position
	TickUpperIndex int32

	// The amount of liquidity minted to the position range
	Liquidity binary.Uint128

	// The amount of token_0 was deposit for the liquidity
	DepositAmount0 uint64

	// The amount of token_1 was deposit for the liquidity
	DepositAmount1 uint64

	// The token transfer fee for deposit_amount_0
	DepositAmount0TransferFee uint64

	// The token transfer fee for deposit_amount_1
	DepositAmount1TransferFee uint64
}

func (obj CreatePersonalPositionEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_CreatePersonalPositionEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolState`:
	if err = encoder.Encode(obj.PoolState); err != nil {
		return fmt.Errorf("error while marshaling PoolState:%w", err)
	}
	// Serialize `Minter`:
	if err = encoder.Encode(obj.Minter); err != nil {
		return fmt.Errorf("error while marshaling Minter:%w", err)
	}
	// Serialize `NftOwner`:
	if err = encoder.Encode(obj.NftOwner); err != nil {
		return fmt.Errorf("error while marshaling NftOwner:%w", err)
	}
	// Serialize `TickLowerIndex`:
	if err = encoder.Encode(obj.TickLowerIndex); err != nil {
		return fmt.Errorf("error while marshaling TickLowerIndex:%w", err)
	}
	// Serialize `TickUpperIndex`:
	if err = encoder.Encode(obj.TickUpperIndex); err != nil {
		return fmt.Errorf("error while marshaling TickUpperIndex:%w", err)
	}
	// Serialize `Liquidity`:
	if err = encoder.Encode(obj.Liquidity); err != nil {
		return fmt.Errorf("error while marshaling Liquidity:%w", err)
	}
	// Serialize `DepositAmount0`:
	if err = encoder.Encode(obj.DepositAmount0); err != nil {
		return fmt.Errorf("error while marshaling DepositAmount0:%w", err)
	}
	// Serialize `DepositAmount1`:
	if err = encoder.Encode(obj.DepositAmount1); err != nil {
		return fmt.Errorf("error while marshaling DepositAmount1:%w", err)
	}
	// Serialize `DepositAmount0TransferFee`:
	if err = encoder.Encode(obj.DepositAmount0TransferFee); err != nil {
		return fmt.Errorf("error while marshaling DepositAmount0TransferFee:%w", err)
	}
	// Serialize `DepositAmount1TransferFee`:
	if err = encoder.Encode(obj.DepositAmount1TransferFee); err != nil {
		return fmt.Errorf("error while marshaling DepositAmount1TransferFee:%w", err)
	}
	return nil
}

func (obj CreatePersonalPositionEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding CreatePersonalPositionEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *CreatePersonalPositionEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_CreatePersonalPositionEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_CreatePersonalPositionEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolState`:
	if err = decoder.Decode(&obj.PoolState); err != nil {
		return fmt.Errorf("error while unmarshaling PoolState:%w", err)
	}
	// Deserialize `Minter`:
	if err = decoder.Decode(&obj.Minter); err != nil {
		return fmt.Errorf("error while unmarshaling Minter:%w", err)
	}
	// Deserialize `NftOwner`:
	if err = decoder.Decode(&obj.NftOwner); err != nil {
		return fmt.Errorf("error while unmarshaling NftOwner:%w", err)
	}
	// Deserialize `TickLowerIndex`:
	if err = decoder.Decode(&obj.TickLowerIndex); err != nil {
		return fmt.Errorf("error while unmarshaling TickLowerIndex:%w", err)
	}
	// Deserialize `TickUpperIndex`:
	if err = decoder.Decode(&obj.TickUpperIndex); err != nil {
		return fmt.Errorf("error while unmarshaling TickUpperIndex:%w", err)
	}
	// Deserialize `Liquidity`:
	if err = decoder.Decode(&obj.Liquidity); err != nil {
		return fmt.Errorf("error while unmarshaling Liquidity:%w", err)
	}
	// Deserialize `DepositAmount0`:
	if err = decoder.Decode(&obj.DepositAmount0); err != nil {
		return fmt.Errorf("error while unmarshaling DepositAmount0:%w", err)
	}
	// Deserialize `DepositAmount1`:
	if err = decoder.Decode(&obj.DepositAmount1); err != nil {
		return fmt.Errorf("error while unmarshaling DepositAmount1:%w", err)
	}
	// Deserialize `DepositAmount0TransferFee`:
	if err = decoder.Decode(&obj.DepositAmount0TransferFee); err != nil {
		return fmt.Errorf("error while unmarshaling DepositAmount0TransferFee:%w", err)
	}
	// Deserialize `DepositAmount1TransferFee`:
	if err = decoder.Decode(&obj.DepositAmount1TransferFee); err != nil {
		return fmt.Errorf("error while unmarshaling DepositAmount1TransferFee:%w", err)
	}
	return nil
}

func (obj *CreatePersonalPositionEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling CreatePersonalPositionEvent: %w", err)
	}
	return nil
}

func UnmarshalCreatePersonalPositionEvent(buf []byte) (*CreatePersonalPositionEvent, error) {
	obj := new(CreatePersonalPositionEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Emitted when liquidity is decreased.
type DecreaseLiquidityEvent struct {
	// The ID of the token for which liquidity was decreased
	PositionNftMint solanago.PublicKey

	// The amount by which liquidity for the position was decreased
	Liquidity binary.Uint128

	// The amount of token_0 that was paid for the decrease in liquidity
	DecreaseAmount0 uint64

	// The amount of token_1 that was paid for the decrease in liquidity
	DecreaseAmount1 uint64
	FeeAmount0      uint64

	// The amount of token_1 fee
	FeeAmount1 uint64

	// The amount of rewards
	RewardAmounts [3]uint64

	// The amount of token_0 transfer fee
	TransferFee0 uint64

	// The amount of token_1 transfer fee
	TransferFee1 uint64
}

func (obj DecreaseLiquidityEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_DecreaseLiquidityEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `PositionNftMint`:
	if err = encoder.Encode(obj.PositionNftMint); err != nil {
		return fmt.Errorf("error while marshaling PositionNftMint:%w", err)
	}
	// Serialize `Liquidity`:
	if err = encoder.Encode(obj.Liquidity); err != nil {
		return fmt.Errorf("error while marshaling Liquidity:%w", err)
	}
	// Serialize `DecreaseAmount0`:
	if err = encoder.Encode(obj.DecreaseAmount0); err != nil {
		return fmt.Errorf("error while marshaling DecreaseAmount0:%w", err)
	}
	// Serialize `DecreaseAmount1`:
	if err = encoder.Encode(obj.DecreaseAmount1); err != nil {
		return fmt.Errorf("error while marshaling DecreaseAmount1:%w", err)
	}
	// Serialize `FeeAmount0`:
	if err = encoder.Encode(obj.FeeAmount0); err != nil {
		return fmt.Errorf("error while marshaling FeeAmount0:%w", err)
	}
	// Serialize `FeeAmount1`:
	if err = encoder.Encode(obj.FeeAmount1); err != nil {
		return fmt.Errorf("error while marshaling FeeAmount1:%w", err)
	}
	// Serialize `RewardAmounts`:
	if err = encoder.Encode(obj.RewardAmounts); err != nil {
		return fmt.Errorf("error while marshaling RewardAmounts:%w", err)
	}
	// Serialize `TransferFee0`:
	if err = encoder.Encode(obj.TransferFee0); err != nil {
		return fmt.Errorf("error while marshaling TransferFee0:%w", err)
	}
	// Serialize `TransferFee1`:
	if err = encoder.Encode(obj.TransferFee1); err != nil {
		return fmt.Errorf("error while marshaling TransferFee1:%w", err)
	}
	return nil
}

func (obj DecreaseLiquidityEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding DecreaseLiquidityEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *DecreaseLiquidityEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_DecreaseLiquidityEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_DecreaseLiquidityEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PositionNftMint`:
	if err = decoder.Decode(&obj.PositionNftMint); err != nil {
		return fmt.Errorf("error while unmarshaling PositionNftMint:%w", err)
	}
	// Deserialize `Liquidity`:
	if err = decoder.Decode(&obj.Liquidity); err != nil {
		return fmt.Errorf("error while unmarshaling Liquidity:%w", err)
	}
	// Deserialize `DecreaseAmount0`:
	if err = decoder.Decode(&obj.DecreaseAmount0); err != nil {
		return fmt.Errorf("error while unmarshaling DecreaseAmount0:%w", err)
	}
	// Deserialize `DecreaseAmount1`:
	if err = decoder.Decode(&obj.DecreaseAmount1); err != nil {
		return fmt.Errorf("error while unmarshaling DecreaseAmount1:%w", err)
	}
	// Deserialize `FeeAmount0`:
	if err = decoder.Decode(&obj.FeeAmount0); err != nil {
		return fmt.Errorf("error while unmarshaling FeeAmount0:%w", err)
	}
	// Deserialize `FeeAmount1`:
	if err = decoder.Decode(&obj.FeeAmount1); err != nil {
		return fmt.Errorf("error while unmarshaling FeeAmount1:%w", err)
	}
	// Deserialize `RewardAmounts`:
	if err = decoder.Decode(&obj.RewardAmounts); err != nil {
		return fmt.Errorf("error while unmarshaling RewardAmounts:%w", err)
	}
	// Deserialize `TransferFee0`:
	if err = decoder.Decode(&obj.TransferFee0); err != nil {
		return fmt.Errorf("error while unmarshaling TransferFee0:%w", err)
	}
	// Deserialize `TransferFee1`:
	if err = decoder.Decode(&obj.TransferFee1); err != nil {
		return fmt.Errorf("error while unmarshaling TransferFee1:%w", err)
	}
	return nil
}

func (obj *DecreaseLiquidityEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling DecreaseLiquidityEvent: %w", err)
	}
	return nil
}

func UnmarshalDecreaseLiquidityEvent(buf []byte) (*DecreaseLiquidityEvent, error) {
	obj := new(DecreaseLiquidityEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Emitted when liquidity is increased.
type IncreaseLiquidityEvent struct {
	// The ID of the token for which liquidity was increased
	PositionNftMint solanago.PublicKey

	// The amount by which liquidity for the NFT position was increased
	Liquidity binary.Uint128

	// The amount of token_0 that was paid for the increase in liquidity
	Amount0 uint64

	// The amount of token_1 that was paid for the increase in liquidity
	Amount1 uint64

	// The token transfer fee for amount_0
	Amount0TransferFee uint64

	// The token transfer fee for amount_1
	Amount1TransferFee uint64
}

func (obj IncreaseLiquidityEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_IncreaseLiquidityEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `PositionNftMint`:
	if err = encoder.Encode(obj.PositionNftMint); err != nil {
		return fmt.Errorf("error while marshaling PositionNftMint:%w", err)
	}
	// Serialize `Liquidity`:
	if err = encoder.Encode(obj.Liquidity); err != nil {
		return fmt.Errorf("error while marshaling Liquidity:%w", err)
	}
	// Serialize `Amount0`:
	if err = encoder.Encode(obj.Amount0); err != nil {
		return fmt.Errorf("error while marshaling Amount0:%w", err)
	}
	// Serialize `Amount1`:
	if err = encoder.Encode(obj.Amount1); err != nil {
		return fmt.Errorf("error while marshaling Amount1:%w", err)
	}
	// Serialize `Amount0TransferFee`:
	if err = encoder.Encode(obj.Amount0TransferFee); err != nil {
		return fmt.Errorf("error while marshaling Amount0TransferFee:%w", err)
	}
	// Serialize `Amount1TransferFee`:
	if err = encoder.Encode(obj.Amount1TransferFee); err != nil {
		return fmt.Errorf("error while marshaling Amount1TransferFee:%w", err)
	}
	return nil
}

func (obj IncreaseLiquidityEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding IncreaseLiquidityEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *IncreaseLiquidityEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_IncreaseLiquidityEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_IncreaseLiquidityEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PositionNftMint`:
	if err = decoder.Decode(&obj.PositionNftMint); err != nil {
		return fmt.Errorf("error while unmarshaling PositionNftMint:%w", err)
	}
	// Deserialize `Liquidity`:
	if err = decoder.Decode(&obj.Liquidity); err != nil {
		return fmt.Errorf("error while unmarshaling Liquidity:%w", err)
	}
	// Deserialize `Amount0`:
	if err = decoder.Decode(&obj.Amount0); err != nil {
		return fmt.Errorf("error while unmarshaling Amount0:%w", err)
	}
	// Deserialize `Amount1`:
	if err = decoder.Decode(&obj.Amount1); err != nil {
		return fmt.Errorf("error while unmarshaling Amount1:%w", err)
	}
	// Deserialize `Amount0TransferFee`:
	if err = decoder.Decode(&obj.Amount0TransferFee); err != nil {
		return fmt.Errorf("error while unmarshaling Amount0TransferFee:%w", err)
	}
	// Deserialize `Amount1TransferFee`:
	if err = decoder.Decode(&obj.Amount1TransferFee); err != nil {
		return fmt.Errorf("error while unmarshaling Amount1TransferFee:%w", err)
	}
	return nil
}

func (obj *IncreaseLiquidityEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling IncreaseLiquidityEvent: %w", err)
	}
	return nil
}

func UnmarshalIncreaseLiquidityEvent(buf []byte) (*IncreaseLiquidityEvent, error) {
	obj := new(IncreaseLiquidityEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Emitted when liquidity decreased or increase.
type LiquidityCalculateEvent struct {
	// The pool liquidity before decrease or increase
	PoolLiquidity binary.Uint128

	// The pool price when decrease or increase in liquidity
	PoolSqrtPriceX64 binary.Uint128

	// The pool tick when decrease or increase in liquidity
	PoolTick int32

	// The amount of token_0 that was calculated for the decrease or increase in liquidity
	CalcAmount0 uint64

	// The amount of token_1 that was calculated for the decrease or increase in liquidity
	CalcAmount1   uint64
	TradeFeeOwed0 uint64

	// The amount of token_1 fee
	TradeFeeOwed1 uint64

	// The amount of token_0 transfer fee without trade_fee_amount_0
	TransferFee0 uint64

	// The amount of token_1 transfer fee without trade_fee_amount_0
	TransferFee1 uint64
}

func (obj LiquidityCalculateEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_LiquidityCalculateEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolLiquidity`:
	if err = encoder.Encode(obj.PoolLiquidity); err != nil {
		return fmt.Errorf("error while marshaling PoolLiquidity:%w", err)
	}
	// Serialize `PoolSqrtPriceX64`:
	if err = encoder.Encode(obj.PoolSqrtPriceX64); err != nil {
		return fmt.Errorf("error while marshaling PoolSqrtPriceX64:%w", err)
	}
	// Serialize `PoolTick`:
	if err = encoder.Encode(obj.PoolTick); err != nil {
		return fmt.Errorf("error while marshaling PoolTick:%w", err)
	}
	// Serialize `CalcAmount0`:
	if err = encoder.Encode(obj.CalcAmount0); err != nil {
		return fmt.Errorf("error while marshaling CalcAmount0:%w", err)
	}
	// Serialize `CalcAmount1`:
	if err = encoder.Encode(obj.CalcAmount1); err != nil {
		return fmt.Errorf("error while marshaling CalcAmount1:%w", err)
	}
	// Serialize `TradeFeeOwed0`:
	if err = encoder.Encode(obj.TradeFeeOwed0); err != nil {
		return fmt.Errorf("error while marshaling TradeFeeOwed0:%w", err)
	}
	// Serialize `TradeFeeOwed1`:
	if err = encoder.Encode(obj.TradeFeeOwed1); err != nil {
		return fmt.Errorf("error while marshaling TradeFeeOwed1:%w", err)
	}
	// Serialize `TransferFee0`:
	if err = encoder.Encode(obj.TransferFee0); err != nil {
		return fmt.Errorf("error while marshaling TransferFee0:%w", err)
	}
	// Serialize `TransferFee1`:
	if err = encoder.Encode(obj.TransferFee1); err != nil {
		return fmt.Errorf("error while marshaling TransferFee1:%w", err)
	}
	return nil
}

func (obj LiquidityCalculateEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding LiquidityCalculateEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *LiquidityCalculateEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_LiquidityCalculateEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_LiquidityCalculateEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolLiquidity`:
	if err = decoder.Decode(&obj.PoolLiquidity); err != nil {
		return fmt.Errorf("error while unmarshaling PoolLiquidity:%w", err)
	}
	// Deserialize `PoolSqrtPriceX64`:
	if err = decoder.Decode(&obj.PoolSqrtPriceX64); err != nil {
		return fmt.Errorf("error while unmarshaling PoolSqrtPriceX64:%w", err)
	}
	// Deserialize `PoolTick`:
	if err = decoder.Decode(&obj.PoolTick); err != nil {
		return fmt.Errorf("error while unmarshaling PoolTick:%w", err)
	}
	// Deserialize `CalcAmount0`:
	if err = decoder.Decode(&obj.CalcAmount0); err != nil {
		return fmt.Errorf("error while unmarshaling CalcAmount0:%w", err)
	}
	// Deserialize `CalcAmount1`:
	if err = decoder.Decode(&obj.CalcAmount1); err != nil {
		return fmt.Errorf("error while unmarshaling CalcAmount1:%w", err)
	}
	// Deserialize `TradeFeeOwed0`:
	if err = decoder.Decode(&obj.TradeFeeOwed0); err != nil {
		return fmt.Errorf("error while unmarshaling TradeFeeOwed0:%w", err)
	}
	// Deserialize `TradeFeeOwed1`:
	if err = decoder.Decode(&obj.TradeFeeOwed1); err != nil {
		return fmt.Errorf("error while unmarshaling TradeFeeOwed1:%w", err)
	}
	// Deserialize `TransferFee0`:
	if err = decoder.Decode(&obj.TransferFee0); err != nil {
		return fmt.Errorf("error while unmarshaling TransferFee0:%w", err)
	}
	// Deserialize `TransferFee1`:
	if err = decoder.Decode(&obj.TransferFee1); err != nil {
		return fmt.Errorf("error while unmarshaling TransferFee1:%w", err)
	}
	return nil
}

func (obj *LiquidityCalculateEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling LiquidityCalculateEvent: %w", err)
	}
	return nil
}

func UnmarshalLiquidityCalculateEvent(buf []byte) (*LiquidityCalculateEvent, error) {
	obj := new(LiquidityCalculateEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Emitted pool liquidity change when increase and decrease liquidity
type LiquidityChangeEvent struct {
	// The pool for swap
	PoolState solanago.PublicKey

	// The tick of the pool
	Tick int32

	// The tick lower of position
	TickLower int32

	// The tick lower of position
	TickUpper int32

	// The liquidity of the pool before liquidity change
	LiquidityBefore binary.Uint128

	// The liquidity of the pool after liquidity change
	LiquidityAfter binary.Uint128
}

func (obj LiquidityChangeEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_LiquidityChangeEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolState`:
	if err = encoder.Encode(obj.PoolState); err != nil {
		return fmt.Errorf("error while marshaling PoolState:%w", err)
	}
	// Serialize `Tick`:
	if err = encoder.Encode(obj.Tick); err != nil {
		return fmt.Errorf("error while marshaling Tick:%w", err)
	}
	// Serialize `TickLower`:
	if err = encoder.Encode(obj.TickLower); err != nil {
		return fmt.Errorf("error while marshaling TickLower:%w", err)
	}
	// Serialize `TickUpper`:
	if err = encoder.Encode(obj.TickUpper); err != nil {
		return fmt.Errorf("error while marshaling TickUpper:%w", err)
	}
	// Serialize `LiquidityBefore`:
	if err = encoder.Encode(obj.LiquidityBefore); err != nil {
		return fmt.Errorf("error while marshaling LiquidityBefore:%w", err)
	}
	// Serialize `LiquidityAfter`:
	if err = encoder.Encode(obj.LiquidityAfter); err != nil {
		return fmt.Errorf("error while marshaling LiquidityAfter:%w", err)
	}
	return nil
}

func (obj LiquidityChangeEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding LiquidityChangeEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *LiquidityChangeEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_LiquidityChangeEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_LiquidityChangeEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolState`:
	if err = decoder.Decode(&obj.PoolState); err != nil {
		return fmt.Errorf("error while unmarshaling PoolState:%w", err)
	}
	// Deserialize `Tick`:
	if err = decoder.Decode(&obj.Tick); err != nil {
		return fmt.Errorf("error while unmarshaling Tick:%w", err)
	}
	// Deserialize `TickLower`:
	if err = decoder.Decode(&obj.TickLower); err != nil {
		return fmt.Errorf("error while unmarshaling TickLower:%w", err)
	}
	// Deserialize `TickUpper`:
	if err = decoder.Decode(&obj.TickUpper); err != nil {
		return fmt.Errorf("error while unmarshaling TickUpper:%w", err)
	}
	// Deserialize `LiquidityBefore`:
	if err = decoder.Decode(&obj.LiquidityBefore); err != nil {
		return fmt.Errorf("error while unmarshaling LiquidityBefore:%w", err)
	}
	// Deserialize `LiquidityAfter`:
	if err = decoder.Decode(&obj.LiquidityAfter); err != nil {
		return fmt.Errorf("error while unmarshaling LiquidityAfter:%w", err)
	}
	return nil
}

func (obj *LiquidityChangeEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling LiquidityChangeEvent: %w", err)
	}
	return nil
}

func UnmarshalLiquidityChangeEvent(buf []byte) (*LiquidityChangeEvent, error) {
	obj := new(LiquidityChangeEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Emitted when a pool is created and initialized with a starting price
type PoolCreatedEvent struct {
	// The first token of the pool by address sort order
	TokenMint0 solanago.PublicKey

	// The second token of the pool by address sort order
	TokenMint1 solanago.PublicKey

	// The minimum number of ticks between initialized ticks
	TickSpacing uint16

	// The address of the created pool
	PoolState solanago.PublicKey

	// The initial sqrt price of the pool, as a Q64.64
	SqrtPriceX64 binary.Uint128

	// The initial tick of the pool, i.e. log base 1.0001 of the starting price of the pool
	Tick int32

	// Vault of token_0
	TokenVault0 solanago.PublicKey

	// Vault of token_1
	TokenVault1 solanago.PublicKey
}

func (obj PoolCreatedEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_PoolCreatedEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `TokenMint0`:
	if err = encoder.Encode(obj.TokenMint0); err != nil {
		return fmt.Errorf("error while marshaling TokenMint0:%w", err)
	}
	// Serialize `TokenMint1`:
	if err = encoder.Encode(obj.TokenMint1); err != nil {
		return fmt.Errorf("error while marshaling TokenMint1:%w", err)
	}
	// Serialize `TickSpacing`:
	if err = encoder.Encode(obj.TickSpacing); err != nil {
		return fmt.Errorf("error while marshaling TickSpacing:%w", err)
	}
	// Serialize `PoolState`:
	if err = encoder.Encode(obj.PoolState); err != nil {
		return fmt.Errorf("error while marshaling PoolState:%w", err)
	}
	// Serialize `SqrtPriceX64`:
	if err = encoder.Encode(obj.SqrtPriceX64); err != nil {
		return fmt.Errorf("error while marshaling SqrtPriceX64:%w", err)
	}
	// Serialize `Tick`:
	if err = encoder.Encode(obj.Tick); err != nil {
		return fmt.Errorf("error while marshaling Tick:%w", err)
	}
	// Serialize `TokenVault0`:
	if err = encoder.Encode(obj.TokenVault0); err != nil {
		return fmt.Errorf("error while marshaling TokenVault0:%w", err)
	}
	// Serialize `TokenVault1`:
	if err = encoder.Encode(obj.TokenVault1); err != nil {
		return fmt.Errorf("error while marshaling TokenVault1:%w", err)
	}
	return nil
}

func (obj PoolCreatedEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding PoolCreatedEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *PoolCreatedEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_PoolCreatedEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_PoolCreatedEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TokenMint0`:
	if err = decoder.Decode(&obj.TokenMint0); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMint0:%w", err)
	}
	// Deserialize `TokenMint1`:
	if err = decoder.Decode(&obj.TokenMint1); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMint1:%w", err)
	}
	// Deserialize `TickSpacing`:
	if err = decoder.Decode(&obj.TickSpacing); err != nil {
		return fmt.Errorf("error while unmarshaling TickSpacing:%w", err)
	}
	// Deserialize `PoolState`:
	if err = decoder.Decode(&obj.PoolState); err != nil {
		return fmt.Errorf("error while unmarshaling PoolState:%w", err)
	}
	// Deserialize `SqrtPriceX64`:
	if err = decoder.Decode(&obj.SqrtPriceX64); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtPriceX64:%w", err)
	}
	// Deserialize `Tick`:
	if err = decoder.Decode(&obj.Tick); err != nil {
		return fmt.Errorf("error while unmarshaling Tick:%w", err)
	}
	// Deserialize `TokenVault0`:
	if err = decoder.Decode(&obj.TokenVault0); err != nil {
		return fmt.Errorf("error while unmarshaling TokenVault0:%w", err)
	}
	// Deserialize `TokenVault1`:
	if err = decoder.Decode(&obj.TokenVault1); err != nil {
		return fmt.Errorf("error while unmarshaling TokenVault1:%w", err)
	}
	return nil
}

func (obj *PoolCreatedEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling PoolCreatedEvent: %w", err)
	}
	return nil
}

func UnmarshalPoolCreatedEvent(buf []byte) (*PoolCreatedEvent, error) {
	obj := new(PoolCreatedEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Emitted by when a swap is performed for a pool
type SwapEvent struct {
	// The pool for which token_0 and token_1 were swapped
	PoolState solanago.PublicKey

	// The address that initiated the swap call, and that received the callback
	Sender solanago.PublicKey

	// The payer token account in zero for one swaps, or the recipient token account
	// in one for zero swaps
	TokenAccount0 solanago.PublicKey

	// The payer token account in one for zero swaps, or the recipient token account
	// in zero for one swaps
	TokenAccount1 solanago.PublicKey

	// The real delta amount of the token_0 of the pool or user
	Amount0 uint64

	// The transfer fee charged by the withheld_amount of the token_0
	TransferFee0 uint64

	// The real delta of the token_1 of the pool or user
	Amount1 uint64

	// The transfer fee charged by the withheld_amount of the token_1
	TransferFee1 uint64

	// if true, amount_0 is negtive and amount_1 is positive
	ZeroForOne bool

	// The sqrt(price) of the pool after the swap, as a Q64.64
	SqrtPriceX64 binary.Uint128

	// The liquidity of the pool after the swap
	Liquidity binary.Uint128

	// The log base 1.0001 of price of the pool after the swap
	Tick int32
}

func (obj SwapEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_SwapEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolState`:
	if err = encoder.Encode(obj.PoolState); err != nil {
		return fmt.Errorf("error while marshaling PoolState:%w", err)
	}
	// Serialize `Sender`:
	if err = encoder.Encode(obj.Sender); err != nil {
		return fmt.Errorf("error while marshaling Sender:%w", err)
	}
	// Serialize `TokenAccount0`:
	if err = encoder.Encode(obj.TokenAccount0); err != nil {
		return fmt.Errorf("error while marshaling TokenAccount0:%w", err)
	}
	// Serialize `TokenAccount1`:
	if err = encoder.Encode(obj.TokenAccount1); err != nil {
		return fmt.Errorf("error while marshaling TokenAccount1:%w", err)
	}
	// Serialize `Amount0`:
	if err = encoder.Encode(obj.Amount0); err != nil {
		return fmt.Errorf("error while marshaling Amount0:%w", err)
	}
	// Serialize `TransferFee0`:
	if err = encoder.Encode(obj.TransferFee0); err != nil {
		return fmt.Errorf("error while marshaling TransferFee0:%w", err)
	}
	// Serialize `Amount1`:
	if err = encoder.Encode(obj.Amount1); err != nil {
		return fmt.Errorf("error while marshaling Amount1:%w", err)
	}
	// Serialize `TransferFee1`:
	if err = encoder.Encode(obj.TransferFee1); err != nil {
		return fmt.Errorf("error while marshaling TransferFee1:%w", err)
	}
	// Serialize `ZeroForOne`:
	if err = encoder.Encode(obj.ZeroForOne); err != nil {
		return fmt.Errorf("error while marshaling ZeroForOne:%w", err)
	}
	// Serialize `SqrtPriceX64`:
	if err = encoder.Encode(obj.SqrtPriceX64); err != nil {
		return fmt.Errorf("error while marshaling SqrtPriceX64:%w", err)
	}
	// Serialize `Liquidity`:
	if err = encoder.Encode(obj.Liquidity); err != nil {
		return fmt.Errorf("error while marshaling Liquidity:%w", err)
	}
	// Serialize `Tick`:
	if err = encoder.Encode(obj.Tick); err != nil {
		return fmt.Errorf("error while marshaling Tick:%w", err)
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
	// Deserialize `PoolState`:
	if err = decoder.Decode(&obj.PoolState); err != nil {
		return fmt.Errorf("error while unmarshaling PoolState:%w", err)
	}
	// Deserialize `Sender`:
	if err = decoder.Decode(&obj.Sender); err != nil {
		return fmt.Errorf("error while unmarshaling Sender:%w", err)
	}
	// Deserialize `TokenAccount0`:
	if err = decoder.Decode(&obj.TokenAccount0); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAccount0:%w", err)
	}
	// Deserialize `TokenAccount1`:
	if err = decoder.Decode(&obj.TokenAccount1); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAccount1:%w", err)
	}
	// Deserialize `Amount0`:
	if err = decoder.Decode(&obj.Amount0); err != nil {
		return fmt.Errorf("error while unmarshaling Amount0:%w", err)
	}
	// Deserialize `TransferFee0`:
	if err = decoder.Decode(&obj.TransferFee0); err != nil {
		return fmt.Errorf("error while unmarshaling TransferFee0:%w", err)
	}
	// Deserialize `Amount1`:
	if err = decoder.Decode(&obj.Amount1); err != nil {
		return fmt.Errorf("error while unmarshaling Amount1:%w", err)
	}
	// Deserialize `TransferFee1`:
	if err = decoder.Decode(&obj.TransferFee1); err != nil {
		return fmt.Errorf("error while unmarshaling TransferFee1:%w", err)
	}
	// Deserialize `ZeroForOne`:
	if err = decoder.Decode(&obj.ZeroForOne); err != nil {
		return fmt.Errorf("error while unmarshaling ZeroForOne:%w", err)
	}
	// Deserialize `SqrtPriceX64`:
	if err = decoder.Decode(&obj.SqrtPriceX64); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtPriceX64:%w", err)
	}
	// Deserialize `Liquidity`:
	if err = decoder.Decode(&obj.Liquidity); err != nil {
		return fmt.Errorf("error while unmarshaling Liquidity:%w", err)
	}
	// Deserialize `Tick`:
	if err = decoder.Decode(&obj.Tick); err != nil {
		return fmt.Errorf("error while unmarshaling Tick:%w", err)
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

// Emitted when Reward are updated for a pool
type UpdateRewardInfosEvent struct {
	// Reward info
	RewardGrowthGlobalX64 [3]binary.Uint128
}

func (obj UpdateRewardInfosEvent) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Event_UpdateRewardInfosEvent[:], false)
	if err != nil {
		return err
	}
	// Serialize `RewardGrowthGlobalX64`:
	if err = encoder.Encode(obj.RewardGrowthGlobalX64); err != nil {
		return fmt.Errorf("error while marshaling RewardGrowthGlobalX64:%w", err)
	}
	return nil
}

func (obj UpdateRewardInfosEvent) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding UpdateRewardInfosEvent: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *UpdateRewardInfosEvent) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Event_UpdateRewardInfosEvent[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Event_UpdateRewardInfosEvent[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `RewardGrowthGlobalX64`:
	if err = decoder.Decode(&obj.RewardGrowthGlobalX64); err != nil {
		return fmt.Errorf("error while unmarshaling RewardGrowthGlobalX64:%w", err)
	}
	return nil
}

func (obj *UpdateRewardInfosEvent) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling UpdateRewardInfosEvent: %w", err)
	}
	return nil
}

func UnmarshalUpdateRewardInfosEvent(buf []byte) (*UpdateRewardInfosEvent, error) {
	obj := new(UpdateRewardInfosEvent)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
