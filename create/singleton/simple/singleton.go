package simple

type Singleton struct {
	data string
}

var s = &Singleton{data: "ss"}

func Get() *Singleton {
	return s
}
