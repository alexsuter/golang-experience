package main

import ("fmt"
	"reflect"
	)

func main() {
	i := 10
	v := reflect.ValueOf(&i)
	e := v.Elem()
	e.SetInt(1)

	fmt.Println("t")
}
