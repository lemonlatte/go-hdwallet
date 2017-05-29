package hdwallet

type CoinType string

const (
	BTC CoinType = "BTC"
	LTC CoinType = "LTC"
)

var PubKeyAddrPrefix = map[CoinType]byte{
	BTC: 0,
	LTC: 48,
}

var PubKeyAddrTestPrefix = map[CoinType]byte{
	BTC: 111,
	LTC: 111,
}

var ScriptAddrPrefix = map[CoinType]byte{
	BTC: 05,
	LTC: 50,
}

var ScriptAddrTestPrefix = map[CoinType]byte{
	BTC: 196,
	LTC: 196,
}
