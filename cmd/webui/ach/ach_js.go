package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/describe"
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

func parseReadable(file *ach.File) (string, error) {
	var buf bytes.Buffer
	opts := describe.Opts{MaskNames: false, MaskAccountNumbers: false}
	describe.File(&buf, file, &opts)
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
			return "Invalid number of arguments passed"
		}

		if json.Valid([]byte(args[0].String())) {
			return args[0].String()
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
			return "Invalid number of arguments passed"
		}

		if !json.Valid([]byte(args[0].String())) {
			return args[0].String()
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

func printReadable() js.Func {
	return js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments passed"
		}

		file, err := parseFile(args[0].String())
		if err != nil {
			msg := fmt.Sprintf("unable to parse ach file: %v", err)
			fmt.Println(msg)
			return msg
		}

		parsed, err := parseReadable(file)
		if err != nil {
			fmt.Printf("unable to convert ach file to human-readable format %s\n", err)
			return "There was an error formatting the output"
		}

		return parsed
	})
}

// Parses input, either JSON or Nacha format to an ach.File
func parseFile(input string) (*ach.File, error) {
	if json.Valid([]byte(input)) {
		file, err := ach.FileFromJSON([]byte(input))
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	file, err := ach.NewReader(strings.NewReader(input)).Read()
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func writeVersion() {
	span := js.Global().Get("document").Call("querySelector", "#version")
	span.Set("innerHTML", fmt.Sprintf("Version: %s", ach.Version))
}

func main() {
	js.Global().Set("parseACH", prettyPrintJSON())
	js.Global().Set("parseJSON", printACH())
	js.Global().Set("parseReadable", printReadable())

	writeVersion()

	<-make(chan bool)
}
