package meteora_damm

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
	// Vault config key
	VaultConfigKey solanago.PublicKey

	// Only pool_creator_authority can use the current config to initialize new pool. When it's Pubkey::default, it's a public config.
	PoolCreatorAuthority solanago.PublicKey

	// Pool fee
	PoolFees PoolFeesConfig

	// Activation type
	ActivationType uint8

	// Collect fee mode
	CollectFeeMode uint8

	// Config type mode, 0 for static, 1 for dynamic
	ConfigType uint8

	// padding 0
	Padding0 [5]uint8

	// config index
	Index uint64

	// sqrt min price
	SqrtMinPrice binary.Uint128

	// sqrt max price
	SqrtMaxPrice binary.Uint128

	// Fee curve point
	// Padding for further use
	Padding1 [10]uint64
}

func (obj Config) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Config[:], false)
	if err != nil {
		return err
	}
	// Serialize `VaultConfigKey`:
	if err = encoder.Encode(obj.VaultConfigKey); err != nil {
		return fmt.Errorf("error while marshaling VaultConfigKey:%w", err)
	}
	// Serialize `PoolCreatorAuthority`:
	if err = encoder.Encode(obj.PoolCreatorAuthority); err != nil {
		return fmt.Errorf("error while marshaling PoolCreatorAuthority:%w", err)
	}
	// Serialize `PoolFees`:
	if err = encoder.Encode(obj.PoolFees); err != nil {
		return fmt.Errorf("error while marshaling PoolFees:%w", err)
	}
	// Serialize `ActivationType`:
	if err = encoder.Encode(obj.ActivationType); err != nil {
		return fmt.Errorf("error while marshaling ActivationType:%w", err)
	}
	// Serialize `CollectFeeMode`:
	if err = encoder.Encode(obj.CollectFeeMode); err != nil {
		return fmt.Errorf("error while marshaling CollectFeeMode:%w", err)
	}
	// Serialize `ConfigType`:
	if err = encoder.Encode(obj.ConfigType); err != nil {
		return fmt.Errorf("error while marshaling ConfigType:%w", err)
	}
	// Serialize `Padding0`:
	if err = encoder.Encode(obj.Padding0); err != nil {
		return fmt.Errorf("error while marshaling Padding0:%w", err)
	}
	// Serialize `Index`:
	if err = encoder.Encode(obj.Index); err != nil {
		return fmt.Errorf("error while marshaling Index:%w", err)
	}
	// Serialize `SqrtMinPrice`:
	if err = encoder.Encode(obj.SqrtMinPrice); err != nil {
		return fmt.Errorf("error while marshaling SqrtMinPrice:%w", err)
	}
	// Serialize `SqrtMaxPrice`:
	if err = encoder.Encode(obj.SqrtMaxPrice); err != nil {
		return fmt.Errorf("error while marshaling SqrtMaxPrice:%w", err)
	}
	// Serialize `Padding1`:
	if err = encoder.Encode(obj.Padding1); err != nil {
		return fmt.Errorf("error while marshaling Padding1:%w", err)
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
	// Deserialize `VaultConfigKey`:
	if err = decoder.Decode(&obj.VaultConfigKey); err != nil {
		return fmt.Errorf("error while unmarshaling VaultConfigKey:%w", err)
	}
	// Deserialize `PoolCreatorAuthority`:
	if err = decoder.Decode(&obj.PoolCreatorAuthority); err != nil {
		return fmt.Errorf("error while unmarshaling PoolCreatorAuthority:%w", err)
	}
	// Deserialize `PoolFees`:
	if err = decoder.Decode(&obj.PoolFees); err != nil {
		return fmt.Errorf("error while unmarshaling PoolFees:%w", err)
	}
	// Deserialize `ActivationType`:
	if err = decoder.Decode(&obj.ActivationType); err != nil {
		return fmt.Errorf("error while unmarshaling ActivationType:%w", err)
	}
	// Deserialize `CollectFeeMode`:
	if err = decoder.Decode(&obj.CollectFeeMode); err != nil {
		return fmt.Errorf("error while unmarshaling CollectFeeMode:%w", err)
	}
	// Deserialize `ConfigType`:
	if err = decoder.Decode(&obj.ConfigType); err != nil {
		return fmt.Errorf("error while unmarshaling ConfigType:%w", err)
	}
	// Deserialize `Padding0`:
	if err = decoder.Decode(&obj.Padding0); err != nil {
		return fmt.Errorf("error while unmarshaling Padding0:%w", err)
	}
	// Deserialize `Index`:
	if err = decoder.Decode(&obj.Index); err != nil {
		return fmt.Errorf("error while unmarshaling Index:%w", err)
	}
	// Deserialize `SqrtMinPrice`:
	if err = decoder.Decode(&obj.SqrtMinPrice); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtMinPrice:%w", err)
	}
	// Deserialize `SqrtMaxPrice`:
	if err = decoder.Decode(&obj.SqrtMaxPrice); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtMaxPrice:%w", err)
	}
	// Deserialize `Padding1`:
	if err = decoder.Decode(&obj.Padding1); err != nil {
		return fmt.Errorf("error while unmarshaling Padding1:%w", err)
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

type Pool struct {
	// Pool fee
	PoolFees PoolFeesStruct

	// token a mint
	TokenAMint solanago.PublicKey

	// token b mint
	TokenBMint solanago.PublicKey

	// token a vault
	TokenAVault solanago.PublicKey

	// token b vault
	TokenBVault solanago.PublicKey

	// Whitelisted vault to be able to buy pool before activation_point
	WhitelistedVault solanago.PublicKey

	// partner
	Partner solanago.PublicKey

	// liquidity share
	Liquidity binary.Uint128

	// padding, previous reserve amount, be careful to use that field
	Padding binary.Uint128

	// protocol a fee
	ProtocolAFee uint64

	// protocol b fee
	ProtocolBFee uint64

	// partner a fee
	PartnerAFee uint64

	// partner b fee
	PartnerBFee uint64

	// min price
	SqrtMinPrice binary.Uint128

	// max price
	SqrtMaxPrice binary.Uint128

	// current price
	SqrtPrice binary.Uint128

	// Activation point, can be slot or timestamp
	ActivationPoint uint64

	// Activation type, 0 means by slot, 1 means by timestamp
	ActivationType uint8

	// pool status, 0: enable, 1 disable
	PoolStatus uint8

	// token a flag
	TokenAFlag uint8

	// token b flag
	TokenBFlag uint8

	// 0 is collect fee in both token, 1 only collect fee in token a, 2 only collect fee in token b
	CollectFeeMode uint8

	// pool type
	PoolType uint8

	// padding
	Padding0 [2]uint8

	// cumulative
	FeeAPerLiquidity [32]uint8

	// cumulative
	FeeBPerLiquidity       [32]uint8
	PermanentLockLiquidity binary.Uint128

	// metrics
	Metrics PoolMetrics

	// pool creator
	Creator solanago.PublicKey

	// Padding for further use
	Padding1 [6]uint64

	// Farming reward information
	RewardInfos [2]RewardInfo
}

func (obj Pool) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Pool[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolFees`:
	if err = encoder.Encode(obj.PoolFees); err != nil {
		return fmt.Errorf("error while marshaling PoolFees:%w", err)
	}
	// Serialize `TokenAMint`:
	if err = encoder.Encode(obj.TokenAMint); err != nil {
		return fmt.Errorf("error while marshaling TokenAMint:%w", err)
	}
	// Serialize `TokenBMint`:
	if err = encoder.Encode(obj.TokenBMint); err != nil {
		return fmt.Errorf("error while marshaling TokenBMint:%w", err)
	}
	// Serialize `TokenAVault`:
	if err = encoder.Encode(obj.TokenAVault); err != nil {
		return fmt.Errorf("error while marshaling TokenAVault:%w", err)
	}
	// Serialize `TokenBVault`:
	if err = encoder.Encode(obj.TokenBVault); err != nil {
		return fmt.Errorf("error while marshaling TokenBVault:%w", err)
	}
	// Serialize `WhitelistedVault`:
	if err = encoder.Encode(obj.WhitelistedVault); err != nil {
		return fmt.Errorf("error while marshaling WhitelistedVault:%w", err)
	}
	// Serialize `Partner`:
	if err = encoder.Encode(obj.Partner); err != nil {
		return fmt.Errorf("error while marshaling Partner:%w", err)
	}
	// Serialize `Liquidity`:
	if err = encoder.Encode(obj.Liquidity); err != nil {
		return fmt.Errorf("error while marshaling Liquidity:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	// Serialize `ProtocolAFee`:
	if err = encoder.Encode(obj.ProtocolAFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolAFee:%w", err)
	}
	// Serialize `ProtocolBFee`:
	if err = encoder.Encode(obj.ProtocolBFee); err != nil {
		return fmt.Errorf("error while marshaling ProtocolBFee:%w", err)
	}
	// Serialize `PartnerAFee`:
	if err = encoder.Encode(obj.PartnerAFee); err != nil {
		return fmt.Errorf("error while marshaling PartnerAFee:%w", err)
	}
	// Serialize `PartnerBFee`:
	if err = encoder.Encode(obj.PartnerBFee); err != nil {
		return fmt.Errorf("error while marshaling PartnerBFee:%w", err)
	}
	// Serialize `SqrtMinPrice`:
	if err = encoder.Encode(obj.SqrtMinPrice); err != nil {
		return fmt.Errorf("error while marshaling SqrtMinPrice:%w", err)
	}
	// Serialize `SqrtMaxPrice`:
	if err = encoder.Encode(obj.SqrtMaxPrice); err != nil {
		return fmt.Errorf("error while marshaling SqrtMaxPrice:%w", err)
	}
	// Serialize `SqrtPrice`:
	if err = encoder.Encode(obj.SqrtPrice); err != nil {
		return fmt.Errorf("error while marshaling SqrtPrice:%w", err)
	}
	// Serialize `ActivationPoint`:
	if err = encoder.Encode(obj.ActivationPoint); err != nil {
		return fmt.Errorf("error while marshaling ActivationPoint:%w", err)
	}
	// Serialize `ActivationType`:
	if err = encoder.Encode(obj.ActivationType); err != nil {
		return fmt.Errorf("error while marshaling ActivationType:%w", err)
	}
	// Serialize `PoolStatus`:
	if err = encoder.Encode(obj.PoolStatus); err != nil {
		return fmt.Errorf("error while marshaling PoolStatus:%w", err)
	}
	// Serialize `TokenAFlag`:
	if err = encoder.Encode(obj.TokenAFlag); err != nil {
		return fmt.Errorf("error while marshaling TokenAFlag:%w", err)
	}
	// Serialize `TokenBFlag`:
	if err = encoder.Encode(obj.TokenBFlag); err != nil {
		return fmt.Errorf("error while marshaling TokenBFlag:%w", err)
	}
	// Serialize `CollectFeeMode`:
	if err = encoder.Encode(obj.CollectFeeMode); err != nil {
		return fmt.Errorf("error while marshaling CollectFeeMode:%w", err)
	}
	// Serialize `PoolType`:
	if err = encoder.Encode(obj.PoolType); err != nil {
		return fmt.Errorf("error while marshaling PoolType:%w", err)
	}
	// Serialize `Padding0`:
	if err = encoder.Encode(obj.Padding0); err != nil {
		return fmt.Errorf("error while marshaling Padding0:%w", err)
	}
	// Serialize `FeeAPerLiquidity`:
	if err = encoder.Encode(obj.FeeAPerLiquidity); err != nil {
		return fmt.Errorf("error while marshaling FeeAPerLiquidity:%w", err)
	}
	// Serialize `FeeBPerLiquidity`:
	if err = encoder.Encode(obj.FeeBPerLiquidity); err != nil {
		return fmt.Errorf("error while marshaling FeeBPerLiquidity:%w", err)
	}
	// Serialize `PermanentLockLiquidity`:
	if err = encoder.Encode(obj.PermanentLockLiquidity); err != nil {
		return fmt.Errorf("error while marshaling PermanentLockLiquidity:%w", err)
	}
	// Serialize `Metrics`:
	if err = encoder.Encode(obj.Metrics); err != nil {
		return fmt.Errorf("error while marshaling Metrics:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
	}
	// Serialize `Padding1`:
	if err = encoder.Encode(obj.Padding1); err != nil {
		return fmt.Errorf("error while marshaling Padding1:%w", err)
	}
	// Serialize `RewardInfos`:
	if err = encoder.Encode(obj.RewardInfos); err != nil {
		return fmt.Errorf("error while marshaling RewardInfos:%w", err)
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
	// Deserialize `PoolFees`:
	if err = decoder.Decode(&obj.PoolFees); err != nil {
		return fmt.Errorf("error while unmarshaling PoolFees:%w", err)
	}
	// Deserialize `TokenAMint`:
	if err = decoder.Decode(&obj.TokenAMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAMint:%w", err)
	}
	// Deserialize `TokenBMint`:
	if err = decoder.Decode(&obj.TokenBMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBMint:%w", err)
	}
	// Deserialize `TokenAVault`:
	if err = decoder.Decode(&obj.TokenAVault); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAVault:%w", err)
	}
	// Deserialize `TokenBVault`:
	if err = decoder.Decode(&obj.TokenBVault); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBVault:%w", err)
	}
	// Deserialize `WhitelistedVault`:
	if err = decoder.Decode(&obj.WhitelistedVault); err != nil {
		return fmt.Errorf("error while unmarshaling WhitelistedVault:%w", err)
	}
	// Deserialize `Partner`:
	if err = decoder.Decode(&obj.Partner); err != nil {
		return fmt.Errorf("error while unmarshaling Partner:%w", err)
	}
	// Deserialize `Liquidity`:
	if err = decoder.Decode(&obj.Liquidity); err != nil {
		return fmt.Errorf("error while unmarshaling Liquidity:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	// Deserialize `ProtocolAFee`:
	if err = decoder.Decode(&obj.ProtocolAFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolAFee:%w", err)
	}
	// Deserialize `ProtocolBFee`:
	if err = decoder.Decode(&obj.ProtocolBFee); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolBFee:%w", err)
	}
	// Deserialize `PartnerAFee`:
	if err = decoder.Decode(&obj.PartnerAFee); err != nil {
		return fmt.Errorf("error while unmarshaling PartnerAFee:%w", err)
	}
	// Deserialize `PartnerBFee`:
	if err = decoder.Decode(&obj.PartnerBFee); err != nil {
		return fmt.Errorf("error while unmarshaling PartnerBFee:%w", err)
	}
	// Deserialize `SqrtMinPrice`:
	if err = decoder.Decode(&obj.SqrtMinPrice); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtMinPrice:%w", err)
	}
	// Deserialize `SqrtMaxPrice`:
	if err = decoder.Decode(&obj.SqrtMaxPrice); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtMaxPrice:%w", err)
	}
	// Deserialize `SqrtPrice`:
	if err = decoder.Decode(&obj.SqrtPrice); err != nil {
		return fmt.Errorf("error while unmarshaling SqrtPrice:%w", err)
	}
	// Deserialize `ActivationPoint`:
	if err = decoder.Decode(&obj.ActivationPoint); err != nil {
		return fmt.Errorf("error while unmarshaling ActivationPoint:%w", err)
	}
	// Deserialize `ActivationType`:
	if err = decoder.Decode(&obj.ActivationType); err != nil {
		return fmt.Errorf("error while unmarshaling ActivationType:%w", err)
	}
	// Deserialize `PoolStatus`:
	if err = decoder.Decode(&obj.PoolStatus); err != nil {
		return fmt.Errorf("error while unmarshaling PoolStatus:%w", err)
	}
	// Deserialize `TokenAFlag`:
	if err = decoder.Decode(&obj.TokenAFlag); err != nil {
		return fmt.Errorf("error while unmarshaling TokenAFlag:%w", err)
	}
	// Deserialize `TokenBFlag`:
	if err = decoder.Decode(&obj.TokenBFlag); err != nil {
		return fmt.Errorf("error while unmarshaling TokenBFlag:%w", err)
	}
	// Deserialize `CollectFeeMode`:
	if err = decoder.Decode(&obj.CollectFeeMode); err != nil {
		return fmt.Errorf("error while unmarshaling CollectFeeMode:%w", err)
	}
	// Deserialize `PoolType`:
	if err = decoder.Decode(&obj.PoolType); err != nil {
		return fmt.Errorf("error while unmarshaling PoolType:%w", err)
	}
	// Deserialize `Padding0`:
	if err = decoder.Decode(&obj.Padding0); err != nil {
		return fmt.Errorf("error while unmarshaling Padding0:%w", err)
	}
	// Deserialize `FeeAPerLiquidity`:
	if err = decoder.Decode(&obj.FeeAPerLiquidity); err != nil {
		return fmt.Errorf("error while unmarshaling FeeAPerLiquidity:%w", err)
	}
	// Deserialize `FeeBPerLiquidity`:
	if err = decoder.Decode(&obj.FeeBPerLiquidity); err != nil {
		return fmt.Errorf("error while unmarshaling FeeBPerLiquidity:%w", err)
	}
	// Deserialize `PermanentLockLiquidity`:
	if err = decoder.Decode(&obj.PermanentLockLiquidity); err != nil {
		return fmt.Errorf("error while unmarshaling PermanentLockLiquidity:%w", err)
	}
	// Deserialize `Metrics`:
	if err = decoder.Decode(&obj.Metrics); err != nil {
		return fmt.Errorf("error while unmarshaling Metrics:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
	}
	// Deserialize `Padding1`:
	if err = decoder.Decode(&obj.Padding1); err != nil {
		return fmt.Errorf("error while unmarshaling Padding1:%w", err)
	}
	// Deserialize `RewardInfos`:
	if err = decoder.Decode(&obj.RewardInfos); err != nil {
		return fmt.Errorf("error while unmarshaling RewardInfos:%w", err)
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

type Position struct {
	Pool solanago.PublicKey

	// nft mint
	NftMint solanago.PublicKey

	// fee a checkpoint
	FeeAPerTokenCheckpoint [32]uint8

	// fee b checkpoint
	FeeBPerTokenCheckpoint [32]uint8

	// fee a pending
	FeeAPending uint64

	// fee b pending
	FeeBPending uint64

	// unlock liquidity
	UnlockedLiquidity binary.Uint128

	// vesting liquidity
	VestedLiquidity binary.Uint128

	// permanent locked liquidity
	PermanentLockedLiquidity binary.Uint128

	// metrics
	Metrics PositionMetrics

	// Farming reward information
	RewardInfos [2]UserRewardInfo

	// padding for future usage
	Padding [6]binary.Uint128
}

func (obj Position) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Position[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `NftMint`:
	if err = encoder.Encode(obj.NftMint); err != nil {
		return fmt.Errorf("error while marshaling NftMint:%w", err)
	}
	// Serialize `FeeAPerTokenCheckpoint`:
	if err = encoder.Encode(obj.FeeAPerTokenCheckpoint); err != nil {
		return fmt.Errorf("error while marshaling FeeAPerTokenCheckpoint:%w", err)
	}
	// Serialize `FeeBPerTokenCheckpoint`:
	if err = encoder.Encode(obj.FeeBPerTokenCheckpoint); err != nil {
		return fmt.Errorf("error while marshaling FeeBPerTokenCheckpoint:%w", err)
	}
	// Serialize `FeeAPending`:
	if err = encoder.Encode(obj.FeeAPending); err != nil {
		return fmt.Errorf("error while marshaling FeeAPending:%w", err)
	}
	// Serialize `FeeBPending`:
	if err = encoder.Encode(obj.FeeBPending); err != nil {
		return fmt.Errorf("error while marshaling FeeBPending:%w", err)
	}
	// Serialize `UnlockedLiquidity`:
	if err = encoder.Encode(obj.UnlockedLiquidity); err != nil {
		return fmt.Errorf("error while marshaling UnlockedLiquidity:%w", err)
	}
	// Serialize `VestedLiquidity`:
	if err = encoder.Encode(obj.VestedLiquidity); err != nil {
		return fmt.Errorf("error while marshaling VestedLiquidity:%w", err)
	}
	// Serialize `PermanentLockedLiquidity`:
	if err = encoder.Encode(obj.PermanentLockedLiquidity); err != nil {
		return fmt.Errorf("error while marshaling PermanentLockedLiquidity:%w", err)
	}
	// Serialize `Metrics`:
	if err = encoder.Encode(obj.Metrics); err != nil {
		return fmt.Errorf("error while marshaling Metrics:%w", err)
	}
	// Serialize `RewardInfos`:
	if err = encoder.Encode(obj.RewardInfos); err != nil {
		return fmt.Errorf("error while marshaling RewardInfos:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj Position) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Position: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Position) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Position[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Position[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `NftMint`:
	if err = decoder.Decode(&obj.NftMint); err != nil {
		return fmt.Errorf("error while unmarshaling NftMint:%w", err)
	}
	// Deserialize `FeeAPerTokenCheckpoint`:
	if err = decoder.Decode(&obj.FeeAPerTokenCheckpoint); err != nil {
		return fmt.Errorf("error while unmarshaling FeeAPerTokenCheckpoint:%w", err)
	}
	// Deserialize `FeeBPerTokenCheckpoint`:
	if err = decoder.Decode(&obj.FeeBPerTokenCheckpoint); err != nil {
		return fmt.Errorf("error while unmarshaling FeeBPerTokenCheckpoint:%w", err)
	}
	// Deserialize `FeeAPending`:
	if err = decoder.Decode(&obj.FeeAPending); err != nil {
		return fmt.Errorf("error while unmarshaling FeeAPending:%w", err)
	}
	// Deserialize `FeeBPending`:
	if err = decoder.Decode(&obj.FeeBPending); err != nil {
		return fmt.Errorf("error while unmarshaling FeeBPending:%w", err)
	}
	// Deserialize `UnlockedLiquidity`:
	if err = decoder.Decode(&obj.UnlockedLiquidity); err != nil {
		return fmt.Errorf("error while unmarshaling UnlockedLiquidity:%w", err)
	}
	// Deserialize `VestedLiquidity`:
	if err = decoder.Decode(&obj.VestedLiquidity); err != nil {
		return fmt.Errorf("error while unmarshaling VestedLiquidity:%w", err)
	}
	// Deserialize `PermanentLockedLiquidity`:
	if err = decoder.Decode(&obj.PermanentLockedLiquidity); err != nil {
		return fmt.Errorf("error while unmarshaling PermanentLockedLiquidity:%w", err)
	}
	// Deserialize `Metrics`:
	if err = decoder.Decode(&obj.Metrics); err != nil {
		return fmt.Errorf("error while unmarshaling Metrics:%w", err)
	}
	// Deserialize `RewardInfos`:
	if err = decoder.Decode(&obj.RewardInfos); err != nil {
		return fmt.Errorf("error while unmarshaling RewardInfos:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *Position) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Position: %w", err)
	}
	return nil
}

func UnmarshalPosition(buf []byte) (*Position, error) {
	obj := new(Position)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Parameter that set by the protocol
type TokenBadge struct {
	// token mint
	TokenMint solanago.PublicKey

	// Reserve
	Padding [128]uint8
}

func (obj TokenBadge) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_TokenBadge[:], false)
	if err != nil {
		return err
	}
	// Serialize `TokenMint`:
	if err = encoder.Encode(obj.TokenMint); err != nil {
		return fmt.Errorf("error while marshaling TokenMint:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj TokenBadge) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding TokenBadge: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *TokenBadge) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_TokenBadge[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_TokenBadge[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TokenMint`:
	if err = decoder.Decode(&obj.TokenMint); err != nil {
		return fmt.Errorf("error while unmarshaling TokenMint:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *TokenBadge) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling TokenBadge: %w", err)
	}
	return nil
}

func UnmarshalTokenBadge(buf []byte) (*TokenBadge, error) {
	obj := new(TokenBadge)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type Vesting struct {
	Position               solanago.PublicKey
	CliffPoint             uint64
	PeriodFrequency        uint64
	CliffUnlockLiquidity   binary.Uint128
	LiquidityPerPeriod     binary.Uint128
	TotalReleasedLiquidity binary.Uint128
	NumberOfPeriod         uint16
	Padding                [14]uint8
	Padding2               [4]binary.Uint128
}

func (obj Vesting) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_Vesting[:], false)
	if err != nil {
		return err
	}
	// Serialize `Position`:
	if err = encoder.Encode(obj.Position); err != nil {
		return fmt.Errorf("error while marshaling Position:%w", err)
	}
	// Serialize `CliffPoint`:
	if err = encoder.Encode(obj.CliffPoint); err != nil {
		return fmt.Errorf("error while marshaling CliffPoint:%w", err)
	}
	// Serialize `PeriodFrequency`:
	if err = encoder.Encode(obj.PeriodFrequency); err != nil {
		return fmt.Errorf("error while marshaling PeriodFrequency:%w", err)
	}
	// Serialize `CliffUnlockLiquidity`:
	if err = encoder.Encode(obj.CliffUnlockLiquidity); err != nil {
		return fmt.Errorf("error while marshaling CliffUnlockLiquidity:%w", err)
	}
	// Serialize `LiquidityPerPeriod`:
	if err = encoder.Encode(obj.LiquidityPerPeriod); err != nil {
		return fmt.Errorf("error while marshaling LiquidityPerPeriod:%w", err)
	}
	// Serialize `TotalReleasedLiquidity`:
	if err = encoder.Encode(obj.TotalReleasedLiquidity); err != nil {
		return fmt.Errorf("error while marshaling TotalReleasedLiquidity:%w", err)
	}
	// Serialize `NumberOfPeriod`:
	if err = encoder.Encode(obj.NumberOfPeriod); err != nil {
		return fmt.Errorf("error while marshaling NumberOfPeriod:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	// Serialize `Padding2`:
	if err = encoder.Encode(obj.Padding2); err != nil {
		return fmt.Errorf("error while marshaling Padding2:%w", err)
	}
	return nil
}

func (obj Vesting) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Vesting: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Vesting) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_Vesting[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_Vesting[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Position`:
	if err = decoder.Decode(&obj.Position); err != nil {
		return fmt.Errorf("error while unmarshaling Position:%w", err)
	}
	// Deserialize `CliffPoint`:
	if err = decoder.Decode(&obj.CliffPoint); err != nil {
		return fmt.Errorf("error while unmarshaling CliffPoint:%w", err)
	}
	// Deserialize `PeriodFrequency`:
	if err = decoder.Decode(&obj.PeriodFrequency); err != nil {
		return fmt.Errorf("error while unmarshaling PeriodFrequency:%w", err)
	}
	// Deserialize `CliffUnlockLiquidity`:
	if err = decoder.Decode(&obj.CliffUnlockLiquidity); err != nil {
		return fmt.Errorf("error while unmarshaling CliffUnlockLiquidity:%w", err)
	}
	// Deserialize `LiquidityPerPeriod`:
	if err = decoder.Decode(&obj.LiquidityPerPeriod); err != nil {
		return fmt.Errorf("error while unmarshaling LiquidityPerPeriod:%w", err)
	}
	// Deserialize `TotalReleasedLiquidity`:
	if err = decoder.Decode(&obj.TotalReleasedLiquidity); err != nil {
		return fmt.Errorf("error while unmarshaling TotalReleasedLiquidity:%w", err)
	}
	// Deserialize `NumberOfPeriod`:
	if err = decoder.Decode(&obj.NumberOfPeriod); err != nil {
		return fmt.Errorf("error while unmarshaling NumberOfPeriod:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	// Deserialize `Padding2`:
	if err = decoder.Decode(&obj.Padding2); err != nil {
		return fmt.Errorf("error while unmarshaling Padding2:%w", err)
	}
	return nil
}

func (obj *Vesting) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Vesting: %w", err)
	}
	return nil
}

func UnmarshalVesting(buf []byte) (*Vesting, error) {
	obj := new(Vesting)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
