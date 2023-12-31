package decoder_vector

import "testing"

func TestParser(t *testing.T) {
	t.Run("json", testStage)
	t.Run("xml", testStage)
	// // t.Run("yaml", testStage) // todo uncomment after implement yamlvector
	t.Run("url", testStage)
	t.Run("hal", testStage)
}

func BenchmarkParser(b *testing.B) {
	b.Run("json", benchStage)
	b.Run("xml", benchStage)
	// b.Run("yaml", benchStage)  // todo uncomment after implement yamlvector
	b.Run("url", benchStage)
	b.Run("hal", benchStage)
}
