// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package testsuite

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"storj.io/storj/storage"
)

// RunQueueTests runs common storage.Queue tests
func RunQueueTests(t *testing.T, q storage.Queue) {
	t.Run("basic", func(t *testing.T) { testBasic(t, q) })
}

func testBasic(t *testing.T, q storage.Queue) {
	err := q.Enqueue(storage.Value("hello"))
	assert.NoError(t, err)
	err = q.Enqueue(storage.Value("world"))
	assert.NoError(t, err)
	out, err := q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, out, storage.Value("hello"))
	out, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, out, storage.Value("world"))
	out, err = q.Dequeue()
	assert.Nil(t, out)
	assert.Error(t, err)
}
