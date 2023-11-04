package decoder_vector

import "testing"

func TestMod(t *testing.T) {
	t.Run("coalesce", testStage)
	t.Run("marshal", testStage)
}

func BenchmarkMod(b *testing.B) {
	b.Run("coalesce", benchStage)
	b.Run("marshal", benchStage)
}
