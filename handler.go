// Copyright 2016 someonegg. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package golog

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"
	"time"
)

var LevelName = [...]string{
	LevelDebug: "Debug",
	LevelInfo:  "Info",
	LevelWarn:  "Warn",
	LevelError: "Error",
	LevelPanic: "Panic",
	LevelFatal: "Fatal",
}

var start = time.Now()

// Field represents a log field.
type Field struct {
	K string
	V interface{}
}

// ByKey sorts fields by name.
type byKey []Field

func (a byKey) Len() int           { return len(a) }
func (a byKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byKey) Less(i, j int) bool { return a[i].K < a[j].K }

type handler struct {
	mu sync.Mutex
	w  io.Writer
}

// NewHandler create a handler outputting to w.
// The handler is a development-friendly textual handler.
func NewHandler(w io.Writer) Handler {
	return &handler{
		w: w,
	}
}

func (h *handler) ProcessLog(l *Log) {
	level := LevelName[l.Level]
	level = strings.ToUpper(level[0:4])

	fields := make([]Field, 0, len(l.Fields))

	for k, v := range l.Fields {
		fields = append(fields, Field{k, v})
	}

	sort.Sort(byKey(fields))

	h.mu.Lock()
	defer h.mu.Unlock()

	ts := l.Time.Sub(start) / time.Second
	fmt.Fprintf(h.w, "%4s[%04d] %-20s", level, ts, l.Message)

	for _, f := range fields {
		fmt.Fprintf(h.w, " %s=%v", f.K, f.V)
	}

	fmt.Fprintln(h.w)
}
