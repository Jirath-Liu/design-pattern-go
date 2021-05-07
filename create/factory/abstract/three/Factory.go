package three

import (
	"design-pattern-go/create/factory/abstract"
	"design-pattern-go/create/factory/simple"
)

type Factory struct {
}

func (f Factory) buildGoods(name string) simple.Goods {
	b := abstract.NewBaseInfo(name)
	return abstract.Three{BaseInfo: b}
}
