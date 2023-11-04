# Decoder vector bindings

Provide [vector](https://github.com/koykov/vector) and [vector_inspector](https://github.com/koykov/vector_inspector)
features to use in [dyntpl](https://github.com/koykov/decoder).

### Usage

```go
package main

import (
	"github.com/koykov/decoder"
	"github.com/koykov/inspector/testobj"
	"github.com/koykov/inspector/testobj_ins"

	_ "github.com/koykov/decoder_vector"   // register vector bindings
	_ "github.com/koykov/vector_inspector" // register vector inspector
)

const (
	dec = `ctx.data = vector::parseJSON(source).(vector)
obj.Name = data.x.y.z`
	json = `{"x":{"y":{"z":"foobar"}}}`
)

func main() {
	ruleset, _ := decoder.Parse([]byte(dec))
	decoder.RegisterDecoder("example", ruleset)
	ctx := decoder.NewCtx()
	var obj testobj.TestObject
	ctx.SetStatic("source", json)
	ctx.Set("obj", &obj, testobj_ins.TestObjectInspector{})
	_ = decoder.Decode("example", ctx)
	println(string(obj.Name)) // output: foobar
}

```
