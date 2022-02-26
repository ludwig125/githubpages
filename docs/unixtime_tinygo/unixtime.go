package main

import (
	"errors"
	"fmt"
	"strconv"
	"syscall/js"
	"time"
)

func main() {}

//export setTimeZone
func setTimeZone() {
	t := time.Now()
	zone, _ := t.Zone()
	setJSValue("time_zone", fmt.Sprintf("(%s)", zone))
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

//export clock
func clock() {
	nowStr, nowUnix := getNow(time.Now())

	getElementByID("clock").Set("textContent", nowStr)
	getElementByID("clock_unixtime").Set("textContent", nowUnix)
}

//export convTime
func convTime() {
	in := getElementByID("in").Get("value").String()
	date, err := unixtimeToDate(in)
	if err != nil {
		getElementByID("out").Set("value", js.ValueOf("不正な時刻です"))
		return
	}
	getElementByID("out").Set("value", js.ValueOf(date))
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
