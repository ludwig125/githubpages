package main

import (
	"fmt"
	"strconv"
	"syscall/js"
	"time"
)

func unixtime() {
	js.Global().Call("setInterval", js.FuncOf(clock), "200")

	getElementByID("in").Call("addEventListener", "input", js.FuncOf(convTime))

}

func clock(this js.Value, args []js.Value) interface{} {
	nowStr, nowUnix := getNow(time.Now())

	getElementByID("clock").Set("textContent", nowStr)
	getElementByID("clock_unixtime").Set("textContent", nowUnix)
	return nil
}

func convTime(this js.Value, args []js.Value) interface{} {
	in := getElementByID("in").Get("value").String()
	date, err := unixtimeToDate(in))
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
