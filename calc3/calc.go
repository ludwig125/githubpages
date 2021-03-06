package main

import (
	"errors"
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	registerCallbacks()
	<-make(chan struct{})
}

func registerCallbacks() {
	js.Global().Set("calcAdd", calculatorWrapper("add"))
	js.Global().Set("calcSubtract", calculatorWrapper("subtract"))
}

func calculatorWrapper(ope string) js.Func {
	calcFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		value1, err := getJSValue(args[0].String())
		if err != nil {
			return wrapResult("", err)
		}
		value2, err := getJSValue(args[1].String())
		if err != nil {
			return wrapResult("", err)
		}
		fmt.Println("value1:", value1, " value2:", value2)

		int1, err := strconv.Atoi(value1)
		if err != nil {
			return wrapResult("", fmt.Errorf("failed to convert value1 to int: %v", err))
		}
		int2, err := strconv.Atoi(value2)
		if err != nil {
			return wrapResult("", fmt.Errorf("failed to convert value2 to int: %v", err))
		}

		var ans int
		switch ope {
		case "add":
			ans = int1 + int2
		case "subtract":
			ans = int1 - int2
		default:
			return wrapResult("", fmt.Errorf("invalid operation: %s", ope))
		}
		fmt.Println("Answer:", ans)

		if err := setJSValue("answer", ans); err != nil {
			return wrapResult("", err)
		}
		return nil
	})
	return calcFunc
}

func getJSValue(elemID string) (string, error) {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return "", errors.New("failed to get document object")
	}

	jsElement := jsDoc.Call("getElementById", elemID)
	if !jsElement.Truthy() {
		return "", fmt.Errorf("failed to getElementById: %s", elemID)
	}

	jsValue := jsElement.Get("value")
	if !jsValue.Truthy() {
		return "", fmt.Errorf("failed to Get value: %s", elemID)
	}
	return jsValue.String(), nil
}

func setJSValue(elemID string, value interface{}) error {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return errors.New("failed to get document object")
	}

	jsElement := jsDoc.Call("getElementById", elemID)
	if !jsElement.Truthy() {
		return fmt.Errorf("failed to getElementById: %s", elemID)
	}
	jsElement.Set("innerHTML", value)
	return nil
}

func wrapResult(result string, err error) map[string]interface{} {
	return map[string]interface{}{
		"error":    err.Error(),
		"response": result,
	}
}
