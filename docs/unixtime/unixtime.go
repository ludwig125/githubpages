package main

import (
	"errors"
	"fmt"
	"strconv"
	"syscall/js"
	"time"
)

func main() {
	unixtime()

	<-make(chan struct{})
}

func unixtime() {
	// js.Global().Set("add", js.FuncOf(addFunction))
	// js.Global().Set("setTimeZoneFunc", js.FuncOf(setTimeZone))

	// time zoneを最初に表示させる
	js.Global().Call("queueMicrotask", js.FuncOf(setTimeZone))
	js.FuncOf(setTimeZone).Release()

	// 一定時間おきにclockを呼び出す
	js.Global().Call("setInterval", js.FuncOf(clock), "200")

	getElementByID("in").Call("addEventListener", "input", js.FuncOf(convTime))
}

func setTimeZone(this js.Value, args []js.Value) interface{} {
	t := time.Now()
	zone, _ := t.Zone()
	return setJSValue("time_zone", fmt.Sprintf("(%s)", zone))
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

func getElementByID(targetID string) js.Value {
	return js.Global().Get("document").Call("getElementById", targetID)
}

func clock(this js.Value, args []js.Value) interface{} {
	nowStr, nowUnix := getNow(time.Now())

	getElementByID("clock").Set("textContent", nowStr)
	getElementByID("clock_unixtime").Set("textContent", nowUnix)
	return nil
}

func convTime(this js.Value, args []js.Value) interface{} {
	in := getElementByID("in").Get("value").String()
	date, err := unixtimeToDate(in)
	if err != nil {
		getElementByID("out").Set("value", js.ValueOf("不正な時刻です"))
		return nil
	}
	getElementByID("out").Set("value", js.ValueOf(date))
	return nil
}

func getNow(now time.Time) (string, string) {
	s := now.Format("2006-01-02 15:04:05")
	unix := now.Unix()
	return s, fmt.Sprintf("%d", unix)
}

func unixtimeToDate(s string) (string, error) {
	unixtime, err := strconv.Atoi(s)
	if err != nil {
		return "", err
	}
	date := time.Unix(int64(unixtime), 0)
	layout := "2006-01-02 15:04:05" // Goの時刻フォーマットではこれで時分秒まで取れる
	return date.Format(layout), nil
}
