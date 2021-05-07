package simple

type BaseInfo struct {
	name string
}

func (b BaseInfo) GetName() string {
	return b.name
}

