# Decoder vector bindings

Provide [vector](https://github.com/koykov/vector) and [vector_inspector](https://github.com/koykov/vector_inspector)
features to use in [dyntpl](https://github.com/koykov/decoder).

### Usage

```go
import _ "github.com/koykov/decoder_vector"

const dec = `ctx.data = vector::parseJSON(source) as vector
dst.Name = data.name`

const source = `{"name":"foobar"}`

func main() {
	ctx := decoder.AcquireCtx()
	ctx.SetStatic("source", source)
	ctx.Set("dst", dstObj, ObjInspector{})
	decoder.Decode("example", ctx)
	println(dstObj.Name) // foobar
}
```
