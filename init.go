package decoder_vector

import "github.com/koykov/decoder"

func init() {
	_ = decoder.RegisterPool("jsonvector", &ipool{fmtJSON})
	_ = decoder.RegisterPool("xmlvector", &ipool{fmtXML})
	_ = decoder.RegisterPool("yamlvector", &ipool{fmtYAML})
	_ = decoder.RegisterPool("urlvector", &ipool{fmtURL})
	_ = decoder.RegisterPool("halvector", &ipool{fmtHAL})

	decoder.RegisterModFnNS("vector", "parseJSON", "", modParseJSON)
	decoder.RegisterModFnNS("vector", "parseXML", "", modParseXML)
	decoder.RegisterModFnNS("vector", "parseYAML", "", modParseYAML)
	decoder.RegisterModFnNS("vector", "parseURL", "", modParseURL)
	decoder.RegisterModFnNS("vector", "parseHAL", "", modParseHAL)

	decoder.RegisterModFnNS("vector", "coalesce", "", modCoalesce)
	decoder.RegisterModFnNS("vector", "marshal", "serialize", modMarshal)
}
