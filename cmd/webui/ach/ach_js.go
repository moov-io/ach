package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/moov-io/ach"
)

func parseACH(input string) (string, error) {
	r := strings.NewReader(input)
	file, err := ach.NewReader(r).Read()
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(file); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func prettyJson(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func prettyPrintJSON() js.Func {
	return js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		parsed, err := parseACH(args[0].String())
		if err != nil {
			msg := fmt.Sprintf("unable to parse ach file: %v", err)
			fmt.Println(msg)
			return msg
		}
		pretty, err := prettyJson(parsed)
		if err != nil {
			fmt.Printf("unable to convert ach file to json %s\n", err)
			return "There was an error converting the json"
		}

		return pretty
	})
}

func printACH() js.Func {
	return js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		file, err := ach.FileFromJSON([]byte(args[0].String()))
		if err != nil {
			msg := fmt.Sprintf("unable to parse json ACH file: %v", err)
			fmt.Println(msg)
			return msg
		}
		var buf bytes.Buffer
		if err := ach.NewWriter(&buf).Write(file); err != nil {
			msg := fmt.Sprintf("problem writing ACH file: %v", err)
			fmt.Println(msg)
			return msg
		}
		return buf.String()
	})
}

func main() {
	js.Global().Set("parseACH", prettyPrintJSON())
	js.Global().Set("parseJSON", printACH())
	<-make(chan bool)
}
