package meteora_pools

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type Config struct {
	PoolFees           PoolFees
	ActivationDuration uint64
	VaultConfigKey     solanago.PublicKey

	// Only pool_creator_authority can use the current config to initialize new pool. When it's Pubkey::default, it's a public config.
	PoolCreatorAuthority solanago.PublicKey

	// Activation type
	ActivationType      uint8
	PartnerFeeNumerator uint64
	Padding             [219]uint8
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
	// Pool address
	Pool solanago.PublicKey

	// Owner address
	Owner solanago.PublicKey

	// Vault address, store the lock user lock
	EscrowVault solanago.PublicKey

	// bump, used to sign
	Bump uint8

	// Total locked amount
	TotalLockedAmount uint64

	// Lp per token, virtual price of lp token
	LpPerToken binary.Uint128

	// Unclaimed fee pending
	UnclaimedFeePending uint64

	// Total a fee claimed so far
	AFee uint64

	// Total b fee claimed so far
	BFee uint64
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

// State of pool account
type Pool struct {
	// LP token mint of the pool
	LpMint solanago.PublicKey

	// Token A mint of the pool. Eg: USDT
	TokenAMint solanago.PublicKey

	// Token B mint of the pool. Eg: USDC
	TokenBMint solanago.PublicKey

	// Vault account for token A. Token A of the pool will be deposit / withdraw from this vault account.
	AVault solanago.PublicKey

	// Vault account for token B. Token B of the pool will be deposit / withdraw from this vault account.
	BVault solanago.PublicKey

	// LP token account of vault A. Used to receive/burn the vault LP upon deposit/withdraw from the vault.
	AVaultLp solanago.PublicKey

	// LP token account of vault B. Used to receive/burn the vault LP upon deposit/withdraw from the vault.
	BVaultLp solanago.PublicKey

	// "A" vault lp bump. Used to create signer seeds.
	AVaultLpBump uint8

	// Flag to determine whether the pool is enabled, or disabled.
	Enabled bool

	// Protocol fee token account for token A. Used to receive trading fee.
	ProtocolTokenAFee solanago.PublicKey

	// Protocol fee token account for token B. Used to receive trading fee.
	ProtocolTokenBFee solanago.PublicKey

	// Fee last updated timestamp
	FeeLastUpdatedAt uint64
	Padding0         [24]uint8

	// Store the fee charges setting.
	Fees PoolFees

	// Pool type
	PoolType PoolType

	// Stake pubkey of SPL stake pool
	Stake solanago.PublicKey

	// Total locked lp token
	TotalLockedLp uint64

	// bootstrapping config
	Bootstrapping Bootstrapping
	PartnerInfo   PartnerInfo

	// Padding for future pool field
	Padding Padding

	// The type of the swap curve supported by the pool.
	CurveType CurveType
}

func (obj Pool) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Pool[:], false)
	if err != nil {
		return err
	}
	// Serialize `LpMint`:
	if err = encoder.Encode(obj.LpMint); err != nil {
		return fmt.Errorf("error while marshaling LpMint:%w", err)
	}
	// Serialize `TokenAMint`:
	if err = encoder.Encode(obj.TokenAMint); err != nil {
		return fmt.Errorf("error while marshaling TokenAMint:%w", err)
	}
	// Serialize `TokenBMint`:
	if err = encoder.Encode(obj.TokenBMint); err != nil {
		return fmt.Errorf("error while marshaling TokenBMint:%w", err)
	}
	// Serialize `AVault`:
	if err = encoder.Encode(obj.AVault); err != nil {
		return fmt.Errorf("error while marshaling AVault:%w", err)
	}
	// Serialize `BVault`:
	if err = encoder.Encode(obj.BVault); err != nil {
		return fmt.Errorf("error while marshaling BVault:%w", err)
	}
	// Serialize `AVaultLp`:
	if err = encoder.Encode(obj.AVaultLp); err != nil {
		return fmt.Errorf("error while marshaling AVaultLp:%w", err)
	}
	// Serialize `BVaultLp`:
	if err = encoder.Encode(obj.BVaultLp); err != nil {
		return fmt.Errorf("error while marshaling BVaultLp:%w", err)
	}
	// Serialize `AVaultLpBump`:
	if err = encoder.Encode(obj.AVaultLpBump); err != nil {
		return fmt.Errorf("error while marshaling AVaultLpBump:%w", err)
	}
	// Serialize `Enabled`:
	if err = encoder.Encode(obj.Enabled); err != nil {
		return fmt.Errorf("error while marshaling Enabled:%w", err)
	}
	// Serialize `ProtocolTokenAFee`:
	if err = encoder.Encode(obj.ProtocolTokenAFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolTokenAFee:%w", err)
	}
	// Serialize `ProtocolTokenBFee`:
	if err = encoder.Encode(obj.ProtocolTokenBFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolTokenBFee:%w", err)
	}
	// Serialize `FeeLastUpdatedAt`:
	if err = encoder.Encode(obj.FeeLastUpdatedAt); err != nil {
		return fmt.Errorf("error while marshaling FeeLastUpdatedAt:%w", err)
	}
	// Serialize `Padding0`:
	if err = encoder.Encode(obj.Padding0); err != nil {
		return fmt.Errorf("error while marshaling Padding0:%w", err)
	}
	// Serialize `Fees`:
	if err = encoder.Encode(obj.Fees); err != nil {
		return fmt.Errorf("error while marshaling Fees:%w", err)
	}
	// Serialize `PoolType`:
	if err = encoder.Encode(obj.PoolType); err != nil {
		return fmt.Errorf("error while marshaling PoolType:%w", err)
	}
	// Serialize `Stake`:
	if err = encoder.Encode(obj.Stake); err != nil {
		return fmt.Errorf("error while marshaling Stake:%w", err)
	}
	// Serialize `TotalLockedLp`:
	if err = encoder.Encode(obj.TotalLockedLp); err != nil {
		return fmt.Errorf("error while marshaling TotalLockedLp:%w", err)
	}
	// Serialize `Bootstrapping`:
	if err = encoder.Encode(obj.Bootstrapping); err != nil {
		return fmt.Errorf("error while marshaling Bootstrapping:%w", err)
	}
	// Serialize `PartnerInfo`:
	if err = encoder.Encode(obj.PartnerInfo); err != nil {
		return fmt.Errorf("error while marshaling PartnerInfo:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	// Serialize `CurveType`:
	{
		if err = EncodeCurveType(encoder, obj.CurveType); err != nil {
			return fmt.Errorf("error while marshalingCurveType:%w", err)
		}
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
	// Deserialize `LpMint`:
	if err = decoder.Decode(&obj.LpMint); err != nil {
		return fmt.Errorf("error while unmarshaling LpMint:%w", err)
	}
	// Deserialize `TokenAMint`:
	if err = decoder.Decode(&obj.TokenAMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAMint:%w", err)
	}
	// Deserialize `TokenBMint`:
	if err = decoder.Decode(&obj.TokenBMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBMint:%w", err)
	}
	// Deserialize `AVault`:
	if err = decoder.Decode(&obj.AVault); err != nil {
		return fmt.Errorf("error while unmarshaling AVault:%w", err)
	}
	// Deserialize `BVault`:
	if err = decoder.Decode(&obj.BVault); err != nil {
		return fmt.Errorf("error while unmarshaling BVault:%w", err)
	}
	// Deserialize `AVaultLp`:
	if err = decoder.Decode(&obj.AVaultLp); err != nil {
		return fmt.Errorf("error while unmarshaling AVaultLp:%w", err)
	}
	// Deserialize `BVaultLp`:
	if err = decoder.Decode(&obj.BVaultLp); err != nil {
		return fmt.Errorf("error while unmarshaling BVaultLp:%w", err)
	}
	// Deserialize `AVaultLpBump`:
	if err = decoder.Decode(&obj.AVaultLpBump); err != nil {
		return fmt.Errorf("error while unmarshaling AVaultLpBump:%w", err)
	}
	// Deserialize `Enabled`:
	if err = decoder.Decode(&obj.Enabled); err != nil {
		return fmt.Errorf("error while unmarshaling Enabled:%w", err)
	}
	// Deserialize `ProtocolTokenAFee`:
	if err = decoder.Decode(&obj.ProtocolTokenAFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolTokenAFee:%w", err)
	}
	// Deserialize `ProtocolTokenBFee`:
	if err = decoder.Decode(&obj.ProtocolTokenBFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolTokenBFee:%w", err)
	}
	// Deserialize `FeeLastUpdatedAt`:
	if err = decoder.Decode(&obj.FeeLastUpdatedAt); err != nil {
		return fmt.Errorf("error while unmarshaling FeeLastUpdatedAt:%w", err)
	}
	// Deserialize `Padding0`:
	if err = decoder.Decode(&obj.Padding0); err != nil {
		return fmt.Errorf("error while unmarshaling Padding0:%w", err)
	}
	// Deserialize `Fees`:
	if err = decoder.Decode(&obj.Fees); err != nil {
		return fmt.Errorf("error while unmarshaling Fees:%w", err)
	}
	// Deserialize `PoolType`:
	if err = decoder.Decode(&obj.PoolType); err != nil {
		return fmt.Errorf("error while unmarshaling PoolType:%w", err)
	}
	// Deserialize `Stake`:
	if err = decoder.Decode(&obj.Stake); err != nil {
		return fmt.Errorf("error while unmarshaling Stake:%w", err)
	}
	// Deserialize `TotalLockedLp`:
	if err = decoder.Decode(&obj.TotalLockedLp); err != nil {
		return fmt.Errorf("error while unmarshaling TotalLockedLp:%w", err)
	}
	// Deserialize `Bootstrapping`:
	if err = decoder.Decode(&obj.Bootstrapping); err != nil {
		return fmt.Errorf("error while unmarshaling Bootstrapping:%w", err)
	}
	// Deserialize `PartnerInfo`:
	if err = decoder.Decode(&obj.PartnerInfo); err != nil {
		return fmt.Errorf("error while unmarshaling PartnerInfo:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	// Deserialize `CurveType`:
	{
		var err error
		obj.CurveType, err = DecodeCurveType(decoder)
		if err != nil {
			return err
		}
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
