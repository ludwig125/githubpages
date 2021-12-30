package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{})

	fmt.Println("Hello, WebAssembly!")
	registerCallbacks()
	<-c
}

func add(this js.Value, args []js.Value) interface{} {
	println(js.ValueOf(args[0].Int() + args[1].Int()).String())
	return nil
}

func subtract(this js.Value, args []js.Value) interface{} {
	println(js.ValueOf(args[0].Int() - args[1].Int()).String())
	return nil
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}
