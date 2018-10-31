// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package testqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"storj.io/storj/storage"
)

// RunTests runs common storage.Queue tests
func RunTests(t *testing.T, q storage.Queue) {
	t.Run("basic", func(t *testing.T) { testBasic(t, q) })
}

func testBasic(t *testing.T, q storage.Queue) {
	q.Enqueue(storage.Value("hello"))
	q.Enqueue(storage.Value("world"))
	out, err := q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, out, storage.Value("hello"))
	out, err = q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, out, storage.Value("world"))
	out, err = q.Dequeue()
	assert.Nil(t, out)
	assert.NotNil(t, err)
}
