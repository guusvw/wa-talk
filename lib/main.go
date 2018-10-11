package main

import (
	"log"
	"strconv"
	"syscall/js"
)

func getElementById(i []js.Value, id int) js.Value {
	log.Printf("access `getElementById(%d)`", id)
	return js.Global().Get("document").Call("getElementById", i[id].String())
}

func add(i []js.Value) {
	value1 := getElementById(i, 0).Get("value").String()
	value2 := getElementById(i, 1).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)
}

func subtract(i []js.Value) {
	value1 := getElementById(i, 0).Get("value").String()
	value2 := getElementById(i, 1).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1-int2)
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
	js.Global().Set("subtract", js.NewCallback(subtract))
}

func main() {
	c := make(chan struct{}, 0)

	log.Println("Go WebAssembly Initialized")
	registerCallbacks()

	<-c
}
