package main

import (
	"bytes"
	"fmt"
	v1 "gogen-enum-demo/schema/v1"
	v2 "gogen-enum-demo/schema/v2"
)

func main() {
	source := v2.Example{
		Enumfield: v2.MyEnumTHREE,
	}

	fmt.Println(source) // Prints '{THREE}'

	avroBuffer := &bytes.Buffer{}

	err := source.Serialize(avroBuffer)
	if err != nil {
		panic(err)
	}
	avroBytes := avroBuffer.Bytes()

	target, err := v1.DeserializeExampleFromSchema(bytes.NewReader(avroBytes), source.Schema())
	if err != nil {
		panic(err)
	}
	fmt.Println(target) // Prints '{unknown}', expected {ONE} per field default.

}
