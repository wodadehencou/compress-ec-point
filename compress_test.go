package compressecpoint

import (
	"crypto/elliptic"
	"crypto/rand"
	"testing"
)

func TestCompress(t *testing.T) {
	curves := []elliptic.Curve{elliptic.P224(), elliptic.P256(), elliptic.P384(), elliptic.P521()}

	for _, c := range curves {

		_, x, y, err := elliptic.GenerateKey(c, rand.Reader)
		if err != nil {
			t.Error(err)
			return
		}

		compressed := Compress(c, x, y)
		xx, yy := Decompress(c, compressed)

		if xx == nil {
			t.Error("failed to decompress")
			break
		}
		if xx.Cmp(x) != 0 || yy.Cmp(y) != 0 {
			t.Error("Decompress returned different values")
			break
		}
	}
}

func BenchmarkP256Decompress(b *testing.B) {
	c := elliptic.P256()
	_, x, y, _ := elliptic.GenerateKey(c, rand.Reader)
	compressed := Compress(c, x, y)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Decompress(c, compressed)
	}
}
