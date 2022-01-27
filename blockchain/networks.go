package blockchain

type Network uint8

const (
	Ethereum Network = iota
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
