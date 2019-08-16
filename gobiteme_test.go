package gobiteme_test

import (
	"testing"

	gobiteme "github.com/bgokden/go-bite-me"
	"github.com/stretchr/testify/assert"
)

func TestGoBiteMe(t *testing.T) {
	b := gobiteme.New()
	b.Bite([]byte("key1"), []byte("value1"))
	b.Bite([]byte("key2"), []byte("value2"))
	b.Bite([]byte("key3"), []byte("value3"))
	b.Bite([]byte("key2"), []byte("value3"))

	count := 0
	b.ThrowUp(func(e *gobiteme.Bite) {
		t.Logf("Key: %v\n", string(e.Key))
		count++
	})

	assert.Equal(t, 3, count)
}
