package simple

type BaseInfo struct {
	Name string
}

func (b BaseInfo) GetName() string {
	return b.Name
}
