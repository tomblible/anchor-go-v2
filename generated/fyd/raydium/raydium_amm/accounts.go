package raydium_amm

import (
	"bytes"
	"fmt"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

type TargetOrders struct {
	Owner                [4]uint64
	BuyOrders            [50]TargetOrder
	Padding1             [8]uint64
	TargetX              binary.Uint128
	TargetY              binary.Uint128
	PlanXBuy             binary.Uint128
	PlanYBuy             binary.Uint128
	PlanXSell            binary.Uint128
	PlanYSell            binary.Uint128
	PlacedX              binary.Uint128
	PlacedY              binary.Uint128
	CalcPnlX             binary.Uint128
	CalcPnlY             binary.Uint128
	SellOrders           [50]TargetOrder
	Padding2             [6]uint64
	ReplaceBuyClientId   [10]uint64
	ReplaceSellClientId  [10]uint64
	LastOrderNumerator   uint64
	LastOrderDenominator uint64
	PlanOrdersCur        uint64
	PlaceOrdersCur       uint64
	ValidBuyOrderNum     uint64
	ValidSellOrderNum    uint64
	Padding3             [10]uint64
	FreeSlotBits         binary.Uint128
}

func (obj TargetOrders) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Serialize `Owner`:
	if err = encoder.Encode(obj.Owner); err != nil {
		return fmt.Errorf("error while marshaling Owner:%w", err)
	}
	// Serialize `BuyOrders`:
	if err = encoder.Encode(obj.BuyOrders); err != nil {
		return fmt.Errorf("error while marshaling BuyOrders:%w", err)
	}
	// Serialize `Padding1`:
	if err = encoder.Encode(obj.Padding1); err != nil {
		return fmt.Errorf("error while marshaling Padding1:%w", err)
	}
	// Serialize `TargetX`:
	if err = encoder.Encode(obj.TargetX); err != nil {
		return fmt.Errorf("error while marshaling TargetX:%w", err)
	}
	// Serialize `TargetY`:
	if err = encoder.Encode(obj.TargetY); err != nil {
		return fmt.Errorf("error while marshaling TargetY:%w", err)
	}
	// Serialize `PlanXBuy`:
	if err = encoder.Encode(obj.PlanXBuy); err != nil {
		return fmt.Errorf("error while marshaling PlanXBuy:%w", err)
	}
	// Serialize `PlanYBuy`:
	if err = encoder.Encode(obj.PlanYBuy); err != nil {
		return fmt.Errorf("error while marshaling PlanYBuy:%w", err)
	}
	// Serialize `PlanXSell`:
	if err = encoder.Encode(obj.PlanXSell); err != nil {
		return fmt.Errorf("error while marshaling PlanXSell:%w", err)
	}
	// Serialize `PlanYSell`:
	if err = encoder.Encode(obj.PlanYSell); err != nil {
		return fmt.Errorf("error while marshaling PlanYSell:%w", err)
	}
	// Serialize `PlacedX`:
	if err = encoder.Encode(obj.PlacedX); err != nil {
		return fmt.Errorf("error while marshaling PlacedX:%w", err)
	}
	// Serialize `PlacedY`:
	if err = encoder.Encode(obj.PlacedY); err != nil {
		return fmt.Errorf("error while marshaling PlacedY:%w", err)
	}
	// Serialize `CalcPnlX`:
	if err = encoder.Encode(obj.CalcPnlX); err != nil {
		return fmt.Errorf("error while marshaling CalcPnlX:%w", err)
	}
	// Serialize `CalcPnlY`:
	if err = encoder.Encode(obj.CalcPnlY); err != nil {
		return fmt.Errorf("error while marshaling CalcPnlY:%w", err)
	}
	// Serialize `SellOrders`:
	if err = encoder.Encode(obj.SellOrders); err != nil {
		return fmt.Errorf("error while marshaling SellOrders:%w", err)
	}
	// Serialize `Padding2`:
	if err = encoder.Encode(obj.Padding2); err != nil {
		return fmt.Errorf("error while marshaling Padding2:%w", err)
	}
	// Serialize `ReplaceBuyClientId`:
	if err = encoder.Encode(obj.ReplaceBuyClientId); err != nil {
		return fmt.Errorf("error while marshaling ReplaceBuyClientId:%w", err)
	}
	// Serialize `ReplaceSellClientId`:
	if err = encoder.Encode(obj.ReplaceSellClientId); err != nil {
		return fmt.Errorf("error while marshaling ReplaceSellClientId:%w", err)
	}
	// Serialize `LastOrderNumerator`:
	if err = encoder.Encode(obj.LastOrderNumerator); err != nil {
		return fmt.Errorf("error while marshaling LastOrderNumerator:%w", err)
	}
	// Serialize `LastOrderDenominator`:
	if err = encoder.Encode(obj.LastOrderDenominator); err != nil {
		return fmt.Errorf("error while marshaling LastOrderDenominator:%w", err)
	}
	// Serialize `PlanOrdersCur`:
	if err = encoder.Encode(obj.PlanOrdersCur); err != nil {
		return fmt.Errorf("error while marshaling PlanOrdersCur:%w", err)
	}
	// Serialize `PlaceOrdersCur`:
	if err = encoder.Encode(obj.PlaceOrdersCur); err != nil {
		return fmt.Errorf("error while marshaling PlaceOrdersCur:%w", err)
	}
	// Serialize `ValidBuyOrderNum`:
	if err = encoder.Encode(obj.ValidBuyOrderNum); err != nil {
		return fmt.Errorf("error while marshaling ValidBuyOrderNum:%w", err)
	}
	// Serialize `ValidSellOrderNum`:
	if err = encoder.Encode(obj.ValidSellOrderNum); err != nil {
		return fmt.Errorf("error while marshaling ValidSellOrderNum:%w", err)
	}
	// Serialize `Padding3`:
	if err = encoder.Encode(obj.Padding3); err != nil {
		return fmt.Errorf("error while marshaling Padding3:%w", err)
	}
	// Serialize `FreeSlotBits`:
	if err = encoder.Encode(obj.FreeSlotBits); err != nil {
		return fmt.Errorf("error while marshaling FreeSlotBits:%w", err)
	}
	return nil
}

func (obj TargetOrders) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding TargetOrders: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *TargetOrders) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Deserialize `Owner`:
	if err = decoder.Decode(&obj.Owner); err != nil {
		return fmt.Errorf("error while unmarshaling Owner:%w", err)
	}
	// Deserialize `BuyOrders`:
	if err = decoder.Decode(&obj.BuyOrders); err != nil {
		return fmt.Errorf("error while unmarshaling BuyOrders:%w", err)
	}
	// Deserialize `Padding1`:
	if err = decoder.Decode(&obj.Padding1); err != nil {
		return fmt.Errorf("error while unmarshaling Padding1:%w", err)
	}
	// Deserialize `TargetX`:
	if err = decoder.Decode(&obj.TargetX); err != nil {
		return fmt.Errorf("error while unmarshaling TargetX:%w", err)
	}
	// Deserialize `TargetY`:
	if err = decoder.Decode(&obj.TargetY); err != nil {
		return fmt.Errorf("error while unmarshaling TargetY:%w", err)
	}
	// Deserialize `PlanXBuy`:
	if err = decoder.Decode(&obj.PlanXBuy); err != nil {
		return fmt.Errorf("error while unmarshaling PlanXBuy:%w", err)
	}
	// Deserialize `PlanYBuy`:
	if err = decoder.Decode(&obj.PlanYBuy); err != nil {
		return fmt.Errorf("error while unmarshaling PlanYBuy:%w", err)
	}
	// Deserialize `PlanXSell`:
	if err = decoder.Decode(&obj.PlanXSell); err != nil {
		return fmt.Errorf("error while unmarshaling PlanXSell:%w", err)
	}
	// Deserialize `PlanYSell`:
	if err = decoder.Decode(&obj.PlanYSell); err != nil {
		return fmt.Errorf("error while unmarshaling PlanYSell:%w", err)
	}
	// Deserialize `PlacedX`:
	if err = decoder.Decode(&obj.PlacedX); err != nil {
		return fmt.Errorf("error while unmarshaling PlacedX:%w", err)
	}
	// Deserialize `PlacedY`:
	if err = decoder.Decode(&obj.PlacedY); err != nil {
		return fmt.Errorf("error while unmarshaling PlacedY:%w", err)
	}
	// Deserialize `CalcPnlX`:
	if err = decoder.Decode(&obj.CalcPnlX); err != nil {
		return fmt.Errorf("error while unmarshaling CalcPnlX:%w", err)
	}
	// Deserialize `CalcPnlY`:
	if err = decoder.Decode(&obj.CalcPnlY); err != nil {
		return fmt.Errorf("error while unmarshaling CalcPnlY:%w", err)
	}
	// Deserialize `SellOrders`:
	if err = decoder.Decode(&obj.SellOrders); err != nil {
		return fmt.Errorf("error while unmarshaling SellOrders:%w", err)
	}
	// Deserialize `Padding2`:
	if err = decoder.Decode(&obj.Padding2); err != nil {
		return fmt.Errorf("error while unmarshaling Padding2:%w", err)
	}
	// Deserialize `ReplaceBuyClientId`:
	if err = decoder.Decode(&obj.ReplaceBuyClientId); err != nil {
		return fmt.Errorf("error while unmarshaling ReplaceBuyClientId:%w", err)
	}
	// Deserialize `ReplaceSellClientId`:
	if err = decoder.Decode(&obj.ReplaceSellClientId); err != nil {
		return fmt.Errorf("error while unmarshaling ReplaceSellClientId:%w", err)
	}
	// Deserialize `LastOrderNumerator`:
	if err = decoder.Decode(&obj.LastOrderNumerator); err != nil {
		return fmt.Errorf("error while unmarshaling LastOrderNumerator:%w", err)
	}
	// Deserialize `LastOrderDenominator`:
	if err = decoder.Decode(&obj.LastOrderDenominator); err != nil {
		return fmt.Errorf("error while unmarshaling LastOrderDenominator:%w", err)
	}
	// Deserialize `PlanOrdersCur`:
	if err = decoder.Decode(&obj.PlanOrdersCur); err != nil {
		return fmt.Errorf("error while unmarshaling PlanOrdersCur:%w", err)
	}
	// Deserialize `PlaceOrdersCur`:
	if err = decoder.Decode(&obj.PlaceOrdersCur); err != nil {
		return fmt.Errorf("error while unmarshaling PlaceOrdersCur:%w", err)
	}
	// Deserialize `ValidBuyOrderNum`:
	if err = decoder.Decode(&obj.ValidBuyOrderNum); err != nil {
		return fmt.Errorf("error while unmarshaling ValidBuyOrderNum:%w", err)
	}
	// Deserialize `ValidSellOrderNum`:
	if err = decoder.Decode(&obj.ValidSellOrderNum); err != nil {
		return fmt.Errorf("error while unmarshaling ValidSellOrderNum:%w", err)
	}
	// Deserialize `Padding3`:
	if err = decoder.Decode(&obj.Padding3); err != nil {
		return fmt.Errorf("error while unmarshaling Padding3:%w", err)
	}
	// Deserialize `FreeSlotBits`:
	if err = decoder.Decode(&obj.FreeSlotBits); err != nil {
		return fmt.Errorf("error while unmarshaling FreeSlotBits:%w", err)
	}
	return nil
}

func (obj *TargetOrders) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling TargetOrders: %w", err)
	}
	return nil
}

func UnmarshalTargetOrders(buf []byte) (*TargetOrders, error) {
	obj := new(TargetOrders)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type Fees struct {
	MinSeparateNumerator   uint64
	MinSeparateDenominator uint64
	TradeFeeNumerator      uint64
	TradeFeeDenominator    uint64
	PnlNumerator           uint64
	PnlDenominator         uint64
	SwapFeeNumerator       uint64
	SwapFeeDenominator     uint64
}

func (obj Fees) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Serialize `MinSeparateNumerator`:
	if err = encoder.Encode(obj.MinSeparateNumerator); err != nil {
		return fmt.Errorf("error while marshaling MinSeparateNumerator:%w", err)
	}
	// Serialize `MinSeparateDenominator`:
	if err = encoder.Encode(obj.MinSeparateDenominator); err != nil {
		return fmt.Errorf("error while marshaling MinSeparateDenominator:%w", err)
	}
	// Serialize `TradeFeeNumerator`:
	if err = encoder.Encode(obj.TradeFeeNumerator); err != nil {
		return fmt.Errorf("error while marshaling TradeFeeNumerator:%w", err)
	}
	// Serialize `TradeFeeDenominator`:
	if err = encoder.Encode(obj.TradeFeeDenominator); err != nil {
		return fmt.Errorf("error while marshaling TradeFeeDenominator:%w", err)
	}
	// Serialize `PnlNumerator`:
	if err = encoder.Encode(obj.PnlNumerator); err != nil {
		return fmt.Errorf("error while marshaling PnlNumerator:%w", err)
	}
	// Serialize `PnlDenominator`:
	if err = encoder.Encode(obj.PnlDenominator); err != nil {
		return fmt.Errorf("error while marshaling PnlDenominator:%w", err)
	}
	// Serialize `SwapFeeNumerator`:
	if err = encoder.Encode(obj.SwapFeeNumerator); err != nil {
		return fmt.Errorf("error while marshaling SwapFeeNumerator:%w", err)
	}
	// Serialize `SwapFeeDenominator`:
	if err = encoder.Encode(obj.SwapFeeDenominator); err != nil {
		return fmt.Errorf("error while marshaling SwapFeeDenominator:%w", err)
	}
	return nil
}

func (obj Fees) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding Fees: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Fees) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Deserialize `MinSeparateNumerator`:
	if err = decoder.Decode(&obj.MinSeparateNumerator); err != nil {
		return fmt.Errorf("error while unmarshaling MinSeparateNumerator:%w", err)
	}
	// Deserialize `MinSeparateDenominator`:
	if err = decoder.Decode(&obj.MinSeparateDenominator); err != nil {
		return fmt.Errorf("error while unmarshaling MinSeparateDenominator:%w", err)
	}
	// Deserialize `TradeFeeNumerator`:
	if err = decoder.Decode(&obj.TradeFeeNumerator); err != nil {
		return fmt.Errorf("error while unmarshaling TradeFeeNumerator:%w", err)
	}
	// Deserialize `TradeFeeDenominator`:
	if err = decoder.Decode(&obj.TradeFeeDenominator); err != nil {
		return fmt.Errorf("error while unmarshaling TradeFeeDenominator:%w", err)
	}
	// Deserialize `PnlNumerator`:
	if err = decoder.Decode(&obj.PnlNumerator); err != nil {
		return fmt.Errorf("error while unmarshaling PnlNumerator:%w", err)
	}
	// Deserialize `PnlDenominator`:
	if err = decoder.Decode(&obj.PnlDenominator); err != nil {
		return fmt.Errorf("error while unmarshaling PnlDenominator:%w", err)
	}
	// Deserialize `SwapFeeNumerator`:
	if err = decoder.Decode(&obj.SwapFeeNumerator); err != nil {
		return fmt.Errorf("error while unmarshaling SwapFeeNumerator:%w", err)
	}
	// Deserialize `SwapFeeDenominator`:
	if err = decoder.Decode(&obj.SwapFeeDenominator); err != nil {
		return fmt.Errorf("error while unmarshaling SwapFeeDenominator:%w", err)
	}
	return nil
}

func (obj *Fees) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling Fees: %w", err)
	}
	return nil
}

func UnmarshalFees(buf []byte) (*Fees, error) {
	obj := new(Fees)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

type AmmInfo struct {
	Status             uint64
	Nonce              uint64
	OrderNum           uint64
	Depth              uint64
	CoinDecimals       uint64
	PcDecimals         uint64
	State              uint64
	ResetFlag          uint64
	MinSize            uint64
	VolMaxCutRatio     uint64
	AmountWave         uint64
	CoinLotSize        uint64
	PcLotSize          uint64
	MinPriceMultiplier uint64
	MaxPriceMultiplier uint64
	SysDecimalValue    uint64
	Fees               Fees
	OutPut             OutPutData
	TokenCoin          solanago.PublicKey
	TokenPc            solanago.PublicKey
	CoinMint           solanago.PublicKey
	PcMint             solanago.PublicKey
	LpMint             solanago.PublicKey
	OpenOrders         solanago.PublicKey
	Market             solanago.PublicKey
	SerumDex           solanago.PublicKey
	TargetOrders       solanago.PublicKey
	WithdrawQueue      solanago.PublicKey
	TokenTempLp        solanago.PublicKey
	AmmOwner           solanago.PublicKey
	LpAmount           uint64
	ClientOrderId      uint64
	Padding            [2]uint64
}

func (obj AmmInfo) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Serialize `Status`:
	if err = encoder.Encode(obj.Status); err != nil {
		return fmt.Errorf("error while marshaling Status:%w", err)
	}
	// Serialize `Nonce`:
	if err = encoder.Encode(obj.Nonce); err != nil {
		return fmt.Errorf("error while marshaling Nonce:%w", err)
	}
	// Serialize `OrderNum`:
	if err = encoder.Encode(obj.OrderNum); err != nil {
		return fmt.Errorf("error while marshaling OrderNum:%w", err)
	}
	// Serialize `Depth`:
	if err = encoder.Encode(obj.Depth); err != nil {
		return fmt.Errorf("error while marshaling Depth:%w", err)
	}
	// Serialize `CoinDecimals`:
	if err = encoder.Encode(obj.CoinDecimals); err != nil {
		return fmt.Errorf("error while marshaling CoinDecimals:%w", err)
	}
	// Serialize `PcDecimals`:
	if err = encoder.Encode(obj.PcDecimals); err != nil {
		return fmt.Errorf("error while marshaling PcDecimals:%w", err)
	}
	// Serialize `State`:
	if err = encoder.Encode(obj.State); err != nil {
		return fmt.Errorf("error while marshaling State:%w", err)
	}
	// Serialize `ResetFlag`:
	if err = encoder.Encode(obj.ResetFlag); err != nil {
		return fmt.Errorf("error while marshaling ResetFlag:%w", err)
	}
	// Serialize `MinSize`:
	if err = encoder.Encode(obj.MinSize); err != nil {
		return fmt.Errorf("error while marshaling MinSize:%w", err)
	}
	// Serialize `VolMaxCutRatio`:
	if err = encoder.Encode(obj.VolMaxCutRatio); err != nil {
		return fmt.Errorf("error while marshaling VolMaxCutRatio:%w", err)
	}
	// Serialize `AmountWave`:
	if err = encoder.Encode(obj.AmountWave); err != nil {
		return fmt.Errorf("error while marshaling AmountWave:%w", err)
	}
	// Serialize `CoinLotSize`:
	if err = encoder.Encode(obj.CoinLotSize); err != nil {
		return fmt.Errorf("error while marshaling CoinLotSize:%w", err)
	}
	// Serialize `PcLotSize`:
	if err = encoder.Encode(obj.PcLotSize); err != nil {
		return fmt.Errorf("error while marshaling PcLotSize:%w", err)
	}
	// Serialize `MinPriceMultiplier`:
	if err = encoder.Encode(obj.MinPriceMultiplier); err != nil {
		return fmt.Errorf("error while marshaling MinPriceMultiplier:%w", err)
	}
	// Serialize `MaxPriceMultiplier`:
	if err = encoder.Encode(obj.MaxPriceMultiplier); err != nil {
		return fmt.Errorf("error while marshaling MaxPriceMultiplier:%w", err)
	}
	// Serialize `SysDecimalValue`:
	if err = encoder.Encode(obj.SysDecimalValue); err != nil {
		return fmt.Errorf("error while marshaling SysDecimalValue:%w", err)
	}
	// Serialize `Fees`:
	if err = encoder.Encode(obj.Fees); err != nil {
		return fmt.Errorf("error while marshaling Fees:%w", err)
	}
	// Serialize `OutPut`:
	if err = encoder.Encode(obj.OutPut); err != nil {
		return fmt.Errorf("error while marshaling OutPut:%w", err)
	}
	// Serialize `TokenCoin`:
	if err = encoder.Encode(obj.TokenCoin); err != nil {
		return fmt.Errorf("error while marshaling TokenCoin:%w", err)
	}
	// Serialize `TokenPc`:
	if err = encoder.Encode(obj.TokenPc); err != nil {
		return fmt.Errorf("error while marshaling TokenPc:%w", err)
	}
	// Serialize `CoinMint`:
	if err = encoder.Encode(obj.CoinMint); err != nil {
		return fmt.Errorf("error while marshaling CoinMint:%w", err)
	}
	// Serialize `PcMint`:
	if err = encoder.Encode(obj.PcMint); err != nil {
		return fmt.Errorf("error while marshaling PcMint:%w", err)
	}
	// Serialize `LpMint`:
	if err = encoder.Encode(obj.LpMint); err != nil {
		return fmt.Errorf("error while marshaling LpMint:%w", err)
	}
	// Serialize `OpenOrders`:
	if err = encoder.Encode(obj.OpenOrders); err != nil {
		return fmt.Errorf("error while marshaling OpenOrders:%w", err)
	}
	// Serialize `Market`:
	if err = encoder.Encode(obj.Market); err != nil {
		return fmt.Errorf("error while marshaling Market:%w", err)
	}
	// Serialize `SerumDex`:
	if err = encoder.Encode(obj.SerumDex); err != nil {
		return fmt.Errorf("error while marshaling SerumDex:%w", err)
	}
	// Serialize `TargetOrders`:
	if err = encoder.Encode(obj.TargetOrders); err != nil {
		return fmt.Errorf("error while marshaling TargetOrders:%w", err)
	}
	// Serialize `WithdrawQueue`:
	if err = encoder.Encode(obj.WithdrawQueue); err != nil {
		return fmt.Errorf("error while marshaling WithdrawQueue:%w", err)
	}
	// Serialize `TokenTempLp`:
	if err = encoder.Encode(obj.TokenTempLp); err != nil {
		return fmt.Errorf("error while marshaling TokenTempLp:%w", err)
	}
	// Serialize `AmmOwner`:
	if err = encoder.Encode(obj.AmmOwner); err != nil {
		return fmt.Errorf("error while marshaling AmmOwner:%w", err)
	}
	// Serialize `LpAmount`:
	if err = encoder.Encode(obj.LpAmount); err != nil {
		return fmt.Errorf("error while marshaling LpAmount:%w", err)
	}
	// Serialize `ClientOrderId`:
	if err = encoder.Encode(obj.ClientOrderId); err != nil {
		return fmt.Errorf("error while marshaling ClientOrderId:%w", err)
	}
	// Serialize `Padding`:
	if err = encoder.Encode(obj.Padding); err != nil {
		return fmt.Errorf("error while marshaling Padding:%w", err)
	}
	return nil
}

func (obj AmmInfo) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := binary.NewBorshEncoder(buf)
	err := obj.MarshalWithEncoder(encoder)
	if err != nil {
		return nil, fmt.Errorf("error while encoding AmmInfo: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *AmmInfo) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Deserialize `Status`:
	if err = decoder.Decode(&obj.Status); err != nil {
		return fmt.Errorf("error while unmarshaling Status:%w", err)
	}
	// Deserialize `Nonce`:
	if err = decoder.Decode(&obj.Nonce); err != nil {
		return fmt.Errorf("error while unmarshaling Nonce:%w", err)
	}
	// Deserialize `OrderNum`:
	if err = decoder.Decode(&obj.OrderNum); err != nil {
		return fmt.Errorf("error while unmarshaling OrderNum:%w", err)
	}
	// Deserialize `Depth`:
	if err = decoder.Decode(&obj.Depth); err != nil {
		return fmt.Errorf("error while unmarshaling Depth:%w", err)
	}
	// Deserialize `CoinDecimals`:
	if err = decoder.Decode(&obj.CoinDecimals); err != nil {
		return fmt.Errorf("error while unmarshaling CoinDecimals:%w", err)
	}
	// Deserialize `PcDecimals`:
	if err = decoder.Decode(&obj.PcDecimals); err != nil {
		return fmt.Errorf("error while unmarshaling PcDecimals:%w", err)
	}
	// Deserialize `State`:
	if err = decoder.Decode(&obj.State); err != nil {
		return fmt.Errorf("error while unmarshaling State:%w", err)
	}
	// Deserialize `ResetFlag`:
	if err = decoder.Decode(&obj.ResetFlag); err != nil {
		return fmt.Errorf("error while unmarshaling ResetFlag:%w", err)
	}
	// Deserialize `MinSize`:
	if err = decoder.Decode(&obj.MinSize); err != nil {
		return fmt.Errorf("error while unmarshaling MinSize:%w", err)
	}
	// Deserialize `VolMaxCutRatio`:
	if err = decoder.Decode(&obj.VolMaxCutRatio); err != nil {
		return fmt.Errorf("error while unmarshaling VolMaxCutRatio:%w", err)
	}
	// Deserialize `AmountWave`:
	if err = decoder.Decode(&obj.AmountWave); err != nil {
		return fmt.Errorf("error while unmarshaling AmountWave:%w", err)
	}
	// Deserialize `CoinLotSize`:
	if err = decoder.Decode(&obj.CoinLotSize); err != nil {
		return fmt.Errorf("error while unmarshaling CoinLotSize:%w", err)
	}
	// Deserialize `PcLotSize`:
	if err = decoder.Decode(&obj.PcLotSize); err != nil {
		return fmt.Errorf("error while unmarshaling PcLotSize:%w", err)
	}
	// Deserialize `MinPriceMultiplier`:
	if err = decoder.Decode(&obj.MinPriceMultiplier); err != nil {
		return fmt.Errorf("error while unmarshaling MinPriceMultiplier:%w", err)
	}
	// Deserialize `MaxPriceMultiplier`:
	if err = decoder.Decode(&obj.MaxPriceMultiplier); err != nil {
		return fmt.Errorf("error while unmarshaling MaxPriceMultiplier:%w", err)
	}
	// Deserialize `SysDecimalValue`:
	if err = decoder.Decode(&obj.SysDecimalValue); err != nil {
		return fmt.Errorf("error while unmarshaling SysDecimalValue:%w", err)
	}
	// Deserialize `Fees`:
	if err = decoder.Decode(&obj.Fees); err != nil {
		return fmt.Errorf("error while unmarshaling Fees:%w", err)
	}
	// Deserialize `OutPut`:
	if err = decoder.Decode(&obj.OutPut); err != nil {
		return fmt.Errorf("error while unmarshaling OutPut:%w", err)
	}
	// Deserialize `TokenCoin`:
	if err = decoder.Decode(&obj.TokenCoin); err != nil {
		return fmt.Errorf("error while unmarshaling TokenCoin:%w", err)
	}
	// Deserialize `TokenPc`:
	if err = decoder.Decode(&obj.TokenPc); err != nil {
		return fmt.Errorf("error while unmarshaling TokenPc:%w", err)
	}
	// Deserialize `CoinMint`:
	if err = decoder.Decode(&obj.CoinMint); err != nil {
		return fmt.Errorf("error while unmarshaling CoinMint:%w", err)
	}
	// Deserialize `PcMint`:
	if err = decoder.Decode(&obj.PcMint); err != nil {
		return fmt.Errorf("error while unmarshaling PcMint:%w", err)
	}
	// Deserialize `LpMint`:
	if err = decoder.Decode(&obj.LpMint); err != nil {
		return fmt.Errorf("error while unmarshaling LpMint:%w", err)
	}
	// Deserialize `OpenOrders`:
	if err = decoder.Decode(&obj.OpenOrders); err != nil {
		return fmt.Errorf("error while unmarshaling OpenOrders:%w", err)
	}
	// Deserialize `Market`:
	if err = decoder.Decode(&obj.Market); err != nil {
		return fmt.Errorf("error while unmarshaling Market:%w", err)
	}
	// Deserialize `SerumDex`:
	if err = decoder.Decode(&obj.SerumDex); err != nil {
		return fmt.Errorf("error while unmarshaling SerumDex:%w", err)
	}
	// Deserialize `TargetOrders`:
	if err = decoder.Decode(&obj.TargetOrders); err != nil {
		return fmt.Errorf("error while unmarshaling TargetOrders:%w", err)
	}
	// Deserialize `WithdrawQueue`:
	if err = decoder.Decode(&obj.WithdrawQueue); err != nil {
		return fmt.Errorf("error while unmarshaling WithdrawQueue:%w", err)
	}
	// Deserialize `TokenTempLp`:
	if err = decoder.Decode(&obj.TokenTempLp); err != nil {
		return fmt.Errorf("error while unmarshaling TokenTempLp:%w", err)
	}
	// Deserialize `AmmOwner`:
	if err = decoder.Decode(&obj.AmmOwner); err != nil {
		return fmt.Errorf("error while unmarshaling AmmOwner:%w", err)
	}
	// Deserialize `LpAmount`:
	if err = decoder.Decode(&obj.LpAmount); err != nil {
		return fmt.Errorf("error while unmarshaling LpAmount:%w", err)
	}
	// Deserialize `ClientOrderId`:
	if err = decoder.Decode(&obj.ClientOrderId); err != nil {
		return fmt.Errorf("error while unmarshaling ClientOrderId:%w", err)
	}
	// Deserialize `Padding`:
	if err = decoder.Decode(&obj.Padding); err != nil {
		return fmt.Errorf("error while unmarshaling Padding:%w", err)
	}
	return nil
}

func (obj *AmmInfo) Unmarshal(buf []byte) error {
	err := obj.UnmarshalWithDecoder(binary.NewBorshDecoder(buf))
	if err != nil {
		return fmt.Errorf("error while unmarshaling AmmInfo: %w", err)
	}
	return nil
}

func UnmarshalAmmInfo(buf []byte) (*AmmInfo, error) {
	obj := new(AmmInfo)
	err := obj.Unmarshal(buf)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
