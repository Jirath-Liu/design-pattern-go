package abstract

type One struct {
	BaseInfo
}

func (o One) GetName() string {
	return "one" + o.BaseInfo.GetName()
}

type Two struct {
	BaseInfo
}

func (o Two) GetName() string {
	return "one" + o.BaseInfo.GetName()
}

type Three struct {
	BaseInfo
}

func (o Three) GetName() string {
	return "one" + o.BaseInfo.GetName()
}
