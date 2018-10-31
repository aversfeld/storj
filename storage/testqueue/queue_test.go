// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package testqueue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	RunTests(t, New())
}
