package decoder_vector

import (
	"github.com/koykov/decoder"
	_ "github.com/koykov/vector_inspector"
)

func init() {
	_ = decoder.RegisterPool("jsonvector", &ipool{fmtJSON})
	_ = decoder.RegisterPool("xmlvector", &ipool{fmtXML})
	_ = decoder.RegisterPool("yamlvector", &ipool{fmtYAML})
	_ = decoder.RegisterPool("urlvector", &ipool{fmtURL})
	_ = decoder.RegisterPool("halvector", &ipool{fmtHAL})

	decoder.RegisterModFnNS("vector", "parseJSON", "", modParseJSON).
		WithParam("data string", "JSON source to parse.").
		WithDescription("Parse `data` using jsonvector and return vector instance.").
		WithExample(`ctx.SetString("source", "{"x":{"y":{"z":"foobar"}}}")
---
ctx.data = vector::parseJSON(source).(vector)
obj.Name = data.x.y.z // foobar`)
	decoder.RegisterModFnNS("vector", "parseXML", "", modParseXML).
		WithParam("data string", "XML source to parse.").
		WithDescription("Parse `data` using xmlvector and return vector instance.").
		WithExample(`ctx.SetString("source", "<?xml version="1.0" encoding="UTF-8"?><x><y><z>foobar</z></y></x>")
---
ctx.data = vector::parseXML(source).(vector)
obj.Name = data.x.y.z // foobar`)
	decoder.RegisterModFnNS("vector", "parseYAML", "", modParseYAML).
		WithParam("data string", "YAML source to parse.").
		WithDescription("Parse `data` using yamlvector and return vector instance.").
		WithNote("CAUTION! Still not implement.")
	decoder.RegisterModFnNS("vector", "parseURL", "", modParseURL).
		WithParam("data string", "URL to parse.").
		WithDescription("Parse `data` using urlvector and return vector instance.").
		WithExample(`ctx.SetString("source", "http:://127.0.0.1:8080/post?xyz=foobar")
---
ctx.data = vector::parseURL(source).(vector)
obj.Name = data.query.xyz // foobar`)
	decoder.RegisterModFnNS("vector", "parseHAL", "", modParseHAL).
		WithParam("data string", "HAL string to parse.").
		WithDescription("Parse `data` using halvector and return vector instance.").
		WithExample(`ctx.SetString("source", "zh-Hant-cn;q=1, zh-cn;q=0.6, zh;q=0.4")
---
ctx.data = vector::parseHAL(source).(vector)
obj.Name = data.0.code // zh`)

	decoder.RegisterModFnNS("vector", "coalesce", "", modCoalesce).
		WithParam("args ...string", "Keys to read.").
		WithDescription("Return value of first non-empty key in vector object.").
		WithExample(`// source: {"x":{"y":{"z":"foobar"}}}
---
ctx.data = source|vector::parseJSON().(vector)
obj.Name = data|vector::coalesce("x.y.z.a.b.c", "x.y.z.a.b", "x.y.z.a", "x.y.z") // foobar`)
	decoder.RegisterModFnNS("vector", "marshal", "serialize", modMarshal).
		WithParam("data vector|node", "Vector or node object.").
		WithDescription("Serialize vector|node to string according format.").
		WithExample(`// source: {"x":{"y":{"z":"foobar"}}}
---
ctx.data = source|vector::parseJSON().(vector)
obj.Name = vector::marshal(data.x.y) // "foobar"`)
}
