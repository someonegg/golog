// Copyright 2016 someonegg. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hjson implements a JSON handler.
package hjson

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"

	. "github.com/someonegg/golog"
)

// Default handler, outputting to Stderr.
var Default = New(os.Stderr)

type handler struct {
	mu sync.Mutex
	tf string //time format
	ec *json.Encoder
}

// New handler, outputting to w.
func New(w io.Writer) Handler {
	return &handler{
		tf: time.RFC3339,
		ec: json.NewEncoder(w),
	}
}

// New handler, outputting to w, with the timeformat.
func New2(w io.Writer, timeformat string) Handler {
	return &handler{
		tf: timeformat,
		ec: json.NewEncoder(w),
	}
}

func (h *handler) ProcessLog(l *Log) {
	h.mu.Lock()
	defer h.mu.Unlock()
	l.Fields["level"] = LevelName[l.Level]
	l.Fields["time"] = l.Time.Format(h.tf)
	l.Fields["message"] = l.Message
	h.ec.Encode(l.Fields)
	delete(l.Fields, "level")
	delete(l.Fields, "time")
	delete(l.Fields, "message")
}
