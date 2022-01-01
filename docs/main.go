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
	value1 := textBoxDataToStr(args[0].String())
	value2 := textBoxDataToStr(args[1].String())

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)
	fmt.Println("int1:", int1, " int2:", int2)
	ans := int1 + int2

	printAnswer(ans)
	return nil
}

func subtract(this js.Value, args []js.Value) interface{} {
	value1 := textBoxDataToStr(args[0].String())
	value2 := textBoxDataToStr(args[1].String())

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)
	fmt.Println("int1:", int1, " int2:", int2)
	ans := int1 - int2

	printAnswer(ans)
	return nil
}

func printAnswer(ans int) {
	println(ans)
	js.Global().Get("document").Call("getElementById", "answer").Set("innerHTML", ans)
}

func textBoxDataToStr(s string) string {
	return js.Global().Get("document").Call("getElementById", s).Get("value").String()
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}
