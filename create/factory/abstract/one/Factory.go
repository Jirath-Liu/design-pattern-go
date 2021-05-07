package one

import (
	"design-pattern-go/create/factory/simple"
)

type Factory struct {
}

func (f Factory) buildGoods(name string) simple.Goods {
	b := simple.NewBaseInfo(name)
	return simple.One{BaseInfo: b}
}
