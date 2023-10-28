package decoder_vector

import (
	"github.com/koykov/halvector"
	"github.com/koykov/jsonvector"
	"github.com/koykov/urlvector"
	"github.com/koykov/xmlvector"
	"github.com/koykov/yamlvector"
)

type fmt_ uint8

const (
	fmtJSON fmt_ = iota
	fmtXML
	fmtYAML
	fmtURL
	fmtHAL
)

type ipool struct {
	fmt_ fmt_
}

func (p *ipool) Get() any {
	switch p.fmt_ {
	case fmtXML:
		return xmlvector.Acquire()
	case fmtYAML:
		return yamlvector.Acquire()
	case fmtURL:
		return urlvector.Acquire()
	case fmtHAL:
		return halvector.Acquire()
	case fmtJSON:
		fallthrough
	default:
		return jsonvector.Acquire()
	}
}

func (p *ipool) Reset(x any) {
	switch p.fmt_ {
	case fmtJSON:
		vec, ok := x.(*jsonvector.Vector)
		if !ok {
			return
		}
		vec.Reset()
	case fmtXML:
		vec, ok := x.(*xmlvector.Vector)
		if !ok {
			return
		}
		vec.Reset()
	case fmtYAML:
		vec, ok := x.(*yamlvector.Vector)
		if !ok {
			return
		}
		vec.Reset()
	case fmtURL:
		vec, ok := x.(*urlvector.Vector)
		if !ok {
			return
		}
		vec.Reset()
	case fmtHAL:
		vec, ok := x.(*halvector.Vector)
		if !ok {
			return
		}
		vec.Reset()
	}
}

func (p *ipool) Put(x any) {
	switch p.fmt_ {
	case fmtJSON:
		vec, ok := x.(*jsonvector.Vector)
		if !ok {
			return
		}
		jsonvector.Release(vec)
	case fmtXML:
		vec, ok := x.(*xmlvector.Vector)
		if !ok {
			return
		}
		xmlvector.Release(vec)
	case fmtYAML:
		vec, ok := x.(*yamlvector.Vector)
		if !ok {
			return
		}
		yamlvector.Release(vec)
	case fmtURL:
		vec, ok := x.(*urlvector.Vector)
		if !ok {
			return
		}
		urlvector.Release(vec)
	case fmtHAL:
		vec, ok := x.(*halvector.Vector)
		if !ok {
			return
		}
		halvector.Release(vec)
	}
}
