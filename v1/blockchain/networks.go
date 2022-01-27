package blockchain

type Network uint8

const (
	Undefined Network = iota
	Ethereum
	Polygon
)

func (b Network) String() string {
	switch b {
	case Ethereum:
		return "ethereum"
	case Polygon:
		return "polygon"
	}
	return "unknown"
}
