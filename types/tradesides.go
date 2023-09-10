package types

type TradeSide string

const (
	SELL TradeSide = "SELL"
	BUY  TradeSide = "BUY"
)

var TradeSides = []TradeSide{SELL, BUY}
