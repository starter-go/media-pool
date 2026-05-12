package dxo

type URL string

func (u URL) String() string {
	return string(u)
}
