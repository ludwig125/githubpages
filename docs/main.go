package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	c := make(chan struct{})

	fmt.Println("Hello, WebAssembly!")
	registerCallbacks()
	<-c
}

func add(this js.Value, args []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", args[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", args[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	fmt.Println("int1:", int1, " int2:", int2)

	js.Global().Set("output", int1+int2)
	println(int1 + int2)
	answer(int1 + int2)
	return nil
}

func subtract(this js.Value, args []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", args[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", args[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	fmt.Println("int1:", int1, " int2:", int2)

	js.Global().Set("output", int1+int2)
	println(int1 - int2)
	answer(int1 - int2)
	return nil
}

func answer(ans int) {
	js.Global().Get("document").Call("getElementById", "answer").Set("innerHTML", ans)
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}
