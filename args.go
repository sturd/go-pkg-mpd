// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain
// Dedication license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

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
	if v, e := strconv.ParseFloat(this.read(k), 64); e == nil {
		return float32(v)
	}
	return 0
}

func (this Args) F64(k string) float64 {
	if v, e := strconv.ParseFloat(this.read(k), 64); e == nil {
		return v
	}
	return 0
}

func (this Args) I64(k string) int64 {
	if v, e := strconv.ParseInt(this.read(k), 10, 64); e == nil {
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
