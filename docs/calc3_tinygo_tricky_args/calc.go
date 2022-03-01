package main

import (
	"errors"
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/ludwig125/githubpages/docs/calc3_tinygo_tricky_args/args"
	"github.com/mailru/easyjson"
)

func main() {
}

var buf [1024]byte

//export getBuffer
func getBuffer() *byte {
	return &buf[0]
}

//export calcAdd
func calcAdd(x, y string) map[string]interface{} {
	println("yahooooo")
	fmt.Printf("calcAdd x,y: '%s %s'\n", x, y)
	return calculatorWrapper("add", x, y)
}

//export calcAdd2
func calcAdd2(x string) map[string]interface{} {
	println("yahooooo2")
	fmt.Printf("calcAdd2 x,y: '%s'\n", x)
	return calculatorWrapper("add", x, x)
}

func SetUint8ArrayInGo(this js.Value, args []js.Value) interface{} {
	_ = js.CopyBytesToJS(args[0], []byte{0, 9, 21, 32})
	return nil
}

// //export calcAdd3
// func calcAdd3(s string) map[string]interface{} {
// 	fmt.Printf("yahooooo2: '%s'\n", s)
// 	x, y, err := decodeArgs(s)
// 	if err != nil {
// 		return wrapResult("", fmt.Errorf("failed to decodeArgs: %v", err))
// 	}
// 	fmt.Printf("calcAdd3 x,y: '%s %s'\n", x, y)
// 	return calculatorWrapper("add", x, y)
// 	// return nil
// }

//export calcAdd3
func calcAdd3(s string) string {
	fmt.Printf("yahooooo2: '%s'\n", s)
	x, y, err := decodeArgs(s)
	if err != nil {
		return fmt.Sprintf("failed to decodeArgs: %v", err)
	}
	fmt.Printf("calcAdd3 x,y: '%s %s'\n", x, y)
	// return calculatorWrapper("add", x, y)
	calculatorWrapper("add", x, y)
	return "aaab"
	// return nil
}

func decodeArgs(s string) (string, string, error) {
	var a args.Args
	fmt.Printf("decode: '%s'\n", s)
	// s=`{"X":"1","Y":"2"}`
	if err := easyjson.Unmarshal([]byte(s), &a); err != nil {
		return "", "", err
	}
	fmt.Printf("after decode %+v\n", a)
	return a.X, a.Y, nil
}

// //export add
// func add(x int, y int) int {
// 	fmt.Println("add x,y",x, y)
// 	return x + y
//   }

//export calcSubtract
func calcSubtract(x, y string) map[string]interface{} {
	return calculatorWrapper("subtract", x, y)
}

func calculatorWrapper(ope, value1, value2 string) map[string]interface{} {
	fmt.Printf("calculatorWrapper calcAdd2 x,y: '%s %s'\n", value1, value2)

	// func calculatorWrapper(ope string) js.Func {
	// calcFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// value1, err := getJSValue(val1)
	// if err != nil {
	// 	return wrapResult("", err)
	// }
	// value2, err := getJSValue(val2)
	// if err != nil {
	// 	return wrapResult("", err)
	// }
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
	// })
	// return calcFunc
}

func getJSValue(elemID string) (string, error) {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		fmt.Println("!jsDoc.Truthy() ")
		return "", errors.New("failed to get document object")
	}

	jsElement := jsDoc.Call("getElementById", elemID)
	if !jsElement.Truthy() {
		fmt.Println(" !jsElement.Truthy() ")
		return "", fmt.Errorf("failed to getElementById: %s", elemID)
	}

	jsValue := jsElement.Get("value")
	if !jsValue.Truthy() {
		fmt.Println("!jsValue.Truthy() ")
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
