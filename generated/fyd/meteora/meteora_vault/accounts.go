package meteora_vault

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

// Vault struct
type Vault struct {
	// The flag, if admin set enable = false, then the user can only withdraw and cannot deposit in the vault.
	Enabled uint8

	// Vault nonce, to create vault seeds
	Bumps VaultBumps

	// The total liquidity of the vault, including remaining tokens in token_vault and the liquidity in all strategies.
	TotalAmount uint64

	// Token account, hold liquidity in vault reserve
	TokenVault solanago.PublicKey

	// Hold lp token of vault, each time rebalance crank is called, vault calculate performance fee and mint corresponding lp token amount to fee_vault. fee_vault is owned by treasury address
	FeeVault solanago.PublicKey

	// Token mint that vault supports
	TokenMint solanago.PublicKey

	// Lp mint of vault
	LpMint solanago.PublicKey

	// The list of strategy addresses that vault supports, vault can support up to MAX_STRATEGY strategies at the same time.
	Strategies [30]solanago.PublicKey

	// The base address to create vault seeds
	Base solanago.PublicKey

	// Admin of vault
	Admin solanago.PublicKey

	// Person who can send the crank. Operator can only send liquidity to strategies that admin defined, and claim reward to account of treasury address
	Operator solanago.PublicKey

	// Stores information for locked profit.
	LockedProfitTracker LockedProfitTracker
}

func (obj Vault) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Vault[:], false)
	if err != nil {
		return err
	}
	// Serialize `Enabled`:
	if err = encoder.Encode(obj.Enabled); err != nil {
		return fmt.Errorf("error while marshaling Enabled:%w", err)
	}
	// Serialize `Bumps`:
	if err = encoder.Encode(obj.Bumps); err != nil {
		return fmt.Errorf("error while marshaling Bumps:%w", err)
	}
	// Serialize `TotalAmount`:
	if err = encoder.Encode(obj.TotalAmount); err != nil {
		return fmt.Errorf("error while marshaling TotalAmount:%w", err)
	}
	// Serialize `TokenVault`:
	if err = encoder.Encode(obj.TokenVault); err != nil {
		return fmt.Errorf("error while marshaling TokenVault:%w", err)
	}
	// Serialize `FeeVault`:
	if err = encoder.Encode(obj.FeeVault); err != nil {
		return fmt.Errorf("error while marshaling FeeVault:%w", err)
	}
	// Serialize `TokenMint`:
	if err = encoder.Encode(obj.TokenMint); err != nil {
		return fmt.Errorf("error while marshaling TokenMint:%w", err)
	}
	// Serialize `LpMint`:
	if err = encoder.Encode(obj.LpMint); err != nil {
		return fmt.Errorf("error while marshaling LpMint:%w", err)
	}
	// Serialize `Strategies`:
	if err = encoder.Encode(obj.Strategies); err != nil {
		return fmt.Errorf("error while marshaling Strategies:%w", err)
	}
	// Serialize `Base`:
	if err = encoder.Encode(obj.Base); err != nil {
		return fmt.Errorf("error while marshaling Base:%w", err)
	}
	// Serialize `Admin`:
	if err = encoder.Encode(obj.Admin); err != nil {
		return fmt.Errorf("error while marshaling Admin:%w", err)
	}
	// Serialize `Operator`:
	if err = encoder.Encode(obj.Operator); err != nil {
		return fmt.Errorf("error while marshaling Operator:%w", err)
	}
	// Serialize `LockedProfitTracker`:
	if err = encoder.Encode(obj.LockedProfitTracker); err != nil {
		return fmt.Errorf("error while marshaling LockedProfitTracker:%w", err)
	}
	return nil
}

func (obj Vault) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Vault: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Vault) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Vault[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Vault[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Enabled`:
	if err = decoder.Decode(&obj.Enabled); err != nil {
		return fmt.Errorf("error while unmarshaling Enabled:%w", err)
	}
	// Deserialize `Bumps`:
	if err = decoder.Decode(&obj.Bumps); err != nil {
		return fmt.Errorf("error while unmarshaling Bumps:%w", err)
	}
	// Deserialize `TotalAmount`:
	if err = decoder.Decode(&obj.TotalAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TotalAmount:%w", err)
	}
	// Deserialize `TokenVault`:
	if err = decoder.Decode(&obj.TokenVault); err != nil {
		return fmt.Errorf("error while unmarshaling TokenVault:%w", err)
	}
	// Deserialize `FeeVault`:
	if err = decoder.Decode(&obj.FeeVault); err != nil {
		return fmt.Errorf("error while unmarshaling FeeVault:%w", err)
	}
	// Deserialize `TokenMint`:
	if err = decoder.Decode(&obj.TokenMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMint:%w", err)
	}
	// Deserialize `LpMint`:
	if err = decoder.Decode(&obj.LpMint); err != nil {
		return fmt.Errorf("error while unmarshaling LpMint:%w", err)
	}
	// Deserialize `Strategies`:
	if err = decoder.Decode(&obj.Strategies); err != nil {
		return fmt.Errorf("error while unmarshaling Strategies:%w", err)
	}
	// Deserialize `Base`:
	if err = decoder.Decode(&obj.Base); err != nil {
		return fmt.Errorf("error while unmarshaling Base:%w", err)
	}
	// Deserialize `Admin`:
	if err = decoder.Decode(&obj.Admin); err != nil {
		return fmt.Errorf("error while unmarshaling Admin:%w", err)
	}
	// Deserialize `Operator`:
	if err = decoder.Decode(&obj.Operator); err != nil {
		return fmt.Errorf("error while unmarshaling Operator:%w", err)
	}
	// Deserialize `LockedProfitTracker`:
	if err = decoder.Decode(&obj.LockedProfitTracker); err != nil {
		return fmt.Errorf("error while unmarshaling LockedProfitTracker:%w", err)
	}
	return nil
}

func (obj *Vault) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Vault: %w", err)
	}
	return nil
}

func UnmarshalVault(buf []byte) (*Vault, error) {
	obj := new(Vault)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Strategy struct
type Strategy struct {
	// Lending pool address, that the strategy will deposit/withdraw balance
	Reserve solanago.PublicKey

	// The token account, that holds the collateral token
	CollateralVault solanago.PublicKey

	// Specify type of strategy
	StrategyType StrategyType

	// The liquidity in strategy at the time vault deposit/withdraw from a lending protocol
	CurrentLiquidity uint64

	// Hold some bumps, in case the strategy needs to use other seeds to sign a CPI call.
	Bumps [10]uint8

	// Vault address, that the strategy belongs
	Vault solanago.PublicKey

	// If we remove strategy by remove_strategy2 endpoint, this account will be never added again
	IsDisable uint8
}

func (obj Strategy) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Strategy[:], false)
	if err != nil {
		return err
	}
	// Serialize `Reserve`:
	if err = encoder.Encode(obj.Reserve); err != nil {
		return fmt.Errorf("error while marshaling Reserve:%w", err)
	}
	// Serialize `CollateralVault`:
	if err = encoder.Encode(obj.CollateralVault); err != nil {
		return fmt.Errorf("error while marshaling CollateralVault:%w", err)
	}
	// Serialize `StrategyType`:
	if err = encoder.Encode(obj.StrategyType); err != nil {
		return fmt.Errorf("error while marshaling StrategyType:%w", err)
	}
	// Serialize `CurrentLiquidity`:
	if err = encoder.Encode(obj.CurrentLiquidity); err != nil {
		return fmt.Errorf("error while marshaling CurrentLiquidity:%w", err)
	}
	// Serialize `Bumps`:
	if err = encoder.Encode(obj.Bumps); err != nil {
		return fmt.Errorf("error while marshaling Bumps:%w", err)
	}
	// Serialize `Vault`:
	if err = encoder.Encode(obj.Vault); err != nil {
		return fmt.Errorf("error while marshaling Vault:%w", err)
	}
	// Serialize `IsDisable`:
	if err = encoder.Encode(obj.IsDisable); err != nil {
		return fmt.Errorf("error while marshaling IsDisable:%w", err)
	}
	return nil
}

func (obj Strategy) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Strategy: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Strategy) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Strategy[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Strategy[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Reserve`:
	if err = decoder.Decode(&obj.Reserve); err != nil {
		return fmt.Errorf("error while unmarshaling Reserve:%w", err)
	}
	// Deserialize `CollateralVault`:
	if err = decoder.Decode(&obj.CollateralVault); err != nil {
		return fmt.Errorf("error while unmarshaling CollateralVault:%w", err)
	}
	// Deserialize `StrategyType`:
	if err = decoder.Decode(&obj.StrategyType); err != nil {
		return fmt.Errorf("error while unmarshaling StrategyType:%w", err)
	}
	// Deserialize `CurrentLiquidity`:
	if err = decoder.Decode(&obj.CurrentLiquidity); err != nil {
		return fmt.Errorf("error while unmarshaling CurrentLiquidity:%w", err)
	}
	// Deserialize `Bumps`:
	if err = decoder.Decode(&obj.Bumps); err != nil {
		return fmt.Errorf("error while unmarshaling Bumps:%w", err)
	}
	// Deserialize `Vault`:
	if err = decoder.Decode(&obj.Vault); err != nil {
		return fmt.Errorf("error while unmarshaling Vault:%w", err)
	}
	// Deserialize `IsDisable`:
	if err = decoder.Decode(&obj.IsDisable); err != nil {
		return fmt.Errorf("error while unmarshaling IsDisable:%w", err)
	}
	return nil
}

func (obj *Strategy) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Strategy: %w", err)
	}
	return nil
}

func UnmarshalStrategy(buf []byte) (*Strategy, error) {
	obj := new(Strategy)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
