package hashish

import (
	"github.com/google/uuid"
	"testing"
)

func TestUUIDTo32(t *testing.T) {
	id, err := UUIDTo32(uuid.New().String())
	if err != nil {
		t.Error(err)
	}
	t.Log(id)
}

func TestPearson(t *testing.T) {
	id := Pearson("myin3333put")
	t.Log(id)
}

func TestUUIDTo64(t *testing.T) {
	id, err := UUIDTo64(uuid.New().String())
	if err != nil {
		t.Error(err)
	}
	t.Log(UUIDBinary())
	t.Log(id)
}

func BenchmarkUUIDTo32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UUIDTo32(uuid.New().String())
	}
}
