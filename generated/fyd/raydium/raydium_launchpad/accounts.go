package raydium_launchpad

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

// Holds the current owner of the factory
type GlobalConfig struct {
	// Account update epoch
	Epoch uint64

	// 0: Constant Product Curve
	// 1: Fixed Price Curve
	// 2: Linear Price Curve
	CurveType uint8

	// Config index
	Index uint16

	// The fee of migrate to amm
	MigrateFee uint64

	// The trade fee rate, denominated in hundredths of a bip (10^-6)
	TradeFeeRate uint64

	// The maximum share fee rate, denominated in hundredths of a bip (10^-6)
	MaxShareFeeRate uint64

	// The minimum base supply, the value without decimals
	MinBaseSupply uint64

	// The maximum lock rate, denominated in hundredths of a bip (10^-6)
	MaxLockRate uint64

	// The minimum base sell rate, denominated in hundredths of a bip (10^-6)
	MinBaseSellRate uint64

	// The minimum base migrate rate, denominated in hundredths of a bip (10^-6)
	MinBaseMigrateRate uint64

	// The minimum quote fund raising, the value with decimals
	MinQuoteFundRaising uint64

	// Mint information for quote token
	QuoteMint solanago.PublicKey

	// Protocol Fee owner
	ProtocolFeeOwner solanago.PublicKey

	// Migrate Fee owner
	MigrateFeeOwner solanago.PublicKey

	// Migrate to amm control wallet
	MigrateToAmmWallet solanago.PublicKey

	// Migrate to cpswap wallet
	MigrateToCpswapWallet solanago.PublicKey

	// padding for future updates
	Padding [16]uint64
}

func (obj GlobalConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_GlobalConfig[:], false)
	if err != nil {
		return err
	}
	// Serialize `Epoch`:
	if err = encoder.Encode(obj.Epoch); err != nil {
		return fmt.Errorf("error while marshaling Epoch:%w", err)
	}
	// Serialize `CurveType`:
	if err = encoder.Encode(obj.CurveType); err != nil {
		return fmt.Errorf("error while marshaling CurveType:%w", err)
	}
	// Serialize `Index`:
	if err = encoder.Encode(obj.Index); err != nil {
		return fmt.Errorf("error while marshaling Index:%w", err)
	}
	// Serialize `MigrateFee`:
	if err = encoder.Encode(obj.MigrateFee); err != nil {
		return fmt.Errorf("error while marshaling MigrateFee:%w", err)
	}
	// Serialize `TradeFeeRate`:
	if err = encoder.Encode(obj.TradeFeeRate); err != nil {
		return fmt.Errorf("error while marshaling TradeFeeRate:%w", err)
	}
	// Serialize `MaxShareFeeRate`:
	if err = encoder.Encode(obj.MaxShareFeeRate); err != nil {
		return fmt.Errorf("error while marshaling MaxShareFeeRate:%w", err)
	}
	// Serialize `MinBaseSupply`:
	if err = encoder.Encode(obj.MinBaseSupply); err != nil {
		return fmt.Errorf("error while marshaling MinBaseSupply:%w", err)
	}
	// Serialize `MaxLockRate`:
	if err = encoder.Encode(obj.MaxLockRate); err != nil {
		return fmt.Errorf("error while marshaling MaxLockRate:%w", err)
	}
	// Serialize `MinBaseSellRate`:
	if err = encoder.Encode(obj.MinBaseSellRate); err != nil {
		return fmt.Errorf("error while marshaling MinBaseSellRate:%w", err)
	}
	// Serialize `MinBaseMigrateRate`:
	if err = encoder.Encode(obj.MinBaseMigrateRate); err != nil {
		return fmt.Errorf("error while marshaling MinBaseMigrateRate:%w", err)
	}
	// Serialize `MinQuoteFundRaising`:
	if err = encoder.Encode(obj.MinQuoteFundRaising); err != nil {
		return fmt.Errorf("error while marshaling MinQuoteFundRaising:%w", err)
	}
	// Serialize `QuoteMint`:
	if err = encoder.Encode(obj.QuoteMint); err != nil {
		return fmt.Errorf("error while marshaling QuoteMint:%w", err)
	}
	// Serialize `ProtocolFeeOwner`:
	if err = encoder.Encode(obj.ProtocolFeeOwner); err != nil {
		return fmt.Errorf("error while marshaling ProtocolFeeOwner:%w", err)
	}
	// Serialize `MigrateFeeOwner`:
	if err = encoder.Encode(obj.MigrateFeeOwner); err != nil {
		return fmt.Errorf("error while marshaling MigrateFeeOwner:%w", err)
	}
	// Serialize `MigrateToAmmWallet`:
	if err = encoder.Encode(obj.MigrateToAmmWallet); err != nil {
		return fmt.Errorf("error while marshaling MigrateToAmmWallet:%w", err)
	}
	// Serialize `MigrateToCpswapWallet`:
	if err = encoder.Encode(obj.MigrateToCpswapWallet); err != nil {
		return fmt.Errorf("error while marshaling MigrateToCpswapWallet:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj GlobalConfig) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding GlobalConfig: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *GlobalConfig) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_GlobalConfig[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_GlobalConfig[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Epoch`:
	if err = decoder.Decode(&obj.Epoch); err != nil {
		return fmt.Errorf("error while unmarshaling Epoch:%w", err)
	}
	// Deserialize `CurveType`:
	if err = decoder.Decode(&obj.CurveType); err != nil {
		return fmt.Errorf("error while unmarshaling CurveType:%w", err)
	}
	// Deserialize `Index`:
	if err = decoder.Decode(&obj.Index); err != nil {
		return fmt.Errorf("error while unmarshaling Index:%w", err)
	}
	// Deserialize `MigrateFee`:
	if err = decoder.Decode(&obj.MigrateFee); err != nil {
		return fmt.Errorf("error while unmarshaling MigrateFee:%w", err)
	}
	// Deserialize `TradeFeeRate`:
	if err = decoder.Decode(&obj.TradeFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling TradeFeeRate:%w", err)
	}
	// Deserialize `MaxShareFeeRate`:
	if err = decoder.Decode(&obj.MaxShareFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling MaxShareFeeRate:%w", err)
	}
	// Deserialize `MinBaseSupply`:
	if err = decoder.Decode(&obj.MinBaseSupply); err != nil {
		return fmt.Errorf("error while unmarshaling MinBaseSupply:%w", err)
	}
	// Deserialize `MaxLockRate`:
	if err = decoder.Decode(&obj.MaxLockRate); err != nil {
		return fmt.Errorf("error while unmarshaling MaxLockRate:%w", err)
	}
	// Deserialize `MinBaseSellRate`:
	if err = decoder.Decode(&obj.MinBaseSellRate); err != nil {
		return fmt.Errorf("error while unmarshaling MinBaseSellRate:%w", err)
	}
	// Deserialize `MinBaseMigrateRate`:
	if err = decoder.Decode(&obj.MinBaseMigrateRate); err != nil {
		return fmt.Errorf("error while unmarshaling MinBaseMigrateRate:%w", err)
	}
	// Deserialize `MinQuoteFundRaising`:
	if err = decoder.Decode(&obj.MinQuoteFundRaising); err != nil {
		return fmt.Errorf("error while unmarshaling MinQuoteFundRaising:%w", err)
	}
	// Deserialize `QuoteMint`:
	if err = decoder.Decode(&obj.QuoteMint); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteMint:%w", err)
	}
	// Deserialize `ProtocolFeeOwner`:
	if err = decoder.Decode(&obj.ProtocolFeeOwner); err != nil {
		return fmt.Errorf("error while unmarshaling ProtocolFeeOwner:%w", err)
	}
	// Deserialize `MigrateFeeOwner`:
	if err = decoder.Decode(&obj.MigrateFeeOwner); err != nil {
		return fmt.Errorf("error while unmarshaling MigrateFeeOwner:%w", err)
	}
	// Deserialize `MigrateToAmmWallet`:
	if err = decoder.Decode(&obj.MigrateToAmmWallet); err != nil {
		return fmt.Errorf("error while unmarshaling MigrateToAmmWallet:%w", err)
	}
	// Deserialize `MigrateToCpswapWallet`:
	if err = decoder.Decode(&obj.MigrateToCpswapWallet); err != nil {
		return fmt.Errorf("error while unmarshaling MigrateToCpswapWallet:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *GlobalConfig) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling GlobalConfig: %w", err)
	}
	return nil
}

func UnmarshalGlobalConfig(buf []byte) (*GlobalConfig, error) {
	obj := new(GlobalConfig)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type PlatformConfig struct {
	// The epoch for update interval
	Epoch uint64

	// The platform fee wallet
	PlatformFeeWallet solanago.PublicKey

	// The platform nft wallet to receive the platform NFT after migration if platform_scale is not 0(Only support MigrateType::CPSWAP)
	PlatformNftWallet solanago.PublicKey

	// Scale of the platform liquidity quantity rights will be converted into NFT(Only support MigrateType::CPSWAP)
	PlatformScale uint64

	// Scale of the token creator liquidity quantity rights will be converted into NFT(Only support MigrateType::CPSWAP)
	CreatorScale uint64

	// Scale of liquidity directly to burn
	BurnScale uint64

	// The platform fee rate
	FeeRate uint64

	// The platform name
	Name [64]uint8

	// The platform website
	Web [256]uint8

	// The platform img link
	Img [256]uint8

	// The platform specifies the trade fee rate after migration to cp swap
	CpswapConfig solanago.PublicKey

	// Creator fee rate
	CreatorFeeRate uint64

	// If the base token belongs to token2022, then you can choose to support the transferfeeConfig extension, which includes permissions such as `transfer_fee_config_authority“ and `withdraw_withheld_authority`.
	// When initializing mint, `withdraw_withheld_authority` and `transfer_fee_config_authority` both belongs to the contract.
	// Once the token is migrated to AMM, the authorities will be reset to this value
	TransferFeeExtensionAuth solanago.PublicKey

	// padding for future updates
	Padding [180]uint8

	// The parameters for launching the pool
	CurveParams []PlatformCurveParam
}

func (obj PlatformConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_PlatformConfig[:], false)
	if err != nil {
		return err
	}
	// Serialize `Epoch`:
	if err = encoder.Encode(obj.Epoch); err != nil {
		return fmt.Errorf("error while marshaling Epoch:%w", err)
	}
	// Serialize `PlatformFeeWallet`:
	if err = encoder.Encode(obj.PlatformFeeWallet); err != nil {
		return fmt.Errorf("error while marshaling PlatformFeeWallet:%w", err)
	}
	// Serialize `PlatformNftWallet`:
	if err = encoder.Encode(obj.PlatformNftWallet); err != nil {
		return fmt.Errorf("error while marshaling PlatformNftWallet:%w", err)
	}
	// Serialize `PlatformScale`:
	if err = encoder.Encode(obj.PlatformScale); err != nil {
		return fmt.Errorf("error while marshaling PlatformScale:%w", err)
	}
	// Serialize `CreatorScale`:
	if err = encoder.Encode(obj.CreatorScale); err != nil {
		return fmt.Errorf("error while marshaling CreatorScale:%w", err)
	}
	// Serialize `BurnScale`:
	if err = encoder.Encode(obj.BurnScale); err != nil {
		return fmt.Errorf("error while marshaling BurnScale:%w", err)
	}
	// Serialize `FeeRate`:
	if err = encoder.Encode(obj.FeeRate); err != nil {
		return fmt.Errorf("error while marshaling FeeRate:%w", err)
	}
	// Serialize `Name`:
	if err = encoder.Encode(obj.Name); err != nil {
		return fmt.Errorf("error while marshaling Name:%w", err)
	}
	// Serialize `Web`:
	if err = encoder.Encode(obj.Web); err != nil {
		return fmt.Errorf("error while marshaling Web:%w", err)
	}
	// Serialize `Img`:
	if err = encoder.Encode(obj.Img); err != nil {
		return fmt.Errorf("error while marshaling Img:%w", err)
	}
	// Serialize `CpswapConfig`:
	if err = encoder.Encode(obj.CpswapConfig); err != nil {
		return fmt.Errorf("error while marshaling CpswapConfig:%w", err)
	}
	// Serialize `CreatorFeeRate`:
	if err = encoder.Encode(obj.CreatorFeeRate); err != nil {
		return fmt.Errorf("error while marshaling CreatorFeeRate:%w", err)
	}
	// Serialize `TransferFeeExtensionAuth`:
	if err = encoder.Encode(obj.TransferFeeExtensionAuth); err != nil {
		return fmt.Errorf("error while marshaling TransferFeeExtensionAuth:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	// Serialize `CurveParams`:
	if err = encoder.Encode(obj.CurveParams); err != nil {
		return fmt.Errorf("error while marshaling CurveParams:%w", err)
	}
	return nil
}

func (obj PlatformConfig) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding PlatformConfig: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *PlatformConfig) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_PlatformConfig[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_PlatformConfig[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Epoch`:
	if err = decoder.Decode(&obj.Epoch); err != nil {
		return fmt.Errorf("error while unmarshaling Epoch:%w", err)
	}
	// Deserialize `PlatformFeeWallet`:
	if err = decoder.Decode(&obj.PlatformFeeWallet); err != nil {
		return fmt.Errorf("error while unmarshaling PlatformFeeWallet:%w", err)
	}
	// Deserialize `PlatformNftWallet`:
	if err = decoder.Decode(&obj.PlatformNftWallet); err != nil {
		return fmt.Errorf("error while unmarshaling PlatformNftWallet:%w", err)
	}
	// Deserialize `PlatformScale`:
	if err = decoder.Decode(&obj.PlatformScale); err != nil {
		return fmt.Errorf("error while unmarshaling PlatformScale:%w", err)
	}
	// Deserialize `CreatorScale`:
	if err = decoder.Decode(&obj.CreatorScale); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorScale:%w", err)
	}
	// Deserialize `BurnScale`:
	if err = decoder.Decode(&obj.BurnScale); err != nil {
		return fmt.Errorf("error while unmarshaling BurnScale:%w", err)
	}
	// Deserialize `FeeRate`:
	if err = decoder.Decode(&obj.FeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling FeeRate:%w", err)
	}
	// Deserialize `Name`:
	if err = decoder.Decode(&obj.Name); err != nil {
		return fmt.Errorf("error while unmarshaling Name:%w", err)
	}
	// Deserialize `Web`:
	if err = decoder.Decode(&obj.Web); err != nil {
		return fmt.Errorf("error while unmarshaling Web:%w", err)
	}
	// Deserialize `Img`:
	if err = decoder.Decode(&obj.Img); err != nil {
		return fmt.Errorf("error while unmarshaling Img:%w", err)
	}
	// Deserialize `CpswapConfig`:
	if err = decoder.Decode(&obj.CpswapConfig); err != nil {
		return fmt.Errorf("error while unmarshaling CpswapConfig:%w", err)
	}
	// Deserialize `CreatorFeeRate`:
	if err = decoder.Decode(&obj.CreatorFeeRate); err != nil {
		return fmt.Errorf("error while unmarshaling CreatorFeeRate:%w", err)
	}
	// Deserialize `TransferFeeExtensionAuth`:
	if err = decoder.Decode(&obj.TransferFeeExtensionAuth); err != nil {
		return fmt.Errorf("error while unmarshaling TransferFeeExtensionAuth:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	// Deserialize `CurveParams`:
	if err = decoder.Decode(&obj.CurveParams); err != nil {
		return fmt.Errorf("error while unmarshaling CurveParams:%w", err)
	}
	return nil
}

func (obj *PlatformConfig) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling PlatformConfig: %w", err)
	}
	return nil
}

func UnmarshalPlatformConfig(buf []byte) (*PlatformConfig, error) {
	obj := new(PlatformConfig)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Represents the state of a trading pool in the protocol
// Stores all essential information about pool balances, fees, and configuration
type PoolState struct {
	// Account update epoch
	Epoch uint64

	// Bump seed used for PDA address derivation
	AuthBump uint8

	// Current status of the pool
	// * 0: Pool is funding
	// * 1: Pool funding is end, waiting for migration
	// * 2: Pool migration is done
	Status uint8

	// Decimals of the pool base token
	BaseDecimals uint8

	// Decimals of the pool quote token
	QuoteDecimals uint8

	// Migrate to AMM or CpSwap, 0: amm， 1: cpswap
	MigrateType uint8

	// Supply of the pool base token
	Supply uint64

	// Total sell amount of the base token
	TotalBaseSell uint64

	// For different curves, virtual_base and virtual_quote have different meanings
	// For constant product curve, virtual_base and virtual_quote are virtual liquidity, virtual_quote/virtual_base is the initial price
	// For linear price curve, virtual_base is the price slope parameter a, virtual_quote has no effect
	// For fixed price curve, virtual_quote/virtual_base is the initial price
	VirtualBase  uint64
	VirtualQuote uint64

	// Actual base token amount in the pool
	// Represents the real tokens available for trading
	RealBase uint64

	// Actual quote token amount in the pool
	// Represents the real tokens available for trading
	RealQuote uint64

	// The total quote fund raising of the pool
	TotalQuoteFundRaising uint64

	// Accumulated trading fees in quote tokens
	// Can be collected by the protocol fee owner
	QuoteProtocolFee uint64

	// Accumulated platform fees in quote tokens
	// Can be collected by the platform wallet stored in platform config
	PlatformFee uint64

	// The fee of migrate to amm
	MigrateFee uint64

	// Vesting schedule for the base token
	VestingSchedule VestingSchedule

	// Public key of the global configuration account
	// Contains protocol-wide settings this pool adheres to
	GlobalConfig solanago.PublicKey

	// Public key of the platform configuration account
	// Contains platform-wide settings this pool adheres to
	PlatformConfig solanago.PublicKey

	// Public key of the base mint address
	BaseMint solanago.PublicKey

	// Public key of the quote mint address
	QuoteMint solanago.PublicKey

	// Public key of the base token vault
	// Holds the actual base tokens owned by the pool
	BaseVault solanago.PublicKey

	// Public key of the quote token vault
	// Holds the actual quote tokens owned by the pool
	QuoteVault solanago.PublicKey

	// The creator of base token
	Creator solanago.PublicKey

	// token program bits
	// bit0: base token program flag
	// 0: spl_token_program
	// 1: token_program_2022
	//
	// bit1: quote token program flag
	// 0: spl_token_program
	// 1: token_program_2022
	TokenProgramFlag uint8

	// migrate to cpmm fee model
	AmmFeeOn AmmFeeOn

	// padding for future updates
	Padding [62]uint8
}

func (obj PoolState) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_PoolState[:], false)
	if err != nil {
		return err
	}
	// Serialize `Epoch`:
	if err = encoder.Encode(obj.Epoch); err != nil {
		return fmt.Errorf("error while marshaling Epoch:%w", err)
	}
	// Serialize `AuthBump`:
	if err = encoder.Encode(obj.AuthBump); err != nil {
		return fmt.Errorf("error while marshaling AuthBump:%w", err)
	}
	// Serialize `Status`:
	if err = encoder.Encode(obj.Status); err != nil {
		return fmt.Errorf("error while marshaling Status:%w", err)
	}
	// Serialize `BaseDecimals`:
	if err = encoder.Encode(obj.BaseDecimals); err != nil {
		return fmt.Errorf("error while marshaling BaseDecimals:%w", err)
	}
	// Serialize `QuoteDecimals`:
	if err = encoder.Encode(obj.QuoteDecimals); err != nil {
		return fmt.Errorf("error while marshaling QuoteDecimals:%w", err)
	}
	// Serialize `MigrateType`:
	if err = encoder.Encode(obj.MigrateType); err != nil {
		return fmt.Errorf("error while marshaling MigrateType:%w", err)
	}
	// Serialize `Supply`:
	if err = encoder.Encode(obj.Supply); err != nil {
		return fmt.Errorf("error while marshaling Supply:%w", err)
	}
	// Serialize `TotalBaseSell`:
	if err = encoder.Encode(obj.TotalBaseSell); err != nil {
		return fmt.Errorf("error while marshaling TotalBaseSell:%w", err)
	}
	// Serialize `VirtualBase`:
	if err = encoder.Encode(obj.VirtualBase); err != nil {
		return fmt.Errorf("error while marshaling VirtualBase:%w", err)
	}
	// Serialize `VirtualQuote`:
	if err = encoder.Encode(obj.VirtualQuote); err != nil {
		return fmt.Errorf("error while marshaling VirtualQuote:%w", err)
	}
	// Serialize `RealBase`:
	if err = encoder.Encode(obj.RealBase); err != nil {
		return fmt.Errorf("error while marshaling RealBase:%w", err)
	}
	// Serialize `RealQuote`:
	if err = encoder.Encode(obj.RealQuote); err != nil {
		return fmt.Errorf("error while marshaling RealQuote:%w", err)
	}
	// Serialize `TotalQuoteFundRaising`:
	if err = encoder.Encode(obj.TotalQuoteFundRaising); err != nil {
		return fmt.Errorf("error while marshaling TotalQuoteFundRaising:%w", err)
	}
	// Serialize `QuoteProtocolFee`:
	if err = encoder.Encode(obj.QuoteProtocolFee); err != nil {
		return fmt.Errorf("error while marshaling QuoteProtocolFee:%w", err)
	}
	// Serialize `PlatformFee`:
	if err = encoder.Encode(obj.PlatformFee); err != nil {
		return fmt.Errorf("error while marshaling PlatformFee:%w", err)
	}
	// Serialize `MigrateFee`:
	if err = encoder.Encode(obj.MigrateFee); err != nil {
		return fmt.Errorf("error while marshaling MigrateFee:%w", err)
	}
	// Serialize `VestingSchedule`:
	if err = encoder.Encode(obj.VestingSchedule); err != nil {
		return fmt.Errorf("error while marshaling VestingSchedule:%w", err)
	}
	// Serialize `GlobalConfig`:
	if err = encoder.Encode(obj.GlobalConfig); err != nil {
		return fmt.Errorf("error while marshaling GlobalConfig:%w", err)
	}
	// Serialize `PlatformConfig`:
	if err = encoder.Encode(obj.PlatformConfig); err != nil {
		return fmt.Errorf("error while marshaling PlatformConfig:%w", err)
	}
	// Serialize `BaseMint`:
	if err = encoder.Encode(obj.BaseMint); err != nil {
		return fmt.Errorf("error while marshaling BaseMint:%w", err)
	}
	// Serialize `QuoteMint`:
	if err = encoder.Encode(obj.QuoteMint); err != nil {
		return fmt.Errorf("error while marshaling QuoteMint:%w", err)
	}
	// Serialize `BaseVault`:
	if err = encoder.Encode(obj.BaseVault); err != nil {
		return fmt.Errorf("error while marshaling BaseVault:%w", err)
	}
	// Serialize `QuoteVault`:
	if err = encoder.Encode(obj.QuoteVault); err != nil {
		return fmt.Errorf("error while marshaling QuoteVault:%w", err)
	}
	// Serialize `Creator`:
	if err = encoder.Encode(obj.Creator); err != nil {
		return fmt.Errorf("error while marshaling Creator:%w", err)
	}
	// Serialize `TokenProgramFlag`:
	if err = encoder.Encode(obj.TokenProgramFlag); err != nil {
		return fmt.Errorf("error while marshaling TokenProgramFlag:%w", err)
	}
	// Serialize `AmmFeeOn`:
	if err = encoder.Encode(obj.AmmFeeOn); err != nil {
		return fmt.Errorf("error while marshaling AmmFeeOn:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj PoolState) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding PoolState: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *PoolState) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_PoolState[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_PoolState[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Epoch`:
	if err = decoder.Decode(&obj.Epoch); err != nil {
		return fmt.Errorf("error while unmarshaling Epoch:%w", err)
	}
	// Deserialize `AuthBump`:
	if err = decoder.Decode(&obj.AuthBump); err != nil {
		return fmt.Errorf("error while unmarshaling AuthBump:%w", err)
	}
	// Deserialize `Status`:
	if err = decoder.Decode(&obj.Status); err != nil {
		return fmt.Errorf("error while unmarshaling Status:%w", err)
	}
	// Deserialize `BaseDecimals`:
	if err = decoder.Decode(&obj.BaseDecimals); err != nil {
		return fmt.Errorf("error while unmarshaling BaseDecimals:%w", err)
	}
	// Deserialize `QuoteDecimals`:
	if err = decoder.Decode(&obj.QuoteDecimals); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteDecimals:%w", err)
	}
	// Deserialize `MigrateType`:
	if err = decoder.Decode(&obj.MigrateType); err != nil {
		return fmt.Errorf("error while unmarshaling MigrateType:%w", err)
	}
	// Deserialize `Supply`:
	if err = decoder.Decode(&obj.Supply); err != nil {
		return fmt.Errorf("error while unmarshaling Supply:%w", err)
	}
	// Deserialize `TotalBaseSell`:
	if err = decoder.Decode(&obj.TotalBaseSell); err != nil {
		return fmt.Errorf("error while unmarshaling TotalBaseSell:%w", err)
	}
	// Deserialize `VirtualBase`:
	if err = decoder.Decode(&obj.VirtualBase); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualBase:%w", err)
	}
	// Deserialize `VirtualQuote`:
	if err = decoder.Decode(&obj.VirtualQuote); err != nil {
		return fmt.Errorf("error while unmarshaling VirtualQuote:%w", err)
	}
	// Deserialize `RealBase`:
	if err = decoder.Decode(&obj.RealBase); err != nil {
		return fmt.Errorf("error while unmarshaling RealBase:%w", err)
	}
	// Deserialize `RealQuote`:
	if err = decoder.Decode(&obj.RealQuote); err != nil {
		return fmt.Errorf("error while unmarshaling RealQuote:%w", err)
	}
	// Deserialize `TotalQuoteFundRaising`:
	if err = decoder.Decode(&obj.TotalQuoteFundRaising); err != nil {
		return fmt.Errorf("error while unmarshaling TotalQuoteFundRaising:%w", err)
	}
	// Deserialize `QuoteProtocolFee`:
	if err = decoder.Decode(&obj.QuoteProtocolFee); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteProtocolFee:%w", err)
	}
	// Deserialize `PlatformFee`:
	if err = decoder.Decode(&obj.PlatformFee); err != nil {
		return fmt.Errorf("error while unmarshaling PlatformFee:%w", err)
	}
	// Deserialize `MigrateFee`:
	if err = decoder.Decode(&obj.MigrateFee); err != nil {
		return fmt.Errorf("error while unmarshaling MigrateFee:%w", err)
	}
	// Deserialize `VestingSchedule`:
	if err = decoder.Decode(&obj.VestingSchedule); err != nil {
		return fmt.Errorf("error while unmarshaling VestingSchedule:%w", err)
	}
	// Deserialize `GlobalConfig`:
	if err = decoder.Decode(&obj.GlobalConfig); err != nil {
		return fmt.Errorf("error while unmarshaling GlobalConfig:%w", err)
	}
	// Deserialize `PlatformConfig`:
	if err = decoder.Decode(&obj.PlatformConfig); err != nil {
		return fmt.Errorf("error while unmarshaling PlatformConfig:%w", err)
	}
	// Deserialize `BaseMint`:
	if err = decoder.Decode(&obj.BaseMint); err != nil {
		return fmt.Errorf("error while unmarshaling BaseMint:%w", err)
	}
	// Deserialize `QuoteMint`:
	if err = decoder.Decode(&obj.QuoteMint); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteMint:%w", err)
	}
	// Deserialize `BaseVault`:
	if err = decoder.Decode(&obj.BaseVault); err != nil {
		return fmt.Errorf("error while unmarshaling BaseVault:%w", err)
	}
	// Deserialize `QuoteVault`:
	if err = decoder.Decode(&obj.QuoteVault); err != nil {
		return fmt.Errorf("error while unmarshaling QuoteVault:%w", err)
	}
	// Deserialize `Creator`:
	if err = decoder.Decode(&obj.Creator); err != nil {
		return fmt.Errorf("error while unmarshaling Creator:%w", err)
	}
	// Deserialize `TokenProgramFlag`:
	if err = decoder.Decode(&obj.TokenProgramFlag); err != nil {
		return fmt.Errorf("error while unmarshaling TokenProgramFlag:%w", err)
	}
	// Deserialize `AmmFeeOn`:
	if err = decoder.Decode(&obj.AmmFeeOn); err != nil {
		return fmt.Errorf("error while unmarshaling AmmFeeOn:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *PoolState) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling PoolState: %w", err)
	}
	return nil
}

func UnmarshalPoolState(buf []byte) (*PoolState, error) {
	obj := new(PoolState)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type VestingRecord struct {
	// Account update epoch
	Epoch uint64

	// The pool state account
	Pool solanago.PublicKey

	// The beneficiary of the vesting account
	Beneficiary solanago.PublicKey

	// The amount of tokens claimed
	ClaimedAmount uint64

	// The share amount of the token to be vested
	TokenShareAmount uint64

	// padding for future updates
	Padding [8]uint64
}

func (obj VestingRecord) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(Account_VestingRecord[:], false)
	if err != nil {
		return err
	}
	// Serialize `Epoch`:
	if err = encoder.Encode(obj.Epoch); err != nil {
		return fmt.Errorf("error while marshaling Epoch:%w", err)
	}
	// Serialize `Pool`:
	if err = encoder.Encode(obj.Pool); err != nil {
		return fmt.Errorf("error while marshaling Pool:%w", err)
	}
	// Serialize `Beneficiary`:
	if err = encoder.Encode(obj.Beneficiary); err != nil {
		return fmt.Errorf("error while marshaling Beneficiary:%w", err)
	}
	// Serialize `ClaimedAmount`:
	if err = encoder.Encode(obj.ClaimedAmount); err != nil {
		return fmt.Errorf("error while marshaling ClaimedAmount:%w", err)
	}
	// Serialize `TokenShareAmount`:
	if err = encoder.Encode(obj.TokenShareAmount); err != nil {
		return fmt.Errorf("error while marshaling TokenShareAmount:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj VestingRecord) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding VestingRecord: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *VestingRecord) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadDiscriminator()
		if err != nil {
			return err
		}
		if !discriminator.Equal(Account_VestingRecord[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				Account_VestingRecord[:],
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Epoch`:
	if err = decoder.Decode(&obj.Epoch); err != nil {
		return fmt.Errorf("error while unmarshaling Epoch:%w", err)
	}
	// Deserialize `Pool`:
	if err = decoder.Decode(&obj.Pool); err != nil {
		return fmt.Errorf("error while unmarshaling Pool:%w", err)
	}
	// Deserialize `Beneficiary`:
	if err = decoder.Decode(&obj.Beneficiary); err != nil {
		return fmt.Errorf("error while unmarshaling Beneficiary:%w", err)
	}
	// Deserialize `ClaimedAmount`:
	if err = decoder.Decode(&obj.ClaimedAmount); err != nil {
		return fmt.Errorf("error while unmarshaling ClaimedAmount:%w", err)
	}
	// Deserialize `TokenShareAmount`:
	if err = decoder.Decode(&obj.TokenShareAmount); err != nil {
		return fmt.Errorf("error while unmarshaling TokenShareAmount:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *VestingRecord) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling VestingRecord: %w", err)
	}
	return nil
}

func UnmarshalVestingRecord(buf []byte) (*VestingRecord, error) {
	obj := new(VestingRecord)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
