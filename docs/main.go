package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	c := make(chan struct{})

	registerCallbacks()

	// getElementByID("in").Call("addEventListener", "keyup", js.FuncOf(func(js.Value, []js.Value) interface{} {
	// 	getElementByID("out").Set("value", getElementByID("in").Get("value"))
	// 	return nil
	// }))
	unixtime()

	<-c
}

func getElementByID(targetID string) js.Value {
	return js.Global().Get("document").Call("getElementById", targetID)
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}

func add(this js.Value, args []js.Value) interface{} {
	value1 := textToStr(args[0])
	value2 := textToStr(args[1])

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)
	fmt.Println("int1:", int1, " int2:", int2)
	ans := int1 + int2

	printAnswer(ans)
	return nil
}

func subtract(this js.Value, args []js.Value) interface{} {
	value1 := textToStr(args[0])
	value2 := textToStr(args[1])

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)
	fmt.Println("int1:", int1, " int2:", int2)
	ans := int1 - int2

	printAnswer(ans)
	return nil
}

func textToStr(v js.Value) string {
	return js.Global().Get("document").Call("getElementById", v.String()).Get("value").String()
}

func printAnswer(ans int) {
	println(ans)
	js.Global().Get("document").Call("getElementById", "answer").Set("innerHTML", ans)
}
