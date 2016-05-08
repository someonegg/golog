// Copyright 2016 someonegg. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hmulti implements a handler which invokes
// multi handlers.
package hmulti

import (
	. "github.com/someonegg/golog"
)

type handler struct {
	hs []Handler
}

func New(hs ...Handler) Handler {
	return &handler{
		hs: hs,
	}
}

func (h *handler) ProcessLog(l *Log) {
	for _, hh := range h.hs {
		hh.ProcessLog(l)
	}
}
