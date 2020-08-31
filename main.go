package main

import (
	"fmt"
	"syscall/js"
)

var (
	document = js.Global().Get("document")
	body     = document.Get("body")

	react    = js.Global().Get("React")
	reactDOM = js.Global().Get("ReactDOM")

	materialUI = js.Global().Get("MaterialUI")
	muiButton  = materialUI.Get("Button")
)

func render(element, container js.Value) js.Value {
	return reactDOM.Call("render", element, container)
}

func createElement(typ js.Value, props map[string]interface{}, children ...interface{}) js.Value {

	// TODO a TinyGo workaround
	if props == nil {
		props = map[string]interface{}{}
	}

	return react.Call("createElement",
		typ,
		props,
		children,
	)
}

func app() js.Value {
	return createElement(
		muiButton,
		map[string]interface{}{
			"onClick": js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				fmt.Println("asdf")
				return nil
			}),
		},
		"Button",
	)
}

func main() {
	println("howdy~ from goiso")

	body.Set("style", "margin:0")

	render(
		app(),
		document.Call("getElementById", "root"),
	)

	select {}
}
