package abstract

import (
	"design-pattern-go/create/factory/simple"
)

type Factory interface {
	buildGoods(name string) simple.Goods
}
