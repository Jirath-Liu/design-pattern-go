package builder

// DataBuilder Data建造者
type DataBuilder struct {
	_A int
	_B string
	_C float64
}

func (b *DataBuilder) A(a int) *DataBuilder {
	b._A = a
	return b
}

func (b *DataBuilder) B(data string) *DataBuilder {
	b._B = data
	return b
}

func (b *DataBuilder) C(data float64) *DataBuilder {
	b._C = data
	return b
}

func (b DataBuilder) build() Data {
	return Data{
		A: b._A,
		B: b._B,
		C: b._C,
	}
}
