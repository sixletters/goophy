package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

var myvar = 0

func main() {
	fmt.Println("HARRIS WAS HERE DAWG")
	js.Global().Set("formatJSON", jsonWrapper())
	<-make(chan struct{})
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			return "Unable to get document object"
		}

		jsonOuputTextArea := jsDoc.Call("getElementById", "jsonoutput")
		if !jsonOuputTextArea.Truthy() {
			return "Unable to get output text area"
		}
		inputJSON := args[0].String()
		// fmt.Printf("input %s\n", inputJSON)
		pretty, err := prettyJson(inputJSON)
		if err != nil {
			errStr := fmt.Sprintf("unable to parse JSON. Error %s occurred\n", err)
			return errStr
		}
		myvar += 1
		fmt.Println("%d", myvar)
		jsonOuputTextArea.Set("value", pretty)
		return pretty
	})
	return jsonFunc
}

func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}
