// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import "strconv"

type Args map[string]string

func (this Args) B(k string) bool  { return this.I(k) == 1 }
func (this Args) U8(k string) byte { return byte(this.I(k)) }

func (this Args) I(k string) int {
	if v, e := strconv.Atoi(this.read(k)); e == nil {
		return v
	}
	return 0
}

func (this Args) F32(k string) float32 {
	if v, e := strconv.Atof64(this.read(k)); e == nil {
		return float32(v)
	}
	return 0
}

func (this Args) F64(k string) float64 {
	if v, e := strconv.Atof64(this.read(k)); e == nil {
		return v
	}
	return 0
}

func (this Args) I64(k string) int64 {
	if v, e := strconv.Atoi64(this.read(k)); e == nil {
		return v
	}
	return 0
}

func (this Args) S(k string) string {
	if v := this.read(k); v != "" {
		return v
	}
	return ""
}

func (this Args) read(key string) string {
	v, ok := this[key]
	if !ok {
		return ""
	}
	return v
}
