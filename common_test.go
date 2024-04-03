package decoder_vector

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/koykov/decoder"
	"github.com/koykov/inspector/testobj"
	"github.com/koykov/inspector/testobj_ins"
)

type stage struct {
	key                    string
	origin, source, expect []byte
}

var stages []stage

func init() {
	registerTestStages("parser")
	registerTestStages("mod")
}

func registerTestStages(dir string) {
	_ = filepath.Walk("testdata/"+dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".dec" {
			st := stage{}
			st.key = strings.Replace(filepath.Base(path), ".dec", "", 1)

			st.origin, _ = os.ReadFile(path)
			ruleset, err := decoder.Parse(st.origin)
			if err != nil {
				println(err.Error())
			}

			st.source, _ = os.ReadFile(strings.Replace(path, ".dec", ".source.txt", 1))
			st.source = bytes.Trim(st.source, "\n")

			st.expect, _ = os.ReadFile(strings.Replace(path, ".dec", ".expect.txt", 1))
			st.expect = bytes.Trim(st.expect, "\n")
			stages = append(stages, st)

			decoder.RegisterDecoder(st.key, ruleset)
		}
		return nil
	})
}

func getStage(key string) (st *stage) {
	for i := 0; i < len(stages); i++ {
		st1 := &stages[i]
		if st1.key == key {
			st = st1
		}
	}
	return st
}

func getTBName(tb testing.TB) string {
	key := tb.Name()
	return key[strings.Index(key, "/")+1:]
}

func testStage(t *testing.T) {
	key := getTBName(t)
	st := getStage(key)
	if st == nil {
		t.Error("stage not found")
		return
	}
	ctx := decoder.NewCtx()
	var obj testobj.TestObject
	ctx.Set("obj", &obj, testobj_ins.TestObjectInspector{})
	ctx.SetStatic("source", st.source)
	err := decoder.Decode(key, ctx)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(obj.Name, st.expect) {
		t.FailNow()
	}
}

func benchStage(b *testing.B) {
	b.ReportAllocs()
	key := getTBName(b)
	st := getStage(key)
	if st == nil {
		b.Error("stage not found")
		return
	}
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		ctx := decoder.AcquireCtx()
		ctx.SetStatic("source", &st.source)
		_ = decoder.Decode("json", ctx)
		decoder.ReleaseCtx(ctx)
		buf.Reset()
	}
}
