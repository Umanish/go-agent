// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package newrelic

import (
	"bufio"
	"io"
	"net"
	"net/http"

	"github.com/Umanish/go-agent/internal"
)

func (thd *thread) CloseNotify() <-chan bool {
	return thd.txn.getWriter().(http.CloseNotifier).CloseNotify()
}
func (thd *thread) Flush() {
	thd.txn.getWriter().(http.Flusher).Flush()
}
func (thd *thread) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return thd.txn.getWriter().(http.Hijacker).Hijack()
}
func (thd *thread) ReadFrom(r io.Reader) (int64, error) {
	return thd.txn.getWriter().(io.ReaderFrom).ReadFrom(r)
}

type threadWithExtras interface {
	Transaction
	internal.AddAgentAttributer
	internal.AddAgentSpanAttributer
}

func upgradeTxn(thd *thread) Transaction {
	// Note that thd.txn.getWriter() is not used here.  The transaction is
	// locked (or under construction) when this function is used.

	// GENERATED CODE DO NOT MODIFY
	// This code generated by internal/tools/interface-wrapping
	var (
		i0 int32 = 1 << 0
		i1 int32 = 1 << 1
		i2 int32 = 1 << 2
		i3 int32 = 1 << 3
	)
	var interfaceSet int32
	if _, ok := thd.txn.writer.(http.CloseNotifier); ok {
		interfaceSet |= i0
	}
	if _, ok := thd.txn.writer.(http.Flusher); ok {
		interfaceSet |= i1
	}
	if _, ok := thd.txn.writer.(http.Hijacker); ok {
		interfaceSet |= i2
	}
	if _, ok := thd.txn.writer.(io.ReaderFrom); ok {
		interfaceSet |= i3
	}
	switch interfaceSet {
	default: // No optional interfaces implemented
		return struct {
			threadWithExtras
		}{thd}
	case i0:
		return struct {
			threadWithExtras
			http.CloseNotifier
		}{thd, thd}
	case i1:
		return struct {
			threadWithExtras
			http.Flusher
		}{thd, thd}
	case i0 | i1:
		return struct {
			threadWithExtras
			http.CloseNotifier
			http.Flusher
		}{thd, thd, thd}
	case i2:
		return struct {
			threadWithExtras
			http.Hijacker
		}{thd, thd}
	case i0 | i2:
		return struct {
			threadWithExtras
			http.CloseNotifier
			http.Hijacker
		}{thd, thd, thd}
	case i1 | i2:
		return struct {
			threadWithExtras
			http.Flusher
			http.Hijacker
		}{thd, thd, thd}
	case i0 | i1 | i2:
		return struct {
			threadWithExtras
			http.CloseNotifier
			http.Flusher
			http.Hijacker
		}{thd, thd, thd, thd}
	case i3:
		return struct {
			threadWithExtras
			io.ReaderFrom
		}{thd, thd}
	case i0 | i3:
		return struct {
			threadWithExtras
			http.CloseNotifier
			io.ReaderFrom
		}{thd, thd, thd}
	case i1 | i3:
		return struct {
			threadWithExtras
			http.Flusher
			io.ReaderFrom
		}{thd, thd, thd}
	case i0 | i1 | i3:
		return struct {
			threadWithExtras
			http.CloseNotifier
			http.Flusher
			io.ReaderFrom
		}{thd, thd, thd, thd}
	case i2 | i3:
		return struct {
			threadWithExtras
			http.Hijacker
			io.ReaderFrom
		}{thd, thd, thd}
	case i0 | i2 | i3:
		return struct {
			threadWithExtras
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{thd, thd, thd, thd}
	case i1 | i2 | i3:
		return struct {
			threadWithExtras
			http.Flusher
			http.Hijacker
			io.ReaderFrom
		}{thd, thd, thd, thd}
	case i0 | i1 | i2 | i3:
		return struct {
			threadWithExtras
			http.CloseNotifier
			http.Flusher
			http.Hijacker
			io.ReaderFrom
		}{thd, thd, thd, thd, thd}
	}
}
