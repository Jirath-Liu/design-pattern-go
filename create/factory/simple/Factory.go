package simple

import "reflect"

type Factory struct {
	types map[string]reflect.Type
}

func NewFactory() *Factory {
	m := make(map[string]reflect.Type)
	m["one"]=reflect.TypeOf(One{})
	m["two"]=reflect.TypeOf(Two{})
	m["three"]=reflect.TypeOf(Three{})
	return &Factory{types: m}
}

func (f Factory) Build(name string,typeName string) (Goods,bool)  {
	var t reflect.Type
	t,ok := f.types[typeName]
	if !ok {
		return nil,ok
	}
	g := reflect.New(t)
	g.FieldByName("name").SetString(name)
	if value,ok := g.Interface().(Goods);ok{
		return value,true
	}
	return nil,false
}

