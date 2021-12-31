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
	println(args[0].Int() + args[1].Int())
	return nil
}

func subtract(this js.Value, args []js.Value) interface{} {
	println(args[0].Int() - args[1].Int())
	return nil
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}
