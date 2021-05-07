package abstract

type BaseInfo struct {
	name string
}

func NewBaseInfo(name string) BaseInfo {
	return BaseInfo{name: name}
}
func (b BaseInfo) GetName() string {
	return b.name
}
