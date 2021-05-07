package main

import (
	"design-pattern-go/create/builder"
	"design-pattern-go/create/factory/simple"
	"fmt"
	"reflect"
)

func main() {
	o := simple.BaseInfo{Name: "name"}
	tv := reflect.ValueOf(o)
	for i := 0; i < tv.NumField(); i++ {
		ct := tv.Field(i)
		if ct.CanInterface() {
			fmt.Printf(" %s %v %s \n", ct.Type(), ct.Interface(), ct)
		} else {
			fmt.Printf(" %s %s \n", ct.Type(), ct)
		}
	}

	d := builder.NewBuilder().
		A(2).
		B("b").
		C(2.3).
		Build()
	newD := d
	fmt.Printf(" %t ", &newD == &d)
}
