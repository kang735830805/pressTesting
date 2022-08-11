package collector

type tpsCollector struct {
	// time info
	Start int64
	End   int64
	// block info
	BlockNum int
	// transaction info of blockchain
	TxNum int
	// tps of blockchain
	//CTps float64
	// bps of blockchain
	//Bps float64
	// sent transaction info
	SentTx int64
	// missed transaction info
	MissedTx int64
	// tps of system
	Tps float64
}